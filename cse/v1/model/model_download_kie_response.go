package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadKieResponse struct {
	Metadata *DownloadKieResponseBodyMetadata `json:"metadata,omitempty"`

	Data           *[]CreateKieReq `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o DownloadKieResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieResponse struct{}"
	}

	return strings.Join([]string{"DownloadKieResponse", string(data)}, " ")
}
