package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OutputThumbnailPara struct {
	TotalPictures *int32 `json:"total_pictures,omitempty"`

	Width *int32 `json:"width,omitempty"`

	Height *int32 `json:"height,omitempty"`

	FileName *string `json:"file_name,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`
}

func (o OutputThumbnailPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputThumbnailPara struct{}"
	}

	return strings.Join([]string{"OutputThumbnailPara", string(data)}, " ")
}
