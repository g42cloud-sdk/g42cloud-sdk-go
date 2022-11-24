package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowLogFile struct {
	FileName string `json:"file_name"`

	FileSize string `json:"file_size"`
}

func (o SlowLogFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogFile struct{}"
	}

	return strings.Join([]string{"SlowLogFile", string(data)}, " ")
}
