package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizePostPaidServerResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizePostPaidServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerResponse struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerResponse", string(data)}, " ")
}
