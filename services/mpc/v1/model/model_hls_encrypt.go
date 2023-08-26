package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type HlsEncrypt struct {
	Key string `json:"key"`

	Url string `json:"url"`

	Iv *string `json:"iv,omitempty"`

	Algorithm *string `json:"algorithm,omitempty"`
}

func (o HlsEncrypt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HlsEncrypt struct{}"
	}

	return strings.Join([]string{"HlsEncrypt", string(data)}, " ")
}
