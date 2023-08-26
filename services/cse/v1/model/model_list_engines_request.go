package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListEnginesRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *string `json:"limit,omitempty"`
}

func (o ListEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesRequest struct{}"
	}

	return strings.Join([]string{"ListEnginesRequest", string(data)}, " ")
}
