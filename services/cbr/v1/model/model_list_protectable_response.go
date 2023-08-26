package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListProtectableResponse struct {
	Instances      *[]ProtectablesResp `json:"instances,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListProtectableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectableResponse struct{}"
	}

	return strings.Join([]string{"ListProtectableResponse", string(data)}, " ")
}
