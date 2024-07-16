package useCase

import "service-hf-voucher-p5/internal/core/domain/entity/dto"

type VoucherUseCase interface {
	SaveVoucher(reqVoucher dto.RequestVoucher) error
	GetVoucherByID(uuid string) error
	UpdateVoucherByID(uuid string, voucher dto.RequestVoucher) error
}
