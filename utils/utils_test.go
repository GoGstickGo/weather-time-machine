package utils

import (
	"testing"
)

func TestBuildTime(t *testing.T) {
	type args struct {
		y string
		m string
		d string
	}
	tests := []struct {
		name    string
		args    args
		wantU   string
		wantErr bool
	}{
		{"good#1", args{"1945", "03", "02"}, "1945-03-02T12:00:00", false},
		{"good#2", args{"1980", "03", "02"}, "1980-03-02T12:00:00", false},
		{"good#3", args{"2021", "03", "02"}, "2021-03-02T12:00:00", false},
		{"wrongYearPast", args{"1920", "45", "02"}, "", true},
		{"wrongMonth", args{"1945", "48", "02"}, "", true},
		{"wrongDay", args{"1962", "02", "99"}, "", true},
		{"wrongYearFuture", args{"2022", "02", "02"}, "", true},
		{"wrongDateLetters#1", args{"2019", "aa", "b2"}, "", true},
		{"wrongDateLetters#1", args{"201a", "aa", "b2"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, err := BuildTimeStr(tt.args.y, tt.args.m, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotU != tt.wantU {
				t.Errorf("BuildTime() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}
