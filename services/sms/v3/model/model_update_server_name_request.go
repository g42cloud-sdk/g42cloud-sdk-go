package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerNameRequest struct {
	SourceId string `json:"source_id"`

	Body *PutSourceServerBody `json:"body,omitempty"`
}

func (o UpdateServerNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateServerNameRequest", string(data)}, " ")
}
