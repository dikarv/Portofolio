package service

import (
	"apixyz/src/apixyz/domain"

	"math/rand"
	"time"
)

type CustomerServiceImpl struct {
	custRepo domain.CustomerRepository
}

func (c CustomerServiceImpl) Paymenst(customer, contract, installment int, Amount int) (*domain.Response, domain.ErrorData) {

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101)

	currentTime := time.Now()

	//Format the current time using a specific layout
	formattedTime := currentTime.Format("2006-01-02")

	return c.custRepo.Paymenst(randomNumber, customer, contract, installment, formattedTime, Amount)

}

func (c CustomerServiceImpl) UpdateCustomerLimitAmount(customer, tenor, limit, total int) (*domain.Response, domain.ErrorData) {

	data, err := c.custRepo.GetTenorAgent(customer, tenor, limit)
	if err.Message != nil {
		return nil, domain.ErrorData{}
	}

	newAmountCust := data.LimitAmmount - total

	return c.custRepo.UpdateCustomerLimitAmount(customer, tenor, limit, newAmountCust)
}

func (c CustomerServiceImpl) GetTenorAgent(id, tenor, limit int) (*domain.ResponseAgentLimit, domain.ErrorData) {
	return c.custRepo.GetTenorAgent(id, tenor, limit)
}
func (c CustomerServiceImpl) Transaction(customer int, contractNo string, tenor, otr int, assetName, transactionType string) (*domain.Response, domain.ErrorData) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101)

	currentTime := time.Now()

	//Format the current time using a specific layout
	formattedTime := currentTime.Format("2006-01-02")

	newAdminFee := otr / 70
	newInstallment := otr / tenor
	newInterestAmount := otr / 30
	total := newInstallment*tenor + newInterestAmount*tenor + newAdminFee
	cicilan := total / tenor

	return c.custRepo.Transaction(randomNumber, customer, contractNo, total, tenor, otr, newAdminFee, cicilan, newInterestAmount, assetName, transactionType, formattedTime)

}
func (c CustomerServiceImpl) GetGajiAgent(id int) (int, domain.ErrorData) {
	return c.custRepo.GetGajiAgent(id)
}

func (c CustomerServiceImpl) GetDataListTenor(id int) ([]domain.AgentLimit, domain.ErrorData) {

	return c.custRepo.GetLimitAgent(id)
}

func (c CustomerServiceImpl) GetDataLimit(id int) ([]domain.AgentLimit, domain.ErrorData) {
	gaji, err := c.custRepo.GetGajiAgent(id)
	if err.Message != nil {
		return nil, domain.ErrorData{}
	}

	go func() {
		if gaji >= 5000000 {
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(10001)
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 1, 1000000)
			randomNumber++
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 2, 1200000)
			randomNumber++
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 3, 1500000)
			randomNumber++
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 6, 2000000)
		} else {
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(10001)
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 1, 100000)
			randomNumber++
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 2, 200000)
			randomNumber++
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 3, 500000)
			randomNumber++
			c.custRepo.InsertUpdateCustomerLimit(randomNumber, id, 6, 700000)
		}
	}()

	// Tunggu goroutine selesai
	time.Sleep(1 * time.Second)

	return c.custRepo.GetLimitAgent(id)
}

func (c CustomerServiceImpl) InsertUpdateCustomerLimit(id, customer, tenor, limit int) (*domain.Response, domain.ErrorData) {
	return c.custRepo.InsertUpdateCustomerLimit(id, customer, tenor, limit)
}

func (c CustomerServiceImpl) InsertUpdateCustomer(id int, nik, fullname, legalname, placeOfbirth, birthdate string, gaji int, ktp, selfie string) (*domain.Response, domain.ErrorData) {

	return c.custRepo.InsertUpdateCustomer(id, nik, fullname, legalname, placeOfbirth, birthdate, gaji, ktp, selfie)
}

func (c CustomerServiceImpl) GetListCustomer() ([]domain.AgentData, domain.ErrorData) {
	return c.custRepo.GetListCustomer()
}

func CreateServiceImpl(custRepo domain.CustomerRepository) domain.CustomerService {
	return &CustomerServiceImpl{
		custRepo: custRepo,
	}
}
