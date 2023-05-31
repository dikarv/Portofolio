package test

// import (

// 	"net/http"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCustomerController_Payments_Success(t *testing.T) {
// 	mockService := &MockCustomerService{}
// 	controller := CustomerController{
// 		CustomerService: mockService,
// 	}

// 	// Simulasikan request JSON yang diterima oleh controller
// 	request := domain.Requestpayments{
// 		CustomerId:  1,
// 		ContractNo:  2,
// 		Installment: 3,
// 		Amount:      1000,
// 	}
// 	jsonRequest := []byte(`{"CustomerId": 1, "ContractNo": 2, "Installment": 3, "Amount": 1000}`)

// 	// Simulasikan hasil yang diharapkan dari pemanggilan service
// 	mockResponse := &domain.Response{
// 		Data:    &domain.Result{},
// 		Message: "Update Success",
// 		Status:  "200",
// 	}
// 	mockService.On("Paymenst", request.CustomerId, request.ContractNo, request.Installment, request.Amount).Return(mockResponse, nil)

// 	// Buat context dan request HTTP palsu untuk pengujian
// 	router := gin.Default()
// 	router.POST("/payments", controller.Payments)
// 	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonRequest))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	// Periksa apakah hasilnya sesuai dengan yang diharapkan
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"data": {}, "message": "Update Success", "status": "200"}`, w.Body.String())

// 	// Periksa apakah service dipanggil dengan argumen yang sesuai
// 	mockService.AssertCalled(t, "Paymenst", request.CustomerId, request.ContractNo, request.Installment, request.Amount)
// }

// // Unit test untuk fungsi-fungsi lainnya dapat dilakukan dengan cara yang serupa.
