package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadKieResponseBodyMetadata struct {
	Version *string `json:"version,omitempty"`

	Annotations *interface{} `json:"annotations,omitempty"`
}

func (o DownloadKieResponseBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieResponseBodyMetadata struct{}"
	}

	return strings.Join([]string{"DownloadKieResponseBodyMetadata", string(data)}, " ")
}
