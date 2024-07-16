package entity

import (
	vo "service-hf-voucher-p5/internal/core/domain/entity/valueObject"
)

type Voucher struct {
	ID        	int64        		`json:"id,omitempty"`
	Code      	string       		`json:"code,omitempty"`
	Percentage  int64       		`json:"percentage,omitempty"`
	CreatedAt 	vo.CreatedAt 		`json:"createdAt,omitempty"`
	ExpiresAt   vo.ExpiresAt      	`json:"expiresAt,omitempty"`
}
