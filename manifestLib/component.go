package manifestlib

import "fmt"

type Component struct {
	label         string
	placeHolder   string
	inputType     string
	source        ISource
	componentType ComponentType
}

type ISource interface {
	GetData() []string
}

type IComponentType interface {
	FillComponentWithData(source []string)
}

type ComponentType struct {
	dataSource    []string
	selectedValue string
	comType       IComponentType
}

func InitComboBox() Component {
	//ComboBox
	comboBox := Component{label: "DatabaseInstance", placeHolder: "Database Instance", inputType: "string"}
	fmt.Println(comboBox.source)
	return comboBox
	//RadioButton

	//TextField
}
