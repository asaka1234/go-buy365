package go_buy365

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &Buy365InitParams{MERCHANT_ID, ACCESS_KEY, BACK_KEY, IP, DEPOSIT_URL, WITHDRAW_URL, WITHDRAW_CONFIRM_URL, ORDERLIST_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() Buy365WithdrawReq {
	return Buy365WithdrawReq{
		Data: []Buy365WithdrawData{
			{
				UserName:    "你好", //商户uid
				BankCardNo:  "30787",
				SerialNo:    "129090",
				BankAddress: "具体地址",
				Amount:      "15.00", //商户订单号
			},
		},
	}
}
