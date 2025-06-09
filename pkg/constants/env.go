package constants

type AppEnv string

const (
	EnvLocal      AppEnv = "local"
	EnvStaging    AppEnv = "staging"
	EnvSandbox    AppEnv = "sandbox"
	EnvProduction AppEnv = "production"
	EnvTesting    AppEnv = "testing"
)
