package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowPassphraseRequest struct {
	TaskId string `json:"task_id"`
}

func (o ShowPassphraseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPassphraseRequest struct{}"
	}

	return strings.Join([]string{"ShowPassphraseRequest", string(data)}, " ")
}
