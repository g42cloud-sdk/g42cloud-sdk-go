package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpListRequestBody struct {
	Ipgroup *UpdateIpListOption `json:"ipgroup,omitempty"`
}

func (o UpdateIpListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpListRequestBody", string(data)}, " ")
}
