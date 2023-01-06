package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteImageTagRequest struct {
	ImageId string `json:"image_id"`

	Key string `json:"key"`
}

func (o DeleteImageTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteImageTagRequest", string(data)}, " ")
}
