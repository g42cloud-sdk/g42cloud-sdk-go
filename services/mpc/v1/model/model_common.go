package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Common struct {
	Pvc *bool `json:"PVC,omitempty"`

	HlsInterval *int32 `json:"hls_interval,omitempty"`

	DashInterval *int32 `json:"dash_interval,omitempty"`

	PackType int32 `json:"pack_type"`
}

func (o Common) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Common struct{}"
	}

	return strings.Join([]string{"Common", string(data)}, " ")
}
