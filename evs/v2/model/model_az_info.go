package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AzInfo struct {
	ZoneName string `json:"zoneName"`

	ZoneState *ZoneState `json:"zoneState"`
}

func (o AzInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzInfo struct{}"
	}

	return strings.Join([]string{"AzInfo", string(data)}, " ")
}
