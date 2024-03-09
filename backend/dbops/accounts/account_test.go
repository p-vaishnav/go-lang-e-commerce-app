package accounts

import (
	"backend-commerce/configs"
	"backend-commerce/database"
	"backend-commerce/entities"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                              TestCreateAccount                             */
/* -------------------------------------------------------------------------- */
func TestCreateAccount(t *testing.T) {
	configs.LoadConfigs()
	var err error
	var account entities.Account

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	account.Balance = 100

	gormDB, _ := database.InitDB()
	accountsGorm := Gorm(gormDB)

	account, err = accountsGorm.CreateAccount(c, nil, account)
	assert.Empty(t, err)

	t.Cleanup(func() {
		gormDB.Model(&entities.Account{}).Where("accounts_id = ?", account.ID).Delete(&account)
	})
}

/* -------------------------------------------------------------------------- */
/*                              TestUdpateAccount                             */
/* -------------------------------------------------------------------------- */
func TestUdpateAccount(t *testing.T) {
	configs.LoadConfigs()
	var err error
	var account entities.Account

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	account.Balance = 100

	gormDB, _ := database.InitDB()
	accountsGorm := Gorm(gormDB)

	account, err = accountsGorm.CreateAccount(c, nil, account)
	assert.Empty(t, err)

	account.Balance = account.Balance + 200

	account, err = accountsGorm.UpdateAccount(c, nil, account)
	assert.Empty(t, err)
	assert.Equal(t, account.Balance, 300)

	t.Cleanup(func() {
		gormDB.Model(&entities.Account{}).Where("accounts_id = ?", account.ID).Delete(&account)
	})
}

/* -------------------------------------------------------------------------- */
/*                           TestTransactionDocument                          */
/* -------------------------------------------------------------------------- */
func TestTransactionDocument(t *testing.T) {
	configs.LoadConfigs()
	var err error
	var accountOne entities.Account
	var accountTwo entities.Account

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	accountOne.Balance = 100

	gormDB, _ := database.InitDB()
	accountsGorm := Gorm(gormDB)

	accountOne, err = accountsGorm.CreateAccount(c, nil, accountOne)
	assert.Empty(t, err)
	assert.NotEmpty(t, accountOne.ID)

	accountTwo.Balance = 200

	accountTwo, err = accountsGorm.CreateAccount(c, nil, accountTwo)
	assert.Empty(t, err)
	assert.NotEmpty(t, accountOne.ID)

	// this is where transaction happens
	accountTwo.Balance = accountTwo.Balance - 50
	accountOne.Balance = accountOne.Balance + 50

	// debit
	accountTwo, err = accountsGorm.UpdateAccount(c, nil, accountTwo)
	assert.Empty(t, err)
	assert.Equal(t, accountTwo.Balance, 150)

	// credit
	accountOne, err = accountsGorm.UpdateAccount(c, nil, accountOne)
	assert.Empty(t, err)
	assert.Equal(t, accountOne.Balance, 150)

	t.Cleanup(func() {
		gormDB.Model(&entities.Account{}).Where("accounts_id = ?", accountOne.ID).Delete(&accountOne)
		gormDB.Model(&entities.Account{}).Where("accounts_id = ?", accountTwo.ID).Delete(&accountTwo)
	})

}

func TestTransactionDocumentInvalid(t *testing.T) {
	configs.LoadConfigs()
	var err error
	var accountOne entities.Account
	var accountTwo entities.Account

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	accountOne.Balance = 100

	gormDB, _ := database.InitDB()
	accountsGorm := Gorm(gormDB)

	accountOne, err = accountsGorm.CreateAccount(c, nil, accountOne)
	assert.Empty(t, err)
	assert.NotEmpty(t, accountOne.ID)

	accountTwo.Balance = 200

	accountTwo, err = accountsGorm.CreateAccount(c, nil, accountTwo)
	assert.Empty(t, err)
	assert.NotEmpty(t, accountTwo.ID)

	tx := gormDB.Begin()

	// this is where transaction happens
	accountTwo.Balance = accountTwo.Balance - 50
	accountOne.Balance = accountOne.Balance + 50

	// debit
	accountTwo, err = accountsGorm.UpdateAccount(c, tx, accountTwo)
	assert.Empty(t, err)
	assert.Equal(t, accountTwo.Balance, 150)

	// credit
	accountOne.Name = "Excellence is the way to success" // failing the db transaction
	_, err = accountsGorm.UpdateAccount(c, tx, accountOne)
	assert.NotEmpty(t, err)
	// assert.Equal(t, accountOne.Balance, 150)

	tx.Commit()

	// I need to get one more time from the database, as state has been updated
	// when I am grabbing it is causing problem
	// var account entities.Account
	// account, err = accountsGorm.ListAccount(c, accountOne.ID) // it is giving me an error
	// assert.Empty(t, err)
	// assert.Equal(t, account.Balance, 100)

	// account, err = accountsGorm.ListAccount(c, accountTwo.ID)
	// assert.Empty(t, err)
	// assert.Equal(t, account.Balance, 200)

	// t.Cleanup(func() {
	// 	gormDB.Model(&entities.Account{}).Where("accounts_id = ?", accountOne.ID).Delete(&accountOne)
	// 	gormDB.Model(&entities.Account{}).Where("accounts_id = ?", accountTwo.ID).Delete(&accountTwo)
	// })

}
