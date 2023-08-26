package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BandwidthRef struct {
	Id string `json:"id"`
}

func (o BandwidthRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthRef struct{}"
	}

	return strings.Join([]string{"BandwidthRef", string(data)}, " ")
}
