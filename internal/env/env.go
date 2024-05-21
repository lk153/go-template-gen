package env

type (
	EnvVars struct {
		SERVER_PORT uint
	}
)

func InitEnvVars() EnvVars {
	return EnvVars{
		SERVER_PORT: 8888,
	}
}
