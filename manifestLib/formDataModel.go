package manifestlib

type DatabaseForm struct {
	Reason            *string `json:"Reason"`
	DatabaseInstances *string `json:"DatabaseInstances"`
	DatabaseRole      *string `json:"DatabaseRole"`
}

type CloudForm struct {
	Reason   *string `json:"Reason"`
	Accounts *string `json:"Accounts"`
	User     *string `json:"User"`
}

type Validator interface {
	isValidType(sourceJsonStr []byte) bool
	valuesAreValid() bool
}
