package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SlowlogDownloadRequest struct {
	FileName *string `json:"file_name,omitempty"`
}

func (o SlowlogDownloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogDownloadRequest struct{}"
	}

	return strings.Join([]string{"SlowlogDownloadRequest", string(data)}, " ")
}
