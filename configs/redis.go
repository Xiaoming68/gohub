package configs

import (
	"gohub/initPack/config"
)

func init() {

	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{

			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"port":     config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD", ""),

			// 业务类存储使用按照业务类型使用 默认 0 库
			"database": config.Env("REDIS_MAIN_DB", 0),
		}
	})
}
