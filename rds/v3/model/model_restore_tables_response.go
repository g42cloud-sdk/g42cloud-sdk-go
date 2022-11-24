package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreTablesResponse struct {
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTablesResponse struct{}"
	}

	return strings.Join([]string{"RestoreTablesResponse", string(data)}, " ")
}
