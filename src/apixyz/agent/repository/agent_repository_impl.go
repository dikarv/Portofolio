package repository

import (
	"apixyz/src/apixyz/database"
	"apixyz/src/apixyz/domain"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type CustomerRepoImpl struct {
}

func (c CustomerRepoImpl) GetGajiAgent(id int) (int, domain.ErrorData) {

	var (
		result        domain.AgentData
		errorResponse domain.ErrorData
	)

	dbsks, err := database.GetConnectionSKS()
	if err != nil {
		fmt.Println("Error disini ", err)

		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errors.New("DB SKS Connection Closed")
		return 0, errorResponse
	}

	rows := dbsks.QueryRow("SELECT Gaji FROM customers WHERE ID = ?", id).Scan(&result.Gaji)

	if rows != nil {
		fmt.Println("Error disini ", rows)
		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = rows
		return 0, errorResponse
	}

	if result.Gaji == 0 {
		errorResponse.Status = http.StatusNotFound
		errorResponse.Message = errors.New("Data Not Found")
		return 0, errorResponse
	} else {
		return result.Gaji, domain.ErrorData{}
	}
}

func (c CustomerRepoImpl) GetLimitAgent(id int) ([]domain.AgentLimit, domain.ErrorData) {
	var result []domain.AgentLimit
	var errorResponse domain.ErrorData

	dbsks, err := database.GetConnectionSKS()
	if err != nil {
		fmt.Println("Error disini ", err)

		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errors.New("DB SKS Connection Closed")
		return nil, errorResponse
	}

	rows, errquery := dbsks.Query("SELECT * FROM customerlimits WHERE Customer = ?", id)

	if errquery != nil {
		fmt.Println("Error disini ", errquery)
		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errquery
		return nil, errorResponse
	}

	defer rows.Close()

	for rows.Next() {
		var cust domain.AgentLimit
		errloop := rows.Scan(&cust.Id, &cust.Customer, &cust.Tenor, &cust.LimitAmmount)
		if errloop != nil {
			fmt.Println("Error disini ", errquery)

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

func (c CustomerRepoImpl) InsertUpdateCustomerLimit(id, customer, tenor, limit int) (*domain.Response, domain.ErrorData) {
	var (
		result        domain.Response
		errorResponse domain.ErrorData
	)

	dbsks, err := database.GetConnectionSKS()
	if err != nil {
		fmt.Println("Error disini ", err)

		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errors.New("DB SKS Connection Closed")
		return nil, errorResponse
	}

	query := `INSERT INTO customerlimits (ID, Customer, Tenor, LimitAmount)
	VALUES (?, ?, ?, ?);`

	data, err := dbsks.Exec(query, id, customer, tenor, limit)

	if err != nil {
		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = err
		return nil, errorResponse
	}
	result.Data = data
	result.Message = "Data Succesfully Inserted"
	result.Status = "200"

	return &result, errorResponse
}

func (c CustomerRepoImpl) InsertUpdateCustomer(id int, nik, fullname, legalname, placeOfbirth, birthdate string, gaji int, ktp, selfie string) (*domain.Response, domain.ErrorData) {
	var (
		result        domain.Response
		errorResponse domain.ErrorData
	)

	dbsks, err := database.GetConnectionSKS()
	if err != nil {
		fmt.Println("Error disini ", err)

		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errors.New("DB SKS Connection Closed")
		return nil, errorResponse
	}

	query := `INSERT INTO customers (ID, NIK, FullName, LegalName, PlaceOfBirth, BirthDate, Gaji, FotoKTP, FotoSelfie)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	data, err := dbsks.Exec(query, id, nik, fullname, legalname, placeOfbirth, birthdate, gaji, ktp, selfie)

	if err != nil {
		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = err
		return nil, errorResponse
	}
	result.Data = data
	result.Message = "Data Succesfully Inserted"
	result.Status = "200"

	return &result, errorResponse
}

func (c CustomerRepoImpl) GetListCustomer() ([]domain.AgentData, domain.ErrorData) {
	var result []domain.AgentData
	var errorResponse domain.ErrorData
	var rows *sql.Rows
	var errquery error

	dbsks, err := database.GetConnectionSKS()
	if err != nil {
		fmt.Println("Error disini ", errquery)

		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errors.New("DB SKS Connection Closed")
		return nil, errorResponse
	}

	rows, errquery = dbsks.Query("SELECT * FROM customers")

	if errquery != nil {
		fmt.Println("Error disini ", errquery)
		errorResponse.Status = http.StatusInternalServerError
		errorResponse.Message = errquery
		return nil, errorResponse
	}

	defer rows.Close()

	for rows.Next() {
		var cust domain.AgentData
		errloop := rows.Scan(&cust.Id, &cust.NIK, &cust.FullName, &cust.LegalName, &cust.PlaceOfBirth, &cust.BirthDate, &cust.Gaji, &cust.FotoKTP, &cust.FotoSelfie)
		if errloop != nil {
			fmt.Println("Error disini ", errquery)

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
