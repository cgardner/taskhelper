package types

type Report struct {
	Report  string `mapstructure:"report"`
	Filters Filter `mapstructure:"filters"`
}
