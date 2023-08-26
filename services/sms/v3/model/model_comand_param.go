package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ComandParam struct {
	TaskId *string `json:"task_id,omitempty"`

	Bucket *string `json:"bucket,omitempty"`
}

func (o ComandParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComandParam struct{}"
	}

	return strings.Join([]string{"ComandParam", string(data)}, " ")
}
