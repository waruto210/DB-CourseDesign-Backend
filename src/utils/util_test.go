package utils

import (
	"log"
	"testing"
)

func GetHashPasswd(passwd string) string {
	hash, err := HashPasswd(passwd)
	if err != nil {
		log.Fatalln(err)
	}
	return hash
}

func TestCheckPasswd(t *testing.T) {
	type args struct {
		passwd string
		hashed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test_empty", args: args{passwd: "", hashed: GetHashPasswd("")}, want: true},
		{name: "test_normal", args: args{passwd: "123456", hashed: GetHashPasswd("123456")}, want: true},
		{name: "test_notequal", args: args{passwd: "99999999", hashed: GetHashPasswd("777777777")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswd(tt.args.passwd, tt.args.hashed); got != tt.want {
				t.Errorf("CheckPasswd() = %v, want %v", got, tt.want)
			}
		})
	}
}
