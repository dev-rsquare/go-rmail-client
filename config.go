package rmail

import "strings"

type RMailEnv = string

const (
	Development RMailEnv = "devel"
	Production  RMailEnv = "prod"
)

func findRMailEnv(envStr string) RMailEnv {

	var env RMailEnv

	switch strings.ToLower(envStr) {
	case "prod":
	case "production":
		env = Production
	default:
		env = Development
	}

	return env
}

func findRMailHost(env RMailEnv) string {

	var host string

	if env == Development {
		host = "https://api3-dev.rsquare.co.kr/graphql/rmail"
	} else {
		host = "https://api3.rsquare.co.kr/graphql/rmail"
	}

	return host
}

func findRMailAuthorization(env RMailEnv) string {

	var authorization string

	if env == Development {
		authorization = "Bearer 5356aabeb26db59a8afe5310b1a81f01df947611a9710ed2e96299cb42b29f8e"
	} else {
		authorization = "Bearer 3f5bb02d66f7cb38afbfc72d59c4536b950a2023838c29bacd517273946c636b"
	}

	return authorization
}
