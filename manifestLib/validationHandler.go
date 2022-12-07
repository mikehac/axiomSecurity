package manifestlib

import (
	"encoding/json"
	"fmt"
)

func (d *DatabaseForm) validateType(sourceJsonStr []byte) bool {
	err := json.Unmarshal([]byte(sourceJsonStr), &d)
	return err == nil
}

func (c *CloudForm) validateType(sourceJsonStr []byte) bool {
	err := json.Unmarshal([]byte(sourceJsonStr), &c)
	return err == nil
}

func ValidateTypeAndValues(sourceJsonStr []byte) {
	var d DatabaseForm
	var c CloudForm

	items := [...]Validator{&d, &c}
	fmt.Println(items)
	for _, item := range items {
		fmt.Println(item)
		if item.validateType(sourceJsonStr) {
			fmt.Println(item)
		}
	}
}
