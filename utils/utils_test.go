package utils

import (
	"reflect"
	"testing"
	"weather-api/defaults"
)

func TestMapping_ConvertMap(t *testing.T) {
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
		Recieved  map[string]interface{}
		Pass      map[string]string
		MapName   string
		Value     string
		Error     error
		HighTemp  string
		LowTemp   string
		TempField []string
		TempSplit []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{"good", fields{Recieved: testMapInterface, MapName: defaults.DarkSkyMap}, testSlice, false},
		{"emptySlice", fields{Recieved: testMapInterfaceMissingDaily, MapName: defaults.DarkSkyMap}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapping{
				Recieved:  tt.fields.Recieved,
				Pass:      tt.fields.Pass,
				MapName:   tt.fields.MapName,
				Value:     tt.fields.Value,
				Error:     tt.fields.Error,
				HighTemp:  tt.fields.HighTemp,
				LowTemp:   tt.fields.LowTemp,
				TempField: tt.fields.TempField,
				TempSplit: tt.fields.TempSplit,
			}
			got, err := m.ConvertMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapping.ConvertMap() error = %v, wantErr %v", err, tt.wantErr)
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

func TestMapping_GetTempH(t *testing.T) {
	testSlice1 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice2 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "emperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "emperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice3 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMax:5.65", "temperatureMin:-0.84"}
	testSlice4 := []string{"map[data:[map[apparentTemperatureHigh:3.53", "apparentTemperatureHighTime:2.5612578e+08", "apparentTemperatureLow:-2.3", "apparentTemperatureLowTime:2.5619412e+08", "apparentTemperatureMax:3.53", "apparentTemperatureMaxTime:2.5612578e+08", "temperatureHigh:5.65", "temperatureHighTime:2.5612704e+08", "temperatureLow:-0.84", "temperatureMin:-0.84"}

	type fields struct {
		Recieved  map[string]interface{}
		Pass      map[string]string
		MapName   string
		Value     string
		Error     error
		HighTemp  string
		LowTemp   string
		TempField []string
		TempSplit []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"good", fields{TempField: testSlice1}, "5.65", false},
		{"missingTempHigh", fields{TempField: testSlice3}, "5.65", false},
		{"missingTempMax", fields{TempField: testSlice4}, "5.65", false},
		{"error", fields{TempField: testSlice2}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapping{
				Recieved:  tt.fields.Recieved,
				Pass:      tt.fields.Pass,
				MapName:   tt.fields.MapName,
				Value:     tt.fields.Value,
				Error:     tt.fields.Error,
				HighTemp:  tt.fields.HighTemp,
				LowTemp:   tt.fields.LowTemp,
				TempField: tt.fields.TempField,
				TempSplit: tt.fields.TempSplit,
			}
			got, err := m.GetTempH()
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapping.GetTempH() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mapping.GetTempH() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateBuild_BuildDate(t *testing.T) {
	type fields struct {
		Date    string
		YearInt int
		Error   error
		Day     string
		Month   string
		Year    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"good#1", fields{Year: "1941", Month: "03", Day: "02"}, "1941-03-02", false},
		{"good#2", fields{Year: "1980", Month: "12", Day: "02"}, "1980-12-02", false},
		{"good#3", fields{Year: "2021", Month: "09", Day: "07"}, "2021-09-07", false},
		{"wrongYearPast", fields{Year: "1920", Month: "09", Day: "07"}, "1920", true},
		{"wrongYearFuture", fields{Year: "2029", Month: "09", Day: "07"}, "2029", true},
		{"wrongMonth", fields{Year: "1958", Month: "20", Day: "07"}, "1958-20-07", true},
		{"wrongDay", fields{Year: "1956", Month: "03", Day: "99"}, "1956-03-99", true},
		{"wrongDateLetters#1", fields{Year: "1948", Month: "aa", Day: "07"}, "1948-aa-07", true},
		{"wrongDateLetters#1", fields{Year: "1972", Month: "09", Day: "0b"}, "1972-09-0b", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &DateBuild{
				Date:    tt.fields.Date,
				YearInt: tt.fields.YearInt,
				Error:   tt.fields.Error,
				Day:     tt.fields.Day,
				Month:   tt.fields.Month,
				Year:    tt.fields.Year,
			}
			got, err := b.BuildDate()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateBuild.BuildDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DateBuild.BuildDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
