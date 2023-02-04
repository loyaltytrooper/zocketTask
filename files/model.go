package files

type ImageURLs struct {
	URLs []string `json:"urls" form:"urls"`
}

type URL struct {
	File string `json:"file" form:"file"`
}
