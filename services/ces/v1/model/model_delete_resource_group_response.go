package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteResourceGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceGroupResponse", string(data)}, " ")
}
