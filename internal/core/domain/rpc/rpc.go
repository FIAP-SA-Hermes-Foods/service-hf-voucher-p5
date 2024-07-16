package rpc

import "service-hf-voucher-p5/internal/core/domain/entity/dto"

type VoucherRPC interface {
	GetVoucherByID(uuid string) (*dto.OutputVoucher, error)
	SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
	UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
}

type VoucherWorkerRPC interface {
	GetVoucherByID(uuid string) (*dto.OutputVoucher, error)
	SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
	UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
}
