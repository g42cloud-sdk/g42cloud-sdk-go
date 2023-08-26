package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateTemplateGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTemplateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateGroupResponse", string(data)}, " ")
}
