package cache

import (
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type CacheApp struct {
	app.App
	Remote *remote.Service
	Cache  *CacheService
}
