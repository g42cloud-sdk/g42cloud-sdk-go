package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ModifyDomainConfigRequestBody struct {
	Configs *Configs `json:"configs,omitempty"`
}

func (o ModifyDomainConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDomainConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyDomainConfigRequestBody", string(data)}, " ")
}
