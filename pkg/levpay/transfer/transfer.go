package transfer

import (
	"fmt"
	"net/http"

	"github.com/pedidopago/levpay/internal/pkg/ww"
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
	resp, err := api.Config.Do(http.MethodGet, "/instance/levpay/banks/", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("TT - ")

	result := make([]levpay.BankAccount, 0)

	if api.Config.Trace {
		if err := ww.UnmarshalTrace(api.Config.Logger, resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction: " + err.Error())
			return nil, err
		}
	} else {
		if err := ww.Unmarshal(resp, result); err != nil {
			api.Config.Logger.Error("could not unmarshal transaction [Put]: " + err.Error())
			return nil, err
		}
	}
	fmt.Println("resultado - ", result)
	return result, nil
}
