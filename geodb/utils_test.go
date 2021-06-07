package geodb

import "testing"

func TestCityLocation(t *testing.T) {
	testSliceCity1 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity2 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity3 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}

	type args struct {
		m *GeoDB
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{"good", args{&GeoDB{TempField: testSliceCity1}}, "47.498333333", "19.040833333", false},
		{"missingLatitude", args{&GeoDB{TempField: testSliceCity2}}, "", "19.040833333", true},
		{"missingLong&Lat", args{&GeoDB{TempField: testSliceCity3}}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getCityLocation(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("CityLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CityLocation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CityLocation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCountryCode(t *testing.T) {
	testSliceCity1 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity2 := []string{"[map[city:Budapest", "country:Hungary", "countryCode", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	type args struct {
		m *GeoDB
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"good", args{&GeoDB{TempField: testSliceCity1}}, "HU", false},
		{"error", args{&GeoDB{TempField: testSliceCity2}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCountryCode(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountryCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
