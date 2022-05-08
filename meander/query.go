package meander

type Place struct {
	*googleGemetry `json:"gemetry"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Photos []*googlePhoto `json:"photos"`
	Vicinity string `json:"vicinity"`
}

type googleResponse struct {
	Result []*Place `json:"results"`
}

type googleGemetry struct{
	*googleLocation `json:"location"`
}

type googleLocation struct{
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct{
	PhotoRef string `json:"photo_reference"`
	URL string `json:"url"`
}
