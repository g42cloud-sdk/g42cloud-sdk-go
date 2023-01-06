package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ShowLogsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Logs           *[]LogObject `json:"logs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsResponse struct{}"
	}

	return strings.Join([]string{"ShowLogsResponse", string(data)}, " ")
}
