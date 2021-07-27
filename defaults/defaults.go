package defaults

const (
	DarkSkyApi      = "dark-sky.p.rapidapi.com"
	DarkSkyApiUrl   = "https://dark-sky.p.rapidapi.com/"
	DarkSkyApiSortC = "?units=ca&exclude=currently%2Chourly%2Calerts%2Cflags&lang=en"
	DarkSkyApiSortF = "?units=us&exclude=currently%2Chourly%2Calerts%2Cflags&lang=en"
	DarkSkyMap      = "daily"
	DarkSkyMapTemp  = "temperature"

	GeoDBApi     = "wft-geo-db.p.rapidapi.com"
	GeoDBUrl     = "https://wft-geo-db.p.rapidapi.com/v1/geo/cities?limit=1&namePrefix="
	GeoDBUrlSort = "&sort=-population"
	GeoDBMap     = "data"

	RapidApiHeaderKey  = "x-rapidapi-key"
	RapidApiHeaderHost = "x-rapidapi-host"

	GET = "GET"

	TestCity  = "Dublin"
	TestYear  = "1978"
	TestMonth = "12"
	TestDay   = "02"
)
