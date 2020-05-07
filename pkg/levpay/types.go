package levpay

import "time"

// BankAccount represents a bank account item
type BankAccount struct {
	ID             int       `db:"id" json:"id"`
	CompanyID      string    `db:"company_id" json:"company_id"`
	Name           string    `db:"name" json:"name"`
	IsPrimary      bool      `db:"is_primary" json:"is_primary"`
	BankCode       string    `db:"bank_code" json:"bank_code"`
	Agency         string    `db:"agency" json:"agency"`
	AgencyDigit    string    `db:"agency_digit" json:"agency_digit"`
	Account        string    `db:"account" json:"account"`
	AccountDigit   string    `db:"account_digit" json:"account_digit"`
	Type           int       `db:"type" json:"-"`
	DocumentType   string    `db:"document_type" json:"document_type"`
	DocumentNumber string    `db:"document_number" json:"document_number"`
	LegalName      string    `db:"legal_name" json:"legal_name"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}

// LevpayToken represents return of levpay tokens
type LevpayToken struct {
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
}

// LevpayBank represents a return of levpay banks information
type LevpayBank struct {
	UUID                 string      `json:"uuid"`
	Name                 string      `json:"name"`
	Slug                 string      `json:"slug"`
	AccountOwner         string      `json:"account_owner"`
	AccountOwnerDocument string      `json:"account_owner_document"`
	AccountAgency        string      `json:"account_agency"`
	AccountNumber        string      `json:"account_number"`
	Description          interface{} `json:"description"`
	Data                 struct{}    `json:"data"`
	AssetURL             string      `json:"asset_url"`
}

// LevpayOrderData contains all fields that must be sent to Levpay to create a new order
type LevpayOrderData struct {
	PaymentMethod    string             `json:"payment_method"`
	Description      string             `json:"description"`
	PartnerReference string             `json:"partner_reference"`
	BankSlug         string             `json:"bank_slug"`
	Amount           float64            `json:"amount"`
	Expiration       int                `json:"expiration"`
	Data             struct{}           `json:"data"`
	Customer         LevpayCustomerData `json:"customer"`
}

// LevpayCustomerData contains all fields that a customer may have when creating a new order at Levpay
type LevpayCustomerData struct {
	Name           string   `json:"name"`
	DocumentNumber string   `json:"document_number"`
	PhoneNumber    string   `json:"phone_number"`
	Email          string   `json:"email"`
	Data           struct{} `json:"data"`
}

// LevpayOrder represents a created order at Levpay with its payment URL if available
type LevpayOrder struct {
	CreatedAt        string `db:"created" json:"created_at"`
	ExpiresAt        string `db:"expires" json:"expires_at"`
	PartnerReference string `db:"-" json:"partner_reference"`
	URL              string `db:"print_url" json:"url"`
	UUID             string `db:"uuid" json:"uuid"`
}

// LevpayOrderStatus represents a order status return
type LevpayOrderStatus struct {
	UUID   string `db:"uuid" json:"uuid"`
	Status string `db:"status" json:"status"`
}
