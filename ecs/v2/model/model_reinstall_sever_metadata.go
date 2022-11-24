package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReinstallSeverMetadata struct {
	UserData *string `json:"user_data,omitempty"`
}

func (o ReinstallSeverMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallSeverMetadata struct{}"
	}

	return strings.Join([]string{"ReinstallSeverMetadata", string(data)}, " ")
}
