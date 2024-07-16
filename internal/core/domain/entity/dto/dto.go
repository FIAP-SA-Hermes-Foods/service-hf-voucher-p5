package dto

import (
	"service-hf-voucher-p5/internal/core/domain/entity"
	vo "service-hf-voucher-p5/internal/core/domain/entity/valueObject"
	"strconv"
	"time"
)

type VoucherDB struct {
	UUID       string `json:"uuid,omitempty"`
	Code       string `json:"code,omitempty"`
	Percentage string `json:"percentage,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	ExpiresAt  string `json:"expiresAt,omitempty"`
}

type (
	RequestVoucher struct {
		UUID       string `json:"uuid,omitempty"`
		Code       string `json:"code,omitempty"`
		Percentage string `json:"percentage,omitempty"`
		CreatedAt  string `json:"createdAt,omitempty"`
		ExpiresAt  string `json:"expiresAt,omitempty"`
	}

	OutputVoucher struct {
		UUID       string `json:"uuid,omitempty"`
		Code       string `json:"code,omitempty"`
		Percentage string `json:"percentage,omitempty"`
		CreatedAt  string `json:"createdAt,omitempty"`
		ExpiresAt  string `json:"expiresAt,omitempty"`
	}
)

func (r RequestVoucher) Voucher() entity.Voucher {
	expirationTime, _ := time.Parse("02-01-2006 15:04:05", r.ExpiresAt)
	createdAt, _ := time.Parse("02-01-2006 15:04:05", r.CreatedAt)

	percentage, _ := strconv.ParseInt(r.Percentage, 10, 64) // Convert r.Percentage to int64

	return entity.Voucher{
		Code:       r.Code,
		Percentage: percentage, // Use the converted value
		ExpiresAt: vo.ExpiresAt{
			Value: &expirationTime,
		},
		CreatedAt: vo.CreatedAt{
			Value: createdAt,
		},
	}
}
