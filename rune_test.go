package parsercombinator

import (
	"testing"
	"unicode"
)

func TestAnyRune(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in    string
		want1 rune
		want2 int
		want3 bool
	}{
		{"abc", 'a', 1, true},
		{"߷ÁÁ", '߷', 1, true},
		{"あいうえお", 'あ', 1, true},
		{"🍣", '🍣', 1, true},
		{"", 0, 0, false},
	}
	for i, c := range cases {
		got, num, ok := AnyRune()([]rune(c.in))
		if !(got == c.want1 && num == c.want2 && ok == c.want3) {
			t.Error(i, got, num, ok, c)
		}
	}
}

func TestRules(t *testing.T) {
	t.Parallel()

	cases := []struct {
		rule RuneFunc
		in   string
		want bool
	}{
		{Digit(), "000", true},
		{Digit(), "൧", true},
		{Digit(), "１", true},
		{Digit(), "テスト", false},
		{Letter(), "abc", true},
		{Letter(), "000", false},
		{Rune('a'), "abc", true},
		{Rune('あ'), "あいう", true},
		{Rune('い'), "あいう", false},
		{Rune('a'), "", false},
		{Control(), "\u0000", true},
		{Control(), "あ", false},
		{Control(), "", false},
		{Graphic(), "あ", true},
		{Graphic(), "\u0000", false},
		{Graphic(), "", false},
		{Lower(), "a", true},
		{Lower(), "A", false},
		{Lower(), "", false},
		{Mark(), "꙰", true},
		{Mark(), "A", false},
		{Mark(), "", false},
		{Number(), "1", true},
		{Number(), "A", false},
		{Number(), "", false},
		{Print(), "1", true},
		{Print(), "\u0000", false},
		{Print(), "", false},
		{Punct(), "-", true},
		{Punct(), "A", false},
		{Punct(), "", false},
		{Space(), " ", true},
		{Space(), "A", false},
		{Space(), "", false},
		{Symbol(), "¥", true},
		{Symbol(), "A", false},
		{Symbol(), "", false},
		{Title(), "ǋ", true},
		{Title(), "A", false},
		{Title(), "", false},
		{Upper(), "A", true},
		{Upper(), "a", false},
		{Upper(), "", false},
		{In(unicode.Hiragana), "あ", true},
		{In(unicode.Hiragana), "ア", false},
		{In(unicode.Hiragana), "", false},
		{InStr("あいうえお"), "あ", true},
		{InStr("あいうえお"), "ア", false},
		{InStr("あいうえお"), "", false},
	}
	for i, c := range cases {
		got, num, ok := c.rule([]rune(c.in))
		if ok != c.want {
			t.Error(i, got, num, ok, c)
		}
	}
}
