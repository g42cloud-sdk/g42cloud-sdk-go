package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListServersDetailsResponse struct {
	Count *int32 `json:"count,omitempty"`

	Servers        *[]ServerDetail `json:"servers,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListServersDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServersDetailsResponse", string(data)}, " ")
}
