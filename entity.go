package go_buy365

// ----------pre order-------------------------

type Buy365DepositReq struct {
	OrderId     string `json:"order_id"`      // 订单ID
	OrderAmount string `json:"order_amount"`  // 订单金额
	SysNo       string `json:"sys_no"`        // 系统编号
	UserId      string `json:"user_id"`       // 用户ID
	OrderIp     string `json:"order_ip"`      // 订单IP
	OrderTime   string `json:"order_time"`    // 订单时间
	PayUserName string `json:"pay_user_name"` // 付款人姓名
}

//------------------------------------------------------------

type Buy365DepositResponse struct {
	Code   int    `json:"code"`   // 111 是正确
	Status string `json:"status"` //success 是正确
	Msg    string `json:"msg"`
	Data   struct {
		OrderNo string `json:"order_no"` // 订单编号
		SendUrl string `json:"send_url"` // 发送URL
		UserId  string `json:"user_id"`  // 用户ID
	} `json:"data"`
}

// ----------withdraw-------------------------

type Buy365WithdrawReq struct {
	SysNo string `json:"sys_no"` // 申请sys_no唯一标识 610001
}

type Buy365WithdrawData struct {
	UserName    string `json:"user_name"`    // 真实姓名
	BankCardNo  string `json:"bankcard_no"`  // 卡号
	SerialNo    string `json:"serial_no"`    // 订单号
	BankAddress string `json:"bank_address"` // 支行地址
	Amount      string `json:"amount"`       // 金额
}

type Buy365WithdrawResponse struct {
	Code int    `json:"code"` //
	Msg  string `json:"msg"`
}

// ----------withdraw confirm-------------------------

type Buy365WithdrawConfirmReq struct {
	Ids   string `json:"ids"`    //确认收款订单列表接口中获取的id，用英文逗号“,”拼接起来
	SysNo string `json:"sys_no"` //申请sys_no唯一标识
}

type Buy365WithdrawConfirmResponse struct {
	Code string `json:"code"` //
	Msg  string `json:"msg"`
}

// ----------withdraw confirm-------------------------

type Buy365OrderListReq struct {
	SysNo string `json:"sys_no"` //申请sys_no唯一标识
}

type Buy365OrderListRsp struct {
	Code   string                `json:"code"` //
	Msg    string                `json:"msg"`
	Result Buy365OrderPageResult `json:"result"`
}

type Buy365OrderPageResult struct {
	TotalCount string             `json:"totalCount"` // 总记录数
	TotalPage  int64              `json:"totalPage"`  // 总页数
	Page       int64              `json:"page"`       // 当前页码
	Data       []*Buy365OrderData `json:"data"`       // 订单数据列表
}

type Buy365OrderData struct {
	ID                       string `json:"id"`
	SysSerialNo              string `json:"sysSerialNo"`
	Amount                   string `json:"amount"`
	PayType                  string `json:"payType"`
	UserName                 string `json:"userName"`
	BankCardNo               string `json:"bankCardNo"`
	BankAddress              string `json:"bankAddress"`
	ChangeRate               string `json:"changeRate"`
	HandlingFee              string `json:"handlingFee"`
	MerchantSettleUSDTNumber string `json:"merchantSettleUSDTNumber"`
	SerialNo                 string `json:"serialNo"`
	CreateTime               string `json:"createTime"`
	Remark                   string `json:"remark"`
	NumRow                   string `json:"numRow"`
	StatusName               string `json:"statusName"`
}
