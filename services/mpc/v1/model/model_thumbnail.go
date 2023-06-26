package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type Thumbnail struct {
	Tar *int32 `json:"tar,omitempty"`

	Out *ObsObjInfo `json:"out,omitempty"`

	Params *ThumbnailPara `json:"params"`
}

func (o Thumbnail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Thumbnail struct{}"
	}

	return strings.Join([]string{"Thumbnail", string(data)}, " ")
}
