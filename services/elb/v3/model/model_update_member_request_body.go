package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateMemberRequestBody struct {
	Member *UpdateMemberOption `json:"member"`
}

func (o UpdateMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMemberRequestBody", string(data)}, " ")
}
