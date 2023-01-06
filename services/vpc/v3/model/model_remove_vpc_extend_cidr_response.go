package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RemoveVpcExtendCidrResponse struct {
	Vpc *Vpc `json:"vpc,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoveVpcExtendCidrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVpcExtendCidrResponse struct{}"
	}

	return strings.Join([]string{"RemoveVpcExtendCidrResponse", string(data)}, " ")
}
