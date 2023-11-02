package easyjsonsample

//go:generate easyjson main.go

//easyjson:json
type sampleStructure struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float32 `json:"height"`
}

func main() {
}
