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
func (d *DatabaseForm) valuesAreValid() bool{
	return false
}

func (c *CloudForm) valuesAreValid() bool  {
	return false
}
func ValidateTypeAndValues(sorceJsonStr []byte, manifestName string) bool{
	var d DatabaseForm
	var c CloudForm

	// items := [...]Validator{&d, &c}
	// fmt.Println(items)
	// for _, item := range items {
	// 	fmt.Println(item)
	// if item.isValidType(sorceJsonStr) {
	// 		fmt.Println(item)
	// 	} else {
	// 		fmt.Println("hi")
	// 	}
	// }
	switch manifestName {
	case "database":
		if d.isValidType(sorceJsonStr) && d.valuesAreValid() {
			fmt.Println("this is database manifest")
			return true
		}
	case "cloud":
		if c.isValidType(sorceJsonStr) && c.valuesAreValid() {
			fmt.Println("this is cloud manifest")
			return true
		}
	}
	
	return false
}
