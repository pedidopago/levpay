package transfer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pedidopago/levpay/pkg/levpay"
)

// API struct contain levpay config
type API struct {
	Config *levpay.Config
}

// New start the instance of levpay api
func New(cfg *levpay.Config) *API {
	return &API{
		Config: cfg,
	}
}

// LevpayAvailableAccounts return an array of accounts available for the given domain.
// These accounts are fetched from Levpay endpoint using GetLevpayKeys to determine
// which keys should be used for given domain
func (api *API) LevpayAvailableAccounts(companyID string) ([]*levpay.BankAccount, error) {
	response, err := api.Config.Do(http.MethodGet, "/instance/levpay/banks/", nil)
	if err != nil {
		fmt.Println("[LEVPAY] GetLevpayAvailableAccounts e2", companyID, err.Error())
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[LEVPAY] GetLevpayAvailableAccounts e3", companyID, err.Error())
		return nil, err
	}

	var accounts []*levpay.BankAccount
	var banks []levpay.LevpayBank
	err = json.Unmarshal(responseBody, &banks)
	if err != nil {
		fmt.Println("[LEVPAY] GetLevpayAvailableAccounts e4", companyID, err.Error(), string(responseBody))
		return nil, err
	}
	for index, bank := range banks {
		var account levpay.BankAccount
		account.ID = index + 1
		account.CompanyID = companyID
		account.Name = bank.Name
		account.IsPrimary = false
		account.BankCode = bank.Slug
		account.Agency = bank.AccountAgency
		account.AgencyDigit = ""
		account.Account = bank.AccountNumber
		account.AccountDigit = ""
		account.DocumentType = "cnpj"
		account.DocumentNumber = bank.AccountOwnerDocument
		account.LegalName = bank.AccountOwner

		accounts = append(accounts, &account)
	}

	fmt.Println("[LEVPAY] GetLevpayAvailableAccounts", companyID, accounts)

	return accounts, nil
}

// LevpayCreatePayment create a new payment at Levpay and return a LevpayOrder object
// containing order details and the payment URL (if available)
func (api *API) LevpayCreatePayment(domainID int, orderData levpay.LevpayOrderData) (levpay.LevpayOrder, error) {
	var order levpay.LevpayOrder

	response, err := api.Config.Do(http.MethodPost, "/instance/levpay/checkout/", orderData)
	if err != nil {
		fmt.Println("[LEVPAY] CreateLevpayPayment e1", domainID, err.Error())
		return order, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[LEVPAY] CreateLevpayPayment e2", domainID, err.Error())
		return order, err
	}

	err = json.Unmarshal(responseBody, &order)
	if err != nil {
		fmt.Println("[LEVPAY] CreateLevpayPayment e3", domainID, err.Error(), string(responseBody))
		return order, err
	}

	return order, nil
}

// LevpayOrderStatus create a request to obtain a payment order status
// containing order details about status and the uuid of order
func (api *API) LevpayOrderStatus(domainID int, UUID string) (levpay.LevpayOrderStatus, error) {
	var status levpay.LevpayOrderStatus

	response, err := api.Config.Do(http.MethodGet, "/instance/levpay/status/"+UUID, nil)
	if err != nil {
		fmt.Println("[LEVPAY] CreateLevpayPayment e1", domainID, err.Error())
		return status, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[LEVPAY] CreateLevpayPayment e2", domainID, err.Error())
		return status, err
	}

	err = json.Unmarshal(responseBody, &status)
	if err != nil {
		fmt.Println("[LEVPAY] CreateLevpayPayment e3", domainID, err.Error(), string(responseBody))
		return status, err
	}

	return status, nil
}
