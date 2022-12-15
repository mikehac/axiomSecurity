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

	//Because the Reason is a string and the required validation is to check if it is not empty
	//there is no need to check the value against the SourceValueMap
	if *c.Reason != "" &&
		slices.Contains(accountsValues.Values, *c.Accounts) &&
		slices.Contains(userValues.Values, *c.User) {
		return true
	}

	return false
}
func ValidateTypeAndValues(sorceJsonStr []byte, manifestName string) bool {
	var v Validator

	fmt.Println(v)
	switch manifestName {
	case "database":
		v = new(DatabaseForm)
	case "cloud":
		v = new(CloudForm)
	}

	if v.isValidType(sorceJsonStr) && v.valuesAreValid() {
		// fmt.Println("this is database manifest")
		return true
	}

	return false
}
