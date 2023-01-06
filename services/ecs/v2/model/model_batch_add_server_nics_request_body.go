package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BatchAddServerNicsRequestBody struct {
	Nics []BatchAddServerNicOption `json:"nics"`
}

func (o BatchAddServerNicsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerNicsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddServerNicsRequestBody", string(data)}, " ")
}
