// stock.go 股票接口相关数据结构
package model

import "encoding/json"

// api name
// 股票列表 api name
const ApiStockBasic string = "stock_basic"

// 交易日历 api name
const ApiTradeCal string = "trade_cal"

// 股票曾用名 api name
const ApiNameChange string = "namechange"

// 沪深股通成份股 api name
const ApiHsConst string = "hs_const"

// 上市公司基本信息 api name
const ApiStockCompany string = "stock_company"

// 上市公司管理层 api name
const ApiStkManagers string = "stk_managers"

// 管理层薪酬和持股 api name
const ApiStkRewards string = "stk_rewards"

// IPO新股列表 api name
const ApiNewShare string = "new_share"

// 备用列表 api name
const ApiBakBasic string = "bak_basic"

// 日线行情 api name
const ApiDaily string = "daily"

// 周线行情 api name
const ApiWeekly string = "weekly"

// 月线行情 api name
const ApiMonthly string = "monthly"

// 复权因子 api name
const ApiAdjFactor string = "adj_factor"

// 每日指标 api name
const ApiDailyBasic string = "daily_basic"

// 个股资金流向 api name
const ApiMoneyFlow string = "moneyflow"

// 每日涨跌停价格 api name
const ApiStkLimit string = "stk_limit"

// 每日停复牌信息 api name
const ApiSuspendd string = "suspend_d"

// 沪深港通资金流向 api name
const ApiMoneyFlowHsgt string = "moneyflow_hsgt"

// 沪深股通十大成交股 api name
const ApiHsgtTop10 string = "hsgt_top10"

// 港股通十大成交股 api name
const ApiGgtTop10 string = "ggt_top10"

// 港股通每日成交统计 api name
const ApiGgtDaily string = "ggt_daily"

// 港股通每月成交统计 api name
const ApiGgtMonthly string = "ggt_monthly"

// 备用行情 api name
const ApiBakDaily string = "bak_daily"

// 利润表 api name
const ApiIncome string = "income"

// 资产负债表 api name
const ApiBalanceSheet string = "balancesheet"

// 现金流量表 api name
const ApiCashFlow string = "cashflow"

// 业绩预告 api name
const ApiForeCast string = "forecast"

// 业绩快报 api name
const ApiExpress string = "express"

// 分红送股 api name
const ApiDividend string = "dividend"

// 财务指标 api name
const ApiFinaIndicator string = "fina_indicator"

// 财务审计意见 api name
const ApiFinaAudit string = "fina_audit"

// 主营业务构成 api name
const ApiFinaMainbz string = "fina_mainbz"

// 财报披露计划 api name
const ApiDisclosureDate string = "disclosure_date"

// 融资融券交易汇总 api name
const ApiMargin string = "margin"

// 融资融券交易明细 api name
const ApiMarginDetail string = "margin_detail"

// 融资融券标的 api name
const ApiMarginTarget string = "margin_target"

// 融资融券标的(盘前) api name
const ApiMarginSecs string = "margin_secs"

// 前十大股东 api name
const ApiTop10Holders string = "top10_holders"

// 前十大流通股东 api name
const ApiTop10Floatholders string = "top10_floatholders"

// 龙虎榜每日明细 api name
const ApiTopList string = "top_list"

// 龙虎榜机构明细 api name
const ApiTopInst string = "top_inst"

// 股权质押统计数据 api name
const ApiPledgeStat string = "pledge_stat"

// 股权质押明细 api name
const ApiPledgeDetail string = "pledge_detail"

// 股票回购 api name
const ApiRepurchase string = "repurchase"

// 概念股分类 api name
const ApiConcept string = "concept"

// 概念股列表 api name
const ApiConceptDetail string = "concept_detail"

// 限售股解禁 api name
const ApiShareFloat string = "share_float"

// 大宗交易 api name
const ApiBlockTrade string = "block_trade"

// 股东人数 api name
const ApiStkHolderNumber string = "stk_holdernumber"

// 股东增减持 api name
const ApiStkHolderTrade string = "stk_holdertrade"

// 券商盈利预测数据 api name
const ApiReportRc string = "report_rc"

// 每日筹码及胜率 api name
const ApiCyqPerf string = "cyq_perf"

// 每日筹码分布 api name
const ApiCyqChips string = "cyq_chips"

// 股票技术因子 api name
const ApiStkFactor string = "stk_factor"

// 中央结算系统持股统计 api name
const ApiCcassHold string = "ccass_hold"

// 中央结算系统持股明细 api name
const ApiCcasHoldDetail string = "ccass_hold_detail"

// 沪深港股通持股明细 api name
const ApiHkHold string = "hk_hold"

// 涨跌停和炸板数据 api name
const ApiLimitListd string = "limit_list_d"

// 机构调研数据 api name
const ApiStkSurv string = "stk_surv"

// 券商月度金股 api name
const ApiBrokerRecommend string = "broker_recommend"

// 游资名录 api name
const ApiHmList string = "hm_list"

// 游资每日明细 api name
const ApiHmDetail string = "hm_detail"

// 同花顺App热榜 api name
const ApiThsHot string = "ths_hot"

// 东方财富App热榜 api name
const ApiDcHot string = "dc_hot"

// fields

// 股票列表 fields
var FieldsStockBasic = []string{"ts_code", "symbol", "name", "area", "industry", "fullname", "enname", "cnspell", "market", "exchange", "curr_type", "list_status", "list_date", "delist_date", "is_hs", "act_name", "act_ent_type"}

// 交易日历 fields
var FieldsTradeCal = []string{"exchange", "cal_date", "is_open", "pretrade_date"}

// 股票曾用名 fields
var FieldsNameChange = []string{"ts_code", "name", "start_date", "end_date", "ann_date", "change_reason"}

// 沪深股通成份股 fields
var FieldsHsConst = []string{"ts_code", "hs_type", "in_date", "out_date", "is_new"}

// 上市公司基本信息 fields
var FieldsStockCompany = []string{"ts_code", "exchange", "chairman", "manager", "secretary", "reg_capital", "setup_date", "province", "city", "introduction", "website", "email", "office", "employees", "main_business", "business_scope"}

// 上市公司管理层 fields
var FieldsStkManagers = []string{"ts_code", "ann_date", "name", "gender", "lev", "title", "edu", "national", "birthday", "begin_date", "end_date", "resume"}

// 管理层薪酬和持股 fields
var FieldsStkRewards = []string{"ts_code", "ann_date", "end_date", "name", "title", "reward", "hold_vol"}

// IPO新股列表 fields
var FieldsNewShare = []string{"ts_code", "sub_code", "name", "ipo_date", "issue_date", "amount", "market_amount", "price", "pe", "limit_amount", "funds", "ballot"}

// 备用列表 fields
var FieldsBakBasic = []string{"trade_date", "ts_code", "name", "industry", "area", "pe", "float_share", "total_share", "total_assets", "liquid_assets", "fixed_assets", "reserved", "reserved_pershare", "eps", "bvps", "pb", "list_date", "undp", "per_undp", "rev_yoy", "profit_yoy", "gpr", "npr", "holder_num"}

// 日线行情 fields
var FieldsDaily = []string{"ts_code", "trade_date", "open", "high", "low", "close", "pre_close", "change", "pct_chg", "vol", "amount"}

// 周线行情 fields
var FieldsWeekly = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"}

// 月线行情 fields
var FieldsMonthly = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"}

// 复权因子 fields
var FieldsAdjFactor = []string{"ts_code", "trade_date", "adj_factor"}

// 每日指标 fields
var FieldsDailyBasic = []string{"ts_code", "trade_date", "close", "turnover_rate", "turnover_rate_f", "volume_ratio", "pe", "pe_ttm", "pb", "ps", "ps_ttm", "dv_ratio", "dv_ttm", "total_share", "float_share", "free_share", "total_mv", "circ_mv"}

// 个股资金流向 fields
var FieldsMoneyFlow = []string{"ts_code", "trade_date", "buy_sm_vol", "buy_sm_amount", "sell_sm_vol", "sell_sm_amount", "buy_md_vol", "buy_md_amount", "sell_md_vol", "sell_md_amount", "buy_lg_vol", "buy_lg_amount", "sell_lg_vol", "sell_lg_amount", "buy_elg_vol", "buy_elg_amount", "sell_elg_vol", "sell_elg_amount", "net_mf_vol", "net_mf_amount"}

// 每日涨跌停价格 fields
var FieldsStkLimit = []string{"trade_date", "ts_code", "pre_close", "up_limit", "down_limit"}

// 每日停复牌信息 fields
var FieldsSuspendd = []string{"ts_code", "trade_date", "suspend_timing", "suspend_type"}

// 沪深港通资金流向 fields
var FieldsMoneyFlowHsgt = []string{"trade_date", "ggt_ss", "ggt_sz", "hgt", "sgt", "north_money", "south_money"}

// 沪深股通十大成交股 fields
var FieldsHsgtTop10 = []string{"trade_date", "ts_code", "name", "close", "change", "rank", "market_type", "amount", "net_amount", "buy", "sell"}

// 港股通十大成交股 fields
var FieldsGgtTop10 = []string{"trade_date", "ts_code", "name", "close", "p_change", "rank", "market_type", "amount", "net_amount", "sh_amount", "sh_net_amount", "sh_buy", "sh_sell", "sz_amount", "sz_net_amount", "sz_buy", "sz_sell"}

// 港股通每日成交统计 fields
var FieldsGgtDaily = []string{"trade_date", "buy_amount", "buy_volume", "sell_amount", "sell_volume"}

// 港股通每月成交统计 fields
var FieldsGgtMonthly = []string{"month", "day_buy_amt", "day_buy_vol", "day_sell_amt", "day_sell_vol", "total_buy_amt", "total_buy_vol", "total_sell_amt", "total_sell_vol"}

// 备用行情 fields
var FieldsBakDaily = []string{"ts_code", "trade_date", "name", "pct_change", "close", "change", "open", "high", "low", "pre_close", "vol_ratio", "turn_over", "swing", "vol", "amount", "selling", "buying", "total_share", "float_share", "pe", "industry", "area", "float_mv", "total_mv", "avg_price", "strength", "activity", "avg_turnover", "attack", "interval_3", "interval_6"}

// 利润表 fields
var FieldsIncome = []string{"ts_code", "ann_date", "f_ann_date", "end_date", "report_type", "comp_type", "end_type", "basic_eps", "diluted_eps", "total_revenue", "revenue", "int_income", "prem_earned", "comm_income", "n_commis_income", "n_oth_income", "n_oth_b_income", "prem_income", "out_prem", "une_prem_reser", "reins_income", "n_sec_tb_income", "n_sec_uw_income", "n_asset_mg_income", "oth_b_income", "fv_value_chg_gain", "invest_income", "ass_invest_income", "forex_gain", "total_cogs", "oper_cost", "int_exp", "comm_exp", "biz_tax_surchg", "sell_exp", "admin_exp", "fin_exp", "assets_impair_loss", "prem_refund", "compens_payout", "reser_insur_liab", "div_payt", "reins_exp", "oper_exp", "compens_payout_refu", "insur_reser_refu", "reins_cost_refund", "other_bus_cost", "operate_profit", "non_oper_income", "non_oper_exp", "nca_disploss", "total_profit", "income_tax", "n_income", "n_income_attr_p", "minority_gain", "oth_compr_income", "t_compr_income", "compr_inc_attr_p", "compr_inc_attr_m_s", "ebit", "ebitda", "insurance_exp", "undist_profit", "distable_profit", "rd_exp", "fin_exp_int_exp", "fin_exp_int_inc", "transfer_surplus_rese", "transfer_housing_imprest", "transfer_oth", "adj_lossgain", "withdra_legal_surplus", "withdra_legal_pubfund", "withdra_biz_devfund", "withdra_rese_fund", "withdra_oth_ersu", "workers_welfare", "distr_profit_shrhder", "prfshare_payable_dvd", "comshare_payable_dvd", "capit_comstock_div", "net_after_nr_lp_correct", "credit_impa_loss", "net_expo_hedging_benefits", "oth_impair_loss_assets", "total_opcost", "amodcost_fin_assets", "oth_income", "asset_disp_income", "continued_net_profit", "end_net_profit", "update_flag"}

// 资产负债表 fields
var FieldsBalanceSheet = []string{"ts_code", "ann_date", "f_ann_date", "end_date", "report_type", "comp_type", "end_type", "total_share", "cap_rese", "undistr_porfit", "surplus_rese", "special_rese", "money_cap", "trad_asset", "notes_receiv", "accounts_receiv", "oth_receiv", "prepayment", "div_receiv", "int_receiv", "inventories", "amor_exp", "nca_within_1y", "sett_rsrv", "loanto_oth_bank_fi", "premium_receiv", "reinsur_receiv", "reinsur_res_receiv", "pur_resale_fa", "oth_cur_assets", "total_cur_assets", "fa_avail_for_sale", "htm_invest", "lt_eqt_invest", "invest_real_estate", "time_deposits", "oth_assets", "lt_rec", "fix_assets", "cip", "const_materials", "fixed_assets_disp", "produc_bio_assets", "oil_and_gas_assets", "intan_assets", "r_and_d", "goodwill", "lt_amor_exp", "defer_tax_assets", "decr_in_disbur", "oth_nca", "total_nca", "cash_reser_cb", "depos_in_oth_bfi", "prec_metals", "deriv_assets", "rr_reins_une_prem", "rr_reins_outstd_cla", "rr_reins_lins_liab", "rr_reins_lthins_liab", "refund_depos", "ph_pledge_loans", "refund_cap_depos", "indep_acct_assets", "client_depos", "client_prov", "transac_seat_fee", "invest_as_receiv", "total_assets", "lt_borr", "st_borr", "cb_borr", "depos_ib_deposits", "loan_oth_bank", "trading_fl", "notes_payable", "acct_payable", "adv_receipts", "sold_for_repur_fa", "comm_payable", "payroll_payable", "taxes_payable", "int_payable", "div_payable", "oth_payable", "acc_exp", "deferred_inc", "st_bonds_payable", "payable_to_reinsurer", "rsrv_insur_cont", "acting_trading_sec", "acting_uw_sec", "non_cur_liab_due_1y", "oth_cur_liab", "total_cur_liab", "bond_payable", "lt_payable", "specific_payables", "estimated_liab", "defer_tax_liab", "defer_inc_non_cur_liab", "oth_ncl", "total_ncl", "depos_oth_bfi", "deriv_liab", "depos", "agency_bus_liab", "oth_liab", "prem_receiv_adva", "depos_received", "ph_invest", "reser_une_prem", "reser_outstd_claims", "reser_lins_liab", "reser_lthins_liab", "indept_acc_liab", "pledge_borr", "indem_payable", "policy_div_payable", "total_liab", "treasury_share", "ordin_risk_reser", "forex_differ", "invest_loss_unconf", "minority_int", "total_hldr_eqy_exc_min_int", "total_hldr_eqy_inc_min_int", "total_liab_hldr_eqy", "lt_payroll_payable", "oth_comp_income", "oth_eqt_tools", "oth_eqt_tools_p_shr", "lending_funds", "acc_receivable", "st_fin_payable", "payables", "hfs_assets", "hfs_sales", "cost_fin_assets", "fair_value_fin_assets", "cip_total", "oth_pay_total", "long_pay_total", "debt_invest", "oth_debt_invest", "oth_eq_invest", "oth_illiq_fin_assets", "oth_eq_ppbond", "receiv_financing", "use_right_assets", "lease_liab", "contract_assets", "contract_liab", "accounts_receiv_bill", "accounts_pay", "oth_rcv_total", "fix_assets_total", "update_flag"}

// 现金流量表 fields
var FieldsCashFlow = []string{"ts_code", "ann_date", "f_ann_date", "end_date", "comp_type", "report_type", "end_type", "net_profit", "finan_exp", "c_fr_sale_sg", "recp_tax_rends", "n_depos_incr_fi", "n_incr_loans_cb", "n_inc_borr_oth_fi", "prem_fr_orig_contr", "n_incr_insured_dep", "n_reinsur_prem", "n_incr_disp_tfa", "ifc_cash_incr", "n_incr_disp_faas", "n_incr_loans_oth_bank", "n_cap_incr_repur", "c_fr_oth_operate_a", "c_inf_fr_operate_a", "c_paid_goods_s", "c_paid_to_for_empl", "c_paid_for_taxes", "n_incr_clt_loan_adv", "n_incr_dep_cbob", "c_pay_claims_orig_inco", "pay_handling_chrg", "pay_comm_insur_plcy", "oth_cash_pay_oper_act", "st_cash_out_act", "n_cashflow_act", "oth_recp_ral_inv_act", "c_disp_withdrwl_invest", "c_recp_return_invest", "n_recp_disp_fiolta", "n_recp_disp_sobu", "stot_inflows_inv_act", "c_pay_acq_const_fiolta", "c_paid_invest", "n_disp_subs_oth_biz", "oth_pay_ral_inv_act", "n_incr_pledge_loan", "stot_out_inv_act", "n_cashflow_inv_act", "c_recp_borrow", "proc_issue_bonds", "oth_cash_recp_ral_fnc_act", "stot_cash_in_fnc_act", "free_cashflow", "c_prepay_amt_borr", "c_pay_dist_dpcp_int_exp", "incl_dvd_profit_paid_sc_ms", "oth_cashpay_ral_fnc_act", "stot_cashout_fnc_act", "n_cash_flows_fnc_act", "eff_fx_flu_cash", "n_incr_cash_cash_equ", "c_cash_equ_beg_period", "c_cash_equ_end_period", "c_recp_cap_contrib", "incl_cash_rec_saims", "uncon_invest_loss", "prov_depr_assets", "depr_fa_coga_dpba", "amort_intang_assets", "lt_amort_deferred_exp", "decr_deferred_exp", "incr_acc_exp", "loss_disp_fiolta", "loss_scr_fa", "loss_fv_chg", "invest_loss", "decr_def_inc_tax_assets", "incr_def_inc_tax_liab", "decr_inventories", "decr_oper_payable", "incr_oper_payable", "others", "im_net_cashflow_oper_act", "conv_debt_into_cap", "conv_copbonds_due_within_1y", "fa_fnc_leases", "im_n_incr_cash_equ", "net_dism_capital_add", "net_cash_rece_sec", "credit_impa_loss", "use_right_asset_dep", "oth_loss_asset", "end_bal_cash", "beg_bal_cash", "end_bal_cash_equ", "beg_bal_cash_equ", "update_flag"}

// 业绩预告 fields
var FieldsForeCast = []string{"ts_code", "ann_date", "end_date", "type", "p_change_min", "p_change_max", "net_profit_min", "net_profit_max", "last_parent_net", "first_ann_date", "summary", "change_reason"}

// 业绩快报 fields
var FieldsExpress = []string{"ts_code", "ann_date", "end_date", "revenue", "operate_profit", "total_profit", "n_income", "total_assets", "total_hldr_eqy_exc_min_int", "diluted_eps", "diluted_roe", "yoy_net_profit", "bps", "yoy_sales", "yoy_op", "yoy_tp", "yoy_dedu_np", "yoy_eps", "yoy_roe", "growth_assets", "yoy_equity", "growth_bps", "or_last_year", "op_last_year", "tp_last_year", "np_last_year", "eps_last_year", "open_net_assets", "open_bps", "perf_summary", "is_audit", "remark"}

