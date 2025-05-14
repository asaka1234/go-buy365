package go_buy365

import (
	"crypto/tls"
	"github.com/asaka1234/go-buy365/utils"
	"github.com/fatih/structs"
)

func (cli *Client) GetOrderList(req Buy365WithdrawReq) (*Buy365OrderListRsp, error) {

	rawURL := cli.OrderListURL

	//返回值会放到这里
	var result Buy365OrderListRsp

	//构造请求(加签名)
	paramMap := structs.Map(req)
	signStr := utils.SignDeposit(paramMap, cli.AccessKey)
	paramMap["sign"] = signStr

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(paramMap).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}
