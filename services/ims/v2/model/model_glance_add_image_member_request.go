package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type GlanceAddImageMemberRequest struct {
	ImageId string `json:"image_id"`

	Body *GlanceAddImageMemberRequestBody `json:"body,omitempty"`
}

func (o GlanceAddImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceAddImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceAddImageMemberRequest", string(data)}, " ")
}
