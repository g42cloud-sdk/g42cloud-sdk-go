package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteTagsOption struct {
	Key string `json:"key"`
}

func (o DeleteTagsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagsOption struct{}"
	}

	return strings.Join([]string{"DeleteTagsOption", string(data)}, " ")
}
