package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMemberRequest struct {
	PoolId string `json:"pool_id"`

	Body *CreateMemberRequestBody `json:"body,omitempty"`
}

func (o CreateMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberRequest struct{}"
	}

	return strings.Join([]string{"CreateMemberRequest", string(data)}, " ")
}
