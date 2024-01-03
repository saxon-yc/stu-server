package dbsvc

type RedisDataBase interface {
	FindCountByStuBindTag(key string) (str string, err error)
	IncrCountWithStuBindTag(key string) (err error)
	DecrCountWithStuBindTag(key string) (err error)
}

func (b basicService) FindCountByStuBindTag(key string) (str string, err error) {
	str, err = b.rdb.Get(b.ctx, key).Result()
	return
}

func (b basicService) IncrCountWithStuBindTag(key string) (err error) {

	str, err := b.FindCountByStuBindTag(key)
	if str == "" {
		b.rdb.Set(b.ctx, key, 1, 0)
	} else {
		b.rdb.Incr(b.ctx, key)
	}

	return
}
func (b basicService) DecrCountWithStuBindTag(key string) (err error) {

	str, err := b.FindCountByStuBindTag(key)
	if str != "" {
		b.rdb.Decr(b.ctx, key)
	}

	return
}
