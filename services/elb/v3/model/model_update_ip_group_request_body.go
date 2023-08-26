package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpGroupRequestBody struct {
	Ipgroup *UpdateIpGroupOption `json:"ipgroup"`
}

func (o UpdateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequestBody", string(data)}, " ")
}
