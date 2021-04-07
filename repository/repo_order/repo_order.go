package repoorder

import (
	"fmt"
	"log"
	"nuryanto2121/cukur_in_web/models"

	"gorm.io/gorm"
)

type RepoOrder struct {
	Conn *gorm.DB
}

func (db *RepoOrder) DataOrder() (result []*models.OrderPost, err error) {

	sSql := `
	select * from vorder_barber_teguk
	order by order_date	
	`
	query := db.Conn.Raw(sSql).Find(&result)
	log.Printf(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
