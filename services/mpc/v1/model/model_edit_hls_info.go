package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type EditHlsInfo struct {
	Interval *int32 `json:"interval,omitempty"`
}

func (o EditHlsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditHlsInfo struct{}"
	}

	return strings.Join([]string{"EditHlsInfo", string(data)}, " ")
}
