package useCase

import (
	"errors"
	"service-hf-voucher-p5/internal/core/domain/entity/dto"
	"service-hf-voucher-p5/internal/core/domain/useCase"
)

var _ useCase.VoucherUseCase = (*voucherUseCase)(nil)

type voucherUseCase struct {
}

func NewVoucherUseCase() voucherUseCase {
	return voucherUseCase{}
}

func (p voucherUseCase) SaveVoucher(reqVoucher dto.RequestVoucher) error {
	voucher := reqVoucher.Voucher()

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return err
	}

	if len(voucher.Code) == 0 {
		return errors.New("the voucher code is null or not valid")
	}

	if voucher.Percentage < 0 || voucher.Percentage > 101 {
		return errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return nil
}

func (p voucherUseCase) GetVoucherByID(uuid string) error {
	if len(uuid) == 0 {
		return errors.New("the id is not valid for consult")
	}

	return nil
}

func (p voucherUseCase) UpdateVoucherByID(uuid string, reqVoucher dto.RequestVoucher) error {
	if len(uuid) == 0 {
		return errors.New("the id is not valid for consult")
	}

	voucher := reqVoucher.Voucher()

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return err
	}

	if len(voucher.Code) == 0 {
		return errors.New("the voucher code is null or not valid")
	}

	if voucher.Percentage < 0 || voucher.Percentage > 101 {
		return errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return nil
}
