package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteServersRequest struct {
	Body *DeleteIds `json:"body,omitempty"`
}

func (o DeleteServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServersRequest struct{}"
	}

	return strings.Join([]string{"DeleteServersRequest", string(data)}, " ")
}
