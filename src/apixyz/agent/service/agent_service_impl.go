package service

import "apixyz/src/apixyz/domain"

type CustomerServiceImpl struct {
	custRepo domain.CustomerRepository
}

func (c CustomerServiceImpl) GetListCustomer() ([]domain.AgentData, domain.ErrorData) {
	return c.custRepo.GetListCustomer()
}

func CreateServiceImpl(custRepo domain.CustomerRepository) domain.CustomerService {
	return &CustomerServiceImpl{
		custRepo: custRepo,
	}
}
