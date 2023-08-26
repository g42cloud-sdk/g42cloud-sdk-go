package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type VersionItem struct {
	Id string `json:"id"`

	MinVersion string `json:"min_version"`

	Status string `json:"status"`

	Updated string `json:"updated"`

	Version string `json:"version"`

	Links []LinksItem `json:"links"`
}

func (o VersionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionItem struct{}"
	}

	return strings.Join([]string{"VersionItem", string(data)}, " ")
}
