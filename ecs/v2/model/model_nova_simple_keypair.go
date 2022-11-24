package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NovaSimpleKeypair struct {
	Fingerprint string `json:"fingerprint"`

	Name string `json:"name"`

	PublicKey string `json:"public_key"`

	Type *string `json:"type,omitempty"`
}

func (o NovaSimpleKeypair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSimpleKeypair struct{}"
	}

	return strings.Join([]string{"NovaSimpleKeypair", string(data)}, " ")
}
