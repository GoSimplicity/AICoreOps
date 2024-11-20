// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_tree.proto

package server

import (
	"context"

	"aicoreops_tree/aicoreops_tree"
	"aicoreops_tree/internal/logic"
	"aicoreops_tree/internal/svc"
)

type AicoreopsTreeServer struct {
	svcCtx *svc.ServiceContext
	aicoreops_tree.UnimplementedAicoreopsTreeServer
}

func NewAicoreopsTreeServer(svcCtx *svc.ServiceContext) *AicoreopsTreeServer {
	return &AicoreopsTreeServer{
		svcCtx: svcCtx,
	}
}

func (s *AicoreopsTreeServer) Ping(ctx context.Context, in *aicoreops_tree.Request) (*aicoreops_tree.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
