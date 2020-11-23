
JSON agregator

See https://services.swpc.noaa.gov/json/  for the source data

This program takes source data from the NOAA sources. It reads last estimation for 
solar wind, magnetic field and Flux 10 cm. Each reading comes with its own time stamp.

Data are loaded form the URL' directly and parsed into structures. 
Then we take the last timestamp from the three reeadings and consider it as an actualisation timestamp.

Finally, new JSON is created, which contains data from sources  and timestamp, formatted in accordance with RFC-4339.


Example:

Colombina:7 - JSON 1$ go run NOAAA.go 

* Source data, taken from JSON sources
Solar wind :  563
Magnetic field: Bt  3  Bz:  -1
Flux 10 cm:  88 

* TimeStamp -  2020-11-23T08:48:00Z 

* Created data structure -  map[Bt:3 Bz:-1 Flux:88 SolarWind:563 TimeStamp:2020-11-23T08:48:00Z] 

* Final JSON -  {"Bt":"3","Bz":"-1","Flux":"88","SolarWind":"563","TimeStamp":"2020-11-23T08:48:00Z"} 

Colombina:7 - JSON 1$ 