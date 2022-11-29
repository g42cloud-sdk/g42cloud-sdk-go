package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
