package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowMembersDetailResponse struct {
	Members *[]Member `json:"members,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMembersDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMembersDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMembersDetailResponse", string(data)}, " ")
}
