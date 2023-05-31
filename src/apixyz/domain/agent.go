package domain

type CustomerService interface {
	GetDataListTenor(id int) ([]AgentLimit, ErrorData)
	GetGajiAgent(id int) (int, ErrorData)
	GetDataLimit(id int) ([]AgentLimit, ErrorData)
	GetListCustomer() ([]AgentData, ErrorData)
	InsertUpdateCustomer(id int, nik, fullname, legalname, placeOfbirth, birthdate string, gaji int, ktp, selfie string) (*Response, ErrorData)
	InsertUpdateCustomerLimit(id, customer, tenor, limit int) (*Response, ErrorData)
}

type CustomerRepository interface {
	GetGajiAgent(id int) (int, ErrorData)
	GetLimitAgent(id int) ([]AgentLimit, ErrorData)
	InsertUpdateCustomerLimit(id, customer, tenor, limit int) (*Response, ErrorData)
	GetListCustomer() ([]AgentData, ErrorData)
	InsertUpdateCustomer(id int, nik, fullname, legalname, placeOfbirth, birthdate string, gaji int, ktp, selfie string) (*Response, ErrorData)
}

type RequestIdTenor struct {
	Id int `json:"id"`
}

type AgentData struct {
	Id           int    `json:"id"`
	NIK          string `json:"nik"`
	FullName     string `json:"full_name"`
	LegalName    string `json:"legal_name"`
	PlaceOfBirth string `json:"place_of_birth"`
	BirthDate    string `json:"birth_date"`
	Gaji         int    `json:"gaji"`
	FotoKTP      string `json:"foto_ktp"`
	FotoSelfie   string `json:"foto_selfie"`
}

type AgentLimit struct {
	Id           int `json:"id"`
	Customer     int `json:"customer_id"`
	Tenor        int `json:"tenor"`
	LimitAmmount int `json:"limit_ammount"`
}

type Listpayments struct {
	Id          int    `json:"id"`
	CustomerId  int    `json:"customer_id"`
	ContractNo  int    `json:"contract_no"`
	Installment int    `json:"installment"`
	PaymentDate string `json:"payment_date"`
	Ammount     int    `json:"ammount"`
}

type Transactions struct {
	Id              int    `json:"id"`
	CustomerId      int    `json:"customer_id"`
	ContractNo      string `json:"contract_no"`
	Otr             int    `json:"otr"`
	AdminFee        int    `json:"admin_fee"`
	Installment     int    `json:"installment"`
	InterestAmmount int    `json:"interest_ammount"`
	Assetname       string `json:"asset_name"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
}
