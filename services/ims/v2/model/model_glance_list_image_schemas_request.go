package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type GlanceListImageSchemasRequest struct {
}

func (o GlanceListImageSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceListImageSchemasRequest", string(data)}, " ")
}
