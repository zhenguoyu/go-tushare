# gotushare

A high-performance, elegant, and extensible Golang SDK for Tushare financial data APIs.  
Empower your quantitative research, data analysis, and financial engineering with ease and reliability.

高性能、优雅且可扩展的 Tushare 金融数据 Golang SDK。  
助力您的量化研究、数据分析与金融工程开发，轻松高效，值得信赖。

---

## Features | 功能特性

- 🚀 Fast and robust HTTP client with customizable timeout and connection pooling  
  快速且健壮的 HTTP 客户端，支持自定义超时与连接池
- 📦 Simple and intuitive API design  
  简洁直观的 API 设计
- 🔒 Secure and reliable data access  
  安全可靠的数据访问
- 🛠️ Easy integration with your own storage or analytics pipeline  
  便于集成到自有存储或分析流程

## Quick Start | 快速开始

```go
import "github.com/zhenguoyu/gotushare"

client := gotushare.NewTushare("your_token")
resp, err := client.RequestApi("api_name", map[string]string{"param": "value"}, "field1,field2")
if err != nil {
    // handle error
}
// use resp.Data
```

## License | 许可证

MIT License

---

欢迎 Star & PR，期待与全球开发者共建高质量金融数据生态！
Welcome contributions and stars from developers worldwide!