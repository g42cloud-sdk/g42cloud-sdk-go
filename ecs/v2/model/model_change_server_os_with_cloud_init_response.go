package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeServerOsWithCloudInitResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeServerOsWithCloudInitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithCloudInitResponse struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithCloudInitResponse", string(data)}, " ")
}
