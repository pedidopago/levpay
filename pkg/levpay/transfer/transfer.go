package transfer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pedidopago/levpay/pkg/levpay"
)

type API struct {
	Config *levpay.Config
}

func New(cfg *levpay.Config) *API {
	return &API{
		Config: cfg,
	}
}

// GetLevpayAvailableAccounts return an array of accounts available for the given domain.
// These accounts are fetched from Levpay endpoint using GetLevpayKeys to determine
// which keys should be used for given domain
func (api *API) GetLevpayAvailableAccounts(domainID int) ([]levpay.BankAccount, error) {
	fmt.Println("Deu bom")
	response, err := api.Config.Do(http.MethodGet, "/instance/levpay/banks/", nil)
	if err != nil {
		fmt.Println("[LEVPAY] GetLevpayAvailableAccounts e2", domainID, err.Error())
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[LEVPAY] GetLevpayAvailableAccounts e3", domainID, err.Error())
		return nil, err
	}

	fmt.Println("RESPOSTA HEADER - ", response.Header)
	fmt.Println("RESPOSTA BODY - ", response.Body)
	var accounts []levpay.BankAccount
	var banks []levpay.LevpayBank
	err = json.Unmarshal(responseBody, &banks)
	if err != nil {
		fmt.Println("[LEVPAY] GetLevpayAvailableAccounts e4", domainID, err.Error(), string(responseBody))
		return nil, err
	}
	for index, bank := range banks {
		var account levpay.BankAccount
		account.ID = index + 1
		account.DomainID = domainID
		account.Name = bank.Name
		account.IsPrimary = false
		account.BankCode = bank.Slug
		account.Agencia = bank.AccountAgency
		account.AgenciaDv = ""
		account.Conta = bank.AccountNumber
		account.ContaDv = ""
		account.DocumentType = "cnpj"
		account.DocumentNumber = bank.AccountOwnerDocument
		account.LegalName = bank.AccountOwner

		accounts = append(accounts, account)
	}

	fmt.Println("[LEVPAY] GetLevpayAvailableAccounts", domainID, accounts)

	return accounts, nil
}
