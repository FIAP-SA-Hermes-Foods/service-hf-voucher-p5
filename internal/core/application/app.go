package application

import (
	"context"
	"errors"
	l "service-hf-voucher-p5/external/logger"
	ps "service-hf-voucher-p5/external/strings"
	"service-hf-voucher-p5/internal/core/domain/entity/dto"
	"service-hf-voucher-p5/internal/core/domain/rpc"
)

type Application interface {
	GetVoucherByID(msgID string, uuid string) (*dto.OutputVoucher, error)
	SaveVoucher(msgID string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
	UpdateVoucherByID(msgID string, id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
}

type application struct {
	ctx              context.Context
	voucherRPC       rpc.VoucherRPC
	voucherWorkerRPC rpc.VoucherWorkerRPC
}

func (app application) setMessageIDCtx(msgID string) {
	// implementation of setMessageIDCtx
}

func NewApplication(ctx context.Context, voucherRPC rpc.VoucherRPC, voucherWorkerRPC rpc.VoucherWorkerRPC) Application {
	return application{
		ctx:              ctx,
		voucherRPC:       voucherRPC,
		voucherWorkerRPC: voucherWorkerRPC,
	}
}

func (app application) GetVoucherByID(msgID string, uuid string) (*dto.OutputVoucher, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "GetVoucherByIDApp: ", " | ", uuid)

	go app.voucherRPC.GetVoucherByID(uuid) // pub

	voucherRpc, err := app.voucherWorkerRPC.GetVoucherByID(uuid)	

	if err != nil {
		l.Errorf(msgID, "GetVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	if voucherRpc == nil {
		l.Infof(msgID, "GetVoucherByIDApp output: ", " | ", nil)
		return nil, nil
	}

	voucher := &dto.OutputVoucher{
		UUID:       voucherRpc.UUID,
		Code:       voucherRpc.Code,
		Percentage: voucherRpc.Percentage,
		CreatedAt:  voucherRpc.CreatedAt,
		ExpiresAt:  voucherRpc.ExpiresAt,
	}

	l.Infof(msgID, "GetVoucherByIDApp output: ", " | ", voucher)
	return voucher, nil
}

func (app application) SaveVoucher(msgID string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "SaveVoucherApp: ", " | ", ps.MarshalString(voucher))

	go app.voucherRPC.SaveVoucher(voucher) // pub

	pRepo, err := app.voucherWorkerRPC.SaveVoucher(voucher)

	if err != nil {
		l.Errorf(msgID, "SaveVoucherApp error: ", " | ", err)
		return nil, err
	}

	if pRepo == nil {
		l.Infof(msgID, "SaveVoucherApp output: ", " | ", nil)
		return nil, errors.New("is not possible to save voucher because it's null")
	}

	out := &dto.OutputVoucher{
		UUID:       pRepo.UUID,
		Code:       pRepo.Code,
		Percentage: pRepo.Percentage,
		CreatedAt:  pRepo.CreatedAt,
		ExpiresAt:  pRepo.ExpiresAt,
	}

	l.Infof(msgID, "SaveVoucherApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) UpdateVoucherByID(msgID string, id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	app.setMessageIDCtx(msgID)

	l.Infof(msgID, "UpdateVoucherByIDApp: ", " | ", id, " | ", ps.MarshalString(voucher))

	go app.voucherRPC.UpdateVoucherByID(id, voucher)

	p, err := app.voucherWorkerRPC.UpdateVoucherByID(id, voucher)

	if err != nil {
		l.Errorf(msgID, "UpdateVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	if p == nil {
		voucherNullErr := errors.New("is not possible to save voucher because it's null")
		l.Errorf(msgID, "UpdateVoucherByIDApp output: ", " | ", voucherNullErr)
		return nil, voucherNullErr
	}

	out := &dto.OutputVoucher{
		UUID:       p.UUID,
		Code:       p.Code,
		Percentage: p.Percentage,
		CreatedAt:  p.CreatedAt,
		ExpiresAt:  p.ExpiresAt,
	}

	l.Infof(msgID, "UpdateVoucherByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}
