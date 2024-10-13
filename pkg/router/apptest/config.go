package apptest

type Config struct {
	ParamNames []string
	ParamValues []string
}
func NewConfig() Config {
    return Config{
        ParamNames:  make([]string, 0),
        ParamValues: make([]string, 0),
    }
}

type UseFn func(*Config)

func UseParam(name string, value string) UseFn {
	return func(c *Config) {
		c.ParamNames = append(c.ParamNames, name)
		c.ParamValues = append(c.ParamValues, value)
	}
}
