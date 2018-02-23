package rmail

type RMailClient struct {
	Env RMailEnv
}

func findRMailEnv(envStr string) RMailEnv {

	var env RMailEnv

	switch envStr {
	case "prod":
	case "production":
		env = Development
	default:
		env = Production
	}

	return env
}

func InitRMailClient(envStr string) *RMailClient {
	return &RMailClient{
		Env: findRMailEnv(envStr),
	}
}
