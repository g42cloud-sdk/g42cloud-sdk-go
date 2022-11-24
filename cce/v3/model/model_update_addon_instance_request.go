package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAddonInstanceRequest struct {
	Id string `json:"id"`

	Body *InstanceRequest `json:"body,omitempty"`
}

func (o UpdateAddonInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateAddonInstanceRequest", string(data)}, " ")
}
