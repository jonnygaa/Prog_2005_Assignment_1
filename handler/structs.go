package handler

type Info struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

type Pop struct {
	Mean   int                      `json:"mean"`
	Values []map[string]interface{} `json:"values"`
}

type Status struct {
	Countriesnowapi  int    `json:"countriesnowapi"`
	Testcountriesapi int    `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}
