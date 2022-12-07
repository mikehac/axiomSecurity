package manifestlib

import (
	"encoding/json"
	"fmt"
)

func (d *DatabaseForm) isValidType(sourceJsonStr []byte) bool {
	err := json.Unmarshal([]byte(sourceJsonStr), &d)
	return err == nil && d.Reason != nil && d.DatabaseInstances != nil && d.DatabaseRole != nil
}

func (c *CloudForm) isValidType(sourceJsonStr []byte) bool {
	err := json.Unmarshal([]byte(sourceJsonStr), &c)
	return err == nil && c.Reason != nil && c.Accounts != nil && c.User != nil
}

func ValidateTypeAndValues(sourceJsonStr []byte) {
	var d DatabaseForm
	var c CloudForm

	items := [...]Validator{&d, &c}
	fmt.Println(items)
	for _, item := range items {
		fmt.Println(item)
		if item.isValidType(sourceJsonStr) {
			fmt.Println(item)
		} else {
			fmt.Println("hi")
		}
	}
}
