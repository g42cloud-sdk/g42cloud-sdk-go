package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"strings"
)

type PostgresqlDatabaseForCreation struct {
	Name string `json:"name"`

	CharacterSet *string `json:"character_set,omitempty"`

	Owner *string `json:"owner,omitempty"`

	Template *string `json:"template,omitempty"`

	LcCollate *string `json:"lc_collate,omitempty"`

	LcCtype *string `json:"lc_ctype,omitempty"`
}

func (o PostgresqlDatabaseForCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlDatabaseForCreation struct{}"
	}

	return strings.Join([]string{"PostgresqlDatabaseForCreation", string(data)}, " ")
}
