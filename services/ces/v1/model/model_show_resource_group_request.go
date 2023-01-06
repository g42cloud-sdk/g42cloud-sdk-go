package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowResourceGroupRequest struct {
	ContentType string `json:"Content-Type"`

	GroupId string `json:"group_id"`

	Status *string `json:"status,omitempty"`

	Namespace *string `json:"namespace,omitempty"`

	Dname *string `json:"dname,omitempty"`

	Start *string `json:"start,omitempty"`

	Limit *string `json:"limit,omitempty"`
}

func (o ShowResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupRequest", string(data)}, " ")
}
