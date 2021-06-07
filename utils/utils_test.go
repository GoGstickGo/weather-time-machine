package utils

import (
	"reflect"
	"testing"
	"weather-api/defaults"
)

func TestBuildDate(t *testing.T) {
	type args struct {
		b *DateBuild
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"good#1", args{&DateBuild{Year: "1941", Month: "03", Day: "02"}}, "1941-03-02T12:00:00", false},
		{"good#2", args{&DateBuild{Year: "1980", Month: "12", Day: "02"}}, "1980-12-02T12:00:00", false},
		{"good#3", args{&DateBuild{Year: "2021", Month: "09", Day: "07"}}, "2021-09-07T12:00:00", false},
		{"wrongYearPast", args{&DateBuild{Year: "1920", Month: "09", Day: "07"}}, "1920", true},
		{"wrongYearFuture", args{&DateBuild{Year: "2029", Month: "09", Day: "07"}}, "2029", true},
		{"wrongMonth", args{&DateBuild{Year: "1958", Month: "20", Day: "07"}}, "1958-20-07", true},
		{"wrongDay", args{&DateBuild{Year: "1956", Month: "03", Day: "99"}}, "1956-03-99", true},
		{"wrongDateLetters#1", args{&DateBuild{Year: "1948", Month: "aa", Day: "07"}}, "1948-aa-07", true},
		{"wrongDateLetters#1", args{&DateBuild{Year: "1972", Month: "09", Day: "0b"}}, "1972-09-0b", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildDate(tt.args.b)
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

func TestGetTempH(t *testing.T) {
	testSlice1 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice2 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "emperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "emperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice3 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice4 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMin:-0.84"}

	type args struct {
		m *Mapping
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"good", args{&Mapping{TempField: testSlice1}}, "5.65", false},
		{"missingTempHigh", args{&Mapping{TempField: testSlice3}}, "5.65", false},
		{"missingTempMax", args{&Mapping{TempField: testSlice4}}, "5.65", false},
		{"error", args{&Mapping{TempField: testSlice2}}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTempH(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTempH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTempH() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertMap(t *testing.T) {
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
	type args struct {
		m *Mapping
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"good", args{&Mapping{Recieved: testMapInterface, MapName: defaults.DarkSkyMap}}, testSlice, false},
		{"emptySlice", args{&Mapping{Recieved: testMapInterfaceMissingDaily, MapName: defaults.DarkSkyMap}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertMap(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertMap() error = %v, wantErr %v", err, tt.wantErr)
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
