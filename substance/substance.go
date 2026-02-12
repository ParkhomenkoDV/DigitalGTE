package substance

type Substance struct {
	Name       string
	Parameters map[string]float64
	Functions  map[string]interface{}
}
