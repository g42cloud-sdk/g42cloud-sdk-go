package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTemplateGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateGroupResponse", string(data)}, " ")
}
