package namespace

import (
	"github.com/lucidstacklabs/namefinder/internal/pkg/apikey"
)

type ApiKeyAccessResponse struct {
	Access *ApiKeyAccess  `json:"access"`
	ApiKey *apikey.ApiKey `json:"api_key"`
}
