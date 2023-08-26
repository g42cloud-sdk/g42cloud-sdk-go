package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RefreshTaskRequest struct {
	RefreshTask *RefreshTaskRequestBody `json:"refresh_task"`
}

func (o RefreshTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshTaskRequest struct{}"
	}

	return strings.Join([]string{"RefreshTaskRequest", string(data)}, " ")
}
