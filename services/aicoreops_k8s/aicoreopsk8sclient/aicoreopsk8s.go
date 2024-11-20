// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_k8s.proto

package aicoreopsk8sclient

import (
	"context"

	"aicoreops_k8s/aicoreops_k8s"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = aicoreops_k8s.Request
	Response = aicoreops_k8s.Response

	AicoreopsK8s interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAicoreopsK8s struct {
		cli zrpc.Client
	}
)

func NewAicoreopsK8s(cli zrpc.Client) AicoreopsK8s {
	return &defaultAicoreopsK8s{
		cli: cli,
	}
}

func (m *defaultAicoreopsK8s) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := aicoreops_k8s.NewAicoreopsK8SClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
