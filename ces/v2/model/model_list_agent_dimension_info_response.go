package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAgentDimensionInfoResponse struct {
	Dimensions *[]AgentDimension `json:"dimensions,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentDimensionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentDimensionInfoResponse struct{}"
	}

	return strings.Join([]string{"ListAgentDimensionInfoResponse", string(data)}, " ")
}
