package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaShowKeypairRequest struct {
	KeypairName string `json:"keypair_name"`

	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`
}

func (o NovaShowKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaShowKeypairRequest", string(data)}, " ")
}
