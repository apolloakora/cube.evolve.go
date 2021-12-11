package redis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var rPool *redis.Pool
