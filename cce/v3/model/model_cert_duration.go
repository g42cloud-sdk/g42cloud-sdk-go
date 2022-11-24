package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertDuration struct {
	Duration int32 `json:"duration"`
}

func (o CertDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertDuration struct{}"
	}

	return strings.Join([]string{"CertDuration", string(data)}, " ")
}
