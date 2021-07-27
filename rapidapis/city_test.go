package rapidapis

import (
	"flag"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	apiKey string
)

func init() {
	flag.StringVar(&apiKey, "Apikey", "", "Please provide RapidAPI key with GeoDB cities subscription to run integration test")
}

func Test_geoDBClient(t *testing.T) {
	if apiKey == "" {
		t.Skip("No ApiKey provided")
	}
	data := map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"city":        "Dublin",
				"country":     "Ireland",
				"countryCode": "IE",
				"id":          3.453097e+06,
				"latitude":    53.349722222,
				"longitude":   -6.260277777,
				"name":        "Dublin",
				"population":  1.173179e+06,
				"region":      "Leinster",
				"regionCode":  "L",
				"type":        "CITY",
				"wikiDataId":  "Q1761",
			},
		},
		"links": []interface{}{
			map[string]interface{}{
				"href": "/v1/geo/cities?offset=0&limit=1&namePrefix=Dublin&sort=-population",
				"rel":  "first",
			},
			map[string]interface{}{
				"href": "/v1/geo/cities?offset=1&limit=1&namePrefix=Dublin&sort=-population",
				"rel":  "next",
			},
			map[string]interface{}{
				"href": "/v1/geo/cities?offset=26&limit=1&namePrefix=Dublin&sort=-population",
				"rel":  "last",
			},
		},
		"metadata": map[string]interface{}{
			"currentOffset": 0.0,
			"totalCount":    27.0,
		},
	}

	type args struct {
		p Params
	}
	tests := []struct {
		name    string
		args    args
		want    *GeoDBClient
		wantErr bool
	}{
		{"good", args{p: Params{"23", "03", "1988", apiKey, "Dublin", "", "", false, nil}}, &GeoDBClient{data: data, params: Params{"23", "03", "1988", apiKey, "Dublin", "", "", false, nil}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gdClient(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("gdClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got.data, tt.want.data); diff != "" {
				t.Errorf("test failed, diff ==> %v\n,", diff)
			}
		})
	}
}
