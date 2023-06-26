package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteExtractTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExtractTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExtractTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteExtractTaskResponse", string(data)}, " ")
}
