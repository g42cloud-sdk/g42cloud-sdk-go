package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListApiVersionResponse struct {
	Versions       *[]ApiVersion `json:"versions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionResponse", string(data)}, " ")
}
