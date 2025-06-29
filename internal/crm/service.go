package crm

import "github.com/go-resty/resty/v2"

type service struct {
	baseUrl string
	client  *resty.Client
}

type Interface interface {
}

func NewCrmService(baseUrl string) Interface {
	return &service{
		baseUrl: baseUrl,
		client:  resty.New(),
	}
}
