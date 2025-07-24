package go_buy365

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-buy365/utils"
	jsoniter "github.com/json-iterator/go"
)

func (cli *Client) GetOrderList() (*Buy365OrderListRsp, error) {

	rawURL := cli.Params.OrderListUrl

	params := map[string]interface{}{
		"sys_no": cli.Params.MerchantId,
	}

	//签名
	signStr := utils.SignWithdraw(params, cli.Params.AccessKey)
	params["sign"] = signStr

	//返回值会放到这里
	var result Buy365OrderListRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetMultipartFormData(utils.ConvertToStringMap(params)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#buy365#orderlist->%+v", string(restLog))

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
