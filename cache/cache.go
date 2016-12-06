package cache

import (
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
	"github.com/kkserver/kk-ping/ping"
)

type CacheApp struct {
	app.App
	Remote *remote.Service
	Ping   *ping.PingService
	Cache  *CacheService
}
