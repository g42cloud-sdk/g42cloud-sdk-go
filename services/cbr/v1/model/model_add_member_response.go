package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type AddMemberResponse struct {
	Members *[]Member `json:"members,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o AddMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMemberResponse struct{}"
	}

	return strings.Join([]string{"AddMemberResponse", string(data)}, " ")
}
