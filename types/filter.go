package types

type Filter struct {
	Project string   `mapstructure:"project"`
	Tags    []string `mapstructure:"tags"`
	Report  string   `mapstructure:"report"`
}
