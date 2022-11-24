package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAddonInstanceResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *InstanceSpec `json:"spec,omitempty"`

	Status         *AddonInstanceStatus `json:"status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateAddonInstanceResponse", string(data)}, " ")
}
