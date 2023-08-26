package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListResourceInstancesRequest struct {
	ResourceType string `json:"resource_type"`

	Body *ListInstanceRequestBody `json:"body,omitempty"`
}

func (o ListResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesRequest", string(data)}, " ")
}
