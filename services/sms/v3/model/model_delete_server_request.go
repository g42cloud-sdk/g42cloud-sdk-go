package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteServerRequest struct {
	SourceId string `json:"source_id"`
}

func (o DeleteServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerRequest", string(data)}, " ")
}
