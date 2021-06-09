package geodb

import (
	"testing"
)

func TestGeoDB_getCityLocation(t *testing.T) {
	testSliceCity1 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity2 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity3 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}

	type fields struct {
		Recieved     map[string]interface{}
		MapNameGeoDB string
		Error        error
		TempField    []string
		TempSplit    []string
		Latitude     string
		Longitude    string
		CountryCode  string
		Pass         map[string]string
		Value        string
		City         string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		want1   string
		wantErr bool
	}{
		{"good", fields{TempField: testSliceCity1}, "47.498333333", "19.040833333", false},
		{"missingLatitude", fields{TempField: testSliceCity2}, "", "19.040833333", true},
		{"missingLong&Lat", fields{TempField: testSliceCity3}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GeoDB{
				Recieved:     tt.fields.Recieved,
				MapNameGeoDB: tt.fields.MapNameGeoDB,
				Error:        tt.fields.Error,
				TempField:    tt.fields.TempField,
				TempSplit:    tt.fields.TempSplit,
				Latitude:     tt.fields.Latitude,
				Longitude:    tt.fields.Longitude,
				CountryCode:  tt.fields.CountryCode,
				Pass:         tt.fields.Pass,
				Value:        tt.fields.Value,
				City:         tt.fields.City,
			}
			got, got1, err := m.getCityLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDB.getCityLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeoDB.getCityLocation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GeoDB.getCityLocation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGeoDB_getCountryCode(t *testing.T) {
	testSliceCity1 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity2 := []string{"[map[city:Budapest", "country:Hungary", "countryCode", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}

	type fields struct {
		Recieved     map[string]interface{}
		MapNameGeoDB string
		Error        error
		TempField    []string
		TempSplit    []string
		Latitude     string
		Longitude    string
		CountryCode  string
		Pass         map[string]string
		Value        string
		City         string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"good", fields{TempField: testSliceCity1}, "HU", false},
		{"error", fields{TempField: testSliceCity2}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GeoDB{
				Recieved:     tt.fields.Recieved,
				MapNameGeoDB: tt.fields.MapNameGeoDB,
				Error:        tt.fields.Error,
				TempField:    tt.fields.TempField,
				TempSplit:    tt.fields.TempSplit,
				Latitude:     tt.fields.Latitude,
				Longitude:    tt.fields.Longitude,
				CountryCode:  tt.fields.CountryCode,
				Pass:         tt.fields.Pass,
				Value:        tt.fields.Value,
				City:         tt.fields.City,
			}
			got, err := m.getCountryCode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDB.getCountryCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeoDB.getCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
