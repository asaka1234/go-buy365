package go_buy365

import (
	"crypto/tls"
	"github.com/asaka1234/go-buy365/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// pre-order
func (cli *Client) Deposit(req Buy365DepositReq) (*Buy365DepositResponse, error) {

	rawURL := cli.DepositURL

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["sys_no"] = cli.MerchantID
	params["order_time"] = time.Now().Format("2006-01-02 15:04:05")

	//签名
	signStr := utils.SignDeposit(params, cli.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result Buy365DepositResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}
