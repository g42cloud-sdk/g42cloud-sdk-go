package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TagsResp struct {
	Key *string `json:"key,omitempty"`

	Values *string `json:"values,omitempty"`
}

func (o TagsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsResp struct{}"
	}

	return strings.Join([]string{"TagsResp", string(data)}, " ")
}
