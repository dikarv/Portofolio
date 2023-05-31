package test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCustomerServiceImpl_Paymenst(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockResponse := &domain.Response{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi Paymenst di repository
// 	mockRepo.On("Paymenst", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.Paymenst(1, 2, 3, 1000)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "Paymenst", mock.Anything, 1, 2, 3, mock.Anything, 1000)
// }

// func TestCustomerServiceImpl_UpdateCustomerLimitAmount(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockResponse := &domain.Response{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi GetTenorAgent di repository
// 	mockRepo.On("GetTenorAgent", mock.Anything, mock.Anything, mock.Anything).Return(&domain.ResponseAgentLimit{}, mockErrorData)

// 	// Simulasikan pemanggilan fungsi UpdateCustomerLimitAmount di repository
// 	mockRepo.On("UpdateCustomerLimitAmount", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.UpdateCustomerLimitAmount(1, 2, 3, 1000)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "GetTenorAgent", mock.Anything, 2, 3)
// 	mockRepo.AssertCalled(t, "UpdateCustomerLimitAmount", mock.Anything, 1, 2, mock.Anything)
// }

// func TestCustomerServiceImpl_GetTenorAgent(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockResponse := &domain.ResponseAgentLimit{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi GetTenorAgent di repository
// 	mockRepo.On("GetTenorAgent", mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.GetTenorAgent(1, 2, 3)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "GetTenorAgent", mock.Anything, 1, 2, 3)
// }

// func TestCustomerServiceImpl_Transaction(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockResponse := &domain.Response{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi Transaction di repository
// 	mockRepo.On("Transaction", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.Transaction(1, "ABC123", 12, 50000000, "Car", "Buy")

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "Transaction", mock.Anything, 1, "ABC123", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, "Car", "Buy", mock.Anything)
// }

// func TestCustomerServiceImpl_GetGajiAgent(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockGaji := 5000000
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi GetGajiAgent di repository
// 	mockRepo.On("GetGajiAgent", mock.Anything).Return(mockGaji, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	gaji, err := service.GetGajiAgent(1)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockGaji, gaji)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "GetGajiAgent", mock.Anything, 1)
// }

// func TestCustomerServiceImpl_GetDataListTenor(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockResponse := []domain.AgentLimit{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi GetLimitAgent di repository
// 	mockRepo.On("GetLimitAgent", mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.GetDataListTenor(1)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "GetLimitAgent", mock.Anything, 1)
// }

// func TestCustomerServiceImpl_GetDataLimit(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockGaji := 5000000
// 	mockResponse := []domain.AgentLimit{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi GetGajiAgent di repository
// 	mockRepo.On("GetGajiAgent", mock.Anything).Return(mockGaji, mockErrorData)

// 	// Simulasikan pemanggilan fungsi InsertUpdateCustomerLimit di repository
// 	mockRepo.On("InsertUpdateCustomerLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&domain.Response{}, domain.ErrorData{})

// 	// Simulasikan pemanggilan fungsi GetLimitAgent di repository
// 	mockRepo.On("GetLimitAgent", mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.GetDataLimit(1)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "GetGajiAgent", mock.Anything, 1)
// 	mockRepo.AssertCalled(t, "InsertUpdateCustomerLimit", mock.Anything, mock.Anything, 1, 1000000)
// 	mockRepo.AssertCalled(t, "InsertUpdateCustomerLimit", mock.Anything, mock.Anything, 2, 1200000)
// 	mockRepo.AssertCalled(t, "InsertUpdateCustomerLimit", mock.Anything, mock.Anything, 3, 1500000)
// 	mockRepo.AssertCalled(t, "InsertUpdateCustomerLimit", mock.Anything, mock.Anything, 6, 2000000)
// 	mockRepo.AssertCalled(t, "GetLimitAgent", mock.Anything, 1)
// }

// func TestCustomerServiceImpl_InsertUpdateCustomerLimit(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockResponse := &domain.Response{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi InsertUpdateCustomerLimit di repository
// 	mockRepo.On("InsertUpdateCustomerLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.InsertUpdateCustomerLimit(1, 2, 3, 1000)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "InsertUpdateCustomerLimit", mock.Anything, 1, 2, 3, 1000)
// }

// func TestCustomerServiceImpl_InsertUpdateCustomer(t *testing.T) {
// 	mockRepo := &MockCustomerRepository{}
// 	service := CreateServiceImpl(mockRepo)

// 	// Simulasikan hasil yang diharapkan dari repository
// 	mockResponse := &domain.Response{}
// 	mockErrorData := domain.ErrorData{}

// 	// Simulasikan pemanggilan fungsi InsertUpdateCustomer di repository
// 	mockRepo.On("InsertUpdateCustomer", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockErrorData)

// 	// Panggil fungsi yang akan diuji
// 	response, err := service.InsertUpdateCustomer(1, "1234567890", "John Doe", "john@example.com", "1234567890", "address", "city", "state", "12345")

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, mockResponse, response)
// 	assert.Equal(t, mockErrorData, err)

// 	// Periksa apakah fungsi di repository dipanggil dengan argumen yang sesuai
// 	mockRepo.AssertCalled(t, "InsertUpdateCustomer", mock.Anything, 1, "1234567890", "John Doe", "john@example.com", "1234567890", "address", "city", "state", "12345")
// }
