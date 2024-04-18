package global

import (
	datebase "git.virjar.com/Junhiee/anilismei/database"
	"go.uber.org/zap"
)

var (
	G_QRY *datebase.Store
	ZLOG  *zap.Logger
)
