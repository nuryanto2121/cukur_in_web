package useimportredem

import (
	"context"
	iimportredem "nuryanto2121/cukur_in_web/interface/import_redem"
	"nuryanto2121/cukur_in_web/models"
	reporedem "nuryanto2121/cukur_in_web/repository/redem"
	"time"

	"gorm.io/gorm"
)

type useImportRedem struct {
	conn           *gorm.DB
	contextTimeOut time.Duration
}

func NewImportRedem(Conn *gorm.DB, timeout time.Duration) iimportredem.Usecase {
	return &useImportRedem{
		contextTimeOut: timeout,
		conn:           Conn,
	}
}
func (u *useImportRedem) CreateRedemTeguk(ctx context.Context, data []*models.RedemTeguk) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()
	var (
		err       error
		reporedem = &reporedem.RepoRedem{
			Conn: u.conn,
		}
	)

	err = reporedem.Delete()
	if err != nil {
		return err
	}

	err = reporedem.Insert(data)
	if err != nil {
		return err
	}

	return nil
}
