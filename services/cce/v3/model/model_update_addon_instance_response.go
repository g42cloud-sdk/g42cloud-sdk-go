package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateAddonInstanceResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *InstanceSpec `json:"spec,omitempty"`

	Status         *AddonInstanceStatus `json:"status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateAddonInstanceResponse", string(data)}, " ")
}
