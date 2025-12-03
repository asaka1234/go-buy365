package go_buy365

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-buy365/utils"
	jsoniter "github.com/json-iterator/go"
)

// withdraw
func (cli *Client) Withdraw(req Buy365WithdrawReq) (*Buy365WithdrawResponse, error) {

	rawURL := cli.Params.WithdrawUrl

	jsonData, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	params := make(map[string]interface{})
	params["data"] = string(jsonData)
	params["sys_no"] = cli.Params.MerchantId

	//签名
	signStr := utils.SignWithdraw(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result Buy365WithdrawResponse

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetBody(params).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#buy365#withdraw->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, err
}
