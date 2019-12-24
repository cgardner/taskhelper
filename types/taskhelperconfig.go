package types

type TaskHelperConfig struct {
	Templates map[string]Template `mapstructure:"templates"`
}
