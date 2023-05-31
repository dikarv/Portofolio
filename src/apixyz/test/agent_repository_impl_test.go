package test

// import (

// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCustomerRepoImpl_Payments(t *testing.T) {
// 	mockDB := &MockDB{}
// 	repo := CustomerRepoImpl{
// 		DB: mockDB,
// 	}

// 	// Simulasikan hasil yang diharapkan dari eksekusi query
// 	mockResult := &MockResult{}
// 	mockResult.On("LastInsertId").Return(int64(1), nil)

// 	mockDB.On("GetConnectionSKS").Return(mockConn, nil)
// 	mockConn.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResult, nil)

// 	// Panggil fungsi yang akan diuji
// 	response, err := repo.Payments(1, 2, 3, 4, "2023-05-31", 1000)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.NoError(t, err)
// 	assert.Equal(t, &domain.Response{
// 		Data:    mockResult,
// 		Message: "Data Succesfully Updated",
// 		Status:  "200",
// 	}, response)

// 	// Periksa apakah fungsi di database dipanggil dengan argumen yang sesuai
// 	mockDB.AssertCalled(t, "GetConnectionSKS")
// 	mockConn.AssertCalled(t, "Exec", "INSERT INTO payments(ID, Customer, ContractNo, Installment, PaymentDate, Amount) VALUES (?,?,?,?,?,?)", 1, 2, 3, 4, "2023-05-31", 1000)
// 	mockResult.AssertCalled(t, "LastInsertId")
// }

// func TestCustomerRepoImpl_UpdateCustomerLimitAmount(t *testing.T) {
// 	mockDB := &MockDB{}
// 	repo := CustomerRepoImpl{
// 		DB: mockDB,
// 	}

// 	// Simulasikan hasil yang diharapkan dari eksekusi query
// 	mockResult := &MockResult{}
// 	mockResult.On("RowsAffected").Return(int64(1), nil)

// 	mockDB.On("GetConnectionSKS").Return(mockConn, nil)
// 	mockConn.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResult, nil)

// 	// Panggil fungsi yang akan diuji
// 	response, err := repo.UpdateCustomerLimitAmount(1, 2, 3, 1000)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.NoError(t, err)
// 	assert.Equal(t, &domain.Response{
// 		Data:    mockResult,
// 		Message: "Data Succesfully Updated",
// 		Status:  "200",
// 	}, response)

// 	// Periksa apakah fungsi di database dipanggil dengan argumen yang sesuai
// 	mockDB.AssertCalled(t, "GetConnectionSKS")
// 	mockConn.AssertCalled(t, "Exec", "UPDATE customerlimits SET LimitAmount=? WHERE Customer = ? AND Tenor = ? AND LimitAmount=?", 1000, 1, 2, 3)
// 	mockResult.AssertCalled(t, "RowsAffected")
// }

// func TestCustomerRepoImpl_GetTenorAgent(t *testing.T) {
// 	mockDB := &MockDB{}
// 	repo := CustomerRepoImpl{
// 		DB: mockDB,
// 	}

// 	// Simulasikan hasil yang diharapkan dari eksekusi query
// 	mockResult := &MockResult{}
// 	mockResult.On("Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

// 	mockDB.On("GetConnectionSKS").Return(mockConn, nil)
// 	mockConn.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResult, nil)

// 	// Panggil fungsi yang akan diuji
// 	response, err := repo.GetTenorAgent(1, 2, 3)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.NoError(t, err)
// 	assert.Equal(t, &domain.ResponseAgentLimit{
// 		Id:           0,
// 		Customer:     "",
// 		Tenor:        0,
// 		LimitAmmount: 0,
// 	}, response)

// 	// Periksa apakah fungsi di database dipanggil dengan argumen yang sesuai
// 	mockDB.AssertCalled(t, "GetConnectionSKS")
// 	mockConn.AssertCalled(t, "QueryRow", "SELECT * FROM customerlimits WHERE Customer = ? AND Tenor = ? AND LimitAmount = ?", 1, 2, 3)
// 	mockResult.AssertCalled(t, "Scan", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
// }

// // Unit test untuk fungsi-fungsi lainnya dapat dilakukan dengan cara yang serupa.
