package iorder

import (
	"context"
	"nuryanto2121/cukur_in_web/models"
)

type Repository interface {
	GetDataOrderWithTeguk(ctx context.Context) (result []*models.OrderPost, err error)
	GetDataOrderStatusArriveOnTime(ctx context.Context) (result []*models.OrderNotif, err error)
}

type Usecase interface {
	// GetDataBy(ctx context.Context, Claims util.Claims, ID int, GeoUser models.GeoBarber) (result interface{}, err error)
	// GetList(ctx context.Context, Claims util.Claims, queryparam models.ParamListGeo) (result models.ResponseModelList, err error)
	// Create(ctx context.Context, Claims util.Claims, data *models.OrderPost) error
	// Update(ctx context.Context, Claims util.Claims, ID int, data models.OrderStatus) (err error)
	// Delete(ctx context.Context, Claims util.Claims, ID int) (err error)
}
