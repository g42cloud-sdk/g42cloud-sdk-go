package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListServersResponse struct {
	Count *int32 `json:"count,omitempty"`

	SourceServers  *[]SourceServersResponseBody `json:"source_servers,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersResponse struct{}"
	}

	return strings.Join([]string{"ListServersResponse", string(data)}, " ")
}
