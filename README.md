# Prog2005 Assignment - Country Information Service

Run main.go and go to http://localhost:8080 plus the paths below

Paths:
* /countryinfo/v1/info/{two_letter_country_code}
  - two_letter_country_code is the two letter iso of countries
  - Optinal ?limit={number} for the amount of cities shown
  - Example: /countryinfo/v1/info/no?limit=20
 
* const POPULATION_PATH = /countryinfo/v1/population/{two_letter_country_code}
  - two_letter_country_code is the two letter iso of countries
  - Optinal ?limit=Startyear-Endyear for the amount of cities shown (4 digit year)
  - Example: /countryinfo/v1/population/so?limit=2000-2005
 
* const STATUS_PATH = "/countryinfo/v1/status/