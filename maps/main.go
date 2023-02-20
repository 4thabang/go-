package main

const (
	ErrInvalidWord         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists          = DictionaryErr("this word exists already")
	ErrCannotUpdateNilWord = DictionaryErr("cannot update non-existent word")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Add(key, value string) error {
	if _, err := d.Search(key); err == nil {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrInvalidWord
	}
	return word, nil
}

func (d Dictionary) Update(key, update string) error {
	if _, err := d.Search(key); err != nil {
		return ErrCannotUpdateNilWord
	}
	d[key] = update
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
