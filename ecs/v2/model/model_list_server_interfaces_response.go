package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListServerInterfacesResponse struct {
	AttachableQuantity *InterfaceAttachableQuantity `json:"attachableQuantity,omitempty"`

	InterfaceAttachments *[]InterfaceAttachment `json:"interfaceAttachments,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ListServerInterfacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerInterfacesResponse struct{}"
	}

	return strings.Join([]string{"ListServerInterfacesResponse", string(data)}, " ")
}
