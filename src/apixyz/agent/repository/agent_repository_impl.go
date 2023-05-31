package repository

import (
	"apixyz/src/apixyz/database"
	"apixyz/src/apixyz/domain"
	"database/sql"
	"errors"
	"net/http"
)

type CustomerRepoImpl struct {
}

func (c CustomerRepoImpl) GetListCustomer(customerId int) ([]domain.AgentData, domain.ErrorData) {
	var result []domain.AgentData
	var errorResponse domain.ErrorData
	var rows *sql.Rows
	var errquery error

	dbsks, err := database.GetConnectionSKS()
	if err != nil {
		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errors.New("DB SKS Connection Closed")
		return nil, errorResponse
	}

	if customerId != 0 {
		rows, errquery = dbsks.Query("SELECT * FROM customers WHERE ID = ?", customerId)

	} else {
		rows, errquery = dbsks.Query("SELECT * FROM customers")

	}
	if errquery != nil {
		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errquery
		return nil, errorResponse
	}

	defer rows.Close()

	for rows.Next() {
		var cust domain.AgentData
		errloop := rows.Scan(&cust.Id, &cust.NIK, &cust.FullName, &cust.LegalName, &cust.PlaceOfBirth, &cust.BirthDate, &cust.Gaji, &cust.FotoKTP, &cust.FotoSelfie)
		if errloop != nil {
			errorResponse.Status = http.StatusInternalServerError
			errorResponse.Message = errloop
			return nil, errorResponse
		}

		result = append(result, cust)

	}
	if result == nil {
		errorResponse.Status = http.StatusNotFound
		errorResponse.Message = errors.New("Data Not Found")
		return nil, errorResponse
	} else {
		return result, domain.ErrorData{}
	}
}

func CreateCustomerRepoImpl() domain.CustomerRepository {
	return &CustomerRepoImpl{}
}
