package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEngineResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	JobId          *int32 `json:"job_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEngineResponse struct{}"
	}

	return strings.Join([]string{"CreateEngineResponse", string(data)}, " ")
}
