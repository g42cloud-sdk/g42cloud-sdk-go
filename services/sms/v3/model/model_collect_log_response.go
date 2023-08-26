package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CollectLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CollectLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectLogResponse struct{}"
	}

	return strings.Join([]string{"CollectLogResponse", string(data)}, " ")
}
