package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

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
