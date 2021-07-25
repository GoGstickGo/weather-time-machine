
![GoVersion](https://img.shields.io/github/go-mod/go-version/gogstickgo/weather-time-machine)
![RepoSize](https://img.shields.io/github/repo-size/gogstickgo/weather-time-machine)
[![Go](https://github.com/GoGstickGo/weather-time-machine/actions/workflows/go.yml/badge.svg)](https://github.com/GoGstickGo/weather-time-machine/actions/workflows/go.yml)
[![MIT Licence](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)


# weather-time-machine

## motivation

Have you ever wondered when people say too hot / too cold for this time of the year. I don't believe in opinions, I believe in data. So I created this small tool to check temperature for specific city and date.

## rapidApi

Tool uses two APIS for retrieving coordinates for city and temperature for city. This API flow is not by design rather limitations of DarkSky API as it can only understand coordinates. First API call to get coordinates for given city then temperatures.

- <https://rapidapi.com/wirefreethought/api/geodb-cities/>
- <https://rapidapi.com/user/darkskyapis>

## under the hood

Tool optimized for highest population cities. For example if one search for Dublin, temperature will be reflected on Dublin (IE), not on Dublin (USA). Temperature will be shown in Celsius.

## CLI

### Requirement

    - valid rapidapi API key

### Functions

    - search via city name
    - search via city coordinates

## Web

### Work In Progress
