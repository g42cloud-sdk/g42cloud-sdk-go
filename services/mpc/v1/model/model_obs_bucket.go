package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ObsBucket struct {
	Bucket *string `json:"bucket,omitempty"`

	CreationDate *string `json:"creation_date,omitempty"`

	IsAuthorized *int32 `json:"is_authorized,omitempty"`
}

func (o ObsBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBucket struct{}"
	}

	return strings.Join([]string{"ObsBucket", string(data)}, " ")
}
