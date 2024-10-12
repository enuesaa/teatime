package apptest

type Config struct {
    Route string
    Invoke string
}

type Option func(*Config)

func UseRoute(route string, path string) Option {
    return func(c *Config) {
        c.Route = route
        c.Invoke = path
    }
}
