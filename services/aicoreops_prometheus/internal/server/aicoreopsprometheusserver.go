// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_prometheus.proto

package server

import (
	"context"

	"aicoreops_prometheus/aicoreops_prometheus"
	"aicoreops_prometheus/internal/logic"
	"aicoreops_prometheus/internal/svc"
)

type AicoreopsPrometheusServer struct {
	svcCtx *svc.ServiceContext
	aicoreops_prometheus.UnimplementedAicoreopsPrometheusServer
}

func NewAicoreopsPrometheusServer(svcCtx *svc.ServiceContext) *AicoreopsPrometheusServer {
	return &AicoreopsPrometheusServer{
		svcCtx: svcCtx,
	}
}

func (s *AicoreopsPrometheusServer) Ping(ctx context.Context, in *aicoreops_prometheus.Request) (*aicoreops_prometheus.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}