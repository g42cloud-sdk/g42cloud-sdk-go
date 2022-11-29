package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRedirectPoolsConfig struct {
	PoolId string `json:"pool_id"`

	Weight int32 `json:"weight"`
}

func (o CreateRedirectPoolsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedirectPoolsConfig struct{}"
	}

	return strings.Join([]string{"CreateRedirectPoolsConfig", string(data)}, " ")
}
