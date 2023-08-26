package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type DeleteSnapshotResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotResponse struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotResponse", string(data)}, " ")
}
