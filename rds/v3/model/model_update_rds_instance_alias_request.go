package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRdsInstanceAliasRequest struct {
	Alias string `json:"alias"`
}

func (o UpdateRdsInstanceAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRdsInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateRdsInstanceAliasRequest", string(data)}, " ")
}
