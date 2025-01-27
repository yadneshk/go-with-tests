package maps

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrMissingKey       = DictionaryErr("missing key")
	ErrKeyAlreadyExists = DictionaryErr("key already exists")
)

func (err DictionaryErr) Error() string {
	return string(err)
}

func (dict Dictionary) Search(key string) (string, error) {
	value, ok := dict[key]
	if !ok {
		return "", ErrMissingKey
	}
	return value, nil
}

func (dict Dictionary) Add(key, value string) error {
	if _, ok := dict[key]; ok {
		return ErrKeyAlreadyExists
	} else {
		dict[key] = value
	}
	return nil
}

func (dict Dictionary) Update(key, value string) error {
	if _, err := dict.Search(key); err != nil {
		return ErrMissingKey
	} else {
		dict[key] = value
	}
	return nil
}

func (dict Dictionary) Delete(key string) error {
	if _, err := dict.Search(key); err != nil {
		return ErrMissingKey
	} else {
		delete(dict, key)
	}
	return nil
}
