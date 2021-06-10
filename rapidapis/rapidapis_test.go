package rapidapis

import (
	"reflect"
	"testing"
)

func TestGeoDBClient_convertMap(t *testing.T) {
	dailyInterface := map[string]interface{}{}
	dailyInterface["apparentTemperatureHigh"] = 3.53
	dailyInterface["apparentTemperatureHighTime"] = 2.5612578e+08
	dailyInterface["apparentTemperatureLow"] = -2.3
	dailyInterface["apparentTemperatureLowTime"] = 2.5619412e+08
	dailyInterface["apparentTemperatureMax"] = 3.53
	dailyInterface["apparentTemperatureMaxTime"] = 2.5612578e+08
	dailyInterface["apparentTemperatureMin"] = -3.26
	dailyInterface["apparentTemperatureMinTime"] = 2.5610766e+08
	dailyInterface["dewPoint"] = 1.05
	dailyInterface["humidity"] = 0.9
	dailyInterface["moonPhase"] = 0.18
	dailyInterface["sunriseTime"] = 2.5611096e+08
	dailyInterface["sunsetTime"] = 2.5614738e+08
	dailyInterface["temperatureHigh"] = 5.65
	dailyInterface["temperatureHighTime"] = 2.5612704e+08
	dailyInterface["temperatureLow"] = -0.84
	dailyInterface["temperatureLowTime"] = 2.561652e+08
	dailyInterface["temperatureMax"] = 5.65
	dailyInterface["temperatureMaxTime"] = 2.5612704e+08
	dailyInterface["temperatureMin"] = -0.84
	dailyInterface["temperatureMinTime"] = 2.561652e+08
	dailyInterface["time"] = 2.56086e+08
	dailyInterface["uvIndex"] = 0
	dailyInterface["uvIndexTime"] = 2.561436e+08
	dailyInterface["windBearing"] = 66
	dailyInterface["windSpeed"] = 10.34
	dataInterface := map[string]interface{}{}
	dataInterface["data"] = dailyInterface
	testMapInterface := map[string]interface{}{}
	testMapInterface["offset"] = 1.00
	testMapInterface["latitude"] = 47.498333333
	testMapInterface["longtidue"] = 19.040833333
	testMapInterface["daily"] = dataInterface
	testMapInterfaceMissingDaily := map[string]interface{}{}
	testMapInterfaceMissingDaily["offset"] = 1.00
	testMapInterfaceMissingDaily["latitude"] = 47.498333333
	testSlice := []string{"map[data:map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "apparentTemperatureMin:-3.26", "apparentTemperatureMinTime:2.5610766e+08", "dewPoint:1.05", "humidity:0.9", "moonPhase:0.18", "sunriseTime:2.5611096e+08", "sunsetTime:2.5614738e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureLowTime:2.561652e+08", "temperatureMax:5.65 temperatureMaxTime:2.5612704e+08", "temperatureMin:-0.84 temperatureMinTime:2.561652e+08", "time:2.56086e+08", "uvIndex:0", "uvIndexTime:2.561436e+08", "windBearing:66", "windSpeed:10.34]]"}
	type fields struct {
		data    map[string]interface{}
		err     error
		mapping Mapping
		params  Params
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{"good", fields{data: testMapInterface}, testSlice, false},
		{"emptySlice", fields{data: testMapInterfaceMissingDaily}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GeoDBClient{
				data:    tt.fields.data,
				err:     tt.fields.err,
				mapping: tt.fields.mapping,
				params:  tt.fields.params,
			}
			got, err := c.convertMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeoDBClient.convertMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "emptySlice" && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertMap() =  got: %v\n want: %v", got, tt.want)
			}
			if tt.name == "good" && len(got) == len(tt.want) {
				t.Errorf("ConvertMap() =  got: %v\n want: %v", got, tt.want)
			}
		})
	}
}

func TestDarkSkyClient_getTempH(t *testing.T) {
	testSlice1 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice2 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "emperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "emperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice3 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice4 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMin:-0.84"}
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
