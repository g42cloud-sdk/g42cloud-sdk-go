package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteMembersResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Members        *[]BatchDeleteMembersState `json:"members,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o BatchDeleteMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersResponse", string(data)}, " ")
}
