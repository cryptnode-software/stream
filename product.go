package stream

type ProductConfig struct {
}

func (config ProductConfig) Step() string {
	return "ProductConfig"
}

func (config ProductConfig) Select(cursor int) error {
	return nil
}
