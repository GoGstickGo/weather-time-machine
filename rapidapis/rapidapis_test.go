package rapidapis

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDarkSkyClient_convertMap(t *testing.T) {
	data := map[string]interface{}{
		"daily": []interface{}{
			map[string]interface{}{
				"apparentTemperatureHigh":     3.53,
				"apparentTemperatureHighTime": 2.5612578e+08,
				"apparentTemperatureLow":      -2.3,
				"apparentTemperatureLowTime":  2.5619412e+08,
				"apparentTemperatureMax":      3.53,
				"apparentTemperatureMaxTime":  2.5612578e+08,
				"apparentTemperatureMin":      -3.26,
				"apparentTemperatureMinTime":  2.5610766e+08,
				"dewPoint":                    1.05,
				"humidity":                    0.9,
				"moonPhase":                   0.18,
				"sunriseTime":                 2.5611096e+08,
				"sunsetTime":                  2.5614738e+08,
				"temperatureHigh":             5.65,
				"temperatureHighTime":         2.5612704e+08,
				"temperatureLow":              -0.84,
				"temperatureLowTime":          2.561652e+08,
				"temperatureMax":              5.65,
				"temperatureMaxTime":          2.5612704e+08,
				"temperatureMin":              -0.84,
				"temperatureMinTime":          2.561652e+08,
				"time":                        2.56086e+08,
				"uvIndex":                     0,
				"uvIndexTime":                 2.561436e+08,
				"windBearing":                 66,
				"windSpeed":                   10.34,
			},
		},
	}
	missingData := map[string]interface{}{
		"none": []interface{}{
			map[string]interface{}{},
		},
	}
	testSlice := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "apparentTemperatureMin:-3.26", "apparentTemperatureMinTime:2.5610766e+08", "dewPoint:1.05", "humidity:0.9", "moonPhase:0.18", "sunriseTime:2.5611096e+08", "sunsetTime:2.5614738e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureLowTime:2.561652e+08", "temperatureMax:5.65", "temperatureMaxTime:2.5612704e+08", "temperatureMin:-0.84", "temperatureMinTime:2.561652e+08", "time:2.56086e+08", "uvIndex:0", "uvIndexTime:2.561436e+08", "windBearing:66", "windSpeed:10.34]]"}
	type fields struct {
		data        map[string]interface{}
		err         error
		date        string
		mapping     Mapping
		countryCode string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{"good", fields{mapping: Mapping{recieved: data}}, testSlice, false},
		{"emptySlice", fields{mapping: Mapping{recieved: missingData}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DarkSkyClient{
				data:        tt.fields.data,
				err:         tt.fields.err,
				mapping:     tt.fields.mapping,
				date:        tt.fields.date,
				countryCode: tt.fields.countryCode,
			}
			got, err := c.convertMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("DarkSkyClient.convertMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("test failed, diff ==> %v\n,", diff)
			}
		})
	}
}

