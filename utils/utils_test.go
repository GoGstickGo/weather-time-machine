package utils

import (
	"testing"
)

func TestBuildDate(t *testing.T) {
	type args struct {
		year  string
		month string
		day   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"good#1", args{year: "1941", month: "03", day: "02"}, "1941-03-02", false},
		{"good#2", args{year: "1980", month: "12", day: "02"}, "1980-12-02", false},
		{"good#3", args{year: "2021", month: "09", day: "07"}, "2021-09-07", false},
		{"wrongYearPast", args{year: "1920", month: "09", day: "07"}, "1920", true},
		{"wrongYearFuture", args{year: "2029", month: "09", day: "07"}, "2029", true},
		{"wrongMonth", args{year: "1958", month: "20", day: "07"}, "1958-20-07", true},
		{"wrongDay", args{year: "1956", month: "03", day: "99"}, "1956-03-99", true},
		{"wrongDateLetters#1", args{year: "1948", month: "aa", day: "07"}, "1948-aa-07", true},
		{"wrongDateLetters#1", args{year: "1972", month: "09", day: "0b"}, "1972-09-0b", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildDate(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
