package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListVersionResponse struct {
	Version        *VersionItem `json:"version,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionResponse struct{}"
	}

	return strings.Join([]string{"ListVersionResponse", string(data)}, " ")
}
