package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowPoolRequest struct {
	PoolId string `json:"pool_id"`
}

func (o ShowPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolRequest", string(data)}, " ")
}
