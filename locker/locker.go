package locker

import (
	"github.com/bsm/redislock"
	"github.com/fishjar/gin-boilerplate/db"
)

// 业务锁常量
const (
	PING = "lock:ping" // ping
)

// Locker redis锁客户端
var Locker *redislock.Client = redislock.New(db.Redis)
