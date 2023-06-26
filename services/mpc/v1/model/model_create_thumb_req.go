package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type CreateThumbReq struct {
	Input *ObsObjInfo `json:"input"`

	Output *ObsObjInfo `json:"output"`

	UserData *string `json:"user_data,omitempty"`

	ThumbnailPara *ThumbnailPara `json:"thumbnail_para"`

	Tar *int32 `json:"tar,omitempty"`

	Sync *int32 `json:"sync,omitempty"`

	OriginalDir *int32 `json:"original_dir,omitempty"`
}

func (o CreateThumbReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateThumbReq struct{}"
	}

	return strings.Join([]string{"CreateThumbReq", string(data)}, " ")
}
