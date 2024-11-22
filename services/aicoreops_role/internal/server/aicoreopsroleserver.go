// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_role.proto

package server

import (
	"context"

	"aicoreops_role/aicoreops_role"
	"aicoreops_role/internal/logic"
	"aicoreops_role/internal/svc"
)

type AicoreopsRoleServer struct {
	svcCtx *svc.ServiceContext
	aicoreops_role.UnimplementedAicoreopsRoleServer
}

func NewAicoreopsRoleServer(svcCtx *svc.ServiceContext) *AicoreopsRoleServer {
	return &AicoreopsRoleServer{
		svcCtx: svcCtx,
	}
}

func (s *AicoreopsRoleServer) Ping(ctx context.Context, in *aicoreops_role.Request) (*aicoreops_role.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
