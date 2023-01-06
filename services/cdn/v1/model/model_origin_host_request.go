package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type OriginHostRequest struct {
	OriginHost *OriginHostBody `json:"origin_host"`
}

func (o OriginHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginHostRequest struct{}"
	}

	return strings.Join([]string{"OriginHostRequest", string(data)}, " ")
}
