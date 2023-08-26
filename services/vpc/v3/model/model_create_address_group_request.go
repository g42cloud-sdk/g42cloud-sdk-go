package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateAddressGroupRequest struct {
	Body *CreateAddressGroupRequestBody `json:"body,omitempty"`
}

func (o CreateAddressGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddressGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateAddressGroupRequest", string(data)}, " ")
}
