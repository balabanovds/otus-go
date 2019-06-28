package main

import (
	"reflect"
	"testing"
	"words"
)

func TestCounter(t *testing.T) {
	text := `Duis ea eu nisi proident eu cillum sint ad in ea minim excepteur. Consequat cillum in ullamco dolore deserunt irure laborum aute exercitation. Et culpa enim et laborum non enim laborum. Nisi ut aliquip laboris consectetur magna magna ullamco. Officia consectetur laboris ad irure ut non eiusmod. Exercitation exercitation duis occaecat commodo sit aute dolore sit eu cupidatat cupidatat. Deserunt labore nulla culpa nisi excepteur occaecat amet magna.
Adipisicing irure officia in irure mollit ex culpa est nulla deserunt aliqua tempor. Do dolore voluptate duis cillum consectetur ad sint aute voluptate mollit consectetur ullamco tempor cillum. Proident fugiat nulla adipisicing aliqua cupidatat sit.
Non velit enim sit in voluptate laboris officia cillum qui. Non fugiat magna do eiusmod reprehenderit. Nostrud aliquip voluptate cillum ex. Dolor ea irure ullamco pariatur reprehenderit sint excepteur.
Cupidatat nisi mollit laborum aliquip pariatur laborum veniam tempor proident officia ad duis voluptate. Eu reprehenderit eu enim excepteur dolor officia ipsum sint do elit id consectetur. Aliqua pariatur incididunt in esse qui mollit. Nostrud aliqua ex excepteur dolor sunt esse esse anim culpa mollit labore culpa. Deserunt non ipsum commodo amet qui do officia enim.
Magna ex consectetur eiusmod ea proident veniam irure labore. Cupidatat occaecat laboris id magna non elit. Consectetur nulla fugiat Lorem ex dolor nostrud sit voluptate et ullamco deserunt adipisicing. Culpa laborum ad laboris Lorem incididunt ad ex eiusmod ex. Elit officia consequat eu est laborum aliqua consequat quis. Cillum ad anim irure in reprehenderit mollit laborum magna. Ad laborum ea duis id voluptate anim qui adipisicing Lorem proident.
Amet nulla consectetur fugiat aute et proident qui ad officia eiusmod nostrud qui id tempor. Et amet ex nostrud aliqua. Dolore amet mollit ipsum in nisi velit ipsum consectetur sit amet laboris cupidatat voluptate commodo.
Commodo labore consequat quis duis reprehenderit dolore nulla cillum nisi incididunt ex. Dolor cillum elit quis id irure enim anim laboris adipisicing do. Nostrud quis irure et fugiat quis occaecat tempor esse veniam elit ex.
Eu eiusmod velit quis proident Lorem consectetur consectetur quis commodo aliqua non nulla duis. Est dolor voluptate ad ullamco reprehenderit officia sunt esse exercitation dolore minim Lorem minim laboris. Enim nostrud nulla proident reprehenderit adipisicing sint voluptate irure veniam pariatur aute.`
	expected := []string{"consectetur", "ad", "irure", "voluptate", "cillum", "ex", "laborum", "officia", "laboris", "nulla"}
	got := words.Count(text)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\nexpected:\t%+v\nreceived:\t%+v", expected, got)
	}
}
