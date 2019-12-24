package types

type TaskTemplate struct {
	Description string   `json:"description"`
	Project     string   `mapstructure:"project" json:"project"`
	Tags        []string `mapstructure:"tags" json:"tags"`
	Priority    string   `mapstructure:"priority" json:"priority"`
}
