package repoorder

import (
	"context"
	iorder "nuryanto2121/cukur_in_web/interface/order"
	"nuryanto2121/cukur_in_web/models"

	"gorm.io/gorm"
)

type repoOrder struct {
	Conn *gorm.DB
}

func NewRepoOrder(conn *gorm.DB) iorder.Repository {
	return &repoOrder{conn}
}

func (r *repoOrder) GetDataOrderWithTeguk(ctx context.Context) (result []*models.OrderPost, err error) {

	sSql := `
	select * from vorder_barber_teguk
	order by order_date
	`
	if err := r.Conn.WithContext(ctx).Raw(sSql).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

// GetDataOrderStatusArriveOnTime implements iorder.Repository.
func (r *repoOrder) GetDataOrderStatusArriveOnTime(ctx context.Context) (result []*models.OrderNotif, err error) {

	intervalNotif30, intervalNotif15 := -30, -15
	// orders := []*models.OrderH{}
	if err := r.Conn.WithContext(ctx).Select(`(EXTRACT(EPOCH FROM (localtimestamp - order_date)) / 60)::int as timeArrive,order_id,order_no,status,user_id,customer_name,capster_id,order_date`).
		Where("status = ?", "N").
		Where(`
			((EXTRACT(EPOCH FROM (localtimestamp - order_date)) / 60)::int =?
			or (EXTRACT(EPOCH FROM (localtimestamp - order_date)) / 60)::int =?)`,
			intervalNotif30, intervalNotif15).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil

}
