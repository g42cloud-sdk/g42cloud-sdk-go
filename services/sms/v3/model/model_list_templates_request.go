package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListTemplatesRequest struct {
	Name *string `json:"name,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Region *string `json:"region,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
