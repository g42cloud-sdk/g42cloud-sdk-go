package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountPreoccupyIpNumResponse struct {
	PreoccupyIp *PreoccupyIp `json:"preoccupy_ip,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountPreoccupyIpNumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPreoccupyIpNumResponse struct{}"
	}

	return strings.Join([]string{"CountPreoccupyIpNumResponse", string(data)}, " ")
}
