package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateCopyStateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCopyStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCopyStateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCopyStateResponse", string(data)}, " ")
}
