package repository

type RedisRepositoryMock struct {
	records map[string]string
}

func NewRedisRepositoryMock() *RedisRepositoryMock {
	return &RedisRepositoryMock{
		records: make(map[string]string),
	}
}

func (repo *RedisRepositoryMock) Append(key string, value string) {
	repo.records[key] = value
}
func (repo *RedisRepositoryMock) Keys(pattern string) []string {
	keys := make([]string, 0)
	for k := range repo.records {
		keys = append(keys, k)
	}
	return keys
}
func (repo *RedisRepositoryMock) Delete(key string) {
}
func (repo *RedisRepositoryMock) JsonMget(ids []string) [][]byte {
	values := make([][]byte, 0)
	for k := range repo.records {
		values = append(values, []byte(repo.records[k]))
	}
	return values
}
func (repo *RedisRepositoryMock) JsonGet(key string) []byte {
	return []byte(repo.records[key])
}
func (repo *RedisRepositoryMock) JsonSet(key string, value interface{}) {}
