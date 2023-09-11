package cache

import "errors"

// Derive 驱动
type Derive interface {
	Init(config map[string]string) error
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Del(key string) error
	IsExist(key string) bool
	Close() error
	Clear() error
}

var (
	defaultDerive = "file"
	derives       = make(map[string]Derive)
)

func GetDerive(name string) (Derive, error) {
	if name == "" {
		name = defaultDerive
	}

	// 判断
	if _, ok := derives[name]; !ok {
		return nil, errors.New("cache derive not found")
	}

	return derives[name], nil
}

func RegisterDerive(name string, derive Derive) {

	if _, ok := derives[name]; ok {
		panic("derive already registered")
	}

	derives[name] = derive
}
