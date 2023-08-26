package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UrlAuthGetBody struct {
	Status string `json:"status"`

	Type *string `json:"type,omitempty"`

	TimeFormat *string `json:"time_format,omitempty"`

	ExpireTime *int32 `json:"expire_time,omitempty"`
}

func (o UrlAuthGetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlAuthGetBody struct{}"
	}

	return strings.Join([]string{"UrlAuthGetBody", string(data)}, " ")
}
