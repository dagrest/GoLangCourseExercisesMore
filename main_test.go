package main

import "testing"

func Test_test(t *testing.T) {
	type args struct {
		arg []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "first", args: args{}, want: 0, wantErr: false},
		// ===========================>
		// QUESTION - HOW TO FIX THIS?
		//{name: "second", args: args{arg: {1, 2}}, want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := test(tt.args.arg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("channel_try() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("channel_try() got = %v, want %v", got, tt.want)
			}
		})
	}
}
