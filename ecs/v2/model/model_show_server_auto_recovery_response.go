package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowServerAutoRecoveryResponse struct {
	SupportAutoRecovery *string `json:"support_auto_recovery,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowServerAutoRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerAutoRecoveryResponse struct{}"
	}

	return strings.Join([]string{"ShowServerAutoRecoveryResponse", string(data)}, " ")
}
