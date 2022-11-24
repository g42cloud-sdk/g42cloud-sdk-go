package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAddonInstancesRequest struct {
	AddonTemplateName *string `json:"addon_template_name,omitempty"`

	ClusterId string `json:"cluster_id"`
}

func (o ListAddonInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAddonInstancesRequest", string(data)}, " ")
}
