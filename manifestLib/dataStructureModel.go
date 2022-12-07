package manifestlib

type Component struct {
	Source        string `json:"source"`
	Label         string `json:"label"`
	ComponentType string `json:"component_type"`
	PlaceHolder   string `json:"placeHolder"`
	InputType     string `json:"input_type"`
}

type Manifest struct {
	Components []Component `json:"components"`
}

var ManifestsMap = make(map[string]Manifest)

type SourceValues struct {
	Values []string `json:"values"`
}

var SourceValueMap = make(map[string]SourceValues)
