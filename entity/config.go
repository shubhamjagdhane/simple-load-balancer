package entity

// Config ...
type Config struct {
	ENV         string `json:"env" mapstructure:"env" validate:"required"`
	LogLevel    string `json:"log_level" mapstructure:"log_level"`
	HTTPAddress string `json:"http_address" mapstructure:"http_address" validate:"required"`
	Port        string `json:"port" mapstructure:"port" validate:"required"`

	Tracer Tracer `json:"tracer" mapstructure:"tracer"`
}

type Tracer struct {
	// ProjectID Google Project ID
	ProjectID  string `json:"project_id" mapstructure:"project_id"`
	TracerName string `json:"name" mapstructure:"name"`
	Enable     bool   `json:"enable" mapstructure:"enable"`
}
