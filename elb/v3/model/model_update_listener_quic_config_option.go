package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateListenerQuicConfigOption struct {
	QuicListenerId *string `json:"quic_listener_id,omitempty"`

	EnableQuicUpgrade *bool `json:"enable_quic_upgrade,omitempty"`
}

func (o UpdateListenerQuicConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerQuicConfigOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerQuicConfigOption", string(data)}, " ")
}
