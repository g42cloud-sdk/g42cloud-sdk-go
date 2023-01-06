package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdatePoliciesPriorityRequest struct {
	Body *BatchUpdatePoliciesPriorityRequestBody `json:"body,omitempty"`
}

func (o BatchUpdatePoliciesPriorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityRequest", string(data)}, " ")
}
