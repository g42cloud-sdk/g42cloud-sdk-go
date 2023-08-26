package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Encryption struct {
	HlsEncrypt *HlsEncrypt `json:"hls_encrypt,omitempty"`
}

func (o Encryption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Encryption struct{}"
	}

	return strings.Join([]string{"Encryption", string(data)}, " ")
}
