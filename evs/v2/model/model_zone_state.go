package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ZoneState struct {
	Available *bool `json:"available,omitempty"`
}

func (o ZoneState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ZoneState struct{}"
	}

	return strings.Join([]string{"ZoneState", string(data)}, " ")
}
