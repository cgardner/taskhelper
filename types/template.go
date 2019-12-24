package types

type Template struct {
	Add      TaskTemplate `mapstructure:"add"`
	Report   Report       `mapstructure:"report"`
	Priority string       `mapstructure:"priority"`
}
