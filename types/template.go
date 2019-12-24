package types

type Template struct {
	Add      TaskTemplate `mapstructure:"add"`
	Report   []string     `mapstructure:"report"`
	Priority string       `mapstructure:"priority"`
}
