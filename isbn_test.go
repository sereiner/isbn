package isbn

import (
	"fmt"
	"testing"
)

func TestGetCountry(t *testing.T) {
	got := NewIsbn("9789888012695")
	fmt.Printf("%+v\n", got)
	t.Log(got)
}
