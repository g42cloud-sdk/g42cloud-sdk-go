package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ImageData struct {
	ImageId *string `json:"image_id,omitempty"`
}

func (o ImageData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageData struct{}"
	}

	return strings.Join([]string{"ImageData", string(data)}, " ")
}
