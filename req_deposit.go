package go_buy365

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-buy365/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"time"
)

// pre-order
func (cli *Client) Deposit(req Buy365DepositReq) (*Buy365DepositResponse, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["sys_no"] = cli.Params.MerchantId
	params["order_time"] = time.Now().Format("2006-01-02 15:04:05")

	//签名
	signStr := utils.SignDeposit(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result Buy365DepositCommonResponse

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
	cli.logger.Infof("PSPResty#buy365#deposit->%+v", string(restLog))

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

	//------------------------------------------------------
	if result.Code == 111 && result.Status == "success" {
		//说明成功

		//step-1
		var data map[string]interface{}
		if err := json.Unmarshal(resp2.Body(), &data); err != nil {
			return nil, err
		}

		//step-2
		var resp3 Buy365DepositResponse
		if err := mapstructure.Decode(data, &resp3); err != nil {
			return nil, err
		}

		return &resp3, nil
	}

	return &Buy365DepositResponse{
		Code:   result.Code,
		Status: result.Status,
		Msg:    result.Msg,
	}, nil
}
