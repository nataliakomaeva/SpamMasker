package main

import "testing"

func TestSpam(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "Вот ссылка: http://spam.com приходи",
			output: "Вот ссылка: http://******** приходи",
		},
		{
			input:  "Без ссылок вообще",
			output: "Без ссылок вообще",
		},
		{
			input:  "Несколько ссылок http://one.com и http://two.com",
			output: "Несколько ссылок http://******* и http://*******",
		},
		{
			input:  "http://superlongurl.com/with/stuff в начале строки",
			output: "http://*************************** в начале строки",
		},
	}
	for _, test := range tests {
		result := Spam(test.input)
		if result != test.output {
			t.Errorf("Spam(%q) = %q; want %q", test.input, result, test.output)
		}
	}
}
