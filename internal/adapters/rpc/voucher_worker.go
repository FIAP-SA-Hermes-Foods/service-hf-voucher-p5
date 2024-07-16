package rpc

import (
	"context"
	"fmt"
	"service-hf-voucher-p5/internal/core/domain/entity/dto"
	"service-hf-voucher-p5/internal/core/domain/rpc"
	op "service-hf-voucher-p5/voucher_api_proto"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ rpc.VoucherWorkerRPC = (*voucherWorkerRPC)(nil)

type voucherWorkerRPC struct {
	ctx  context.Context
	host string
	port string
}

func NewVoucherWorkerRPC(ctx context.Context, host, port string) rpc.VoucherWorkerRPC {
	return voucherWorkerRPC{ctx: ctx, host: host, port: port}
}

func (p voucherWorkerRPC) GetVoucherByID(uuid string) (*dto.OutputVoucher, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	input := op.GetVoucherByIDRequest{
		Uuid: uuid,
	}

	cc := op.NewVoucherClient(conn)

	resp, err := cc.GetVoucherByID(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	out := &dto.OutputVoucher{
		UUID:       resp.Uuid,
		Code:       resp.Code,
		Percentage: strconv.FormatInt(int64(resp.Percentage), 10),
		CreatedAt:  resp.CreatedAt,
		ExpiresAt:  resp.ExpiresAt,
	}

	return out, nil
}

func (p voucherWorkerRPC) SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	percentage, err := strconv.ParseInt(voucher.Percentage, 10, 64)
	if err != nil {
		return nil, err
	}
	
	input := op.CreateVoucherRequest{
		Code:       voucher.Code,
		Percentage: percentage,
		ExpiresAt:  voucher.ExpiresAt,
	}

	cc := op.NewVoucherClient(conn)

	resp, err := cc.CreateVoucher(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputVoucher{
		UUID:       resp.Uuid,
		Code:       resp.Code,
		Percentage: strconv.FormatInt(int64(resp.Percentage), 10),
		CreatedAt:  resp.CreatedAt,
		ExpiresAt:  resp.ExpiresAt,
	}

	return &out, nil
}

func (p voucherWorkerRPC) UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", p.host, p.port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	percentage, err := strconv.ParseInt(voucher.Percentage, 10, 64)
	if err != nil {
		return nil, err
	}

	input := op.UpdateVoucherByIDRequest{
		Uuid:       voucher.UUID,
		Code:       voucher.Code,
		Percentage: percentage,
		CreatedAt:  voucher.CreatedAt,
		ExpiresAt:  voucher.ExpiresAt,
	}

	cc := op.NewVoucherClient(conn)

	resp, err := cc.UpdateVoucherByID(p.ctx, &input)

	if err != nil {
		return nil, err
	}

	var out = dto.OutputVoucher{
		UUID:       resp.Uuid,
		Code:       resp.Code,
		Percentage: strconv.FormatInt(int64(resp.Percentage), 10),
		CreatedAt:  resp.CreatedAt,
		ExpiresAt:  resp.ExpiresAt,
	}

	return &out, nil
}
