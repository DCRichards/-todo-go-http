package environment

import (
	"github.com/kelseyhightower/envconfig"
)

// Environment represents a set of validated environment variables.
type Environment struct {
	Port             int    `required:"true"`
	PostgresHost     string `required:"true" split_words:"true"`
	PostgresPort     string `required:"true" split_words:"true"`
	PostgresUser     string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
	PostgresDb       string `required:"true" split_words:"true"`
}

// Get returns the current environment.
func Get() (*Environment, error) {
	var e Environment
	err := envconfig.Process("", &e)
	return &e, err
}
