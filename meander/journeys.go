package meander

type j struct {
	Name       string
	placeTypes []string
}

var Journeys = []any{
	&j{Name: "ロマンティック", placeTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"}},
	&j{Name: "ショッピング", placeTypes: []string{"department_store", "cafe", "clothing_store", "jewelry_store", "shoe_store"}},
	&j{Name: "ナイトライフ", placeTypes: []string{"bar", "casino", "food", "bar", "night_club", "bar", "bar", "hospital"}},
	&j{Name: "カルチャー", placeTypes: []string{"museum", "cafe", "cemetery", "library", "art_gallery"}},
	&j{Name: "リラックス", placeTypes: []string{"hair_care", "beauty_salon", "cafe", "spa"}},
}
