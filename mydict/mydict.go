package mydict

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

//Dictionary Search
func (d Dictionary) Search(word string) error {

	v, err := d[word]

	if err {
		fmt.Println(word, "의 값은 : ", v)
	} else {
		return errors.New("해당 단어를 찾을수 없습니다.")
	}
	return nil
} //조회

func (d Dictionary) Insert(word, def string) error {

	err := d.Search(word)
	if err != nil {
		d[word] = def
		fmt.Println(d)
	} else {
		return errors.New("해당 단어가 이미존재합니다.")
	}
	return nil
} //추가

func (d Dictionary) Delete(word string) {
	delete(d, word)
	fmt.Println(d)
} //삭제

func (d Dictionary) Update(word, def string) error {

	err := d.Search(word)

	if err == nil {
		d[word] = def
		fmt.Println(d)
	} else {
		return errors.New("해당 단어 업데이트 실패")
	}

	return nil
}
