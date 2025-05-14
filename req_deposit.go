package go_buy365

import (
	"crypto/tls"
	"github.com/asaka1234/go-buy365/utils"
	"github.com/fatih/structs"
)

// pre-order
func (cli *Client) Deposit(req Buy365DepositReq) (*Buy365DepositResponse, error) {

	rawURL := cli.BaseURL

	//返回值会放到这里
	var result Buy365DepositResponse

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
