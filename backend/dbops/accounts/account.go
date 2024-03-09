package accounts

import (
	"backend-commerce/entities"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  Interface                                 */
/* -------------------------------------------------------------------------- */
type Interface interface {
	CreateAccount(ctx *gin.Context, account entities.Account) (entities.Account, error)
	UpdateAccount(ctx *gin.Context, account entities.Account) (entities.Account, error)
	ListAccount(ctx *gin.Context, id int) (entities.Account, error)
}

type accounts struct {
	DB *gorm.DB
}

func Gorm(DB *gorm.DB) *accounts {
	return &accounts{
		DB: DB,
	}
}

/* -------------------------------------------------------------------------- */
/*                                CreateAccount                               */
/* -------------------------------------------------------------------------- */
func (g *accounts) CreateAccount(ctx *gin.Context, tx *gorm.DB, account entities.Account) (entities.Account, error) {
	var err error

	if tx != nil {
		db := tx.Session(&gorm.Session{}).Create(&account)
		err = db.Error
	} else {
		db := g.DB.Session(&gorm.Session{}).Create(&account)
		err = db.Error
	}

	if err != nil {
		g.DB.Rollback()
		return account, errors.Wrap(err, "[CreateAccount]")
	}

	return account, err
}

/* -------------------------------------------------------------------------- */
/*                                UpdateAccount                               */
/* -------------------------------------------------------------------------- */
func (g *accounts) UpdateAccount(ctx *gin.Context, tx *gorm.DB, account entities.Account) (entities.Account, error) {
	var err error

	if tx != nil {
		db := tx.Model(&entities.Account{}).Where("accounts_id = ?", account.ID).Updates(&account)
		err = db.Error
	} else {
		db := g.DB.Model(&entities.Account{}).Where("accounts_id = ?", account.ID).Updates(&account)
		err = db.Error
	}

	if err != nil {
		g.DB.Rollback()
		return account, errors.Wrap(err, "[UpdateDocument]")
	}

	return account, err
}

/* -------------------------------------------------------------------------- */
/*                                 ListAccount                                */
/* -------------------------------------------------------------------------- */
func (g *accounts) ListAccount(ctx *gin.Context, id int) (entities.Account, error) {
	var err error
	var account entities.Account

	db := g.DB.Model(&entities.Account{}).Where("accounts_id = ?", id).Take(&account)
	err = db.Error
	if err != nil {
		return account, errors.Wrap(err, "[ListAccount]")
	}

	return account, err
}
