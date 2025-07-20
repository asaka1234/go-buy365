package go_buy365

import (
	"fmt"
	"testing"
)

func TestWithdrawConfirm(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &Buy365InitParams{MERCHANT_ID, ACCESS_KEY, BACK_KEY, IP, DEPOSIT_URL, WITHDRAW_URL, WITHDRAW_CONFIRM_URL, ORDERLIST_URL})

	//发请求
	resp, err := cli.WithdrawConfirm(GenWithdrawConfirmRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawConfirmRequestDemo() Buy365WithdrawConfirmReq {
	return Buy365WithdrawConfirmReq{
		Ids: "1", //,2,3",
	}
}
