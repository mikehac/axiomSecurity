package manifestlib

type DatabaseForm struct {
	Reason            *string `json:"reason"`
	DatabaseInstances *string `json:"databaseInstances"`
	DatabaseRole      *string `json:"databaseRole"`
}

type CloudForm struct {
	Reason   *string `json:"reason"`
	Accounts *string `json:"accounts"`
	User     *string `json:"user"`
}

type Validator interface {
	isValidType(sourceJsonStr []byte) bool
	valuesAreValid() bool
}
