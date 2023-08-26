package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SourceServerByTask struct {
	Id string `json:"id"`
}

func (o SourceServerByTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceServerByTask struct{}"
	}

	return strings.Join([]string{"SourceServerByTask", string(data)}, " ")
}
