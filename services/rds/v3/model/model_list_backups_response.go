package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListBackupsResponse struct {
	Backups *[]BackupForList `json:"backups,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListBackupsResponse", string(data)}, " ")
}
