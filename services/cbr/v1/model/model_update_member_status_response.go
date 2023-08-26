package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateMemberStatusResponse struct {
	Member         *Member `json:"member,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMemberStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateMemberStatusResponse", string(data)}, " ")
}
