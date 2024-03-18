package config

import "ownapihub/pkg"

func init() {
	pkg.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			"username": pkg.Env("DATABASE_USERNAME", "iu"),
			"password":  pkg.Env("DATABASE_PASSWORD"),
			"host": pkg.Env("DATABASE_HOST"),
			"port":  pkg.Env("DATABASE_PORT"),
			"Dbname":  pkg.Env("DATABASE_DBNAME"),
			"timeout": pkg.Env("DATABASE_TIMEOUT"),
		}
	})
}
