package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ExportImageRequest struct {
	ImageId string `json:"image_id"`

	Body *ExportImageRequestBody `json:"body,omitempty"`
}

func (o ExportImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageRequest struct{}"
	}

	return strings.Join([]string{"ExportImageRequest", string(data)}, " ")
}
