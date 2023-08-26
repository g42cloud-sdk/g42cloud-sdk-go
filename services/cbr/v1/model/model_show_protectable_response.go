package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowProtectableResponse struct {
	Instance       *ProtectablesResp `json:"instance,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowProtectableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectableResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectableResponse", string(data)}, " ")
}
