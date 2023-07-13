package controller

import "context"

type Interface interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type Service interface {
}

type Config interface {
	BindAddr() string
}
