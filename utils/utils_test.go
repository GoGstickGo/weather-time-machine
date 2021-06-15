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

func TestValidateRapidApiKey(t *testing.T) {
	testData1 := map[string]interface{}{}
	testData1["message"] = "dfsfdsfdsfdsYou are not subscribed to this API"
	testData2 := map[string]interface{}{}
	testData2["nomessage"] = "dfsfdsfdsfdsYou are not subscribed to this API"
	testData3 := map[string]interface{}{}
	testData3["message"] = "It's all good"

	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{"#good", args{data: testData1}, true},
		{"bad#1", args{data: testData2}, false},
		{"bad#2", args{data: testData3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateRapidApiKey(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateRapidApiKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateParamsCity(t *testing.T) {
	type args struct {
		city string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"good#1", args{city: "San Francisco"}, false},
		{"good#2", args{city: "New York City"}, false},
		{"good#3", args{city: "Niagara-on-the-Lake"}, false},
		{"bad#1", args{city: "toronto"}, true},
		{"bad#2", args{city: "Toronto12312"}, true},
		{"bad#3", args{city: "aRuBa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateParamsCity(tt.args.city); (err != nil) != tt.wantErr {
				t.Errorf("ValidateParamsCity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateParamsApikey(t *testing.T) {
	type args struct {
		apikey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"good#1", args{apikey: "c9KbT3WfMNDrLj7Tyx7XQKz5j6zKWjWPdebfKf2nKMFQn386N7"}, false},
		{"bad#2", args{apikey: "c9KbT3WfMNDrLj7Tyx7XQKz5j6zKWjWPdebfKf2nKMFQn3"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateParamsApikey(tt.args.apikey); (err != nil) != tt.wantErr {
				t.Errorf("ValidateParamsApikey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGeoDBBuildBaseURL(t *testing.T) {
	type args struct {
		city string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{city: "New York City"}, "https://wft-geo-db.p.rapidapi.com/v1/geo/cities?limit=1&namePrefix=New%20York%20City&sort=-population"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeoDBBuildBaseURL(tt.args.city); got != tt.want {
				t.Errorf("GeoDBBuildBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
