package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteMergeChannelsTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMergeChannelsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMergeChannelsTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteMergeChannelsTaskResponse", string(data)}, " ")
}
