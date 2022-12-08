package manifestlib

type Component struct {
	Source             string `json:"souce"`
	Label              string `json:"lael"`
	ComponentType      string `json:"component_type"`
	PlaceHolder        string `json:"placeHolder"`
	InputType          string `json:"input_tpe"`
	MultiValuesAllowed bool   `json:"multiValues_Allowed"`
}

type Manifest struct {
	Components []Component `json:"components"`
}

var ManifestsMap = make(map[string]Manifest)

type SourceValues struct {
	Values []string `json:"values"`
}

var SourceValueMap = make(map[string]SourceValues)
