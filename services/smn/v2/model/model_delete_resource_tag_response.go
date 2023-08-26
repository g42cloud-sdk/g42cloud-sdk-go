package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagResponse", string(data)}, " ")
}
