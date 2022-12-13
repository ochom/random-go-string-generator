package main

import "testing"

func Test_parseToken(t *testing.T) {
	type args struct {
		token  string
		repeat int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				token:  "a",
				repeat: 2,
			},
			want: "aa",
		},
		{
			name: "test2",
			args: args{
				token:  "[a]",
				repeat: 1,
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseToken(tt.args.token, tt.args.repeat); got != tt.want {
				t.Errorf("parseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
