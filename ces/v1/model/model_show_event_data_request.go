package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowEventDataRequest struct {
	ContentType string `json:"Content-Type"`

	Namespace string `json:"namespace"`

	Dim0 string `json:"dim.0"`

	Dim1 *string `json:"dim.1,omitempty"`

	Dim2 *string `json:"dim.2,omitempty"`

	Dim3 *string `json:"dim.3,omitempty"`

	Type string `json:"type"`

	From int64 `json:"from"`

	To int64 `json:"to"`
}

func (o ShowEventDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventDataRequest struct{}"
	}

	return strings.Join([]string{"ShowEventDataRequest", string(data)}, " ")
}
