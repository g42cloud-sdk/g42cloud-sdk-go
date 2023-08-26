package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type HeaderMap struct {
	ContentDisposition *string `json:"Content-Disposition,omitempty"`

	ContentLanguage *string `json:"Content-Language,omitempty"`

	AccessControlAllowOrigin *string `json:"Access-Control-Allow-Origin,omitempty"`

	AccessControlAllowMethods *string `json:"Access-Control-Allow-Methods,omitempty"`

	AccessControlMaxAge *string `json:"Access-Control-Max-Age,omitempty"`

	AccessControlExposeHeaders *string `json:"Access-Control-Expose-Headers,omitempty"`
}

func (o HeaderMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeaderMap struct{}"
	}

	return strings.Join([]string{"HeaderMap", string(data)}, " ")
}
