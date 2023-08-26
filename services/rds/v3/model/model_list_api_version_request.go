package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type ListApiVersionRequest struct {
}

func (o ListApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionRequest", string(data)}, " ")
}
