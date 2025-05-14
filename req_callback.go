package go_buy365

import (
	"errors"
	"github.com/asaka1234/go-buy365/utils"
	"github.com/fatih/structs"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCancelCallback(req Buy365DepositCancelBackReq, processor func(Buy365DepositCancelBackReq) error) error {
	//验证签名
	paramMap := structs.Map(req)
	verifyResult := utils.VerifySignDeposit(paramMap, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}

// 充值的回调处理(传入一个处理函数)
func (cli *Client) depositSucceedBack(req Buy365DepositSucceedBackReq, processor func(Buy365DepositSucceedBackReq) error) error {
	//验证签名
	paramMap := structs.Map(req)
	verifyResult := utils.VerifySignDeposit(paramMap, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}

// 充值的回调处理(传入一个处理函数)
func (cli *Client) withdrawCancelBack(req Buy365WithdrawCancelBackReq, processor func(Buy365WithdrawCancelBackReq) error) error {
	//验证签名
	paramMap := structs.Map(req)
	verifyResult := utils.VerifySignDeposit(paramMap, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}

// 充值的回调处理(传入一个处理函数)
func (cli *Client) withdrawSucceedBack(req Buy365WithdrawSucceedBackReq, processor func(Buy365WithdrawSucceedBackReq) error) error {
	//验证签名
	paramMap := structs.Map(req)
	verifyResult := utils.VerifySignDeposit(paramMap, cli.BackKey)
	if !verifyResult {
		//验签失败
		return errors.New("verify sign error!")
	}

	//开始处理
	return processor(req)
}
