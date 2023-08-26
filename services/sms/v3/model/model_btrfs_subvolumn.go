package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BtrfsSubvolumn struct {
	Uuid string `json:"uuid"`

	IsSnapshot string `json:"is_snapshot"`

	SubvolId string `json:"subvol_id"`

	ParentId string `json:"parent_id"`

	SubvolName string `json:"subvol_name"`

	SubvolMountPath string `json:"subvol_mount_path"`
}

func (o BtrfsSubvolumn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BtrfsSubvolumn struct{}"
	}

	return strings.Join([]string{"BtrfsSubvolumn", string(data)}, " ")
}
