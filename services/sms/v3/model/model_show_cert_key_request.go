package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowCertKeyRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowCertKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowCertKeyRequest", string(data)}, " ")
}
