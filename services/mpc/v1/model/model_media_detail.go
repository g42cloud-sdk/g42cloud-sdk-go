package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type MediaDetail struct {
	Features *[]string `json:"features,omitempty"`

	OriginPara *OriginPara `json:"origin_para,omitempty"`

	OutputVideoParas *[]OutputVideoPara `json:"output_video_paras,omitempty"`

	OutputThumbnailPara *OutputThumbnailPara `json:"output_thumbnail_para,omitempty"`

	OutputWatermarkParas *OutputWatermarkPara `json:"output_watermark_paras,omitempty"`
}

func (o MediaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MediaDetail struct{}"
	}

	return strings.Join([]string{"MediaDetail", string(data)}, " ")
}
