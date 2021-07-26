
![GoVersion](https://img.shields.io/github/go-mod/go-version/gogstickgo/weather-time-machine)
![RepoSize](https://img.shields.io/github/repo-size/gogstickgo/weather-time-machine)
[![Go](https://github.com/GoGstickGo/weather-time-machine/actions/workflows/go.yml/badge.svg)](https://github.com/GoGstickGo/weather-time-machine/actions/workflows/go.yml)
[![MIT Licence](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)


# weather-time-machine

## motivation

Have you ever wondered when people say too hot / too cold for this time of the year. I don't believe in opinions, I believe in data. So I created this small tool to check temperature for specific city and date.

Golang is so much fun!

## rapidApi

Tool uses two APIs for retrieving coordinates for city and temperature for city. This API flow is not by design rather limitations of DarkSky API as it can only understand coordinates. First API call to get coordinates for given city then temperatures.

- <https://rapidapi.com/wirefreethought/api/geodb-cities/>
- <https://rapidapi.com/user/darkskyapis>

## cli

### requirement

    valid rapidapi API key subscribed two APIs mentioned above

### usage

    Use:   "wtm",
    Short: "Weather time machine provides temperature for specific date and city"
    Long: `Weather time machine gets temperatures for specific date and city.
    You must have valid Rapidapi APIkey, please see: <https://docs.rapidapi.com/docs/keys>.
    Cities with more than 1 name it must in quotations as --city "San Francisco".

    Examples:
    # Long version
    wtm city --year 1972 --month 01 --day 12 --city "San Francisco" --apikey 23lk4jh234jkl23h5dsfh345
    # Shorthand version
    wtm city -y 1972 -m 01 -d 12 -c Dublin --apikey 23lk4jh234jkl23h5dsfh345

## web

### work in progress
