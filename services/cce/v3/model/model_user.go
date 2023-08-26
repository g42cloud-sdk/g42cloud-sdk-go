package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type User struct {
	ClientCertificateData *string `json:"client-certificate-data,omitempty"`

	ClientKeyData *string `json:"client-key-data,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}
