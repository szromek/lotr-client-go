package lotr

type ChacarterList struct {
	Docs   []Character `json:"docs"`
	Total  int         `json:"total"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Page   int         `json:"page"`
	Pages  int         `json:"pages"`
}

type Character struct {
	ID      string `json:"_id"`
	Height  string `json:"height"`
	Race    string `jsob:"race"`
	Gender  string `jsob:"gender"`
	Birth   string `jsob:"birth"`
	Spouse  string `jsob:"spouse"`
	Death   string `jsob:"death"`
	Realm   string `jsob:"realm"`
	Hair    string `jsob:"hair"`
	Name    string `jsob:"name"`
	WikiUrl string `jsob:"wikiUrl"`
}
