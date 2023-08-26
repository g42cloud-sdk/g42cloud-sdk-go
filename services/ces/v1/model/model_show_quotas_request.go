package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowQuotasRequest struct {
	ContentType string `json:"Content-Type"`
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
