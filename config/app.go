package config

import "ownapihub/pkg"

func init() {
	pkg.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"port": pkg.Env("APP_PORT", "3000"),
			"url":  pkg.Env("APP_URL", "http://localhost:3000"),
		}
	})
}
