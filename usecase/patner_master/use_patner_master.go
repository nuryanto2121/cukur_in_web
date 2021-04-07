package usepatnermaster

import (
	"context"
	"fmt"
	"math"
	ipatnermaster "nuryanto2121/cukur_in_web/interface/patner_master"
	"nuryanto2121/cukur_in_web/models"
	util "nuryanto2121/cukur_in_web/pkg/utils"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

type usePatnerMaster struct {
	repoPatnerMaster ipatnermaster.Repository
	contextTimeOut   time.Duration
}

func NewUsePatnerMaster(a ipatnermaster.Repository, timeout time.Duration) ipatnermaster.Usecase {
	return &usePatnerMaster{repoPatnerMaster: a, contextTimeOut: timeout}
}

func (u *usePatnerMaster) GetDataBy(ctx context.Context, Claims util.Claims, ID int) (result *models.PatnerMaster, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	result, err = u.repoPatnerMaster.GetDataBy(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *usePatnerMaster) GetList(ctx context.Context, Claims util.Claims, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	if queryparam.Search != "" {
		queryparam.Search = strings.ToLower(fmt.Sprintf("%%%s%%", queryparam.Search))
	}

	if queryparam.InitSearch != "" {

	}
	result.Data, err = u.repoPatnerMaster.GetList(queryparam)
	if err != nil {
		return result, err
	}

	result.Total, err = u.repoPatnerMaster.Count(queryparam)
	if err != nil {
		return result, err
	}

	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}

func (u *usePatnerMaster) Create(ctx context.Context, Claims util.Claims, data *models.AddPatnerMaster) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()
	var (
		mPatnerMaster = models.PatnerMaster{}
	)

	// mapping to struct model saRole
	err = mapstructure.Decode(data, &mPatnerMaster.AddPatnerMaster)
	if err != nil {
		return err
	}

	mPatnerMaster.UserEdit = "Martin"  //Claims.UserID
	mPatnerMaster.UserInput = "Martin" //Claims.UserID

	err = u.repoPatnerMaster.Create(&mPatnerMaster)
	if err != nil {
		return err
	}
	return nil

}

func (u *usePatnerMaster) Update(ctx context.Context, Claims util.Claims, ID int, data *models.AddPatnerMaster) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	myMap := structs.Map(data)
	myMap["user_edit"] = "Martin" //Claims.UserID
	fmt.Println(myMap)
	err = u.repoPatnerMaster.Update(ID, myMap)
	if err != nil {
		return err
	}
	return nil
}

func (u *usePatnerMaster) Delete(ctx context.Context, Claims util.Claims, ID int) (err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeOut)
	defer cancel()

	err = u.repoPatnerMaster.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
