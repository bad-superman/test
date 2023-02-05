package eastmoney

import (
	"net/http"
	"time"
)

type EastmoneyClient struct {
	client http.Client
}

func NewEastmoneyClient() *EastmoneyClient {
	return &EastmoneyClient{
		client: http.Client{
			Timeout: 2 * time.Second,
		},
	}
}
