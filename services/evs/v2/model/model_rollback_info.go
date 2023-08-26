package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RollbackInfo struct {
	VolumeId string `json:"volume_id"`
}

func (o RollbackInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackInfo struct{}"
	}

	return strings.Join([]string{"RollbackInfo", string(data)}, " ")
}
