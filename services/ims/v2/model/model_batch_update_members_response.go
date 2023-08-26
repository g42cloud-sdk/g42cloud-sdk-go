package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateMembersResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpdateMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersResponse", string(data)}, " ")
}
