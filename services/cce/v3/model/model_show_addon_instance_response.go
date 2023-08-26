package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowAddonInstanceResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *InstanceSpec `json:"spec,omitempty"`

	Status         *AddonInstanceStatus `json:"status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowAddonInstanceResponse", string(data)}, " ")
}
