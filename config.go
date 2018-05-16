package rmail

import "strings"

// Env : Environment deployed RMail Servers
type Env = string

const (
	// Development : Dev
	Development = "devel"
	// Production : Prod
	Production = "prod"
)

func findRMailEnv(envStr string) Env {

	var env Env

	switch strings.ToLower(envStr) {
	case "prod":
	case "production":
		env = Production
	default:
		env = Development
	}

	return env
}

func findRMailHost(env Env) string {

	var host string

	if env == Development {
		host = "https://api3-dev.rsquare.co.kr/graphql/rmail"
	} else {
		host = "https://api3.rsquare.co.kr/graphql/rmail"
	}

	return host
}