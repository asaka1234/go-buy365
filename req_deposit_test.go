package go_buy365

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &Buy365InitParams{MERCHANT_ID, ACCESS_KEY, BACK_KEY, IP, DEPOSIT_URL, WITHDRAW_URL, WITHDRAW_CONFIRM_URL, ORDERLIST_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() Buy365DepositReq {
	return Buy365DepositReq{
		OrderId:     "32342", //商户uid
		UserId:      "30779639363",
		OrderIp:     "18.29.120.32",
		OrderAmount: "30.00",
		PayUserName: "你好", //商户订单号
	}
}
