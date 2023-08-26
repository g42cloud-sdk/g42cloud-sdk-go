package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type TagsReq struct {
	Key string `json:"key"`

	Values []string `json:"values"`
}

func (o TagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsReq struct{}"
	}

	return strings.Join([]string{"TagsReq", string(data)}, " ")
}
