package parsercombinator

import (
	"testing"
	"unicode"
)

func TestAnyRune(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, want1 string
		want2     int
		want3     bool
	}{
		{"abc", "a", 1, true},
		{"߷ÁÁ", "߷", 2, true},
		{"あいうえお", "あ", 3, true},
		{"🍣", "🍣", 4, true},
		{"", "", 0, false},
	}
	for i, c := range cases {
		got, num, succeed := AnyRune()(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
		}
	}
}

func TestDigit(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, want1 string
		want2     int
		want3     bool
	}{
		{"000", "0", 1, true},
		{"テスト", "", 0, false},
		{"", "", 0, false},
	}
	for i, c := range cases {
		got, num, succeed := Digit()(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
		}
	}
}

func TestLetter(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in, want1 string
		want2     int
		want3     bool
	}{
		{"abc", "a", 1, true},
		{"テスト", "テ", 3, true},
		{"000", "", 0, false},
		{"", "", 0, false},
	}
	for i, c := range cases {
		got, num, succeed := Letter()(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
		}
	}
}

func TestRune(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in    string
		test  rune
		want1 string
		want2 int
		want3 bool
	}{
		{"abc", 'a', "a", 1, true},
		{"あいう", 'あ', "あ", 3, true},
		{"あいう", 'い', "", 0, false},
		{"", 'a', "", 0, false},
	}
	for i, c := range cases {
		got, num, succeed := Rune(c.test)(c.in)
		if !(got == c.want1 && num == c.want2 && succeed == c.want3) {
			t.Error(i, got, num, succeed, c)
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
		got, num, succeed := c.rule(c.in)
		if succeed != c.want {
			t.Error(i, got, num, succeed, c)
		}
	}
}
