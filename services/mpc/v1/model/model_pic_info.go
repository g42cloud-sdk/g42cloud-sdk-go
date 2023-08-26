package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PicInfo struct {
	PicName *string `json:"pic_name,omitempty"`
}

func (o PicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PicInfo struct{}"
	}

	return strings.Join([]string{"PicInfo", string(data)}, " ")
}
