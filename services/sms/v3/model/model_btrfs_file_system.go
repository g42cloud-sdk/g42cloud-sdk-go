package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BtrfsFileSystem struct {
	Name string `json:"name"`

	Label string `json:"label"`

	Uuid string `json:"uuid"`

	Device string `json:"device"`

	Size int64 `json:"size"`

	Nodesize int64 `json:"nodesize"`

	Sectorsize int32 `json:"sectorsize"`

	DataProfile string `json:"data_profile"`

	SystemProfile string `json:"system_profile"`

	MetadataProfile string `json:"metadata_profile"`

	GlobalReserve1 string `json:"global_reserve1"`

	GVolUsedSize int64 `json:"g_vol_used_size"`

	DefaultSubvolid string `json:"default_subvolid"`

	DefaultSubvolName string `json:"default_subvol_name"`

	DefaultSubvolMountpath string `json:"default_subvol_mountpath"`

	Subvolumn []BtrfsSubvolumn `json:"subvolumn"`
}

func (o BtrfsFileSystem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BtrfsFileSystem struct{}"
	}

	return strings.Join([]string{"BtrfsFileSystem", string(data)}, " ")
}
