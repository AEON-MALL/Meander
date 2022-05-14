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

var APIKey string

func (p *Place) Public() any{
	return map[string]any{
		"name": p.Name,
		"icon": p.Icon,
		"photos": p.Photos,
		"vicinity": p.Vicinity,
		"lat": p.Lat,
		"lng": p.Lng,
	}
}