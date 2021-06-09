package geodb

import (
	"fmt"
	"strings"
)

type GeoDB struct {
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

func (m *GeoDB) convertMap() ([]string, error) {
	m.Pass = make(map[string]string)
	for k, v := range m.Recieved {
		if k == m.MapNameGeoDB {
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

func (m *GeoDB) getCityLocation() (string, string, error) {
	for _, v := range m.TempField {
		switch {
		case strings.HasPrefix(v, "latitude:"):
			m.TempSplit = strings.SplitN(v, ":", 2)
			m.Latitude = m.TempSplit[1]
		case strings.HasPrefix(v, "longitude:"):
			m.TempSplit = strings.SplitN(v, ":", 2)
			m.Longitude = m.TempSplit[1]
		}
	}
	if m.Latitude == "" || m.Longitude == "" {
		return m.Latitude, m.Longitude, fmt.Errorf("failed to get city cordinates")
	}
	return m.Latitude, m.Longitude, nil
}

func (m *GeoDB) getCountryCode() (string, error) {
	for _, v := range m.TempField {
		if strings.HasPrefix(v, "countryCode:") {
			m.TempSplit = strings.SplitN(v, ":", 2)
			m.CountryCode = m.TempSplit[1]
		}
	}
	if m.CountryCode == "" {
		return m.CountryCode, fmt.Errorf("failed to get lowest temperature")
	}
	return m.CountryCode, nil
}
