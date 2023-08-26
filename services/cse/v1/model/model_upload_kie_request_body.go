package model

import (
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/utils"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/g42cloud-sdk/g42cloud-sdk-go/core/def"
	"os"
	"reflect"

	"strings"
)

type UploadKieRequestBody struct {
	UploadFile *def.FilePart `json:"upload_file"`
}

func (o UploadKieRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKieRequestBody struct{}"
	}

	return strings.Join([]string{"UploadKieRequestBody", string(data)}, " ")
}

func (o *UploadKieRequestBody) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if m[jsonName] == nil && strings.Contains(jsonTag, "omitempty") {
			continue
		}
		field := v.FieldByName(utils.UnderscoreToCamel(jsonName))
		switch v.Field(i).Interface().(type) {
		case *def.FilePart:
			filePath := m[jsonName].(string)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(def.NewFilePart(file)))
		case *def.MultiPart:
			field.Set(reflect.ValueOf(def.NewMultiPart(m[jsonName])))
		default:
			return errors.New(fmt.Sprintf("unmarshal %s failed", m[jsonName]))
		}
	}
	return nil
}
