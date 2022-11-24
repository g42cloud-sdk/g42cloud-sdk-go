package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAddonTemplatesResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Items          *[]AddonTemplate `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAddonTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAddonTemplatesResponse", string(data)}, " ")
}
