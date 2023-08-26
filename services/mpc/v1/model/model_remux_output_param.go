package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type RemuxOutputParam struct {
	Format *string `json:"format,omitempty"`

	SegmentDuration *int32 `json:"segment_duration,omitempty"`

	RemoveMeta *bool `json:"remove_meta,omitempty"`
}

func (o RemuxOutputParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemuxOutputParam struct{}"
	}

	return strings.Join([]string{"RemuxOutputParam", string(data)}, " ")
}
