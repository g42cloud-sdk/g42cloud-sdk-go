package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type BackupSync struct {
	BackupId string `json:"backup_id"`

	BackupName string `json:"backup_name"`

	BucketName string `json:"bucket_name"`

	ImagePath string `json:"image_path"`

	ResourceId string `json:"resource_id"`

	ResourceName string `json:"resource_name"`

	ResourceType string `json:"resource_type"`

	CreatedAt int32 `json:"created_at"`
}

func (o BackupSync) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupSync struct{}"
	}

	return strings.Join([]string{"BackupSync", string(data)}, " ")
}
