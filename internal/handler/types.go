package handler

import (
	"github.com/online_marketplace/helper/locale"
	"github.com/online_marketplace/helper/server/core"
)

func Providers() []core.Service {
	return []core.Service{
		locale.NewLocalizer(),
	}
}
