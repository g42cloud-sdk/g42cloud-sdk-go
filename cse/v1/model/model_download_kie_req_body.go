package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadKieReqBody struct {
	Ids []string `json:"ids"`
}

func (o DownloadKieReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieReqBody struct{}"
	}

	return strings.Join([]string{"DownloadKieReqBody", string(data)}, " ")
}
