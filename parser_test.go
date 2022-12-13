package main

import (
	"regexp"
	"testing"
)

func TestPattern_Generate(t *testing.T) {
	tests := []struct {
		name string
		p    Pattern
	}{
		{

			name: "test1",
			p:    "/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{8}/",
		},
		{
			name: "test2",
			p:    "/[-+]?[0-9]{1,16}[.][0-9]{1,6}/",
		},
		{
			name: "test3",
			p:    "/.{8,12}/",
		},
		{
			name: "test4",
			p:    "/[^aeiouAEIOU0-9]{5}/",
		},
		{
			name: "test2",
			p:    "/[a-f-]{5}/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Generate()
			ok := regexp.MustCompile(tt.p.String()).MatchString(got)
			if !ok {
				t.Errorf("Pattern.Generate() = %v, want %v", got, tt.p)
			}
		})
	}
}
