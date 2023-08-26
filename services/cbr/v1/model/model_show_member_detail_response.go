package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowMemberDetailResponse struct {
	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMemberDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMemberDetailResponse", string(data)}, " ")
}
