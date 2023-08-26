package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CinderListQuotasResponse struct {
	QuotaSet       *QuotaList `json:"quota_set,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CinderListQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListQuotasResponse struct{}"
	}

	return strings.Join([]string{"CinderListQuotasResponse", string(data)}, " ")
}
