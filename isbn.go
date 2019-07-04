package isbn

import (
	"fmt"
	"github.com/sereiner/lib/types"
	"strings"
)

type Isbn struct {
	isbn        string
	product     string // GS1 Product Code (978 or 979 for books)
	country     string // 国家、语言或区位代码；
	publisher   string // 出版社代码；
	publication string // 书序码；
}

func NewIsbn(isbn string) *Isbn {
	if len(isbn) != 13 {
		panic("只支持13位isbn")
	}
	i := &Isbn{
		isbn: isbn,
	}

	code := i.removeChecksum()
	fmt.Println(code)
	code = i.removeProductCode(code)
	fmt.Println(code)
	code = i.removeCountryCode(code)
	fmt.Println(code)
	return i
}

func (i *Isbn) removeChecksum() string {

	return i.isbn[:len(i.isbn)-1]

}

func (i *Isbn) removeProductCode(code string) string {
	first3 := code[0:3]

	if first3 == "978" || first3 == "979" {
		i.product = first3
		code = code[3:]
	}

	return code
}

func (i *Isbn) removeCountryCode(code string) string {
	first7 := code[0:7]
	var rules []*Rule
	var length int

	for _, v := range prefixes {
		if v.Prefix == i.product {
			fmt.Println(v)
			rules = v.Rules
			break
		}
	}

	for _, v := range rules {
		ra := strings.Split(v.Range, "-")
		if types.GetInt(first7) >= types.GetInt(ra[0]) && types.GetInt(first7) <= types.GetInt(ra[1]) {
			fmt.Println(ra)
			length = types.GetInt(v.Length)
		}
	}

	if length == 0 {
		i.country = string(code[length])
	} else {
		i.country = code[0:length]
	}

	return code[length:]

}

func (i *Isbn) GetCountryCode() string {
	return i.country
}
