package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberStatus struct {
	ListenerId string `json:"listener_id"`

	OperatingStatus string `json:"operating_status"`
}

func (o MemberStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberStatus struct{}"
	}

	return strings.Join([]string{"MemberStatus", string(data)}, " ")
}
