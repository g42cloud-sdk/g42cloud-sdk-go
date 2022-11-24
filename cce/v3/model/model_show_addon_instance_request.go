package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAddonInstanceRequest struct {
	Id string `json:"id"`

	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o ShowAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowAddonInstanceRequest", string(data)}, " ")
}
