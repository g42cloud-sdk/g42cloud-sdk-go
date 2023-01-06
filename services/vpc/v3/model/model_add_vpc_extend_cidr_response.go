package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddVpcExtendCidrResponse struct {
	Vpc *Vpc `json:"vpc,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddVpcExtendCidrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVpcExtendCidrResponse struct{}"
	}

	return strings.Join([]string{"AddVpcExtendCidrResponse", string(data)}, " ")
}
