package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PrePaidServerSecurityGroup struct {
	Id *string `json:"id,omitempty"`
}

func (o PrePaidServerSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerSecurityGroup struct{}"
	}

	return strings.Join([]string{"PrePaidServerSecurityGroup", string(data)}, " ")
}
