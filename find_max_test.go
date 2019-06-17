package max

import "testing"

func TestFindMaxSimpleInt(t *testing.T) {
	s := []interface{}{1, 333, 33}
	less := func(i, j int) bool {
		return s[i].(int) < s[j].(int)
	}

	max := FindMax(s, less)

	if max != 333 {
		t.Errorf("Want %v, got %v", 333, max)
	}
}

func TestFindMaxSimpleString(t *testing.T) {
	s := []interface{}{"1", "4", "3"}
	less := func(i, j int) bool {
		return s[i].(string) < s[j].(string)
	}

	max := FindMax(s, less)

	if max != "4" {
		t.Errorf("Want %v, got %v", "4", max)
	}
}

func TestFindMaxSimpleStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	s := []interface{}{
		person{"Vasya", 23},
		person{"Petya", 44},
		person{"Sonya", 22},
	}
	less := func(i, j int) bool {
		return s[i].(person).age < s[j].(person).age
	}

	max := FindMax(s, less)

	if max.(person).name != "Petya" {
		t.Errorf("Want %v, got %v", "Petya", max.(person).name)
	}
}
func TestFindMaxReflInt(t *testing.T) {
	s := []int64{1, 333, 33}
	less := func(i, j int) bool {
		return s[i] < s[j]
	}

	max := FindMaxReflection(s, less)

	if max.(int64) != 333 {
		t.Errorf("Want %v, got %v", 333, max)
	}
}

func TestFindMaxReflString(t *testing.T) {
	s := []string{"1", "4", "3"}
	less := func(i, j int) bool {
		return s[i] < s[j]
	}

	max := FindMaxReflection(s, less)

	if max != "4" {
		t.Errorf("Want %v, got %v", "4", max)
	}
}
