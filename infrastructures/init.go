package infrastructures

import (
	cacheredis "github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/cache-redis"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/config"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/db"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/infrastructures/logs"
)

type Init struct {
	*db.DBimpl
	*cacheredis.RedisImpl
}

func NewInit() *Init {
	config.EnvInit()
	logs.LogInit()
	return &Init{
		db.NewDB(),
		cacheredis.NewRedisConn(),
	}
}
