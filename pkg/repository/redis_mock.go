package repository

type RedisRepositoryMock struct {
	records map[string]string
	keys    []string
}

func NewRedisRepositoryMock() *RedisRepositoryMock {
	return &RedisRepositoryMock{
		records: make(map[string]string),
		keys:    make([]string, 0),
	}
}

func (repo *RedisRepositoryMock) Append(key string, value string) {
	repo.records[key] = value
	repo.keys = append(repo.keys, key)
}
func (repo *RedisRepositoryMock) Keys(pattern string) []string {
	return repo.keys
}
func (repo *RedisRepositoryMock) Delete(key string) {
}
func (repo *RedisRepositoryMock) JsonMget(ids []string) [][]byte {
	values := make([][]byte, 0)
	for _, k := range repo.keys {
		values = append(values, []byte(repo.records[k]))
	}
	return values
}
func (repo *RedisRepositoryMock) JsonGet(key string) []byte {
	return []byte(repo.records[key])
}
func (repo *RedisRepositoryMock) JsonSet(key string, value interface{}) error {
	return nil
}
func (repo *RedisRepositoryMock) JsonDel(key string) error {
	return nil
}
