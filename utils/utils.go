package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func BaseUrl(apiurl, latitude, longitude, date, sort string) string {
	baseurl := apiurl + latitude + "," + longitude + "," + date + "T" + "12:00:00" + sort
	return baseurl
}

type Mapping struct {
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

func (m *Mapping) ConvertMap() ([]string, error) {
	m.Pass = make(map[string]string)
	for k, v := range m.Recieved {
		if k == m.MapName {
			m.Value = fmt.Sprintf("%v", v)
			m.Pass[k] = m.Value
		}
	}
	for _, v := range m.Pass {
		m.TempField = strings.Fields(v)
	}
	if m.TempField == nil {
		return m.TempField, fmt.Errorf("failed to convert the map to slice")
	}
	//fmt.Println(m.TempField)
	return m.TempField, nil
}

func (m *Mapping) GetTempH() (string, error) {

	for _, v := range m.TempField {
		if strings.HasPrefix(v, "temperatureHigh:") || strings.HasPrefix(v, "temperatureMax:") {
			m.TempSplit = strings.SplitN(v, ":", 2)
			m.HighTemp = m.TempSplit[1]
		}
	}
	if m.HighTemp == "" {
		return m.HighTemp, fmt.Errorf("failed to get highest temperature")
	}

	return m.HighTemp, nil
}
func (m *Mapping) GetTempL() (string, error) {
	for _, v := range m.TempField {
		if strings.HasPrefix(v, "temperatureLow:") || strings.HasPrefix(v, "temperatureMin:") {
			m.TempSplit = strings.SplitN(v, ":", 2)
			m.LowTemp = m.TempSplit[1]
		}
	}
	if m.LowTemp == "" {
		return m.LowTemp, fmt.Errorf("failed to get lowest temperature")
	}
	return m.LowTemp, nil
}

type DateBuild struct {
	Date    string
	YearInt int
	Error   error
	Day     string
	Month   string
	Year    string
}

var (
	checkY = regexp.MustCompile(`\b(19[4-9][0-9]|20[0-4][0-9]|2050)\b`).MatchString
	checkD = regexp.MustCompile(`^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$`).MatchString
)

func (b *DateBuild) BuildDate() (string, error) {
	b.YearInt, b.Error = strconv.Atoi(b.Year)
	if b.Error != nil {
		return "", fmt.Errorf("Error:%v", b.Error)
	}
	switch {
	case !checkY(b.Year) || b.YearInt > time.Now().Year():
		return b.Year, fmt.Errorf("invalid value for year: %s,year must between 1940-%d", b.Year, time.Now().Year())
	case !checkD(b.Year + "-" + b.Month + "-" + b.Day):
		return b.Year + "-" + b.Month + "-" + b.Day, fmt.Errorf("date format must be: 1978-02-02,invalid value for date: %s-%s-%s, error: %v", b.Year, b.Month, b.Day, b.Error)
	default:
		b.Date = b.Year + "-" + b.Month + "-" + b.Day
	}
	return b.Date, nil
}
