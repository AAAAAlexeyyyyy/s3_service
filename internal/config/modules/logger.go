package modules

type Logger struct {
	Environment string `env:"LOG_LEVEL" envDefault:"development"`
	Level       string `env:"LOG_LEVEL" envDefault:"info"`
}
