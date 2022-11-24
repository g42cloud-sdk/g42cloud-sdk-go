package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDbPortRequest struct {
	Port int32 `json:"port"`
}

func (o UpdateDbPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbPortRequest struct{}"
	}

	return strings.Join([]string{"UpdateDbPortRequest", string(data)}, " ")
}
