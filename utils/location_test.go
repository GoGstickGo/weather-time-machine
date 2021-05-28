package utils

import "testing"

func TestGetCityLocation(t *testing.T) {
	type args struct {
		cit string
	}
	tests := []struct {
		name      string
		args      args
		wantAlti  string
		wantLongi string
		wantErr   bool
	}{
		{"good", args{"Dublin"}, "53.3425", "-6.265833333", false},
		{"wrongMissingCityName", args{"asdsad"}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAlti, gotLongi, err := GetCityLocation(tt.args.cit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCityLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAlti != tt.wantAlti {
				t.Errorf("GetCityLocation() gotAlti = %v, want %v", gotAlti, tt.wantAlti)
			}
			if gotLongi != tt.wantLongi {
				t.Errorf("GetCityLocation() gotLongi = %v, want %v", gotLongi, tt.wantLongi)
			}
		})
	}
}
