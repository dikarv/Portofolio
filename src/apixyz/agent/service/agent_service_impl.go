package service

import "apixyz/src/apixyz/domain"

type CustomerServiceImpl struct {
	custRepo domain.CustomerRepository
}

func (c CustomerServiceImpl) GetListCustomer(customerId int) ([]domain.AgentData, domain.ErrorData) {
	return c.custRepo.GetListCustomer(customerId)
}

func CreateServiceImpl(custRepo domain.CustomerRepository) domain.CustomerService {
	return &CustomerServiceImpl{
		custRepo: custRepo,
	}
}