func TestDarkSkyClient_getTempH(t *testing.T) {
	testSlice1 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice2 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "emperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice3 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice4 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "emperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMin:-0.84"}
	type fields struct {
		data        map[string]interface{}
		date        string
		err         error
		mapping     Mapping
		countryCode string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"good", fields{mapping: Mapping{tempField: testSlice1}}, "5.65", false},
		{"missingTempHigh", fields{mapping: Mapping{tempField: testSlice2}}, "5.65", false},
		{"missingTempMax", fields{mapping: Mapping{tempField: testSlice3}}, "5.65", false},
		{"error", fields{mapping: Mapping{tempField: testSlice4}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DarkSkyClient{
				data:        tt.fields.data,
				date:        tt.fields.date,
				err:         tt.fields.err,
				mapping:     tt.fields.mapping,
				countryCode: tt.fields.countryCode,
			}
			got, err := c.getTempH()
			if (err != nil) != tt.wantErr {
				t.Errorf("DarkSkyClient.getTempH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DarkSkyClient.getTempH() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDarkSkyClient_getTempL(t *testing.T) {
	testSlice1 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice2 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "emperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice3 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65"}
	testSlice4 := []string{"[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "emperatureHigh:5.65", "temperatureHighTime:2.5612704e+08"}
	type fields struct {
		data        map[string]interface{}
		date        string
		err         error
		mapping     Mapping
		countryCode string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"good", fields{mapping: Mapping{tempField: testSlice1}}, "-0.84", false},
		{"missingTempLow", fields{mapping: Mapping{tempField: testSlice2}}, "-0.84", false},
		{"missingTempMin", fields{mapping: Mapping{tempField: testSlice3}}, "-0.84", false},
		{"error", fields{mapping: Mapping{tempField: testSlice4}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DarkSkyClient{
				data:        tt.fields.data,
				date:        tt.fields.date,
				err:         tt.fields.err,
				mapping:     tt.fields.mapping,
				countryCode: tt.fields.countryCode,
			}
			got, err := c.getTempL()
			if (err != nil) != tt.wantErr {
				t.Errorf("DarkSkyClient.getTempL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DarkSkyClient.getTempL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeoDBClient_getCountryCode(t *testing.T) {
	testSliceCity1 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity2 := []string{"[map[city:Budapest", "country:Hungary", "countryCode", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}

	type fields struct {
		data    map[string]interface{}
		err     error
		mapping Mapping
		params  Params
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"good", fields{mapping: Mapping{tempField: testSliceCity1}}, "HU", false},
		{"error", fields{mapping: Mapping{tempField: testSliceCity2}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GeoDBClient{
				data:    tt.fields.data,
				err:     tt.fields.err,
				mapping: tt.fields.mapping,
				params:  tt.fields.params,
			}
			got, err := c.getCountryCode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDBClient.getCountryCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeoDBClient.getCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeoDBClient_getCityLocation(t *testing.T) {
	testSliceCity1 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:47.498333333", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity2 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "latitude:", "longitude:19.040833333", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}
	testSliceCity3 := []string{"[map[city:Budapest", "country:Hungary", "countryCode:HU", "id:51643", "name:Budapest", "population:1.752286e+06", "region:Budapest", "regionCode:BU type:CITY wikiDataId:Q1781]]"}

	type fields struct {
		data    map[string]interface{}
		err     error
		mapping Mapping
		params  Params
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		want1   string
		wantErr bool
	}{
		{"good", fields{mapping: Mapping{tempField: testSliceCity1}}, "47.498333333", "19.040833333", false},
		{"missingLatitude", fields{mapping: Mapping{tempField: testSliceCity2}}, "", "19.040833333", true},
		{"missingLong&Lat", fields{mapping: Mapping{tempField: testSliceCity3}}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GeoDBClient{
				data:    tt.fields.data,
				err:     tt.fields.err,
				mapping: tt.fields.mapping,
				params:  tt.fields.params,
			}
			got, got1, err := c.getCityLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDBClient.getCityLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeoDBClient.getCityLocation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GeoDBClient.getCityLocation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParams_validateParams(t *testing.T) {
	type fields struct {
		Day    string
		Month  string
		Year   string
		Apikey string
		City   string
		Writer io.Writer
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"good", fields{Day: "02", Month: "01", Year: "1979", Apikey: "ox9d2scNzWRUcYaVeq4EpIZLP58hecbvn7aeSKUcHACXQt26Md", City: "Dublin"}, "1979-01-02", false},
		{"badCity", fields{Day: "02", Month: "01", Year: "1979", Apikey: "ox9d2scNzWRUcYaVeq4EpIZLP58hecbvn7aeSKUcHACXQt26Md", City: "Dublin3423"}, "", true},
		{"badApikey", fields{Day: "02", Month: "01", Year: "1979", Apikey: "ox9d2scNzWRUcYaVeq4EpIZLP58hecbvn7aeSKUcHACXQt26", City: "Dublin"}, "", true},
		{"badDate", fields{Day: "45", Month: "01", Year: "1979", Apikey: "ox9d2scNzWRUcYaVeq4EpIZLP58hecbvn7aeSKUcHACXQt26Md", City: "Dublin"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Params{
				Day:    tt.fields.Day,
				Month:  tt.fields.Month,
				Year:   tt.fields.Year,
				Apikey: tt.fields.Apikey,
				City:   tt.fields.City,
				Writer: tt.fields.Writer,
			}
			got, err := p.validateParams()
			if (err != nil) != tt.wantErr {
				t.Errorf("Params.validateParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Params.validateParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
