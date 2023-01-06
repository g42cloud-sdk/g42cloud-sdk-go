package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateApplicationRequestBody struct {
	PlatformPrincipal string `json:"platform_principal"`

	PlatformCredential string `json:"platform_credential"`
}

func (o UpdateApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationRequestBody", string(data)}, " ")
}
