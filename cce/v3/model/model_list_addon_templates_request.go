package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAddonTemplatesRequest struct {
	AddonTemplateName *string `json:"addon_template_name,omitempty"`

	BaseUpdateAddonVersion *string `json:"base_update_addon_version,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	Newest *string `json:"newest,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (o ListAddonTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAddonTemplatesRequest", string(data)}, " ")
}
