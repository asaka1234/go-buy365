package go_buy365

import (
	"fmt"
	"testing"
)

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &Buy365InitParams{MERCHANT_ID, ACCESS_KEY, BACK_KEY, IP, DEPOSIT_URL, WITHDRAW_URL, WITHDRAW_CONFIRM_URL, ORDERLIST_URL})

	req := Buy365DepositSucceedBackReq{
		BillNo:     "1946948839491506176",
		Amount:     "7",
		AmountUsdt: "0.9695",
		SysNo:      "502326",
		Sign:       "6a3d91a877d1ab801562704db552b8cb",
	}
	//发请求
	err := cli.DepositSucceedCallBack(req, processor)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
}

func processor(Buy365DepositSucceedBackReq) error {
	return nil
}
