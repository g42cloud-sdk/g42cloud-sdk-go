package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type SmnUrn struct {
}

func (o SmnUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnUrn struct{}"
	}

	return strings.Join([]string{"SmnUrn", string(data)}, " ")
}