// 分红送股 fields
var FieldsDividend = []string{"ts_code", "end_date", "ann_date", "div_proc", "stk_div", "stk_bo_rate", "stk_co_rate", "cash_div", "cash_div_tax", "record_date", "ex_date", "pay_date", "div_listdate", "imp_ann_date", "base_date", "base_share"}

// 财务指标 fields
var FieldsFinaIndicator = []string{"ts_code", "ann_date", "end_date", "eps", "dt_eps", "total_revenue_ps", "revenue_ps", "capital_rese_ps", "surplus_rese_ps", "undist_profit_ps", "extra_item", "profit_dedt", "gross_margin", "current_ratio", "quick_ratio", "cash_ratio", "invturn_days", "arturn_days", "inv_turn", "ar_turn", "ca_turn", "fa_turn", "assets_turn", "op_income", "valuechange_income", "interst_income", "daa", "ebit", "ebitda", "fcff", "fcfe", "current_exint", "noncurrent_exint", "interestdebt", "netdebt", "tangible_asset", "working_capital", "networking_capital", "invest_capital", "retained_earnings", "diluted2_eps", "bps", "ocfps", "retainedps", "cfps", "ebit_ps", "fcff_ps", "fcfe_ps", "netprofit_margin", "grossprofit_margin", "cogs_of_sales", "expense_of_sales", "profit_to_gr", "saleexp_to_gr", "adminexp_of_gr", "finaexp_of_gr", "impai_ttm", "gc_of_gr", "op_of_gr", "ebit_of_gr", "roe", "roe_waa", "roe_dt", "roa", "npta", "roic", "roe_yearly", "roa2_yearly", "roe_avg", "opincome_of_ebt", "investincome_of_ebt", "n_op_profit_of_ebt", "tax_to_ebt", "dtprofit_to_profit", "salescash_to_or", "ocf_to_or", "ocf_to_opincome", "capitalized_to_da", "debt_to_assets", "assets_to_eqt", "dp_assets_to_eqt", "ca_to_assets", "nca_to_assets", "tbassets_to_totalassets", "int_to_talcap", "eqt_to_talcapital", "currentdebt_to_debt", "longdeb_to_debt", "ocf_to_shortdebt", "debt_to_eqt", "eqt_to_debt", "eqt_to_interestdebt", "tangibleasset_to_debt", "tangasset_to_intdebt", "tangibleasset_to_netdebt", "ocf_to_debt", "ocf_to_interestdebt", "ocf_to_netdebt", "ebit_to_interest", "longdebt_to_workingcapital", "ebitda_to_debt", "turn_days", "roa_yearly", "roa_dp", "fixed_assets", "profit_prefin_exp", "non_op_profit", "op_to_ebt", "nop_to_ebt", "ocf_to_profit", "cash_to_liqdebt", "cash_to_liqdebt_withinterest", "op_to_liqdebt", "op_to_debt", "roic_yearly", "total_fa_trun", "profit_to_op", "q_opincome", "q_investincome", "q_dtprofit", "q_eps", "q_netprofit_margin", "q_gsprofit_margin", "q_exp_to_sales", "q_profit_to_gr", "q_saleexp_to_gr", "q_adminexp_to_gr", "q_finaexp_to_gr", "q_impair_to_gr_ttm", "q_gc_to_gr", "q_op_to_gr", "q_roe", "q_dt_roe", "q_npta", "q_opincome_to_ebt", "q_investincome_to_ebt", "q_dtprofit_to_profit", "q_salescash_to_or", "q_ocf_to_sales", "q_ocf_to_or", "basic_eps_yoy", "dt_eps_yoy", "cfps_yoy", "op_yoy", "ebt_yoy", "netprofit_yoy", "dt_netprofit_yoy", "ocf_yoy", "roe_yoy", "bps_yoy", "assets_yoy", "eqt_yoy", "tr_yoy", "or_yoy", "q_gr_yoy", "q_gr_qoq", "q_sales_yoy", "q_sales_qoq", "q_op_yoy", "q_op_qoq", "q_profit_yoy", "q_profit_qoq", "q_netprofit_yoy", "q_netprofit_qoq", "equity_yoy", "rd_exp", "update_flag"}

// 财务审计意见 fields
var FieldsFinaAudit = []string{"ts_code", "ann_date", "end_date", "audit_result", "audit_fees", "audit_agency", "audit_sign"}

// 主营业务构成 fields
var FieldsFinaMainbz = []string{"ts_code", "end_date", "bz_item", "bz_sales", "bz_profit", "bz_cost", "curr_type", "update_flag"}

// 财报披露计划 fields
var FieldsDisclosureDate = []string{"ts_code", "ann_date", "end_date", "pre_date", "actual_date", "modify_date"}

// 融资融券交易汇总 fields
var FieldsMargin = []string{"trade_date", "exchange_id", "rzye", "rzmre", "rzche", "rqye", "rqmcl", "rzrqye", "rqyl"}

// 融资融券交易明细 fields
var FieldsMarginDetail = []string{"trade_date", "ts_code", "name", "rzye", "rqye", "rzmre", "rqyl", "rzche", "rqchl", "rqmcl", "rzrqye"}

// 融资融券标的 fields
var FieldsMarginTarget = []string{"ts_code", "mg_type", "is_new", "in_date", "out_date", "ann_date"}

// 融资融券标的(盘前) fields
var FieldsMarginSecs = []string{"ts_code", "trade_date", "exchange", "start_date", "end_date"}

// 前十大股东 fields
var FieldsTop10Holders = []string{"ts_code", "ann_date", "end_date", "holder_name", "hold_amount", "hold_ratio", "hold_float_ratio", "hold_change", "holder_type"}

// 前十大流通股东 fields
var FieldsTop10Floatholders = []string{"ts_code", "ann_date", "end_date", "holder_name", "hold_amount", "hold_ratio", "hold_float_ratio", "hold_change", "holder_type"}

// 龙虎榜每日明细 fields
var FieldsTopList = []string{"trade_date", "ts_code", "name", "close", "pct_change", "turnover_rate", "amount", "l_sell", "l_buy", "l_amount", "net_amount", "net_rate", "amount_rate", "float_values", "reason"}

// 龙虎榜机构明细 fields
var FieldsTopInst = []string{"trade_date", "ts_code", "exalter", "side", "buy", "buy_rate", "sell", "sell_rate", "net_buy", "reason"}

// 股权质押统计数据 fields
var FieldsPledgeStat = []string{"ts_code", "end_date", "pledge_count", "unrest_pledge", "rest_pledge", "total_share", "pledge_ratio"}

// 股权质押明细 fields
var FieldsPledgeDetail = []string{"ts_code", "ann_date", "holder_name", "pledge_amount", "start_date", "end_date", "is_release", "release_date", "pledgor", "holding_amount", "pledged_amount", "p_total_ratio", "h_total_ratio", "is_buyback"}

// 股票回购 fields
var FieldsRepurchase = []string{"ts_code", "ann_date", "end_date", "proc", "exp_date", "vol", "amount", "high_limit", "low_limit"}

// 概念股分类 fields
var FieldsConcept = []string{"code", "name", "src"}

// 概念股列表 fields
var FieldsConceptDetail = []string{"id", "concept_name", "ts_code", "name", "in_date", "out_date"}

// 限售股解禁 fields
var FieldsShareFloat = []string{"ts_code", "ann_date", "float_date", "float_share", "float_ratio", "holder_name", "share_type"}

// 大宗交易 fields
var FieldsBlockTrade = []string{"ts_code", "trade_date", "price", "vol", "amount", "buyer", "seller"}

// 股东人数 fields
var FieldsStkHolderNumber = []string{"ts_code", "ann_date", "end_date", "holder_num"}

// 股东增减持 fields
var FieldsStkHolderTrade = []string{"ts_code", "ann_date", "holder_name", "holder_type", "in_de", "change_vol", "change_ratio", "after_share", "after_ratio", "avg_price", "total_share", "begin_date", "close_date"}

// 券商盈利预测数据 fields
var FieldsReportRc = []string{"ts_code", "name", "report_date", "report_title", "report_type", "classify", "org_name", "author_name", "quarter", "op_rt", "op_pr", "tp", "np", "eps", "pe", "rd", "roe", "ev_ebitda", "rating", "max_price", "min_price", "imp_dg", "create_time"}

// 每日筹码及胜率 fields
var FieldsCyqPerf = []string{"ts_code", "trade_date", "his_low", "his_high", "cost_5pct", "cost_15pct", "cost_50pct", "cost_85pct", "cost_95pct", "weight_avg", "winner_rate"}

// 每日筹码分布 fields
var FieldsCyqChips = []string{"ts_code", "trade_date", "price", "percent"}

// 股票技术因子 fields
var FieldsStkFactor = []string{"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_change", "vol", "amount", "adj_factor", "open_hfq", "open_qfq", "close_hfq", "close_qfq", "high_hfq", "high_qfq", "low_hfq", "low_qfq", "pre_close_hfq", "pre_close_qfq", "macd_dif", "macd_dea", "macd", "kdj_k", "kdj_d", "kdj_j", "rsi_6", "rsi_12", "rsi_24", "boll_upper", "boll_mid", "boll_lower", "cci"}

// 中央结算系统持股统计 fields
var FieldsCcassHold = []string{"trade_date", "ts_code", "name", "shareholding", "hold_nums", "hold_ratio"}

// 中央结算系统持股明细 fields
var FieldsCcasHoldDetail = []string{"trade_date", "ts_code", "name", "col_participant_id", "col_participant_name", "col_shareholding", "col_shareholding_percent"}

// 沪深港股通持股明细 fields
var FieldsHkHold = []string{"code", "trade_date", "ts_code", "name", "vol", "ratio", "exchange"}

// 涨跌停和炸板数据 fields
var FieldsLimitListd = []string{"trade_date", "ts_code", "industry", "name", "close", "pct_chg", "amount", "limit_amount", "float_mv", "total_mv", "turnover_ratio", "fd_amount", "first_time", "last_time", "open_times", "up_stat", "limit_times", "limit"}

// 机构调研数据 fields
var FieldsStkSurv = []string{"ts_code", "name", "surv_date", "fund_visitors", "rece_place", "rece_mode", "rece_org", "org_type", "comp_rece", "content"}

// 券商月度金股 fields
var FieldsBrokerRecommend = []string{"month", "broker", "ts_code", "name"}

// 游资名录 fields
var FieldsHmList = []string{"name", "desc", "orgs"}

// 游资每日明细 fields
var FieldsHmDetail = []string{"trade_date", "ts_code", "ts_name", "buy_amount", "sell_amount", "net_amount", "hm_name", "hm_orgs", "tag"}

// 同花顺App热榜 fields
var FieldsThsHot = []string{"trade_date", "data_type", "ts_code", "ts_name", "rank", "pct_change", "current_price", "concept", "rank_reason", "hot", "rank_time"}

// 东方财富App热榜 fields
var FieldsDcHot = []string{"trade_date", "data_type", "ts_code", "ts_name", "rank", "pct_change", "current_price", "rank_time"}

// struct

// 股票列表|stock_basic
type StockBasic struct {
	TsCode     string `json:"ts_code"`      // TS代码
	Symbol     string `json:"symbol"`       // 股票代码
	Name       string `json:"name"`         // 股票名称
	Area       string `json:"area"`         // 地域
	Industry   string `json:"industry"`     // 所属行业
	Fullname   string `json:"fullname"`     // 股票全称
	Enname     string `json:"enname"`       // 英文全称
	Cnspell    string `json:"cnspell"`      // 拼音缩写
	Market     string `json:"market"`       // 市场类型（主板/创业板/科创板/CDR）
	Exchange   string `json:"exchange"`     // 交易所代码
	CurrType   string `json:"curr_type"`    // 交易货币
	ListStatus string `json:"list_status"`  // 上市状态 L上市 D退市 P暂停上市
	ListDate   string `json:"list_date"`    // 上市日期
	DelistDate string `json:"delist_date"`  // 退市日期
	IsHs       string `json:"is_hs"`        // 是否沪深港通标的，N否 H沪股通 S深股通
	ActName    string `json:"act_name"`     // 实控人名称
	ActEntType string `json:"act_ent_type"` // 实控人企业性质
}

type StockBasicRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     //TS股票代码
	Name       string `json:"name"`        //名称
	Market     string `json:"market"`      //市场类别 （主板/创业板/科创板/CDR/北交所）
	ListStatus string `json:"list_status"` //上市状态 L上市 D退市 P暂停上市，默认是L
	Exchange   string `json:"exchange"`    //交易所 SSE上交所 SZSE深交所 BSE北交所
	IsHs       string `json:"is_hs"`       //是否沪深港通标的，N否 H沪股通 S深股通
}

type StockBasicResponse struct {
	List []*StockBasic `json:"list"`
}

func (x *StockBasicResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 交易日历|trade_cal
type TradeCal struct {
	Exchange     string `json:"exchange"`      // 交易所 SSE上交所 SZSE深交所
	CalDate      string `json:"cal_date"`      // 日历日期
	IsOpen       int32  `json:"is_open"`       // 是否交易 0休市 1交易
	PretradeDate string `json:"pretrade_date"` // 上一个交易日
}

type TradeCalRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	Exchange  string `json:"exchange"`   // 交易所 SSE上交所,SZSE深交所,CFFEX 中金所,SHFE 上期所,CZCE 郑商所,DCE 大商所,INE 上能源
	StartDate string `json:"start_date"` // 开始日期 （格式：YYYYMMDD 下同）
	EndDate   string `json:"end_date"`   // 结束日期
	IsOpen    string `json:"is_open"`    // 是否交易 '0'休市 '1'交易
	CalDate   string `json:"cal_date"`   // 日历日期
}

type TradeCalResponse struct {
	List []*TradeCal `json:"list"`
}

func (x *TradeCalResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 股票曾用名|namechange
type NameChange struct {
	TsCode       string `json:"ts_code"`       // TS代码
	Name         string `json:"name"`          // 证券名称
	StartDate    string `json:"start_date"`    // 开始日期
	EndDate      string `json:"end_date"`      // 结束日期
	AnnDate      string `json:"ann_date"`      // 公告日期
	ChangeReason string `json:"change_reason"` // 变更原因
}

type NameChangeRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	StartDate string `json:"start_date"` // 公告开始日期
	EndDate   string `json:"end_date"`   // 公告结束日期
}

type NameChangeResponse struct {
	List []*NameChange `json:"list"`
}

func (x *NameChangeResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 沪深股通成份股|hs_const
type HsConst struct {
	TsCode  string `json:"ts_code"`  // TS代码
	HsType  string `json:"hs_type"`  // 沪深港通类型SH沪SZ深
	InDate  string `json:"in_date"`  // 纳入日期
	OutDate string `json:"out_date"` // 剔除日期
	IsNew   string `json:"is_new"`   // 是否最新 1是 0否
}

type HsConstRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	HsType string `json:"hs_type"` // 类型SH沪股通SZ深股通
	IsNew  string `json:"is_new"`  // 是否最新 1 是 0 否 (默认1)
	TsCode string `json:"ts_code"` // TS代码
}

type HsConstResponse struct {
	List []*HsConst `json:"list"`
}

