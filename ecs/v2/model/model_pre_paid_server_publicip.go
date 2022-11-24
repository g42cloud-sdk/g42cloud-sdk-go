package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrePaidServerPublicip struct {
	Id *string `json:"id,omitempty"`

	Eip *PrePaidServerEip `json:"eip,omitempty"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o PrePaidServerPublicip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerPublicip struct{}"
	}

	return strings.Join([]string{"PrePaidServerPublicip", string(data)}, " ")
}
