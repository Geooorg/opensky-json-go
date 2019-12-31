# opensky-json-go

This little library helps you to retrieve flight data using the 'states' API of OpenSky-Network.org.
For a working example please have a look at FlightParser_test.go

## Features

 - Retrieve all flights at the current point of time
 - Retrieve all flights within a bounding box (using two WGS-84 coordinates) at the current point of time
 - Supports 'authenticated' requests of an OpenSky-Network.org account  

The features are basically passed by providing parameters as environment variables, as explained below. 

## Retrieving flight states within a bounding box

Export the bounding box parameters as environment parameters:
 - OPENSKY_LATITUDE_MIN
 - OPENSKY_LONGITUDE_MIN
 - OPENSKY_LATITUDE_MAX
 - OPENSKY_LONGITUDE_MAX

## Using an OpenSky-Network.org account

Export the credentials as environment parameters:
  
 - OPENSKY_USER
 - OPENSKY_PASSWORD