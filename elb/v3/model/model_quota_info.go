package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaInfo struct {
	QuotaKey string `json:"quota_key"`

	QuotaLimit int32 `json:"quota_limit"`

	Used int32 `json:"used"`

	Unit string `json:"unit"`
}

func (o QuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaInfo struct{}"
	}

	return strings.Join([]string{"QuotaInfo", string(data)}, " ")
}
