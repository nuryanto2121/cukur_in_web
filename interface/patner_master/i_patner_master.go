
package ipatnermaster

import (
	"context"
	util "nuryanto2121/cukur_in_web/pkg/utils"
	"nuryanto2121/cukur_in_web/models"
	
)
	
type Repository interface {
	GetDataBy(ID int) (result *models.PatnerMaster, err error)
	GetList(queryparam models.ParamList) (result []*models.PatnerMaster, err error)
	Create(data *models.PatnerMaster) (err error)
	Update(ID int, data map[string]interface{}) (err error)
	Delete(ID int) (err error)
	Count(queryparam models.ParamList) (result int, err error)
}

type Usecase interface {
	GetDataBy(ctx context.Context, Claims util.Claims, ID int) (result *models.PatnerMaster, err error)
	GetList(ctx context.Context, Claims util.Claims, queryparam models.ParamList) (result models.ResponseModelList, err error)
	Create(ctx context.Context, Claims util.Claims, data *models.AddPatnerMaster) (err error)
	Update(ctx context.Context, Claims util.Claims, ID int, data *models.AddPatnerMaster) (err error)
	Delete(ctx context.Context, Claims util.Claims, ID int) (err error)
}
	
	