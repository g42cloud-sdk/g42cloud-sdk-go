package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateEncryptTaskRequest struct {
	Body *CreateEncryptReq `json:"body,omitempty"`
}

func (o CreateEncryptTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateEncryptTaskRequest", string(data)}, " ")
}
