package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListVersionRequest struct {
	ApiVersion string `json:"api_version"`
}

func (o ListVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionRequest struct{}"
	}

	return strings.Join([]string{"ListVersionRequest", string(data)}, " ")
}
