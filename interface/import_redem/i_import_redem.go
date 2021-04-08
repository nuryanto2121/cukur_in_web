package iimportredem

import (
	"context"
	"nuryanto2121/cukur_in_web/models"
)

type Usecase interface {
	CreateRedemTeguk(ctx context.Context, data []*models.RedemTeguk) error
}

type Repositry interface {
	CreateRedemTeguk(ctx context.Context, data []*models.RedemTeguk) error
}
