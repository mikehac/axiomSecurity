package manifestlib

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/slices"
)

func (d *DatabaseForm) isValidType(sourceJsonStr []byte) bool {
	err := json.Unmarshal([]byte(sourceJsonStr), &d)
	return err == nil && d.Reason != nil && d.DatabaseInstances != nil && d.DatabaseRole != nil
}

func (c *CloudForm) isValidType(sourceJsonStr []byte) bool {
	err := json.Unmarshal([]byte(sourceJsonStr), &c)
	return err == nil && c.Reason != nil && c.Accounts != nil && c.User != nil
}
func (d *DatabaseForm) valuesAreValid() bool {
	//Reflection approach
	// field := reflect.Indirect(reflect.ValueOf(d))
	// for i := 0; i < field.Type().NumField(); i++ {
	// 	currentFieldName := field.Type().Field(i).Name
	// 	possibleValues := SourceValueMap[currentFieldName]
	// 	containsValue := slices.Contains(possibleValues.Values, *d.DatabaseInstances)
	// 	if !containsValue && currentFieldName != "Reason" {
	// 		fmt.Println("values don't match acording the manifest")
	// 		// return false
	// 	}
	// 	fmt.Println(possibleValues.Values)
	// }
	// r := reflect.ValueOf(*d)
	// values := make([]interface{}, r.NumField())
	// for i := 0; i < r.NumField(); i++ {
	// 	// currentFieldName := field.Type().Field(i).Name
	// 	// possibleValues := SourceValueMap[currentFieldName]
	// 	// currentValue := r.Field(i).Interface()
	// 	// fmt.Println(currentValue)
	// 	// containsValue := slices.Contains(possibleValues.Values, string(&currentValue))
	// 	// if !containsValue && currentFieldName != "Reason" {
	// 	// 	fmt.Println("values don't match acording the manifest")
	// 	// 	return false
	// 	// }
	// 	values[i] = r.Field(i).Interface()
	// 	fmt.Println(values[i])
	// 	fmt.Println(&values[i])
	// }
	// fmt.Println(&values)
	databaseInstancesValues := SourceValueMap["DatabaseInstances"]
	databaseRoleValues := SourceValueMap["DatabaseRole"]
	//Because the Reason is a string and the required validation is to check if it is not empty
	//there is no need to check the value against the SourceValueMap
	if *d.Reason != "" &&
		slices.Contains(databaseInstancesValues.Values, *d.DatabaseInstances) &&
		slices.Contains(databaseRoleValues.Values, *d.DatabaseRole) {
		return true
	}

	return false
}

func (c *CloudForm) valuesAreValid() bool {
	accountsValues := SourceValueMap["Accounts"]
	userValues := SourceValueMap["User"]
	if *c.Reason != "" &&
		slices.Contains(accountsValues.Values, *c.Accounts) &&
		slices.Contains(userValues.Values, *c.User) {
		return true
	}

	return false
}
func ValidateTypeAndValues(sorceJsonStr []byte, manifestName string) bool {
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
