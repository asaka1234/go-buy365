package go_buy365

import "github.com/go-resty/resty/v2"

type Client struct {
	AccessKey string
	BackKey   string

	BaseURL            string
	WithdrawURL        string
	WithdrawConfirmURL string
	OrderListURL       string

	ryClient *resty.Client
}

func NewClient(accessKey string, backKey string, baseURL string, withdrawURL, withdrawConfirmURL, orderListURL string) *Client {
	return &Client{
		AccessKey: accessKey,
		BackKey:   backKey,

		BaseURL:            baseURL,
		WithdrawURL:        withdrawURL,
		WithdrawConfirmURL: withdrawConfirmURL,
		OrderListURL:       orderListURL,

		ryClient: resty.New(), //client实例
	}
}
