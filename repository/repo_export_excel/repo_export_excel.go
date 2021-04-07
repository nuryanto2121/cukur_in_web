package repoexportexcel

import (
	"nuryanto2121/cukur_in_web/models"

	"gorm.io/gorm"
)

type RepoExportExcel struct {
	Conn *gorm.DB
}

func (db *RepoExportExcel) GetListTrx(source int) (result []*models.TrxApp, err error) {
	var sSql = ""
	if source == 1 {
		sSql = `
		select a.order_no,a.order_date,
			b.barber_cd,b.barber_name,
				s.name as capster_name,
			a.customer_name,a.from_apps ,
				x.paket_name,x.price
		from order_h a join order_d x
		on a.order_id = x.order_id
		and a.barber_id = x.barber_id
		join barber b on a.barber_id = b.barber_id
		join ss_user s on s.user_id = a.capster_id
		where a.from_apps = true
		and a.user_id > 0
		order by a.order_date
		`
	} else {
		sSql = `
		select a.order_no,a.order_date,
			b.barber_cd,b.barber_name,
				s.name as capster_name,
			a.customer_name,a.from_apps ,
				x.paket_name,x.price
		from order_h a join order_d x
		on a.order_id = x.order_id
		and a.barber_id = x.barber_id
		join barber b on a.barber_id = b.barber_id
		join ss_user s on s.user_id = a.capster_id
		where a.from_apps = false
		and a.user_id = 0
		order by a.order_date
		`
	}

	query := db.Conn.Raw(sSql).Find(&result)
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
