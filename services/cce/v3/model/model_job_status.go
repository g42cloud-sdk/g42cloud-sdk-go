package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type JobStatus struct {
	Phase *string `json:"phase,omitempty"`

	Reason *string `json:"reason,omitempty"`
}

func (o JobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobStatus struct{}"
	}

	return strings.Join([]string{"JobStatus", string(data)}, " ")
}
