package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateResourceGroupResponse struct {
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceGroupResponse", string(data)}, " ")
}
