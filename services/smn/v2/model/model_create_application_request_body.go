package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateApplicationRequestBody struct {
	Name string `json:"name"`

	Platform string `json:"platform"`

	PlatformPrincipal string `json:"platform_principal"`

	PlatformCredential string `json:"platform_credential"`
}

func (o CreateApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequestBody", string(data)}, " ")
}
