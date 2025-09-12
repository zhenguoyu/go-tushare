package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// HttpClientConfig 定义HTTP客户端配置结构
type HttpClientConfig struct {
	Timeout               time.Duration // 整体请求超时时间
	DialTimeout           time.Duration // TCP连接建立超时
	TLSHandshakeTimeout   time.Duration // TLS握手超时
	ResponseHeaderTimeout time.Duration // 响应头超时
	IdleConnTimeout       time.Duration // 空闲连接超时
	MaxIdleConns          int           // 最大空闲连接数
	MaxIdleConnsPerHost   int           // 每个主机最大空闲连接数
	DisableKeepAlives     bool          // 是否禁用长连接
}

// httpClient 自定义HTTP客户端
type HttpClient struct {
	client    *http.Client
	transport *http.Transport
}

// NewHttpClient 创建新的自定义HTTP客户端
func NewHttpClient(config HttpClientConfig) *HttpClient {
	// 创建自定义Transport
	setDefaultHttpClientConfig(&config)
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   config.DialTimeout,
			KeepAlive: 30 * time.Second, // 默认keep-alive时间
		}).DialContext,
		TLSHandshakeTimeout:   config.TLSHandshakeTimeout,
		ResponseHeaderTimeout: config.ResponseHeaderTimeout,
		IdleConnTimeout:       config.IdleConnTimeout,
		MaxIdleConns:          config.MaxIdleConns,
		MaxIdleConnsPerHost:   config.MaxIdleConnsPerHost,
		DisableKeepAlives:     config.DisableKeepAlives,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false, // 生产环境应为false
		},
		ForceAttemptHTTP2: true, // 尝试使用HTTP/2
	}

	// 创建HTTP客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
	}

	return &HttpClient{
		client:    client,
		transport: transport,
	}
}

// DefaultHttpClientConfig 返回默认配置
func setDefaultHttpClientConfig(config *HttpClientConfig) {
	if config.Timeout == 0 {
		config.Timeout = 5 * time.Second
	}
	if config.DialTimeout == 0 {
		config.DialTimeout = 2 * time.Second
	}
	if config.TLSHandshakeTimeout == 0 {
		config.TLSHandshakeTimeout = 2 * time.Second
	}
	if config.ResponseHeaderTimeout == 0 {
		config.ResponseHeaderTimeout = 2 * time.Second
	}
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 10
	}
}

// DoRequest 执行HTTP请求（支持Context）
func (c *HttpClient) DoRequestWithContext(ctx context.Context, method, url string, body io.Reader, headers map[string]string) ([]byte, int, error) {
	// 创建请求
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, 0, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := c.client.Do(req)
	if err != nil {
		// 检查是否是超时或取消错误
		if ctx.Err() == context.DeadlineExceeded {
			return nil, 0, fmt.Errorf("请求超时: %w", err)
		}
		if ctx.Err() == context.Canceled {
			return nil, 0, fmt.Errorf("请求被取消: %w", err)
		}
		return nil, 0, fmt.Errorf("请求执行失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("读取响应体失败: %v", err)
	}

	return responseBody, resp.StatusCode, nil
}

// GetWithContext 执行GET请求
func (c *HttpClient) GetWithContext(ctx context.Context, url string, headers map[string]string) ([]byte, int, error) {
	return c.DoRequestWithContext(ctx, http.MethodGet, url, nil, headers)
}

// PostWithContext 执行POST请求
func (c *HttpClient) PostWithContext(ctx context.Context, url string, body io.Reader, headers map[string]string) ([]byte, int, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	if _, exists := headers["Content-Type"]; !exists {
		headers["Content-Type"] = "application/json"
	}
	return c.DoRequestWithContext(ctx, http.MethodPost, url, body, headers)
}

// Close 关闭客户端，释放资源
func (c *HttpClient) Close() {
	// 关闭空闲连接
	c.transport.CloseIdleConnections()
}

// Get 执行GET请求
func (c *HttpClient) Get(url string, headers map[string]string) ([]byte, int, error) {
	return c.GetWithContext(context.Background(), url, headers)
}

// Post 执行POST请求
func (c *HttpClient) Post(url string, body io.Reader, headers map[string]string) ([]byte, int, error) {
	return c.PostWithContext(context.Background(), url, body, headers)

}
