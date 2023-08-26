package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateMpeCallBackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMpeCallBackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMpeCallBackResponse struct{}"
	}

	return strings.Join([]string{"CreateMpeCallBackResponse", string(data)}, " ")
}