func (x *HsConstResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 上市公司基本信息|stock_company
type StockCompany struct {
	TsCode        string  `json:"ts_code"`        // 股票代码
	Exchange      string  `json:"exchange"`       // 交易所代码 ，SSE上交所 SZSE深交所
	Chairman      string  `json:"chairman"`       // 法人代表
	Manager       string  `json:"manager"`        // 总经理
	Secretary     string  `json:"secretary"`      // 董秘
	RegCapital    float64 `json:"reg_capital"`    // 注册资本(万元)
	SetupDate     string  `json:"setup_date"`     // 注册日期
	Province      string  `json:"province"`       // 所在省份
	City          string  `json:"city"`           // 所在城市
	Introduction  string  `json:"introduction"`   // 公司介绍
	Website       string  `json:"website"`        // 公司主页
	Email         string  `json:"email"`          // 电子邮件
	Office        string  `json:"office"`         // 办公室
	Employees     int64   `json:"employees"`      // 员工人数
	MainBusiness  string  `json:"main_business"`  // 主要业务及产品
	BusinessScope string  `json:"business_scope"` // 经营范围
}

type StockCompanyRequest struct {
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
	TsCode   string `json:"ts_code"`  // 股票代码
	Exchange string `json:"exchange"` // 交易所代码 ，SSE上交所 SZSE深交所 BSE北交所
	Status   string `json:"status"`   // 状态
}

type StockCompanyResponse struct {
	List []*StockCompany `json:"list"`
}

func (x *StockCompanyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 上市公司管理层|stk_managers
type StkManagers struct {
	TsCode    string `json:"ts_code"`    // TS股票代码
	AnnDate   string `json:"ann_date"`   // 公告日期
	Name      string `json:"name"`       // 姓名
	Gender    string `json:"gender"`     // 性别
	Lev       string `json:"lev"`        // 岗位类别
	Title     string `json:"title"`      // 岗位
	Edu       string `json:"edu"`        // 学历
	National  string `json:"national"`   // 国籍
	Birthday  string `json:"birthday"`   // 出生年月
	BeginDate string `json:"begin_date"` // 上任日期
	EndDate   string `json:"end_date"`   // 离任日期
	Resume    string `json:"resume"`     // 个人简历
}

type StkManagersRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    //股票代码，支持单个或多个股票输入
	AnnDate   string `json:"ann_date"`   //公告日期（YYYYMMDD格式，下同）
	StartDate string `json:"start_date"` //公告开始日期
	EndDate   string `json:"end_date"`   //公告结束日期
}

type StkManagersResponse struct {
	List []*StkManagers `json:"list"`
}

func (x *StkManagersResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 管理层薪酬和持股|stk_rewards
type StkRewards struct {
	TsCode  string  `json:"ts_code"`  // TS股票代码
	AnnDate string  `json:"ann_date"` // 公告日期
	EndDate string  `json:"end_date"` // 截止日期
	Name    string  `json:"name"`     // 姓名
	Title   string  `json:"title"`    // 职务
	Reward  float64 `json:"reward"`   // 报酬
	HoldVol float64 `json:"hold_vol"` // 持股数
}

type StkRewardsRequest struct {
	Limit   string `json:"limit"`
	Offset  string `json:"offset"`
	TsCode  string `json:"ts_code"`  // TS股票代码，支持单个或多个代码输入
	EndDate string `json:"end_date"` // 报告期
}

type StkRewardsResponse struct {
	List []*StkRewards `json:"list"`
}

func (x *StkRewardsResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// IPO新股列表|new_share
type NewShare struct {
	TsCode       string  `json:"ts_code"`       // TS股票代码
	SubCode      string  `json:"sub_code"`      // 申购代码
	Name         string  `json:"name"`          // 名称
	IpoDate      string  `json:"ipo_date"`      // 上网发行日期
	IssueDate    string  `json:"issue_date"`    // 上市日期
	Amount       float64 `json:"amount"`        // 发行总量（万股）
	MarketAmount float64 `json:"market_amount"` // 上网发行总量（万股）
	Price        float64 `json:"price"`         // 发行价格
	Pe           float64 `json:"pe"`            // 市盈率
	LimitAmount  float64 `json:"limit_amount"`  // 个人申购上限（万股）
	Funds        float64 `json:"funds"`         // 募集资金（亿元）
	Ballot       float64 `json:"ballot"`        // 中签率
}

type NewShareRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	StartDate string `json:"start_date"` // 上网发行开始日期
	EndDate   string `json:"end_date"`   // 上网发行结束日期
}

type NewShareResponse struct {
	List []*NewShare `json:"list"`
}

func (x *NewShareResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 备用列表|bak_basic
type BakBasic struct {
	TradeDate        string  `json:"trade_date"`        // 交易日期
	TsCode           string  `json:"ts_code"`           // TS股票代码
	Name             string  `json:"name"`              // 股票名称
	Industry         string  `json:"industry"`          // 行业
	Area             string  `json:"area"`              // 地域
	Pe               float64 `json:"pe"`                // 市盈率（动）
	FloatShare       float64 `json:"float_share"`       // 流通股本（亿）
	TotalShare       float64 `json:"total_share"`       // 总股本（亿）
	TotalAssets      float64 `json:"total_assets"`      // 总资产（亿）
	LiquidAssets     float64 `json:"liquid_assets"`     // 流动资产（亿）
	FixedAssets      float64 `json:"fixed_assets"`      // 固定资产（亿）
	Reserved         float64 `json:"reserved"`          // 公积金
	ReservedPershare float64 `json:"reserved_pershare"` // 每股公积金
	Eps              float64 `json:"eps"`               // 每股收益
	Bvps             float64 `json:"bvps"`              // 每股净资产
	Pb               float64 `json:"pb"`                // 市净率
	ListDate         string  `json:"list_date"`         // 上市日期
	Undp             float64 `json:"undp"`              // 未分配利润
	PerUndp          float64 `json:"per_undp"`          // 每股未分配利润
	RevYoy           float64 `json:"rev_yoy"`           // 收入同比（%）
	ProfitYoy        float64 `json:"profit_yoy"`        // 利润同比（%）
	Gpr              float64 `json:"gpr"`               // 毛利率（%）
	Npr              float64 `json:"npr"`               // 净利润率（%）
	HolderNum        int64   `json:"holder_num"`        // 股东人数
}

type BakBasicRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期
	TsCode    string `json:"ts_code"`    // 股票代码
}

type BakBasicResponse struct {
	List []*BakBasic `json:"list"`
}

func (x *BakBasicResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 日线行情|daily
type Daily struct {
	TsCode    string  `json:"ts_code"`    // 股票代码
	TradeDate string  `json:"trade_date"` // 交易日期
	Open      float64 `json:"open"`       // 开盘价
	High      float64 `json:"high"`       // 最高价
	Low       float64 `json:"low"`        // 最低价
	Close     float64 `json:"close"`      // 收盘价
	PreClose  float64 `json:"pre_close"`  // 昨收价(前复权)
	Change    float64 `json:"change"`     // 涨跌额
	PctChg    float64 `json:"pct_chg"`    // 涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
	Vol       float64 `json:"vol"`        // 成交量 （手）
	Amount    float64 `json:"amount"`     // 成交额 （千元）
}

type DailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码（支持多个股票同时提取，逗号分隔）
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD）
	StartDate string `json:"start_date"` // 开始日期(YYYYMMDD)
	EndDate   string `json:"end_date"`   // 结束日期(YYYYMMDD)
}

type DailyResponse struct {
	List []*Daily `json:"list"`
}

func (x *DailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 周线行情|weekly
type Weekly struct {
	TsCode    string  `json:"ts_code"`    // 股票代码
	TradeDate string  `json:"trade_date"` // 交易日期
	Close     float64 `json:"close"`      // 周收盘价
	Open      float64 `json:"open"`       // 周开盘价
	High      float64 `json:"high"`       // 周最高价
	Low       float64 `json:"low"`        // 周最低价
	PreClose  float64 `json:"pre_close"`  // 上一周收盘价
	Change    float64 `json:"change"`     // 周涨跌额
	PctChg    float64 `json:"pct_chg"`    // 周涨跌幅（未复权，如果是复权请用 通用行情接口 ）
	Vol       float64 `json:"vol"`        // 周成交量
	Amount    float64 `json:"amount"`     // 周成交额
}

type WeeklyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码 （ts_code,trade_date两个参数任选一）
	TradeDate string `json:"trade_date"` // 交易日期 （每周最后一个交易日期，YYYYMMDD格式）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type WeeklyResponse struct {
	List []*Weekly `json:"list"`
}

func (x *WeeklyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 月线行情|monthly
type Monthly struct {
	TsCode    string  `json:"ts_code"`    // 股票代码
	TradeDate string  `json:"trade_date"` // 交易日期
	Close     float64 `json:"close"`      // 月收盘价
	Open      float64 `json:"open"`       // 月开盘价
	High      float64 `json:"high"`       // 月最高价
	Low       float64 `json:"low"`        // 月最低价
	PreClose  float64 `json:"pre_close"`  // 上月收盘价
	Change    float64 `json:"change"`     // 月涨跌额
	PctChg    float64 `json:"pct_chg"`    // 月涨跌幅（未复权，如果是复权请用 通用行情接口 ）
	Vol       float64 `json:"vol"`        // 月成交量
	Amount    float64 `json:"amount"`     // 月成交额
}

type MonthlyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码 （ts_code,trade_date两个参数任选一）
	TradeDate string `json:"trade_date"` // 交易日期 （每月最后一个交易日日期，YYYYMMDD格式）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type MonthlyResponse struct {
	List []*Monthly `json:"list"`
}

func (x *MonthlyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 复权因子|adj_factor
type AdjFactor struct {
	TsCode    string  `json:"ts_code"`    // 股票代码
	TradeDate string  `json:"trade_date"` // 交易日期
	AdjFactor float64 `json:"adj_factor"` // 复权因子
}

type AdjFactorRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期(YYYYMMDD，下同)
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type AdjFactorResponse struct {
	List []*AdjFactor `json:"list"`
}

func (x *AdjFactorResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 每日指标|daily_basic
type DailyBasic struct {
	TsCode        string  `json:"ts_code"`         // TS股票代码
	TradeDate     string  `json:"trade_date"`      // 交易日期
	Close         float64 `json:"close"`           // 当日收盘价
	TurnoverRate  float64 `json:"turnover_rate"`   // 换手率（%）
	TurnoverRateF float64 `json:"turnover_rate_f"` // 换手率（自由流通股）
	VolumeRatio   float64 `json:"volume_ratio"`    // 量比
	Pe            float64 `json:"pe"`              // 市盈率（总市值/净利润，亏损的PE为空）
	PeTtm         float64 `json:"pe_ttm"`          // 市盈率（TTM，亏损的PE为空）
	Pb            float64 `json:"pb"`              // 市净率（总市值/净资产）
	Ps            float64 `json:"ps"`              // 市销率
	PsTtm         float64 `json:"ps_ttm"`          // 市销率（TTM）
	DvRatio       float64 `json:"dv_ratio"`        // 股息率（%）
	DvTtm         float64 `json:"dv_ttm"`          // 股息率（TTM）（%）
	TotalShare    float64 `json:"total_share"`     // 总股本（万股）
	FloatShare    float64 `json:"float_share"`     // 流通股本（万股）
	FreeShare     float64 `json:"free_share"`      // 自由流通股本（万）
	TotalMv       float64 `json:"total_mv"`        // 总市值 （万元）
	CircMv        float64 `json:"circ_mv"`         // 流通市值（万元）
}

type DailyBasicRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码（二选一）
	TradeDate string `json:"trade_date"` // 交易日期 （二选一）
	StartDate string `json:"start_date"` // 开始日期(YYYYMMDD)
	EndDate   string `json:"end_date"`   // 结束日期(YYYYMMDD)
}

type DailyBasicResponse struct {
	List []*DailyBasic `json:"list"`
}

func (x *DailyBasicResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 个股资金流向|moneyflow
type MoneyFlow struct {
	TsCode        string  `json:"ts_code"`         // TS代码
	TradeDate     string  `json:"trade_date"`      // 交易日期
	BuySmVol      int64   `json:"buy_sm_vol"`      // 小单买入量（手）
	BuySmAmount   float64 `json:"buy_sm_amount"`   // 小单买入金额（万元）
	SellSmVol     int64   `json:"sell_sm_vol"`     // 小单卖出量（手）
	SellSmAmount  float64 `json:"sell_sm_amount"`  // 小单卖出金额（万元）
	BuyMdVol      int64   `json:"buy_md_vol"`      // 中单买入量（手）
	BuyMdAmount   float64 `json:"buy_md_amount"`   // 中单买入金额（万元）
	SellMdVol     int64   `json:"sell_md_vol"`     // 中单卖出量（手）
	SellMdAmount  float64 `json:"sell_md_amount"`  // 中单卖出金额（万元）
	BuyLgVol      int64   `json:"buy_lg_vol"`      // 大单买入量（手）
	BuyLgAmount   float64 `json:"buy_lg_amount"`   // 大单买入金额（万元）
	SellLgVol     int64   `json:"sell_lg_vol"`     // 大单卖出量（手）
	SellLgAmount  float64 `json:"sell_lg_amount"`  // 大单卖出金额（万元）
	BuyElgVol     int64   `json:"buy_elg_vol"`     // 特大单买入量（手）
	BuyElgAmount  float64 `json:"buy_elg_amount"`  // 特大单买入金额（万元）
	SellElgVol    int64   `json:"sell_elg_vol"`    // 特大单卖出量（手）
	SellElgAmount float64 `json:"sell_elg_amount"` // 特大单卖出金额（万元）
	NetMfVol      int64   `json:"net_mf_vol"`      // 净流入量（手）
	NetMfAmount   float64 `json:"net_mf_amount"`   // 净流入额（万元）
}

type MoneyFlowRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码（股票和时间参数至少输入一个）
	TradeDate string `json:"trade_date"` // 交易日期
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type MoneyFlowResponse struct {
	List []*MoneyFlow `json:"list"`
}

func (x *MoneyFlowResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 每日涨跌停价格|stk_limit
type StkLimit struct {
	TradeDate string  `json:"trade_date"` // 交易日期
	TsCode    string  `json:"ts_code"`    // TS股票代码
	PreClose  float64 `json:"pre_close"`  // 昨日收盘价
	UpLimit   float64 `json:"up_limit"`   // 涨停价
	DownLimit float64 `json:"down_limit"` // 跌停价
}

type StkLimitRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期(YYYYMMDD，下同)
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type StkLimitResponse struct {
	List []*StkLimit `json:"list"`
}

func (x *StkLimitResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 每日停复牌信息|suspend_d
type Suspendd struct {
	TsCode        string `json:"ts_code"`        // TS代码
	TradeDate     string `json:"trade_date"`     // 停复牌日期
	SuspendTiming string `json:"suspend_timing"` // 日内停牌时间段
	SuspendType   string `json:"suspend_type"`   // 停复牌类型：S-停牌，R-复牌
}

type SuspenddRequest struct {
	Limit       string `json:"limit"`
	Offset      string `json:"offset"`
	TsCode      string `json:"ts_code"`      // 股票代码(可输入多值)
	TradeDate   string `json:"trade_date"`   // 交易日日期
	StartDate   string `json:"start_date"`   // 停复牌查询开始日期
	EndDate     string `json:"end_date"`     // 停复牌查询结束日期
	SuspendType string `json:"suspend_type"` // 停复牌类型：S-停牌,R-复牌
}

type SuspenddResponse struct {
	List []*Suspendd `json:"list"`
}

func (x *SuspenddResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 沪深港通资金流向|moneyflow_hsgt
type MoneyFlowHsgt struct {
	TradeDate  string  `json:"trade_date"`  // 交易日期
	GgtSs      float64 `json:"ggt_ss"`      // 港股通（上海）
	GgtSz      float64 `json:"ggt_sz"`      // 港股通（深圳）
	Hgt        float64 `json:"hgt"`         // 沪股通（百万元）
	Sgt        float64 `json:"sgt"`         // 深股通（百万元）
	NorthMoney float64 `json:"north_money"` // 北向资金（百万元）
	SouthMoney float64 `json:"south_money"` // 南向资金（百万元）
}

type MoneyFlowHsgtRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期 (二选一)
	StartDate string `json:"start_date"` // 开始日期 (二选一)
	EndDate   string `json:"end_date"`   // 结束日期
}

type MoneyFlowHsgtResponse struct {
	List []*MoneyFlowHsgt `json:"list"`
}

func (x *MoneyFlowHsgtResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 沪深股通十大成交股|hsgt_top10
type HsgtTop10 struct {
	TradeDate  string  `json:"trade_date"`  //交易日期
	TsCode     string  `json:"ts_code"`     //股票代码
	Name       string  `json:"name"`        //股票名称
	Close      float64 `json:"close"`       //收盘价
	Change     float64 `json:"change"`      //涨跌额
	Rank       int64   `json:"rank"`        //资金排名
	MarketType int64   `json:"market_type"` //市场类型（1：沪市 3：深市）
	Amount     float64 `json:"amount"`      //成交金额（元）
	NetAmount  float64 `json:"net_amount"`  //净成交金额（元）
	Buy        float64 `json:"buy"`         //买入金额（元）
	Sell       float64 `json:"sell"`        //卖出金额（元）
}

type HsgtTop10Request struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // 股票代码（二选一）
	TradeDate  string `json:"trade_date"`  // 交易日期（二选一）
	StartDate  string `json:"start_date"`  // 开始日期
	EndDate    string `json:"end_date"`    // 结束日期
	MarketType string `json:"market_type"` // 市场类型（1：沪市 3：深市）
}

type HsgtTop10Response struct {
	List []*HsgtTop10 `json:"list"`
}

func (x *HsgtTop10Response) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 港股通十大成交股|ggt_top10
type GgtTop10 struct {
	TradeDate   string  `json:"trade_date"`    // 交易日期
	TsCode      string  `json:"ts_code"`       // 股票代码
	Name        string  `json:"name"`          // 股票名称
	Close       float64 `json:"close"`         // 收盘价
	PChange     float64 `json:"p_change"`      // 涨跌幅
	Rank        int64   `json:"rank"`          // 资金排名
	MarketType  int64   `json:"market_type"`   // 市场类型 2：港股通（沪） 4：港股通（深）
	Amount      float64 `json:"amount"`        // 累计成交金额（元）
	NetAmount   float64 `json:"net_amount"`    // 净买入金额（元）
	ShAmount    float64 `json:"sh_amount"`     // 沪市成交金额（元）
	ShNetAmount float64 `json:"sh_net_amount"` // 沪市净买入金额（元）
	ShBuy       float64 `json:"sh_buy"`        // 沪市买入金额（元）
	ShSell      float64 `json:"sh_sell"`       // 沪市卖出金额
	SzAmount    float64 `json:"sz_amount"`     // 深市成交金额（元）
	SzNetAmount float64 `json:"sz_net_amount"` // 深市净买入金额（元）
	SzBuy       float64 `json:"sz_buy"`        // 深市买入金额（元）
	SzSell      float64 `json:"sz_sell"`       // 深市卖出金额（元）
}

type GgtTop10Request struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // 股票代码（二选一）
	TradeDate  string `json:"trade_date"`  // 交易日期（二选一）
	StartDate  string `json:"start_date"`  // 开始日期
	EndDate    string `json:"end_date"`    // 结束日期
	MarketType string `json:"market_type"` // 市场类型 2：港股通（沪） 4：港股通（深）
}

type GgtTop10Response struct {
	List []*GgtTop10 `json:"list"`
}

func (x *GgtTop10Response) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 港股通每日成交统计|ggt_daily
type GgtDaily struct {
	TradeDate  string  `json:"trade_date"`  // 交易日期
	BuyAmount  float64 `json:"buy_amount"`  // 买入成交金额（亿元）
	BuyVolume  float64 `json:"buy_volume"`  // 买入成交笔数（万笔）
	SellAmount float64 `json:"sell_amount"` // 卖出成交金额（亿元）
	SellVolume float64 `json:"sell_volume"` // 卖出成交笔数（万笔）
}

type GgtDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期 （格式YYYYMMDD，下同。支持单日和多日输入）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type GgtDailyResponse struct {
	List []*GgtDaily `json:"list"`
}

func (x *GgtDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 港股通每月成交统计|ggt_monthly
type GgtMonthly struct {
	Month        string  `json:"month"`          //	交易日期
	DayBuyAmt    float64 `json:"day_buy_amt"`    //	当月日均买入成交金额（亿元）
	DayBuyVol    float64 `json:"day_buy_vol"`    //	当月日均买入成交笔数（万笔）
	DaySellAmt   float64 `json:"day_sell_amt"`   //	当月日均卖出成交金额（亿元）
	DaySellVol   float64 `json:"day_sell_vol"`   //	当月日均卖出成交笔数（万笔）
	TotalBuyAmt  float64 `json:"total_buy_amt"`  //	总买入成交金额（亿元）
	TotalBuyVol  float64 `json:"total_buy_vol"`  //	总买入成交笔数（万笔）
	TotalSellAmt float64 `json:"total_sell_amt"` //	总卖出成交金额（亿元）
	TotalSellVol float64 `json:"total_sell_vol"` //	总卖出成交笔数（万笔）
}

type GgtMonthlyRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	Month      string `json:"month"`       // 月度（格式YYYYMM，下同，支持多个输入）
	StartMonth string `json:"start_month"` // 开始月度
	EndMonth   string `json:"end_month"`   // 结束月度
}

type GgtMonthlyResponse struct {
	List []*GgtMonthly `json:"list"`
}

func (x *GgtMonthlyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 备用行情|bak_daily
type BakDaily struct {
	TsCode      string  `json:"ts_code"`      // 股票代码
	TradeDate   string  `json:"trade_date"`   // 交易日期
	Name        string  `json:"name"`         // 股票名称
	PctChange   float64 `json:"pct_change"`   // 涨跌幅
	Close       float64 `json:"close"`        // 收盘价
	Change      float64 `json:"change"`       // 涨跌额
	Open        float64 `json:"open"`         // 开盘价
	High        float64 `json:"high"`         // 最高价
	Low         float64 `json:"low"`          // 最低价
	PreClose    float64 `json:"pre_close"`    // 昨收价
	VolRatio    float64 `json:"vol_ratio"`    // 量比
	TurnOver    float64 `json:"turn_over"`    // 换手率
	Swing       float64 `json:"swing"`        // 振幅
	Vol         float64 `json:"vol"`          // 成交量
	Amount      float64 `json:"amount"`       // 成交额
	Selling     float64 `json:"selling"`      // 内盘（主动卖，手）
	Buying      float64 `json:"buying"`       // 外盘（主动买， 手）
	TotalShare  float64 `json:"total_share"`  // 总股本(亿)
	FloatShare  float64 `json:"float_share"`  // 流通股本(亿)
	Pe          float64 `json:"pe"`           // 市盈(动)
	Industry    string  `json:"industry"`     // 所属行业
	Area        string  `json:"area"`         // 所属地域
	FloatMv     float64 `json:"float_mv"`     // 流通市值
	TotalMv     float64 `json:"total_mv"`     // 总市值
	AvgPrice    float64 `json:"avg_price"`    // 平均价
	Strength    float64 `json:"strength"`     // 强弱度(%)
	Activity    float64 `json:"activity"`     // 活跃度(%)
	AvgTurnover float64 `json:"avg_turnover"` // 笔换手
	Attack      float64 `json:"attack"`       // 攻击波(%)
	Interval_3  float64 `json:"interval_3"`   // 近3月涨幅
	Interval_6  float64 `json:"interval_6"`   // 近6月涨幅
}

type BakDailyRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type BakDailyResponse struct {
	List []*BakDaily `json:"list"`
}

func (x *BakDailyResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 利润表|income
type Income struct {
	TsCode                 string  `json:"ts_code"`                   // TS代码
	AnnDate                string  `json:"ann_date"`                  // 公告日期
	FAnnDate               string  `json:"f_ann_date"`                // 实际公告日期
	EndDate                string  `json:"end_date"`                  // 报告期
	ReportType             string  `json:"report_type"`               // 报告类型 见底部表
	CompType               string  `json:"comp_type"`                 // 公司类型(1一般工商业2银行3保险4证券)
	EndType                string  `json:"end_type"`                  // 报告期类型
	BasicEps               float64 `json:"basic_eps"`                 // 基本每股收益
	DilutedEps             float64 `json:"diluted_eps"`               // 稀释每股收益
	TotalRevenue           float64 `json:"total_revenue"`             // 营业总收入
	Revenue                float64 `json:"revenue"`                   // 营业收入
	IntIncome              float64 `json:"int_income"`                // 利息收入
	PremEarned             float64 `json:"prem_earned"`               // 已赚保费
	CommIncome             float64 `json:"comm_income"`               // 手续费及佣金收入
	NCommisIncome          float64 `json:"n_commis_income"`           // 手续费及佣金净收入
	NOthIncome             float64 `json:"n_oth_income"`              // 其他经营净收益
	NOthBIncome            float64 `json:"n_oth_b_income"`            // 加:其他业务净收益
	PremIncome             float64 `json:"prem_income"`               // 保险业务收入
	OutPrem                float64 `json:"out_prem"`                  // 减:分出保费
	UnePremReser           float64 `json:"une_prem_reser"`            // 提取未到期责任准备金
	ReinsIncome            float64 `json:"reins_income"`              // 其中:分保费收入
	NSecTbIncome           float64 `json:"n_sec_tb_income"`           // 代理买卖证券业务净收入
	NSecUwIncome           float64 `json:"n_sec_uw_income"`           // 证券承销业务净收入
	NAssetMgIncome         float64 `json:"n_asset_mg_income"`         // 受托客户资产管理业务净收入
	OthBIncome             float64 `json:"oth_b_income"`              // 其他业务收入
	FvValueChgGain         float64 `json:"fv_value_chg_gain"`         // 加:公允价值变动净收益
	InvestIncome           float64 `json:"invest_income"`             // 加:投资净收益
	AssInvestIncome        float64 `json:"ass_invest_income"`         // 其中:对联营企业和合营企业的投资收益
	ForexGain              float64 `json:"forex_gain"`                // 加:汇兑净收益
	TotalCogs              float64 `json:"total_cogs"`                // 营业总成本
	OperCost               float64 `json:"oper_cost"`                 // 减:营业成本
	IntExp                 float64 `json:"int_exp"`                   // 减:利息支出
	CommExp                float64 `json:"comm_exp"`                  // 减:手续费及佣金支出
	BizTaxSurchg           float64 `json:"biz_tax_surchg"`            // 减:营业税金及附加
	SellExp                float64 `json:"sell_exp"`                  // 减:销售费用
	AdminExp               float64 `json:"admin_exp"`                 // 减:管理费用
	FinExp                 float64 `json:"fin_exp"`                   // 减:财务费用
	AssetsImpairLoss       float64 `json:"assets_impair_loss"`        // 减:资产减值损失
	PremRefund             float64 `json:"prem_refund"`               // 退保金
	CompensPayout          float64 `json:"compens_payout"`            // 赔付总支出
	ReserInsurLiab         float64 `json:"reser_insur_liab"`          // 提取保险责任准备金
	DivPayt                float64 `json:"div_payt"`                  // 保户红利支出
	ReinsExp               float64 `json:"reins_exp"`                 // 分保费用
	OperExp                float64 `json:"oper_exp"`                  // 营业支出
	CompensPayoutRefu      float64 `json:"compens_payout_refu"`       // 减:摊回赔付支出
	InsurReserRefu         float64 `json:"insur_reser_refu"`          // 减:摊回保险责任准备金
	ReinsCostRefund        float64 `json:"reins_cost_refund"`         // 减:摊回分保费用
	OtherBusCost           float64 `json:"other_bus_cost"`            // 其他业务成本
	OperateProfit          float64 `json:"operate_profit"`            // 营业利润
	NonOperIncome          float64 `json:"non_oper_income"`           // 加:营业外收入
	NonOperExp             float64 `json:"non_oper_exp"`              // 减:营业外支出
	NcaDisploss            float64 `json:"nca_disploss"`              // 其中:减:非流动资产处置净损失
	TotalProfit            float64 `json:"total_profit"`              // 利润总额
	IncomeTax              float64 `json:"income_tax"`                // 所得税费用
	NIncome                float64 `json:"n_income"`                  // 净利润(含少数股东损益)
	NIncomeAttrP           float64 `json:"n_income_attr_p"`           // 净利润(不含少数股东损益)
	MinorityGain           float64 `json:"minority_gain"`             // 少数股东损益
	OthComprIncome         float64 `json:"oth_compr_income"`          // 其他综合收益
	TComprIncome           float64 `json:"t_compr_income"`            // 综合收益总额
	ComprIncAttrP          float64 `json:"compr_inc_attr_p"`          // 归属于母公司(或股东)的综合收益总额
	ComprIncAttrMS         float64 `json:"compr_inc_attr_m_s"`        // 归属于少数股东的综合收益总额
	Ebit                   float64 `json:"ebit"`                      // 息税前利润
	Ebitda                 float64 `json:"ebitda"`                    // 息税折旧摊销前利润
	InsuranceExp           float64 `json:"insurance_exp"`             // 保险业务支出
	UndistProfit           float64 `json:"undist_profit"`             // 年初未分配利润
	DistableProfit         float64 `json:"distable_profit"`           // 可分配利润
	RdExp                  float64 `json:"rd_exp"`                    // 研发费用
	FinExpIntExp           float64 `json:"fin_exp_int_exp"`           // 财务费用:利息费用
	FinExpIntInc           float64 `json:"fin_exp_int_inc"`           // 财务费用:利息收入
	TransferSurplusRese    float64 `json:"transfer_surplus_rese"`     // 盈余公积转入
	TransferHousingImprest float64 `json:"transfer_housing_imprest"`  // 住房周转金转入
	TransferOth            float64 `json:"transfer_oth"`              // 其他转入
	AdjLossgain            float64 `json:"adj_lossgain"`              // 调整以前年度损益
	WithdraLegalSurplus    float64 `json:"withdra_legal_surplus"`     // 提取法定盈余公积
	WithdraLegalPubfund    float64 `json:"withdra_legal_pubfund"`     // 提取法定公益金
	WithdraBizDevfund      float64 `json:"withdra_biz_devfund"`       // 提取企业发展基金
	WithdraReseFund        float64 `json:"withdra_rese_fund"`         // 提取储备基金
	WithdraOthErsu         float64 `json:"withdra_oth_ersu"`          // 提取任意盈余公积金
	WorkersWelfare         float64 `json:"workers_welfare"`           // 职工奖金福利
	DistrProfitShrhder     float64 `json:"distr_profit_shrhder"`      // 可供股东分配的利润
	PrfsharePayableDvd     float64 `json:"prfshare_payable_dvd"`      // 应付优先股股利
	ComsharePayableDvd     float64 `json:"comshare_payable_dvd"`      // 应付普通股股利
	CapitComstockDiv       float64 `json:"capit_comstock_div"`        // 转作股本的普通股股利
	NetAfterNrLpCorrect    float64 `json:"net_after_nr_lp_correct"`   // 扣除非经常性损益后的净利润（更正前）
	CreditImpaLoss         float64 `json:"credit_impa_loss"`          // 信用减值损失
	NetExpoHedgingBenefits float64 `json:"net_expo_hedging_benefits"` // 净敞口套期收益
	OthImpairLossAssets    float64 `json:"oth_impair_loss_assets"`    // 其他资产减值损失
	TotalOpcost            float64 `json:"total_opcost"`              // 营业总成本（二）
	AmodcostFinAssets      float64 `json:"amodcost_fin_assets"`       // 以摊余成本计量的金融资产终止确认收益
	OthIncome              float64 `json:"oth_income"`                // 其他收益
	AssetDispIncome        float64 `json:"asset_disp_income"`         // 资产处置收益
	ContinuedNetProfit     float64 `json:"continued_net_profit"`      // 持续经营净利润
	EndNetProfit           float64 `json:"end_net_profit"`            // 终止经营净利润
	UpdateFlag             string  `json:"update_flag"`               // 更新标识
}

type IncomeRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // 股票代码
	AnnDate    string `json:"ann_date"`    // 公告日期（YYYYMMDD格式，下同）
	FAnnDate   string `json:"f_ann_date"`  // 实际公告日期
	StartDate  string `json:"start_date"`  // 公告开始日期
	EndDate    string `json:"end_date"`    // 公告结束日期
	Period     string `json:"period"`      // 报告期(每个季度最后一天的日期，比如20171231表示年报)
	ReportType string `json:"report_type"` // 报告类型
	CompType   string `json:"comp_type"`   // 公司类型（1一般工商业2银行3保险4证券）
}

type IncomeResponse struct {
	List []*Income `json:"list"`
}

func (x *IncomeResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 资产负债表|balancesheet
type BalanceSheet struct {
	TsCode                string  `json:"ts_code"`                    // TS股票代码
	AnnDate               string  `json:"ann_date"`                   // 公告日期
	FAnnDate              string  `json:"f_ann_date"`                 // 实际公告日期
	EndDate               string  `json:"end_date"`                   // 报告期
	ReportType            string  `json:"report_type"`                // 报表类型
	CompType              string  `json:"comp_type"`                  // 公司类型(1一般工商业2银行3保险4证券)
	EndType               string  `json:"end_type"`                   // 报告期类型
	TotalShare            float64 `json:"total_share"`                // 期末总股本
	CapRese               float64 `json:"cap_rese"`                   // 资本公积金
	UndistrPorfit         float64 `json:"undistr_porfit"`             // 未分配利润
	SurplusRese           float64 `json:"surplus_rese"`               // 盈余公积金
	SpecialRese           float64 `json:"special_rese"`               // 专项储备
	MoneyCap              float64 `json:"money_cap"`                  // 货币资金
	TradAsset             float64 `json:"trad_asset"`                 // 交易性金融资产
	NotesReceiv           float64 `json:"notes_receiv"`               // 应收票据
	AccountsReceiv        float64 `json:"accounts_receiv"`            // 应收账款
	OthReceiv             float64 `json:"oth_receiv"`                 // 其他应收款
	Prepayment            float64 `json:"prepayment"`                 // 预付款项
	DivReceiv             float64 `json:"div_receiv"`                 // 应收股利
	IntReceiv             float64 `json:"int_receiv"`                 // 应收利息
	Inventories           float64 `json:"inventories"`                // 存货
	AmorExp               float64 `json:"amor_exp"`                   // 待摊费用
	NcaWithin_1Y          float64 `json:"nca_within_1y"`              // 一年内到期的非流动资产
	SettRsrv              float64 `json:"sett_rsrv"`                  // 结算备付金
	LoantoOthBankFi       float64 `json:"loanto_oth_bank_fi"`         // 拆出资金
	PremiumReceiv         float64 `json:"premium_receiv"`             // 应收保费
	ReinsurReceiv         float64 `json:"reinsur_receiv"`             // 应收分保账款
	ReinsurResReceiv      float64 `json:"reinsur_res_receiv"`         // 应收分保合同准备金
	PurResaleFa           float64 `json:"pur_resale_fa"`              // 买入返售金融资产
	OthCurAssets          float64 `json:"oth_cur_assets"`             // 其他流动资产
	TotalCurAssets        float64 `json:"total_cur_assets"`           // 流动资产合计
	FaAvailForSale        float64 `json:"fa_avail_for_sale"`          // 可供出售金融资产
	HtmInvest             float64 `json:"htm_invest"`                 // 持有至到期投资
	LtEqtInvest           float64 `json:"lt_eqt_invest"`              // 长期股权投资
	InvestRealEstate      float64 `json:"invest_real_estate"`         // 投资性房地产
	TimeDeposits          float64 `json:"time_deposits"`              // 定期存款
	OthAssets             float64 `json:"oth_assets"`                 // 其他资产
	LtRec                 float64 `json:"lt_rec"`                     // 长期应收款
	FixAssets             float64 `json:"fix_assets"`                 // 固定资产
	Cip                   float64 `json:"cip"`                        // 在建工程
	ConstMaterials        float64 `json:"const_materials"`            // 工程物资
	FixedAssetsDisp       float64 `json:"fixed_assets_disp"`          // 固定资产清理
	ProducBioAssets       float64 `json:"produc_bio_assets"`          // 生产性生物资产
	OilAndGasAssets       float64 `json:"oil_and_gas_assets"`         // 油气资产
	IntanAssets           float64 `json:"intan_assets"`               // 无形资产
	RAndD                 float64 `json:"r_and_d"`                    // 研发支出
	Goodwill              float64 `json:"goodwill"`                   // 商誉
	LtAmorExp             float64 `json:"lt_amor_exp"`                // 长期待摊费用
	DeferTaxAssets        float64 `json:"defer_tax_assets"`           // 递延所得税资产
	DecrInDisbur          float64 `json:"decr_in_disbur"`             // 发放贷款及垫款
	OthNca                float64 `json:"oth_nca"`                    // 其他非流动资产
	TotalNca              float64 `json:"total_nca"`                  // 非流动资产合计
	CashReserCb           float64 `json:"cash_reser_cb"`              // 现金及存放中央银行款项
	DeposInOthBfi         float64 `json:"depos_in_oth_bfi"`           // 存放同业和其它金融机构款项
	PrecMetals            float64 `json:"prec_metals"`                // 贵金属
	DerivAssets           float64 `json:"deriv_assets"`               // 衍生金融资产
	RrReinsUnePrem        float64 `json:"rr_reins_une_prem"`          // 应收分保未到期责任准备金
	RrReinsOutstdCla      float64 `json:"rr_reins_outstd_cla"`        // 应收分保未决赔款准备金
	RrReinsLinsLiab       float64 `json:"rr_reins_lins_liab"`         // 应收分保寿险责任准备金
	RrReinsLthinsLiab     float64 `json:"rr_reins_lthins_liab"`       // 应收分保长期健康险责任准备金
	RefundDepos           float64 `json:"refund_depos"`               // 存出保证金
	PhPledgeLoans         float64 `json:"ph_pledge_loans"`            // 保户质押贷款
	RefundCapDepos        float64 `json:"refund_cap_depos"`           // 存出资本保证金
	IndepAcctAssets       float64 `json:"indep_acct_assets"`          // 独立账户资产
	ClientDepos           float64 `json:"client_depos"`               // 其中：客户资金存款
	ClientProv            float64 `json:"client_prov"`                // 其中：客户备付金
	TransacSeatFee        float64 `json:"transac_seat_fee"`           // 其中:交易席位费
	InvestAsReceiv        float64 `json:"invest_as_receiv"`           // 应收款项类投资
	TotalAssets           float64 `json:"total_assets"`               // 资产总计
	LtBorr                float64 `json:"lt_borr"`                    // 长期借款
	StBorr                float64 `json:"st_borr"`                    // 短期借款
	CbBorr                float64 `json:"cb_borr"`                    // 向中央银行借款
	DeposIbDeposits       float64 `json:"depos_ib_deposits"`          // 吸收存款及同业存放
	LoanOthBank           float64 `json:"loan_oth_bank"`              // 拆入资金
	TradingFl             float64 `json:"trading_fl"`                 // 交易性金融负债
	NotesPayable          float64 `json:"notes_payable"`              // 应付票据
	AcctPayable           float64 `json:"acct_payable"`               // 应付账款
	AdvReceipts           float64 `json:"adv_receipts"`               // 预收款项
	SoldForRepurFa        float64 `json:"sold_for_repur_fa"`          // 卖出回购金融资产款
	CommPayable           float64 `json:"comm_payable"`               // 应付手续费及佣金
	PayrollPayable        float64 `json:"payroll_payable"`            // 应付职工薪酬
	TaxesPayable          float64 `json:"taxes_payable"`              // 应交税费
	IntPayable            float64 `json:"int_payable"`                // 应付利息
	DivPayable            float64 `json:"div_payable"`                // 应付股利
	OthPayable            float64 `json:"oth_payable"`                // 其他应付款
	AccExp                float64 `json:"acc_exp"`                    // 预提费用
	DeferredInc           float64 `json:"deferred_inc"`               // 递延收益
	StBondsPayable        float64 `json:"st_bonds_payable"`           // 应付短期债券
	PayableToReinsurer    float64 `json:"payable_to_reinsurer"`       // 应付分保账款
	RsrvInsurCont         float64 `json:"rsrv_insur_cont"`            // 保险合同准备金
	ActingTradingSec      float64 `json:"acting_trading_sec"`         // 代理买卖证券款
	ActingUwSec           float64 `json:"acting_uw_sec"`              // 代理承销证券款
	NonCurLiabDue_1Y      float64 `json:"non_cur_liab_due_1y"`        // 一年内到期的非流动负债
	OthCurLiab            float64 `json:"oth_cur_liab"`               // 其他流动负债
	TotalCurLiab          float64 `json:"total_cur_liab"`             // 流动负债合计
	BondPayable           float64 `json:"bond_payable"`               // 应付债券
	LtPayable             float64 `json:"lt_payable"`                 // 长期应付款
	SpecificPayables      float64 `json:"specific_payables"`          // 专项应付款
	EstimatedLiab         float64 `json:"estimated_liab"`             // 预计负债
	DeferTaxLiab          float64 `json:"defer_tax_liab"`             // 递延所得税负债
	DeferIncNonCurLiab    float64 `json:"defer_inc_non_cur_liab"`     // 递延收益-非流动负债
	OthNcl                float64 `json:"oth_ncl"`                    // 其他非流动负债
	TotalNcl              float64 `json:"total_ncl"`                  // 非流动负债合计
	DeposOthBfi           float64 `json:"depos_oth_bfi"`              // 同业和其它金融机构存放款项
	DerivLiab             float64 `json:"deriv_liab"`                 // 衍生金融负债
	Depos                 float64 `json:"depos"`                      // 吸收存款
	AgencyBusLiab         float64 `json:"agency_bus_liab"`            // 代理业务负债
	OthLiab               float64 `json:"oth_liab"`                   // 其他负债
	PremReceivAdva        float64 `json:"prem_receiv_adva"`           // 预收保费
	DeposReceived         float64 `json:"depos_received"`             // 存入保证金
	PhInvest              float64 `json:"ph_invest"`                  // 保户储金及投资款
	ReserUnePrem          float64 `json:"reser_une_prem"`             // 未到期责任准备金
	ReserOutstdClaims     float64 `json:"reser_outstd_claims"`        // 未决赔款准备金
	ReserLinsLiab         float64 `json:"reser_lins_liab"`            // 寿险责任准备金
	ReserLthinsLiab       float64 `json:"reser_lthins_liab"`          // 长期健康险责任准备金
	IndeptAccLiab         float64 `json:"indept_acc_liab"`            // 独立账户负债
	PledgeBorr            float64 `json:"pledge_borr"`                // 其中:质押借款
	IndemPayable          float64 `json:"indem_payable"`              // 应付赔付款
	PolicyDivPayable      float64 `json:"policy_div_payable"`         // 应付保单红利
	TotalLiab             float64 `json:"total_liab"`                 // 负债合计
	TreasuryShare         float64 `json:"treasury_share"`             // 减:库存股
	OrdinRiskReser        float64 `json:"ordin_risk_reser"`           // 一般风险准备
	ForexDiffer           float64 `json:"forex_differ"`               // 外币报表折算差额
	InvestLossUnconf      float64 `json:"invest_loss_unconf"`         // 未确认的投资损失
	MinorityInt           float64 `json:"minority_int"`               // 少数股东权益
	TotalHldrEqyExcMinInt float64 `json:"total_hldr_eqy_exc_min_int"` // 股东权益合计(不含少数股东权益)
	TotalHldrEqyIncMinInt float64 `json:"total_hldr_eqy_inc_min_int"` // 股东权益合计(含少数股东权益)
	TotalLiabHldrEqy      float64 `json:"total_liab_hldr_eqy"`        // 负债及股东权益总计
	LtPayrollPayable      float64 `json:"lt_payroll_payable"`         // 长期应付职工薪酬
	OthCompIncome         float64 `json:"oth_comp_income"`            // 其他综合收益
	OthEqtTools           float64 `json:"oth_eqt_tools"`              // 其他权益工具
	OthEqtToolsPShr       float64 `json:"oth_eqt_tools_p_shr"`        // 其他权益工具(优先股)
	LendingFunds          float64 `json:"lending_funds"`              // 融出资金
	AccReceivable         float64 `json:"acc_receivable"`             // 应收款项
	StFinPayable          float64 `json:"st_fin_payable"`             // 应付短期融资款
	Payables              float64 `json:"payables"`                   // 应付款项
	HfsAssets             float64 `json:"hfs_assets"`                 // 持有待售的资产
	HfsSales              float64 `json:"hfs_sales"`                  // 持有待售的负债
	CostFinAssets         float64 `json:"cost_fin_assets"`            // 以摊余成本计量的金融资产
	FairValueFinAssets    float64 `json:"fair_value_fin_assets"`      // 以公允价值计量且其变动计入其他综合收益的金融资产
	CipTotal              float64 `json:"cip_total"`                  // 在建工程(合计)(元)
	OthPayTotal           float64 `json:"oth_pay_total"`              // 其他应付款(合计)(元)
	LongPayTotal          float64 `json:"long_pay_total"`             // 长期应付款(合计)(元)
	DebtInvest            float64 `json:"debt_invest"`                // 债权投资(元)
	OthDebtInvest         float64 `json:"oth_debt_invest"`            // 其他债权投资(元)
	OthEqInvest           float64 `json:"oth_eq_invest"`              // 其他权益工具投资(元)
	OthIlliqFinAssets     float64 `json:"oth_illiq_fin_assets"`       // 其他非流动金融资产(元)
	OthEqPpbond           float64 `json:"oth_eq_ppbond"`              // 其他权益工具:永续债(元)
	ReceivFinancing       float64 `json:"receiv_financing"`           // 应收款项融资
	UseRightAssets        float64 `json:"use_right_assets"`           // 使用权资产
	LeaseLiab             float64 `json:"lease_liab"`                 // 租赁负债
	ContractAssets        float64 `json:"contract_assets"`            // 合同资产
	ContractLiab          float64 `json:"contract_liab"`              // 合同负债
	AccountsReceivBill    float64 `json:"accounts_receiv_bill"`       // 应收票据及应收账款
	AccountsPay           float64 `json:"accounts_pay"`               // 应付票据及应付账款
	OthRcvTotal           float64 `json:"oth_rcv_total"`              // 其他应收款(合计)（元）
	FixAssetsTotal        float64 `json:"fix_assets_total"`           // 固定资产(合计)(元)
	UpdateFlag            string  `json:"update_flag"`                // 更新标识
}

type BalanceSheetRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // 股票代码
	AnnDate    string `json:"ann_date"`    // 公告日期(YYYYMMDD格式，下同)
	StartDate  string `json:"start_date"`  // 公告开始日期
	EndDate    string `json:"end_date"`    // 公告结束日期
	Period     string `json:"period"`      // 报告期(每个季度最后一天的日期，比如20171231表示年报)
	ReportType string `json:"report_type"` // 报告类型：见下方详细说明
	CompType   string `json:"comp_type"`   // 公司类型：1一般工商业 2银行 3保险 4证券
}

type BalanceSheetResponse struct {
	List []*BalanceSheet `json:"list"`
}

func (x *BalanceSheetResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 现金流量表|cashflow
type CashFlow struct {
	TsCode                   string  `json:"ts_code"`                     // TS股票代码
	AnnDate                  string  `json:"ann_date"`                    // 公告日期
	FAnnDate                 string  `json:"f_ann_date"`                  // 实际公告日期
	EndDate                  string  `json:"end_date"`                    // 报告期
	CompType                 string  `json:"comp_type"`                   // 公司类型(1一般工商业2银行3保险4证券)
	ReportType               string  `json:"report_type"`                 // 报表类型
	EndType                  string  `json:"end_type"`                    // 报告期类型
	NetProfit                float64 `json:"net_profit"`                  // 净利润
	FinanExp                 float64 `json:"finan_exp"`                   // 财务费用
	CFrSaleSg                float64 `json:"c_fr_sale_sg"`                // 销售商品、提供劳务收到的现金
	RecpTaxRends             float64 `json:"recp_tax_rends"`              // 收到的税费返还
	NDeposIncrFi             float64 `json:"n_depos_incr_fi"`             // 客户存款和同业存放款项净增加额
	NIncrLoansCb             float64 `json:"n_incr_loans_cb"`             // 向中央银行借款净增加额
	NIncBorrOthFi            float64 `json:"n_inc_borr_oth_fi"`           // 向其他金融机构拆入资金净增加额
	PremFrOrigContr          float64 `json:"prem_fr_orig_contr"`          // 收到原保险合同保费取得的现金
	NIncrInsuredDep          float64 `json:"n_incr_insured_dep"`          // 保户储金净增加额
	NReinsurPrem             float64 `json:"n_reinsur_prem"`              // 收到再保业务现金净额
	NIncrDispTfa             float64 `json:"n_incr_disp_tfa"`             // 处置交易性金融资产净增加额
	IfcCashIncr              float64 `json:"ifc_cash_incr"`               // 收取利息和手续费净增加额
	NIncrDispFaas            float64 `json:"n_incr_disp_faas"`            // 处置可供出售金融资产净增加额
	NIncrLoansOthBank        float64 `json:"n_incr_loans_oth_bank"`       // 拆入资金净增加额
	NCapIncrRepur            float64 `json:"n_cap_incr_repur"`            // 回购业务资金净增加额
	CFrOthOperateA           float64 `json:"c_fr_oth_operate_a"`          // 收到其他与经营活动有关的现金
	CInfFrOperateA           float64 `json:"c_inf_fr_operate_a"`          // 经营活动现金流入小计
	CPaidGoodsS              float64 `json:"c_paid_goods_s"`              // 购买商品、接受劳务支付的现金
	CPaidToForEmpl           float64 `json:"c_paid_to_for_empl"`          // 支付给职工以及为职工支付的现金
	CPaidForTaxes            float64 `json:"c_paid_for_taxes"`            // 支付的各项税费
	NIncrCltLoanAdv          float64 `json:"n_incr_clt_loan_adv"`         // 客户贷款及垫款净增加额
	NIncrDepCbob             float64 `json:"n_incr_dep_cbob"`             // 存放央行和同业款项净增加额
	CPayClaimsOrigInco       float64 `json:"c_pay_claims_orig_inco"`      // 支付原保险合同赔付款项的现金
	PayHandlingChrg          float64 `json:"pay_handling_chrg"`           // 支付手续费的现金
	PayCommInsurPlcy         float64 `json:"pay_comm_insur_plcy"`         // 支付保单红利的现金
	OthCashPayOperAct        float64 `json:"oth_cash_pay_oper_act"`       // 支付其他与经营活动有关的现金
	StCashOutAct             float64 `json:"st_cash_out_act"`             // 经营活动现金流出小计
	NCashflowAct             float64 `json:"n_cashflow_act"`              // 经营活动产生的现金流量净额
	OthRecpRalInvAct         float64 `json:"oth_recp_ral_inv_act"`        // 收到其他与投资活动有关的现金
	CDispWithdrwlInvest      float64 `json:"c_disp_withdrwl_invest"`      // 收回投资收到的现金
	CRecpReturnInvest        float64 `json:"c_recp_return_invest"`        // 取得投资收益收到的现金
	NRecpDispFiolta          float64 `json:"n_recp_disp_fiolta"`          // 处置固定资产、无形资产和其他长期资产收回的现金净额
	NRecpDispSobu            float64 `json:"n_recp_disp_sobu"`            // 处置子公司及其他营业单位收到的现金净额
	StotInflowsInvAct        float64 `json:"stot_inflows_inv_act"`        // 投资活动现金流入小计
	CPayAcqConstFiolta       float64 `json:"c_pay_acq_const_fiolta"`      // 购建固定资产、无形资产和其他长期资产支付的现金
	CPaidInvest              float64 `json:"c_paid_invest"`               // 投资支付的现金
	NDispSubsOthBiz          float64 `json:"n_disp_subs_oth_biz"`         // 取得子公司及其他营业单位支付的现金净额
	OthPayRalInvAct          float64 `json:"oth_pay_ral_inv_act"`         // 支付其他与投资活动有关的现金
	NIncrPledgeLoan          float64 `json:"n_incr_pledge_loan"`          // 质押贷款净增加额
	StotOutInvAct            float64 `json:"stot_out_inv_act"`            // 投资活动现金流出小计
	NCashflowInvAct          float64 `json:"n_cashflow_inv_act"`          // 投资活动产生的现金流量净额
	CRecpBorrow              float64 `json:"c_recp_borrow"`               // 取得借款收到的现金
	ProcIssueBonds           float64 `json:"proc_issue_bonds"`            // 发行债券收到的现金
	OthCashRecpRalFncAct     float64 `json:"oth_cash_recp_ral_fnc_act"`   // 收到其他与筹资活动有关的现金
	StotCashInFncAct         float64 `json:"stot_cash_in_fnc_act"`        // 筹资活动现金流入小计
	FreeCashflow             float64 `json:"free_cashflow"`               // 企业自由现金流量
	CPrepayAmtBorr           float64 `json:"c_prepay_amt_borr"`           // 偿还债务支付的现金
	CPayDistDpcpIntExp       float64 `json:"c_pay_dist_dpcp_int_exp"`     // 分配股利、利润或偿付利息支付的现金
	InclDvdProfitPaidScMs    float64 `json:"incl_dvd_profit_paid_sc_ms"`  // 其中:子公司支付给少数股东的股利、利润
	OthCashpayRalFncAct      float64 `json:"oth_cashpay_ral_fnc_act"`     // 支付其他与筹资活动有关的现金
	StotCashoutFncAct        float64 `json:"stot_cashout_fnc_act"`        // 筹资活动现金流出小计
	NCashFlowsFncAct         float64 `json:"n_cash_flows_fnc_act"`        // 筹资活动产生的现金流量净额
	EffFxFluCash             float64 `json:"eff_fx_flu_cash"`             // 汇率变动对现金的影响
	NIncrCashCashEqu         float64 `json:"n_incr_cash_cash_equ"`        // 现金及现金等价物净增加额
	CCashEquBegPeriod        float64 `json:"c_cash_equ_beg_period"`       // 期初现金及现金等价物余额
	CCashEquEndPeriod        float64 `json:"c_cash_equ_end_period"`       // 期末现金及现金等价物余额
	CRecpCapContrib          float64 `json:"c_recp_cap_contrib"`          // 吸收投资收到的现金
	InclCashRecSaims         float64 `json:"incl_cash_rec_saims"`         // 其中:子公司吸收少数股东投资收到的现金
	UnconInvestLoss          float64 `json:"uncon_invest_loss"`           // 未确认投资损失
	ProvDeprAssets           float64 `json:"prov_depr_assets"`            // 加:资产减值准备
	DeprFaCogaDpba           float64 `json:"depr_fa_coga_dpba"`           // 固定资产折旧、油气资产折耗、生产性生物资产折旧
	AmortIntangAssets        float64 `json:"amort_intang_assets"`         // 无形资产摊销
	LtAmortDeferredExp       float64 `json:"lt_amort_deferred_exp"`       // 长期待摊费用摊销
	DecrDeferredExp          float64 `json:"decr_deferred_exp"`           // 待摊费用减少
	IncrAccExp               float64 `json:"incr_acc_exp"`                // 预提费用增加
	LossDispFiolta           float64 `json:"loss_disp_fiolta"`            // 处置固定、无形资产和其他长期资产的损失
	LossScrFa                float64 `json:"loss_scr_fa"`                 // 固定资产报废损失
	LossFvChg                float64 `json:"loss_fv_chg"`                 // 公允价值变动损失
	InvestLoss               float64 `json:"invest_loss"`                 // 投资损失
	DecrDefIncTaxAssets      float64 `json:"decr_def_inc_tax_assets"`     // 递延所得税资产减少
	IncrDefIncTaxLiab        float64 `json:"incr_def_inc_tax_liab"`       // 递延所得税负债增加
	DecrInventories          float64 `json:"decr_inventories"`            // 存货的减少
	DecrOperPayable          float64 `json:"decr_oper_payable"`           // 经营性应收项目的减少
	IncrOperPayable          float64 `json:"incr_oper_payable"`           // 经营性应付项目的增加
	Others                   float64 `json:"others"`                      // 其他
	ImNetCashflowOperAct     float64 `json:"im_net_cashflow_oper_act"`    // 经营活动产生的现金流量净额(间接法)
	ConvDebtIntoCap          float64 `json:"conv_debt_into_cap"`          // 债务转为资本
	ConvCopbondsDueWithin_1Y float64 `json:"conv_copbonds_due_within_1y"` // 一年内到期的可转换公司债券
	FaFncLeases              float64 `json:"fa_fnc_leases"`               // 融资租入固定资产
	ImNIncrCashEqu           float64 `json:"im_n_incr_cash_equ"`          // 现金及现金等价物净增加额(间接法)
	NetDismCapitalAdd        float64 `json:"net_dism_capital_add"`        // 拆出资金净增加额
	NetCashReceSec           float64 `json:"net_cash_rece_sec"`           // 代理买卖证券收到的现金净额(元)
	CreditImpaLoss           float64 `json:"credit_impa_loss"`            // 信用减值损失
	UseRightAssetDep         float64 `json:"use_right_asset_dep"`         // 使用权资产折旧
	OthLossAsset             float64 `json:"oth_loss_asset"`              // 其他资产减值损失
	EndBalCash               float64 `json:"end_bal_cash"`                // 现金的期末余额
	BegBalCash               float64 `json:"beg_bal_cash"`                // 减:现金的期初余额
	EndBalCashEqu            float64 `json:"end_bal_cash_equ"`            // 加:现金等价物的期末余额
	BegBalCashEqu            float64 `json:"beg_bal_cash_equ"`            // 减:现金等价物的期初余额
	UpdateFlag               string  `json:"update_flag"`                 // 更新标志(1最新）
}

type CashFlowRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // 股票代码
	AnnDate    string `json:"ann_date"`    // 公告日期（YYYYMMDD格式，下同）
	FAnnDate   string `json:"f_ann_date"`  // 实际公告日期
	StartDate  string `json:"start_date"`  // 公告开始日期
	EndDate    string `json:"end_date"`    // 公告结束日期
	Period     string `json:"period"`      // 报告期(每个季度最后一天的日期，比如20171231表示年报)
	ReportType string `json:"report_type"` // 报告类型：见下方详细说明
	CompType   string `json:"comp_type"`   // 公司类型：1一般工商业 2银行 3保险 4证券
	IsCalc     string `json:"is_calc"`     // 是否计算报表
}

type CashFlowResponse struct {
	List []*CashFlow `json:"list"`
}

func (x *CashFlowResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 业绩预告|forecast
type ForeCast struct {
	TsCode        string  `json:"ts_code"`         // TS股票代码
	AnnDate       string  `json:"ann_date"`        // 公告日期
	EndDate       string  `json:"end_date"`        // 报告期
	Type          string  `json:"type"`            // 业绩预告类型(预增/预减/扭亏/首亏/续亏/续盈/略增/略减)
	PChangeMin    float64 `json:"p_change_min"`    // 预告净利润变动幅度下限（%）
	PChangeMax    float64 `json:"p_change_max"`    // 预告净利润变动幅度上限（%）
	NetProfitMin  float64 `json:"net_profit_min"`  // 预告净利润下限（万元）
	NetProfitMax  float64 `json:"net_profit_max"`  // 预告净利润上限（万元）
	LastParentNet float64 `json:"last_parent_net"` // 上年同期归属母公司净利润
	FirstAnnDate  string  `json:"first_ann_date"`  // 首次公告日
	Summary       string  `json:"summary"`         // 业绩预告摘要
	ChangeReason  string  `json:"change_reason"`   // 业绩变动原因
}

type ForeCastRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码(二选一)
	AnnDate   string `json:"ann_date"`   // 公告日期 (二选一)
	StartDate string `json:"start_date"` // 公告开始日期
	EndDate   string `json:"end_date"`   // 公告结束日期
	Period    string `json:"period"`     // 报告期(每个季度最后一天的日期，比如20171231表示年报)
	Type      string `json:"type"`       // 预告类型(预增/预减/扭亏/首亏/续亏/续盈/略增/略减)
}

type ForeCastResponse struct {
	List []*ForeCast `json:"list"`
}

func (x *ForeCastResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 业绩快报|express
type Express struct {
	TsCode                string  `json:"ts_code"`                    // TS股票代码
	AnnDate               string  `json:"ann_date"`                   // 公告日期
	EndDate               string  `json:"end_date"`                   // 报告期
	Revenue               float64 `json:"revenue"`                    // 营业收入(元)
	OperateProfit         float64 `json:"operate_profit"`             // 营业利润(元)
	TotalProfit           float64 `json:"total_profit"`               // 利润总额(元)
	NIncome               float64 `json:"n_income"`                   // 净利润(元)
	TotalAssets           float64 `json:"total_assets"`               // 总资产(元)
	TotalHldrEqyExcMinInt float64 `json:"total_hldr_eqy_exc_min_int"` // 股东权益合计(不含少数股东权益)(元)
	DilutedEps            float64 `json:"diluted_eps"`                // 每股收益(摊薄)(元)
	DilutedRoe            float64 `json:"diluted_roe"`                // 净资产收益率(摊薄)(%)
	YoyNetProfit          float64 `json:"yoy_net_profit"`             // 去年同期修正后净利润
	Bps                   float64 `json:"bps"`                        // 每股净资产
	YoySales              float64 `json:"yoy_sales"`                  // 同比增长率:营业收入
	YoyOp                 float64 `json:"yoy_op"`                     // 同比增长率:营业利润
	YoyTp                 float64 `json:"yoy_tp"`                     // 同比增长率:利润总额
	YoyDeduNp             float64 `json:"yoy_dedu_np"`                // 同比增长率:归属母公司股东的净利润
	YoyEps                float64 `json:"yoy_eps"`                    // 同比增长率:基本每股收益
	YoyRoe                float64 `json:"yoy_roe"`                    // 同比增减:加权平均净资产收益率
	GrowthAssets          float64 `json:"growth_assets"`              // 比年初增长率:总资产
	YoyEquity             float64 `json:"yoy_equity"`                 // 比年初增长率:归属母公司的股东权益
	GrowthBps             float64 `json:"growth_bps"`                 // 比年初增长率:归属于母公司股东的每股净资产
	OrLastYear            float64 `json:"or_last_year"`               // 去年同期营业收入
	OpLastYear            float64 `json:"op_last_year"`               // 去年同期营业利润
	TpLastYear            float64 `json:"tp_last_year"`               // 去年同期利润总额
	NpLastYear            float64 `json:"np_last_year"`               // 去年同期净利润
	EpsLastYear           float64 `json:"eps_last_year"`              // 去年同期每股收益
	OpenNetAssets         float64 `json:"open_net_assets"`            // 期初净资产
	OpenBps               float64 `json:"open_bps"`                   // 期初每股净资产
	PerfSummary           string  `json:"perf_summary"`               // 业绩简要说明
	IsAudit               int64   `json:"is_audit"`                   // 是否审计： 1是 0否
	Remark                string  `json:"remark"`                     // 备注
}

type ExpressRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	AnnDate   string `json:"ann_date"`   // 公告日期
	StartDate string `json:"start_date"` // 公告开始日期
	EndDate   string `json:"end_date"`   // 公告结束日期
	Period    string `json:"period"`     // 报告期(每个季度最后一天的日期,比如20171231表示年报)
}

type ExpressResponse struct {
	List []*Express `json:"list"`
}

func (x *ExpressResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 分红送股|dividend
type Dividend struct {
	TsCode      string  `json:"ts_code"`      // TS代码
	EndDate     string  `json:"end_date"`     // 分红年度
	AnnDate     string  `json:"ann_date"`     // 预案公告日
	DivProc     string  `json:"div_proc"`     // 实施进度
	StkDiv      float64 `json:"stk_div"`      // 每股送转
	StkBoRate   float64 `json:"stk_bo_rate"`  // 每股送股比例
	StkCoRate   float64 `json:"stk_co_rate"`  // 每股转增比例
	CashDiv     float64 `json:"cash_div"`     // 每股分红（税后）
	CashDivTax  float64 `json:"cash_div_tax"` // 每股分红（税前）
	RecordDate  string  `json:"record_date"`  // 股权登记日
	ExDate      string  `json:"ex_date"`      // 除权除息日
	PayDate     string  `json:"pay_date"`     // 派息日
	DivListdate string  `json:"div_listdate"` // 红股上市日
	ImpAnnDate  string  `json:"imp_ann_date"` // 实施公告日
	BaseDate    string  `json:"base_date"`    // 基准日
	BaseShare   float64 `json:"base_share"`   // 基准股本（万）
}

type DividendRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`      // TS代码
	AnnDate    string `json:"ann_date"`     // 公告日
	RecordDate string `json:"record_date"`  // 股权登记日期
	ExDate     string `json:"ex_date"`      // 除权除息日
	ImpAnnDate string `json:"imp_ann_date"` // 实施公告日
}

type DividendResponse struct {
	List []*Dividend `json:"list"`
}

func (x *DividendResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 财务指标|fina_indicator
type FinaIndicator struct {
	TsCode                    string  `json:"ts_code"`                      // TS代码
	AnnDate                   string  `json:"ann_date"`                     // 公告日期
	EndDate                   string  `json:"end_date"`                     // 报告期
	Eps                       float64 `json:"eps"`                          // 基本每股收益
	DtEps                     float64 `json:"dt_eps"`                       // 稀释每股收益
	TotalRevenuePs            float64 `json:"total_revenue_ps"`             // 每股营业总收入
	RevenuePs                 float64 `json:"revenue_ps"`                   // 每股营业收入
	CapitalResePs             float64 `json:"capital_rese_ps"`              // 每股资本公积
	SurplusResePs             float64 `json:"surplus_rese_ps"`              // 每股盈余公积
	UndistProfitPs            float64 `json:"undist_profit_ps"`             // 每股未分配利润
	ExtraItem                 float64 `json:"extra_item"`                   // 非经常性损益
	ProfitDedt                float64 `json:"profit_dedt"`                  // 扣除非经常性损益后的净利润（扣非净利润）
	GrossMargin               float64 `json:"gross_margin"`                 // 毛利
	CurrentRatio              float64 `json:"current_ratio"`                // 流动比率
	QuickRatio                float64 `json:"quick_ratio"`                  // 速动比率
	CashRatio                 float64 `json:"cash_ratio"`                   // 保守速动比率
	InvturnDays               float64 `json:"invturn_days"`                 // 存货周转天数
	ArturnDays                float64 `json:"arturn_days"`                  // 应收账款周转天数
	InvTurn                   float64 `json:"inv_turn"`                     // 存货周转率
	ArTurn                    float64 `json:"ar_turn"`                      // 应收账款周转率
	CaTurn                    float64 `json:"ca_turn"`                      // 流动资产周转率
	FaTurn                    float64 `json:"fa_turn"`                      // 固定资产周转率
	AssetsTurn                float64 `json:"assets_turn"`                  // 总资产周转率
	OpIncome                  float64 `json:"op_income"`                    // 经营活动净收益
	ValuechangeIncome         float64 `json:"valuechange_income"`           // 价值变动净收益
	InterstIncome             float64 `json:"interst_income"`               // 利息费用
	Daa                       float64 `json:"daa"`                          // 折旧与摊销
	Ebit                      float64 `json:"ebit"`                         // 息税前利润
	Ebitda                    float64 `json:"ebitda"`                       // 息税折旧摊销前利润
	Fcff                      float64 `json:"fcff"`                         // 企业自由现金流量
	Fcfe                      float64 `json:"fcfe"`                         // 股权自由现金流量
	CurrentExint              float64 `json:"current_exint"`                // 无息流动负债
	NoncurrentExint           float64 `json:"noncurrent_exint"`             // 无息非流动负债
	Interestdebt              float64 `json:"interestdebt"`                 // 带息债务
	Netdebt                   float64 `json:"netdebt"`                      // 净债务
	TangibleAsset             float64 `json:"tangible_asset"`               // 有形资产
	WorkingCapital            float64 `json:"working_capital"`              // 营运资金
	NetworkingCapital         float64 `json:"networking_capital"`           // 营运流动资本
	InvestCapital             float64 `json:"invest_capital"`               // 全部投入资本
	RetainedEarnings          float64 `json:"retained_earnings"`            // 留存收益
	Diluted2Eps               float64 `json:"diluted2_eps"`                 // 期末摊薄每股收益
	Bps                       float64 `json:"bps"`                          // 每股净资产
	Ocfps                     float64 `json:"ocfps"`                        // 每股经营活动产生的现金流量净额
	Retainedps                float64 `json:"retainedps"`                   // 每股留存收益
	Cfps                      float64 `json:"cfps"`                         // 每股现金流量净额
	EbitPs                    float64 `json:"ebit_ps"`                      // 每股息税前利润
	FcffPs                    float64 `json:"fcff_ps"`                      // 每股企业自由现金流量
	FcfePs                    float64 `json:"fcfe_ps"`                      // 每股股东自由现金流量
	NetprofitMargin           float64 `json:"netprofit_margin"`             // 销售净利率
	GrossprofitMargin         float64 `json:"grossprofit_margin"`           // 销售毛利率
	CogsOfSales               float64 `json:"cogs_of_sales"`                // 销售成本率
	ExpenseOfSales            float64 `json:"expense_of_sales"`             // 销售期间费用率
	ProfitToGr                float64 `json:"profit_to_gr"`                 // 净利润/营业总收入
	SaleexpToGr               float64 `json:"saleexp_to_gr"`                // 销售费用/营业总收入
	AdminexpOfGr              float64 `json:"adminexp_of_gr"`               // 管理费用/营业总收入
	FinaexpOfGr               float64 `json:"finaexp_of_gr"`                // 财务费用/营业总收入
	ImpaiTtm                  float64 `json:"impai_ttm"`                    // 资产减值损失/营业总收入
	GcOfGr                    float64 `json:"gc_of_gr"`                     // 营业总成本/营业总收入
	OpOfGr                    float64 `json:"op_of_gr"`                     // 营业利润/营业总收入
	EbitOfGr                  float64 `json:"ebit_of_gr"`                   // 息税前利润/营业总收入
	Roe                       float64 `json:"roe"`                          // 净资产收益率
	RoeWaa                    float64 `json:"roe_waa"`                      // 加权平均净资产收益率
	RoeDt                     float64 `json:"roe_dt"`                       // 净资产收益率(扣除非经常损益)
	Roa                       float64 `json:"roa"`                          // 总资产报酬率
	Npta                      float64 `json:"npta"`                         // 总资产净利润
	Roic                      float64 `json:"roic"`                         // 投入资本回报率
	RoeYearly                 float64 `json:"roe_yearly"`                   // 年化净资产收益率
	Roa2Yearly                float64 `json:"roa2_yearly"`                  // 年化总资产报酬率
	RoeAvg                    float64 `json:"roe_avg"`                      // 平均净资产收益率(增发条件)
	OpincomeOfEbt             float64 `json:"opincome_of_ebt"`              // 经营活动净收益/利润总额
	InvestincomeOfEbt         float64 `json:"investincome_of_ebt"`          // 价值变动净收益/利润总额
	NOpProfitOfEbt            float64 `json:"n_op_profit_of_ebt"`           // 营业外收支净额/利润总额
	TaxToEbt                  float64 `json:"tax_to_ebt"`                   // 所得税/利润总额
	DtprofitToProfit          float64 `json:"dtprofit_to_profit"`           // 扣除非经常损益后的净利润/净利润
	SalescashToOr             float64 `json:"salescash_to_or"`              // 销售商品提供劳务收到的现金/营业收入
	OcfToOr                   float64 `json:"ocf_to_or"`                    // 经营活动产生的现金流量净额/营业收入
	OcfToOpincome             float64 `json:"ocf_to_opincome"`              // 经营活动产生的现金流量净额/经营活动净收益
	CapitalizedToDa           float64 `json:"capitalized_to_da"`            // 资本支出/折旧和摊销
	DebtToAssets              float64 `json:"debt_to_assets"`               // 资产负债率
	AssetsToEqt               float64 `json:"assets_to_eqt"`                // 权益乘数
	DpAssetsToEqt             float64 `json:"dp_assets_to_eqt"`             // 权益乘数(杜邦分析)
	CaToAssets                float64 `json:"ca_to_assets"`                 // 流动资产/总资产
	NcaToAssets               float64 `json:"nca_to_assets"`                // 非流动资产/总资产
	TbassetsToTotalassets     float64 `json:"tbassets_to_totalassets"`      // 有形资产/总资产
	IntToTalcap               float64 `json:"int_to_talcap"`                // 带息债务/全部投入资本
	EqtToTalcapital           float64 `json:"eqt_to_talcapital"`            // 归属于母公司的股东权益/全部投入资本
	CurrentdebtToDebt         float64 `json:"currentdebt_to_debt"`          // 流动负债/负债合计
	LongdebToDebt             float64 `json:"longdeb_to_debt"`              // 非流动负债/负债合计
	OcfToShortdebt            float64 `json:"ocf_to_shortdebt"`             // 经营活动产生的现金流量净额/流动负债
	DebtToEqt                 float64 `json:"debt_to_eqt"`                  // 产权比率
	EqtToDebt                 float64 `json:"eqt_to_debt"`                  // 归属于母公司的股东权益/负债合计
	EqtToInterestdebt         float64 `json:"eqt_to_interestdebt"`          // 归属于母公司的股东权益/带息债务
	TangibleassetToDebt       float64 `json:"tangibleasset_to_debt"`        // 有形资产/负债合计
	TangassetToIntdebt        float64 `json:"tangasset_to_intdebt"`         // 有形资产/带息债务
	TangibleassetToNetdebt    float64 `json:"tangibleasset_to_netdebt"`     // 有形资产/净债务
	OcfToDebt                 float64 `json:"ocf_to_debt"`                  // 经营活动产生的现金流量净额/负债合计
	OcfToInterestdebt         float64 `json:"ocf_to_interestdebt"`          // 经营活动产生的现金流量净额/带息债务
	OcfToNetdebt              float64 `json:"ocf_to_netdebt"`               // 经营活动产生的现金流量净额/净债务
	EbitToInterest            float64 `json:"ebit_to_interest"`             // 已获利息倍数(EBIT/利息费用)
	LongdebtToWorkingcapital  float64 `json:"longdebt_to_workingcapital"`   // 长期债务与营运资金比率
	EbitdaToDebt              float64 `json:"ebitda_to_debt"`               // 息税折旧摊销前利润/负债合计
	TurnDays                  float64 `json:"turn_days"`                    // 营业周期
	RoaYearly                 float64 `json:"roa_yearly"`                   // 年化总资产净利率
	RoaDp                     float64 `json:"roa_dp"`                       // 总资产净利率(杜邦分析)
	FixedAssets               float64 `json:"fixed_assets"`                 // 固定资产合计
	ProfitPrefinExp           float64 `json:"profit_prefin_exp"`            // 扣除财务费用前营业利润
	NonOpProfit               float64 `json:"non_op_profit"`                // 非营业利润
	OpToEbt                   float64 `json:"op_to_ebt"`                    // 营业利润／利润总额
	NopToEbt                  float64 `json:"nop_to_ebt"`                   // 非营业利润／利润总额
	OcfToProfit               float64 `json:"ocf_to_profit"`                // 经营活动产生的现金流量净额／营业利润
	CashToLiqdebt             float64 `json:"cash_to_liqdebt"`              // 货币资金／流动负债
	CashToLiqdebtWithinterest float64 `json:"cash_to_liqdebt_withinterest"` // 货币资金／带息流动负债
	OpToLiqdebt               float64 `json:"op_to_liqdebt"`                // 营业利润／流动负债
	OpToDebt                  float64 `json:"op_to_debt"`                   // 营业利润／负债合计
	RoicYearly                float64 `json:"roic_yearly"`                  // 年化投入资本回报率
	TotalFaTrun               float64 `json:"total_fa_trun"`                // 固定资产合计周转率
	ProfitToOp                float64 `json:"profit_to_op"`                 // 利润总额／营业收入
	QOpincome                 float64 `json:"q_opincome"`                   // 经营活动单季度净收益
	QInvestincome             float64 `json:"q_investincome"`               // 价值变动单季度净收益
	QDtprofit                 float64 `json:"q_dtprofit"`                   // 扣除非经常损益后的单季度净利润
	QEps                      float64 `json:"q_eps"`                        // 每股收益(单季度)
	QNetprofitMargin          float64 `json:"q_netprofit_margin"`           // 销售净利率(单季度)
	QGsprofitMargin           float64 `json:"q_gsprofit_margin"`            // 销售毛利率(单季度)
	QExpToSales               float64 `json:"q_exp_to_sales"`               // 销售期间费用率(单季度)
	QProfitToGr               float64 `json:"q_profit_to_gr"`               // 净利润／营业总收入(单季度)
	QSaleexpToGr              float64 `json:"q_saleexp_to_gr"`              // 销售费用／营业总收入 (单季度)
	QAdminexpToGr             float64 `json:"q_adminexp_to_gr"`             // 管理费用／营业总收入 (单季度)
	QFinaexpToGr              float64 `json:"q_finaexp_to_gr"`              // 财务费用／营业总收入 (单季度)
	QImpairToGrTtm            float64 `json:"q_impair_to_gr_ttm"`           // 资产减值损失／营业总收入(单季度)
	QGcToGr                   float64 `json:"q_gc_to_gr"`                   // 营业总成本／营业总收入 (单季度)
	QOpToGr                   float64 `json:"q_op_to_gr"`                   // 营业利润／营业总收入(单季度)
	QRoe                      float64 `json:"q_roe"`                        // 净资产收益率(单季度)
	QDtRoe                    float64 `json:"q_dt_roe"`                     // 净资产单季度收益率(扣除非经常损益)
	QNpta                     float64 `json:"q_npta"`                       // 总资产净利润(单季度)
	QOpincomeToEbt            float64 `json:"q_opincome_to_ebt"`            // 经营活动净收益／利润总额(单季度)
	QInvestincomeToEbt        float64 `json:"q_investincome_to_ebt"`        // 价值变动净收益／利润总额(单季度)
	QDtprofitToProfit         float64 `json:"q_dtprofit_to_profit"`         // 扣除非经常损益后的净利润／净利润(单季度)
	QSalescashToOr            float64 `json:"q_salescash_to_or"`            // 销售商品提供劳务收到的现金／营业收入(单季度)
	QOcfToSales               float64 `json:"q_ocf_to_sales"`               // 经营活动产生的现金流量净额／营业收入(单季度)
	QOcfToOr                  float64 `json:"q_ocf_to_or"`                  // 经营活动产生的现金流量净额／经营活动净收益(单季度)
	BasicEpsYoy               float64 `json:"basic_eps_yoy"`                // 基本每股收益同比增长率(%)
	DtEpsYoy                  float64 `json:"dt_eps_yoy"`                   // 稀释每股收益同比增长率(%)
	CfpsYoy                   float64 `json:"cfps_yoy"`                     // 每股经营活动产生的现金流量净额同比增长率(%)
	OpYoy                     float64 `json:"op_yoy"`                       // 营业利润同比增长率(%)
	EbtYoy                    float64 `json:"ebt_yoy"`                      // 利润总额同比增长率(%)
	NetprofitYoy              float64 `json:"netprofit_yoy"`                // 归属母公司股东的净利润同比增长率(%)
	DtNetprofitYoy            float64 `json:"dt_netprofit_yoy"`             // 归属母公司股东的净利润-扣除非经常损益同比增长率(%)
	OcfYoy                    float64 `json:"ocf_yoy"`                      // 经营活动产生的现金流量净额同比增长率(%)
	RoeYoy                    float64 `json:"roe_yoy"`                      // 净资产收益率(摊薄)同比增长率(%)
	BpsYoy                    float64 `json:"bps_yoy"`                      // 每股净资产相对年初增长率(%)
	AssetsYoy                 float64 `json:"assets_yoy"`                   // 资产总计相对年初增长率(%)
	EqtYoy                    float64 `json:"eqt_yoy"`                      // 归属母公司的股东权益相对年初增长率(%)
	TrYoy                     float64 `json:"tr_yoy"`                       // 营业总收入同比增长率(%)
	OrYoy                     float64 `json:"or_yoy"`                       // 营业收入同比增长率(%)
	QGrYoy                    float64 `json:"q_gr_yoy"`                     // 营业总收入同比增长率(%)(单季度)
	QGrQoq                    float64 `json:"q_gr_qoq"`                     // 营业总收入环比增长率(%)(单季度)
	QSalesYoy                 float64 `json:"q_sales_yoy"`                  // 营业收入同比增长率(%)(单季度)
	QSalesQoq                 float64 `json:"q_sales_qoq"`                  // 营业收入环比增长率(%)(单季度)
	QOpYoy                    float64 `json:"q_op_yoy"`                     // 营业利润同比增长率(%)(单季度)
	QOpQoq                    float64 `json:"q_op_qoq"`                     // 营业利润环比增长率(%)(单季度)
	QProfitYoy                float64 `json:"q_profit_yoy"`                 // 净利润同比增长率(%)(单季度)
	QProfitQoq                float64 `json:"q_profit_qoq"`                 // 净利润环比增长率(%)(单季度)
	QNetprofitYoy             float64 `json:"q_netprofit_yoy"`              // 归属母公司股东的净利润同比增长率(%)(单季度)
	QNetprofitQoq             float64 `json:"q_netprofit_qoq"`              // 归属母公司股东的净利润环比增长率(%)(单季度)
	EquityYoy                 float64 `json:"equity_yoy"`                   // 净资产同比增长率
	RdExp                     float64 `json:"rd_exp"`                       // 研发费用
	UpdateFlag                string  `json:"update_flag"`                  // 更新标识
}

type FinaIndicatorRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS股票代码,e.g. 600001.SH/000001.SZ
	AnnDate   string `json:"ann_date"`   // 公告日期
	StartDate string `json:"start_date"` // 报告期开始日期
	EndDate   string `json:"end_date"`   // 报告期结束日期
	Period    string `json:"period"`     // 报告期(每个季度最后一天的日期,比如20171231表示年报)
}

type FinaIndicatorResponse struct {
	List []*FinaIndicator `json:"list"`
}

func (x *FinaIndicatorResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 财务审计意见|fina_audit
type FinaAudit struct {
	TsCode      string  `json:"ts_code"`      // TS股票代码
	AnnDate     string  `json:"ann_date"`     // 公告日期
	EndDate     string  `json:"end_date"`     // 报告期
	AuditResult string  `json:"audit_result"` // 审计结果
	AuditFees   float64 `json:"audit_fees"`   // 审计总费用（元）
	AuditAgency string  `json:"audit_agency"` // 会计事务所
	AuditSign   string  `json:"audit_sign"`   // 签字会计师
}

type FinaAuditRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	AnnDate   string `json:"ann_date"`   // 公告日期
	StartDate string `json:"start_date"` // 公告开始日期
	EndDate   string `json:"end_date"`   // 公告结束日期
	Period    string `json:"period"`     // 报告期(每个季度最后一天的日期,比如20171231表示年报)
}

type FinaAuditResponse struct {
	List []*FinaAudit `json:"list"`
}

func (x *FinaAuditResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 主营业务构成|fina_mainbz
type FinaMainbz struct {
	TsCode     string  `json:"ts_code"`     // TS代码
	EndDate    string  `json:"end_date"`    // 报告期
	BzItem     string  `json:"bz_item"`     // 主营业务来源
	BzSales    float64 `json:"bz_sales"`    // 主营业务收入(元)
	BzProfit   float64 `json:"bz_profit"`   // 主营业务利润(元)
	BzCost     float64 `json:"bz_cost"`     // 主营业务成本(元)
	CurrType   string  `json:"curr_type"`   // 货币代码
	UpdateFlag string  `json:"update_flag"` // 是否更新
}

type FinaMainbzRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	Period    string `json:"period"`     // 报告期(每个季度最后一天的日期,比如20171231表示年报)
	Type      string `json:"type"`       // 类型：P按产品 D按地区 I按行业（请输入大写字母P或者D）
	StartDate string `json:"start_date"` // 报告期开始日期
	EndDate   string `json:"end_date"`   // 报告期结束日期
}

type FinaMainbzResponse struct {
	List []*FinaMainbz `json:"list"`
}

func (x *FinaMainbzResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 财报披露计划|disclosure_date
type DisclosureDate struct {
	TsCode     string `json:"ts_code"`     // TS代码
	AnnDate    string `json:"ann_date"`    // 最新披露公告日
	EndDate    string `json:"end_date"`    // 报告期
	PreDate    string `json:"pre_date"`    // 预计披露日期
	ActualDate string `json:"actual_date"` // 实际披露日期
	ModifyDate string `json:"modify_date"` // 披露日期修正记录
}

type DisclosureDateRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // TS股票代码
	EndDate    string `json:"end_date"`    // 财报周期（比如20181231表示2018年年报，20180630表示中报)
	PreDate    string `json:"pre_date"`    // 计划披露日期
	ActualDate string `json:"actual_date"` // 实际披露日期
}

type DisclosureDateResponse struct {
	List []*DisclosureDate `json:"list"`
}

func (x *DisclosureDateResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 融资融券交易汇总|margin
type Margin struct {
	TradeDate  string  `json:"trade_date"`  // 交易日期
	ExchangeId string  `json:"exchange_id"` // 交易所代码（SSE上交所SZSE深交所BSE北交所）
	Rzye       float64 `json:"rzye"`        // 融资余额(元)
	Rzmre      float64 `json:"rzmre"`       // 融资买入额(元)
	Rzche      float64 `json:"rzche"`       // 融资偿还额(元)
	Rqye       float64 `json:"rqye"`        // 融券余额(元)
	Rqmcl      float64 `json:"rqmcl"`       // 融券卖出量(股,份,手)
	Rzrqye     float64 `json:"rzrqye"`      // 融资融券余额(元)
	Rqyl       float64 `json:"rqyl"`        // 融券余量(股,份,手)
}

type MarginRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TradeDate  string `json:"trade_date"`  // 交易日期
	ExchangeId string `json:"exchange_id"` // 交易所代码（SSE上交所SZSE深交所BSE北交所）
	StartDate  string `json:"start_date"`  // 开始日期
	EndDate    string `json:"end_date"`    // 结束日期
}

type MarginResponse struct {
	List []*Margin `json:"list"`
}

func (x *MarginResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 融资融券交易明细|margin_detail
type MarginDetail struct {
	TradeDate string  `json:"trade_date"` // 交易日期
	TsCode    string  `json:"ts_code"`    // TS股票代码
	Name      string  `json:"name"`       // 股票名称 （20190910后有数据）
	Rzye      float64 `json:"rzye"`       // 融资余额(元)
	Rqye      float64 `json:"rqye"`       // 融券余额(元)
	Rzmre     float64 `json:"rzmre"`      // 融资买入额(元)
	Rqyl      float64 `json:"rqyl"`       // 融券余量（股）
	Rzche     float64 `json:"rzche"`      // 融资偿还额(元)
	Rqchl     float64 `json:"rqchl"`      // 融券偿还量(股)
	Rqmcl     float64 `json:"rqmcl"`      // 融券卖出量(股,份,手)
	Rzrqye    float64 `json:"rzrqye"`     // 融资融券余额(元)
}

type MarginDetailRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期
	TsCode    string `json:"ts_code"`    // TS代码
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type MarginDetailResponse struct {
	List []*MarginDetail `json:"list"`
}

func (x *MarginDetailResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 融资融券标的|margin_target
type MarginTarget struct {
	TsCode  string `json:"ts_code"`  // 标的代码
	MgType  string `json:"mg_type"`  // 标的类型：B买入标的 S卖出标的
	IsNew   string `json:"is_new"`   // 最新标记：Y是 N否
	InDate  string `json:"in_date"`  // 纳入日期
	OutDate string `json:"out_date"` // 剔除日期
	AnnDate string `json:"ann_date"` // 公布日期
}

type MarginTargetRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	TsCode string `json:"ts_code"` // 股票代码
	IsNew  string `json:"is_new"`  // 是否最新
	MgType string `json:"mg_type"` // 标的类型：B买入标的 S卖出标的
}

type MarginTargetResponse struct {
	List []*MarginTarget `json:"list"`
}

func (x *MarginTargetResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 融资融券标的(盘前)|margin_secs
type MarginSecs struct {
	TsCode    string `json:"ts_code"`    // 标的代码
	TradeDate string `json:"trade_date"` // 交易日
	Exchange  string `json:"exchange"`   // 交易所（SSE上交所 SZSE深交所 BSE北交所）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type MarginSecsRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期
	TsCode    string `json:"ts_code"`    // 标的代码
	Name      string `json:"name"`       // 标的名称
	Exchange  string `json:"exchange"`   // 交易所
}

type MarginSecsResponse struct {
	List []*MarginSecs `json:"list"`
}

func (x *MarginSecsResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 前十大股东|top10_holders
type Top10Holders struct {
	TsCode         string  `json:"ts_code"`          // TS股票代码
	AnnDate        string  `json:"ann_date"`         // 公告日期
	EndDate        string  `json:"end_date"`         // 报告期
	HolderName     string  `json:"holder_name"`      // 股东名称
	HoldAmount     float64 `json:"hold_amount"`      // 持有数量（股）
	HoldRatio      float64 `json:"hold_ratio"`       // 占总股本比例(%)
	HoldFloatRatio float64 `json:"hold_float_ratio"` // 占流通股本比例(%)
	HoldChange     float64 `json:"hold_change"`      // 持股变动
	HolderType     string  `json:"holder_type"`      // 股东类型
}

type Top10HoldersRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	Period    string `json:"period"`     // 报告期（YYYYMMDD格式，一般为每个季度最后一天）
	AnnDate   string `json:"ann_date"`   // 公告日期
	StartDate string `json:"start_date"` // 报告期开始日期
	EndDate   string `json:"end_date"`   // 报告期结束日期
}

type Top10HoldersResponse struct {
	List []*Top10Holders `json:"list"`
}

func (x *Top10HoldersResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 前十大流通股东|top10_floatholders
type Top10Floatholders struct {
	TsCode         string  `json:"ts_code"`          // TS股票代码
	AnnDate        string  `json:"ann_date"`         // 公告日期
	EndDate        string  `json:"end_date"`         // 报告期
	HolderName     string  `json:"holder_name"`      // 股东名称
	HoldAmount     float64 `json:"hold_amount"`      // 持有数量（股）
	HoldRatio      float64 `json:"hold_ratio"`       // 占总股本比例(%)
	HoldFloatRatio float64 `json:"hold_float_ratio"` // 占流通股本比例(%)
	HoldChange     float64 `json:"hold_change"`      // 持股变动
	HolderType     string  `json:"holder_type"`      // 股东类型
}

type Top10FloatholdersRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码
	Period    string `json:"period"`     // 报告期（YYYYMMDD格式，一般为每个季度最后一天）
	AnnDate   string `json:"ann_date"`   // 公告日期
	StartDate string `json:"start_date"` // 报告期开始日期
	EndDate   string `json:"end_date"`   // 报告期结束日期
}

type Top10FloatholdersResponse struct {
	List []*Top10Floatholders `json:"list"`
}

func (x *Top10FloatholdersResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 龙虎榜每日明细|top_list
type TopList struct {
	TradeDate    string  `json:"trade_date"`    // 交易日期
	TsCode       string  `json:"ts_code"`       // TS代码
	Name         string  `json:"name"`          // 名称
	Close        float64 `json:"close"`         // 收盘价
	PctChange    float64 `json:"pct_change"`    // 涨跌幅
	TurnoverRate float64 `json:"turnover_rate"` // 换手率
	Amount       float64 `json:"amount"`        // 总成交额
	LSell        float64 `json:"l_sell"`        // 龙虎榜卖出额
	LBuy         float64 `json:"l_buy"`         // 龙虎榜买入额
	LAmount      float64 `json:"l_amount"`      // 龙虎榜成交额
	NetAmount    float64 `json:"net_amount"`    // 龙虎榜净买入额
	NetRate      float64 `json:"net_rate"`      // 龙虎榜净买额占比
	AmountRate   float64 `json:"amount_rate"`   // 龙虎榜成交额占比
	FloatValues  float64 `json:"float_values"`  // 当日流通市值
	Reason       string  `json:"reason"`        // 上榜理由
}

type TopListRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` //	交易日期
	TsCode    string `json:"ts_code"`    //	股票代码
}

type TopListResponse struct {
	List []*TopList `json:"list"`
}

func (x *TopListResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 龙虎榜机构明细|top_inst
type TopInst struct {
	TradeDate string  `json:"trade_date"` // 交易日期
	TsCode    string  `json:"ts_code"`    // TS代码
	Exalter   string  `json:"exalter"`    // 营业部名称
	Side      string  `json:"side"`       // 买卖类型0：买入金额最大的前5名， 1：卖出金额最大的前5名
	Buy       float64 `json:"buy"`        // 买入额（元）
	BuyRate   float64 `json:"buy_rate"`   // 买入占总成交比例
	Sell      float64 `json:"sell"`       // 卖出额（元）
	SellRate  float64 `json:"sell_rate"`  // 卖出占总成交比例
	NetBuy    float64 `json:"net_buy"`    // 净成交额（元）
	Reason    string  `json:"reason"`     // 上榜理由
}

type TopInstRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` //	交易日期
	TsCode    string `json:"ts_code"`    //	TS代码
}

type TopInstResponse struct {
	List []*TopInst `json:"list"`
}

func (x *TopInstResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 股权质押统计数据|pledge_stat
type PledgeStat struct {
	TsCode       string  `json:"ts_code"`       // TS代码
	EndDate      string  `json:"end_date"`      // 截止日期
	PledgeCount  int64   `json:"pledge_count"`  // 质押次数
	UnrestPledge float64 `json:"unrest_pledge"` // 无限售股质押数量（万）
	RestPledge   float64 `json:"rest_pledge"`   // 限售股份质押数量（万）
	TotalShare   float64 `json:"total_share"`   // 总股本
	PledgeRatio  float64 `json:"pledge_ratio"`  // 质押比例
}

type PledgeStatRequest struct {
	Limit   string `json:"limit"`
	Offset  string `json:"offset"`
	TsCode  string `json:"ts_code"`  // 股票代码
	EndDate string `json:"end_date"` // 截止日期
}

type PledgeStatResponse struct {
	List []*PledgeStat `json:"list"`
}

func (x *PledgeStatResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 股权质押明细|pledge_detail
type PledgeDetail struct {
	TsCode        string  `json:"ts_code"`        // TS股票代码
	AnnDate       string  `json:"ann_date"`       // 公告日期
	HolderName    string  `json:"holder_name"`    // 股东名称
	PledgeAmount  float64 `json:"pledge_amount"`  // 质押数量（万股）
	StartDate     string  `json:"start_date"`     // 质押开始日期
	EndDate       string  `json:"end_date"`       // 质押结束日期
	IsRelease     string  `json:"is_release"`     // 是否已解押
	ReleaseDate   string  `json:"release_date"`   // 解押日期
	Pledgor       string  `json:"pledgor"`        // 质押方
	HoldingAmount float64 `json:"holding_amount"` // 持股总数（万股）
	PledgedAmount float64 `json:"pledged_amount"` // 质押总数（万股）
	PTotalRatio   float64 `json:"p_total_ratio"`  // 本次质押占总股本比例
	HTotalRatio   float64 `json:"h_total_ratio"`  // 持股总数占总股本比例
	IsBuyback     string  `json:"is_buyback"`     // 是否回购
}

type PledgeDetailRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	TsCode string `json:"ts_code"` // 股票代码
}

type PledgeDetailResponse struct {
	List []*PledgeDetail `json:"list"`
}

func (x *PledgeDetailResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 股票回购|repurchase
type Repurchase struct {
	TsCode    string  `json:"ts_code"`    // TS代码
	AnnDate   string  `json:"ann_date"`   // 公告日期
	EndDate   string  `json:"end_date"`   // 截止日期
	Proc      string  `json:"proc"`       // 进度
	ExpDate   string  `json:"exp_date"`   // 过期日期
	Vol       float64 `json:"vol"`        // 回购数量
	Amount    float64 `json:"amount"`     // 回购金额
	HighLimit float64 `json:"high_limit"` // 回购最高价
	LowLimit  float64 `json:"low_limit"`  // 回购最低价
}

type RepurchaseRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	AnnDate   string `json:"ann_date"`   // 公告日期（任意填参数，如果都不填，单次默认返回2000条）
	StartDate string `json:"start_date"` // 公告开始日期
	EndDate   string `json:"end_date"`   // 公告结束日期
}

type RepurchaseResponse struct {
	List []*Repurchase `json:"list"`
}

func (x *RepurchaseResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 概念股分类|concept
type Concept struct {
	Code string `json:"code"` // 概念分类ID
	Name string `json:"name"` // 概念分类名称
	Src  string `json:"src"`  // 来源
}

type ConceptRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Src    string `json:"src"` // 来源，默认为ts
}

type ConceptResponse struct {
	List []*Concept `json:"list"`
}

func (x *ConceptResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 概念股列表|concept_detail
type ConceptDetail struct {
	Id          string `json:"id"`           // 概念代码
	ConceptName string `json:"concept_name"` // 概念名称
	TsCode      string `json:"ts_code"`      // 股票代码
	Name        string `json:"name"`         // 股票名称
	InDate      string `json:"in_date"`      // 纳入日期
	OutDate     string `json:"out_date"`     // 剔除日期
}

type ConceptDetailRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Id     string `json:"id"`      // 概念分类ID （id来自概念股分类接口）
	TsCode string `json:"ts_code"` // 股票代码 （以上参数二选一）
}

type ConceptDetailResponse struct {
	List []*ConceptDetail `json:"list"`
}

func (x *ConceptDetailResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 限售股解禁|share_float
type ShareFloat struct {
	TsCode     string  `json:"ts_code"`     // TS代码
	AnnDate    string  `json:"ann_date"`    // 公告日期
	FloatDate  string  `json:"float_date"`  // 解禁日期
	FloatShare float64 `json:"float_share"` // 流通股份(股)
	FloatRatio float64 `json:"float_ratio"` // 流通股份占总股本比率
	HolderName string  `json:"holder_name"` // 股东名称
	ShareType  string  `json:"share_type"`  // 股份类型
}

type ShareFloatRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS股票代码（至少输入一个参数）
	AnnDate   string `json:"ann_date"`   // 公告日期（日期格式：YYYYMMDD，下同）
	FloatDate string `json:"float_date"` // 解禁日期
	StartDate string `json:"start_date"` // 解禁开始日期
	EndDate   string `json:"end_date"`   // 解禁结束日期
}

type ShareFloatResponse struct {
	List []*ShareFloat `json:"list"`
}

func (x *ShareFloatResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 大宗交易|block_trade
type BlockTrade struct {
	TsCode    string  `json:"ts_code"`    // TS代码
	TradeDate string  `json:"trade_date"` // 交易日历
	Price     float64 `json:"price"`      // 成交价
	Vol       float64 `json:"vol"`        // 成交量（万股）
	Amount    float64 `json:"amount"`     // 成交金额
	Buyer     string  `json:"buyer"`      // 买方营业部
	Seller    string  `json:"seller"`     // 卖方营业部
}

type BlockTradeRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS代码（股票代码和日期至少输入一个参数）
	TradeDate string `json:"trade_date"` // 交易日期（格式：YYYYMMDD，下同）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type BlockTradeResponse struct {
	List []*BlockTrade `json:"list"`
}

func (x *BlockTradeResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 股东人数|stk_holdernumber
type StkHolderNumber struct {
	TsCode    string `json:"ts_code"`    // TS股票代码
	AnnDate   string `json:"ann_date"`   // 公告日期
	EndDate   string `json:"end_date"`   // 截止日期
	HolderNum int64  `json:"holder_num"` // 股东户数
}

type StkHolderNumberRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // TS股票代码
	Endate    string `json:"enddate"`    // 截止日期 (字段有冲突，删掉了个d)
	StartDate string `json:"start_date"` // 公告开始日期
	EndDate   string `json:"end_date"`   // 公告结束日期
}

type StkHolderNumberResponse struct {
	List []*StkHolderNumber `json:"list"`
}

func (x *StkHolderNumberResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 股东增减持|stk_holdertrade
type StkHolderTrade struct {
	TsCode      string  `json:"ts_code"`      // TS代码
	AnnDate     string  `json:"ann_date"`     // 公告日期
	HolderName  string  `json:"holder_name"`  // 股东名称
	HolderType  string  `json:"holder_type"`  // 股东类型G高管P个人C公司
	InDe        string  `json:"in_de"`        // 类型IN增持DE减持
	ChangeVol   float64 `json:"change_vol"`   // 变动数量
	ChangeRatio float64 `json:"change_ratio"` // 占流通比例（%）
	AfterShare  float64 `json:"after_share"`  // 变动后持股
	AfterRatio  float64 `json:"after_ratio"`  // 变动后占流通比例（%）
	AvgPrice    float64 `json:"avg_price"`    // 平均价格
	TotalShare  float64 `json:"total_share"`  // 持股总数
	BeginDate   string  `json:"begin_date"`   // 增减持开始日期
	CloseDate   string  `json:"close_date"`   // 增减持结束日期
}

type StkHolderTradeRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // TS股票代码
	AnnDate    string `json:"ann_date"`    // 公告日期
	StartDate  string `json:"start_date"`  // 公告开始日期
	EndDate    string `json:"end_date"`    // 公告结束日期
	TradeType  string `json:"trade_type"`  // 交易类型IN增持DE减持
	HolderType string `json:"holder_type"` // 股东类型C公司P个人G高管
}

type StkHolderTradeResponse struct {
	List []*StkHolderTrade `json:"list"`
}

func (x *StkHolderTradeResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 券商盈利预测数据|report_rc
type ReportRc struct {
	TsCode      string  `json:"ts_code"`      // 股票代码
	Name        string  `json:"name"`         // 股票名称
	ReportDate  string  `json:"report_date"`  // 研报日期
	ReportTitle string  `json:"report_title"` // 报告标题
	ReportType  string  `json:"report_type"`  // 报告类型
	Classify    string  `json:"classify"`     // 报告分类
	OrgName     string  `json:"org_name"`     // 机构名称
	AuthorName  string  `json:"author_name"`  // 作者
	Quarter     string  `json:"quarter"`      // 预测报告期
	OpRt        float64 `json:"op_rt"`        // 预测营业收入（万元）
	OpPr        float64 `json:"op_pr"`        // 预测营业利润（万元）
	Tp          float64 `json:"tp"`           // 预测利润总额（万元）
	Np          float64 `json:"np"`           // 预测净利润（万元）
	Eps         float64 `json:"eps"`          // 预测每股收益（元）
	Pe          float64 `json:"pe"`           // 预测市盈率
	Rd          float64 `json:"rd"`           // 预测股息率
	Roe         float64 `json:"roe"`          // 预测净资产收益率
	EvEbitda    float64 `json:"ev_ebitda"`    // 预测EV/EBITDA
	Rating      string  `json:"rating"`       // 卖方评级
	MaxPrice    float64 `json:"max_price"`    // 预测最高目标价
	MinPrice    float64 `json:"min_price"`    // 预测最低目标价
	ImpDg       string  `json:"imp_dg"`       // 机构关注度
	CreateTime  string  `json:"create_time"`  // TS数据更新时间
}

type ReportRcRequest struct {
	Limit      string `json:"limit"`
	Offset     string `json:"offset"`
	TsCode     string `json:"ts_code"`     // 股票代码
	ReportDate string `json:"report_date"` // 报告日期
	StartDate  string `json:"start_date"`  // 报告开始日期
	EndDate    string `json:"end_date"`    // 报告结束日期
}

type ReportRcResponse struct {
	List []*ReportRc `json:"list"`
}

func (x *ReportRcResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 每日筹码及胜率|cyq_perf
type CyqPerf struct {
	TsCode     string  `json:"ts_code"`     // 股票代码
	TradeDate  string  `json:"trade_date"`  // 交易日期
	HisLow     float64 `json:"his_low"`     // 历史最低价
	HisHigh    float64 `json:"his_high"`    // 历史最高价
	Cost_5Pct  float64 `json:"cost_5pct"`   // 5分位成本
	Cost_15Pct float64 `json:"cost_15pct"`  // 15分位成本
	Cost_50Pct float64 `json:"cost_50pct"`  // 50分位成本
	Cost_85Pct float64 `json:"cost_85pct"`  // 85分位成本
	Cost_95Pct float64 `json:"cost_95pct"`  // 95分位成本
	WeightAvg  float64 `json:"weight_avg"`  // 加权平均成本
	WinnerRate float64 `json:"winner_rate"` // 胜率
}

type CyqPerfRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type CyqPerfResponse struct {
	List []*CyqPerf `json:"list"`
}

func (x *CyqPerfResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 每日筹码分布|cyq_chips
type CyqChips struct {
	TsCode    string  `json:"ts_code"`    // 股票代码
	TradeDate string  `json:"trade_date"` // 交易日期
	Price     float64 `json:"price"`      // 成本价格
	Percent   float64 `json:"percent"`    // 价格占比（%）
}

type CyqChipsRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期（YYYYMMDD）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type CyqChipsResponse struct {
	List []*CyqChips `json:"list"`
}

func (x *CyqChipsResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 股票技术因子|stk_factor
type StkFactor struct {
	TsCode      string  `json:"ts_code"`       // 股票代码
	TradeDate   string  `json:"trade_date"`    // 交易日期
	Close       float64 `json:"close"`         // 收盘价
	Open        float64 `json:"open"`          // 开盘价
	High        float64 `json:"high"`          // 最高价
	Low         float64 `json:"low"`           // 最低价
	PreClose    float64 `json:"pre_close"`     // 昨收价
	Change      float64 `json:"change"`        // 涨跌额
	PctChange   float64 `json:"pct_change"`    // 涨跌幅
	Vol         float64 `json:"vol"`           // 成交量 （手）
	Amount      float64 `json:"amount"`        // 成交额 （千元）
	AdjFactor   float64 `json:"adj_factor"`    // 复权因子
	OpenHfq     float64 `json:"open_hfq"`      // 开盘价后复权
	OpenQfq     float64 `json:"open_qfq"`      // 开盘价前复权
	CloseHfq    float64 `json:"close_hfq"`     // 收盘价后复权
	CloseQfq    float64 `json:"close_qfq"`     // 收盘价前复权
	HighHfq     float64 `json:"high_hfq"`      // 最高价后复权
	HighQfq     float64 `json:"high_qfq"`      // 最高价前复权
	LowHfq      float64 `json:"low_hfq"`       // 最低价后复权
	LowQfq      float64 `json:"low_qfq"`       // 最低价前复权
	PreCloseHfq float64 `json:"pre_close_hfq"` // 昨收价后复权
	PreCloseQfq float64 `json:"pre_close_qfq"` // 昨收价前复权
	MacdDif     float64 `json:"macd_dif"`      // MCAD_DIF (基于前复权价格计算，下同)
	MacdDea     float64 `json:"macd_dea"`      // MCAD_DEA
	Macd        float64 `json:"macd"`          // MCAD
	KdjK        float64 `json:"kdj_k"`         // KDJ_K
	KdjD        float64 `json:"kdj_d"`         // KDJ_D
	KdjJ        float64 `json:"kdj_j"`         // KDJ_J
	Rsi_6       float64 `json:"rsi_6"`         // RSI_6
	Rsi_12      float64 `json:"rsi_12"`        // RSI_12
	Rsi_24      float64 `json:"rsi_24"`        // RSI_24
	BollUpper   float64 `json:"boll_upper"`    // BOLL_UPPER
	BollMid     float64 `json:"boll_mid"`      // BOLL_MID
	BollLower   float64 `json:"boll_lower"`    // BOLL_LOWER
	Cci         float64 `json:"cci"`           // CCI
}

type StkFactorRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期 （yyyymmdd，下同）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type StkFactorResponse struct {
	List []*StkFactor `json:"list"`
}

func (x *StkFactorResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 中央结算系统持股统计|ccass_hold
type CcassHold struct {
	TradeDate    string `json:"trade_date"`   // 交易日期
	TsCode       string `json:"ts_code"`      // 股票代号
	Name         string `json:"name"`         // 股票名称
	Shareholding string `json:"shareholding"` // 于中央结算系统的持股量(股)
	HoldNums     string `json:"hold_nums"`    // 参与者数目（个）
	HoldRatio    string `json:"hold_ratio"`   // 占于上交所上市及交易的A股总数的百分比（%）
}

type CcassHoldRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码 (e.g. 605009.SH)
	HkCode    string `json:"hk_code"`    // 港交所代码 （e.g. 95009）
	TradeDate string `json:"trade_date"` // 交易日期(YYYYMMDD格式，下同)
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type CcassHoldResponse struct {
	List []*CcassHold `json:"list"`
}

func (x *CcassHoldResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 中央结算系统持股明细|ccass_hold_detail
type CcasHoldDetail struct {
	TradeDate              string `json:"trade_date"`               // 交易日期
	TsCode                 string `json:"ts_code"`                  // 股票代号
	Name                   string `json:"name"`                     // 股票名称
	ColParticipantId       string `json:"col_participant_id"`       // 参与者编号
	ColParticipantName     string `json:"col_participant_name"`     // 机构名称
	ColShareholding        string `json:"col_shareholding"`         // 持股量(股)
	ColShareholdingPercent string `json:"col_shareholding_percent"` // 占已发行股份/权证/单位百分比(%)
}

type CcasHoldDetailRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码 (e.g. 605009.SH)
	HkCode    string `json:"hk_code"`    // 交所代码 （e.g. 95009）
	TradeDate string `json:"trade_date"` // 交易日期(YYYYMMDD格式，下同)
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type CcasHoldDetailResponse struct {
	List []*CcasHoldDetail `json:"list"`
}

func (x *CcasHoldDetailResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 沪深港股通持股明细|hk_hold
type HkHold struct {
	Code      string  `json:"code"`       // 原始代码
	TradeDate string  `json:"trade_date"` // 交易日期
	TsCode    string  `json:"ts_code"`    // TS代码
	Name      string  `json:"name"`       // 股票名称
	Vol       int64   `json:"vol"`        // 持股数量(股)
	Ratio     float64 `json:"ratio"`      // 持股占比（%），占已发行股份百分比
	Exchange  string  `json:"exchange"`   // 类型：SH沪股通SZ深股通HK港股通
}

type HkHoldRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	Code      string `json:"code"`       // 交易所代码
	TsCode    string `json:"ts_code"`    // TS股票代码
	TradeDate string `json:"trade_date"` // 交易日期
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
	Exchange  string `json:"exchange"`   // 类型：SH沪股通（北向）SZ深股通（北向）HK港股通（南向持股）
}

type HkHoldResponse struct {
	List []*HkHold `json:"list"`
}

func (x *HkHoldResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 涨跌停和炸板数据|limit_list_d
type LimitListd struct {
	TradeDate     string  `json:"trade_date"`     // 交易日期
	TsCode        string  `json:"ts_code"`        // 股票代码
	Industry      string  `json:"industry"`       // 所属行业
	Name          string  `json:"name"`           // 股票名称
	Close         float64 `json:"close"`          // 收盘价
	PctChg        float64 `json:"pct_chg"`        // 涨跌幅
	Amount        float64 `json:"amount"`         // 成交额
	LimitAmount   float64 `json:"limit_amount"`   // 板上成交金额(涨停无此数据)
	FloatMv       float64 `json:"float_mv"`       // 流通市值
	TotalMv       float64 `json:"total_mv"`       // 总市值
	TurnoverRatio float64 `json:"turnover_ratio"` // 换手率
	FdAmount      float64 `json:"fd_amount"`      // 封单金额
	FirstTime     string  `json:"first_time"`     // 首次封板时间（跌停无此数据）
	LastTime      string  `json:"last_time"`      // 最后封板时间
	OpenTimes     int64   `json:"open_times"`     // 炸板次数(跌停为开板次数)
	UpStat        string  `json:"up_stat"`        // 涨停统计（N/T T天有N次涨停）
	LimitTimes    int64   `json:"limit_times"`    // 连板数
	Limit         string  `json:"limit"`          // D跌停U涨停Z炸板
}

type LimitListdRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期
	TsCode    string `json:"ts_code"`    // 股票代码
	LimitType string `json:"limit_type"` // 涨跌停类型（U涨停D跌停Z炸板）
	Exchange  string `json:"exchange"`   // 交易所（SH上交所SZ深交所BJ北交所）
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

type LimitListdResponse struct {
	List []*LimitListd `json:"list"`
}

func (x *LimitListdResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 机构调研数据|stk_surv
type StkSurv struct {
	TsCode       string `json:"ts_code"`       // 股票代码
	Name         string `json:"name"`          // 股票名称
	SurvDate     string `json:"surv_date"`     // 调研日期
	FundVisitors string `json:"fund_visitors"` // 机构参与人员
	RecePlace    string `json:"rece_place"`    // 接待地点
	ReceMode     string `json:"rece_mode"`     // 接待方式
	ReceOrg      string `json:"rece_org"`      // 接待的公司
	OrgType      string `json:"org_type"`      // 接待公司类型
	CompRece     string `json:"comp_rece"`     // 上市公司接待人员
	Content      string `json:"content"`       // 调研内容
}

type StkSurvRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TsCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 调研日期
	StartDate string `json:"start_date"` // 调研开始日期
	EndDate   string `json:"end_date"`   // 调研结束日期
}

type StkSurvResponse struct {
	List []*StkSurv `json:"list"`
}

func (x *StkSurvResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 券商月度金股|broker_recommend
type BrokerRecommend struct {
	Month  string `json:"month"`   // 月度
	Broker string `json:"broker"`  // 券商
	TsCode string `json:"ts_code"` // 股票代码
	Name   string `json:"name"`    // 股票简称
}

type BrokerRecommendRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Month  string `json:"month"` // 月度（YYYYMM）
}

type BrokerRecommendResponse struct {
	List []*BrokerRecommend `json:"list"`
}

func (x *BrokerRecommendResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 游资名录|hm_list
type HmList struct {
	Name string `json:"name"` // 游资名称
	Desc string `json:"desc"` // 说明
	Orgs string `json:"orgs"` // 关联机构
}

type HmListRequest struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Name   string `json:"name"` // 游资名称
}

type HmListResponse struct {
	List []*HmList `json:"list"`
}

func (x *HmListResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 游资每日明细|hm_detail
type HmDetail struct {
	TradeDate  string  `json:"trade_date"`  // 交易日期
	TsCode     string  `json:"ts_code"`     // 股票代码
	TsName     string  `json:"ts_name"`     // 股票名称
	BuyAmount  float64 `json:"buy_amount"`  // 买入金额（元）
	SellAmount float64 `json:"sell_amount"` // 卖出金额（元）
	NetAmount  float64 `json:"net_amount"`  // 净买卖（元）
	HmName     string  `json:"hm_name"`     // 游资名称
	HmOrgs     string  `json:"hm_orgs"`     // 关联机构（一般为营业部或机构专用）
	Tag        string  `json:"tag"`         // 标签
}

type HmDetailRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期(YYYYMMDD)
	TsCode    string `json:"ts_code"`    // 股票代码
	HmName    string `json:"hm_name"`    // 游资名称
	StartDate string `json:"start_date"` // 开始日期(YYYYMMDD)
	EndDate   string `json:"end_date"`   // 结束日期(YYYYMMDD)
}

type HmDetailResponse struct {
	List []*HmDetail `json:"list"`
}

func (x *HmDetailResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 同花顺App热榜|ths_hot
type ThsHot struct {
	TradeDate    string  `json:"trade_date"`    // 交易日期
	DataType     string  `json:"data_type"`     // 数据类型
	TsCode       string  `json:"ts_code"`       // 股票代码
	TsName       string  `json:"ts_name"`       // 股票名称
	Rank         int64   `json:"rank"`          // 排行
	PctChange    float64 `json:"pct_change"`    // 涨跌幅%
	CurrentPrice float64 `json:"current_price"` // 当前价格
	Concept      string  `json:"concept"`       // 标签
	RankReason   string  `json:"rank_reason"`   // 上榜解读
	Hot          float64 `json:"hot"`           // 热度值
	RankTime     string  `json:"rank_time"`     // 排行榜获取时间
}

type ThsHotRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期
	TsCode    string `json:"ts_code"`    // TS代码
	Market    string `json:"market"`     // 热榜类型(热股、ETF、可转债、行业板块、概念板块、期货、港股、热基、美股)
	IsNew     string `json:"is_new"`     // 是否最新（默认Y，如果为N则为盘中和盘后阶段采集，具体时间可参考rank_time字段）
}

type ThsHotResponse struct {
	List []*ThsHot `json:"list"`
}

func (x *ThsHotResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}

// 东方财富App热榜|dc_hot
type DcHot struct {
	TradeDate    string  `json:"trade_date"`    // 交易日期
	DataType     string  `json:"data_type"`     // 数据类型
	TsCode       string  `json:"ts_code"`       // 股票代码
	TsName       string  `json:"ts_name"`       // 股票名称
	Rank         int64   `json:"rank"`          // 排行或者热度
	PctChange    float64 `json:"pct_change"`    // 涨跌幅%
	CurrentPrice float64 `json:"current_price"` // 当前价
	RankTime     string  `json:"rank_time"`     // 排行榜获取时间
}

type DcHotRequest struct {
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	TradeDate string `json:"trade_date"` // 交易日期
	TsCode    string `json:"ts_code"`    // TS代码
	Market    string `json:"market"`     // 类型(A股市场、ETF基金、港股市场、美股市场)
	HotType   string `json:"hot_type"`   // 热点类型(人气榜、飙升榜)
	IsNew     string `json:"is_new"`     // 是否最新（默认Y，如果为N则为盘中和盘后阶段采集，具体时间可参考rank_time字段）
}

type DcHotResponse struct {
	List []*DcHot `json:"list"`
}

func (x *DcHotResponse) String() string {
	bytes, _ := json.Marshal(x)
	return string(bytes)
}
