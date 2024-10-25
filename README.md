# Go Language Practice Repository

This repo is dedicated to experimenting Go language with small examples...

## 01 Trying to build HTTP server that shows weather of Istanbul

Learn how to build HTTP server listens a port (8080 in my case). Improve the code and instead of printing a basic text, receive JSON from weather API and show the weather conditions to the user. 

>Weather API key can be retrieved from https://www.weatherapi.com/

I followed these steps with example code pieces:

- Started with creating a basic http server that shows a basic text on two different paths: foo and bar.
    - [**Example Code Piece - 1**](https://github.com/tberk-s/go-practices/blob/main/src/01-HTTP-STUFF/basic-http-server/main.go)

- Then I used empty interface (any) and type assertion to receive location, city name, temperature,... and others from an API that returns nested JSON. 

    - [**Example Code Piece - 2**](https://github.com/tberk-s/go-practices/blob/main/src/01-HTTP-STUFF/weather-api-practice-without-struct/main.go)

- This time I used nested structs (not the best practice) and cast the json values into the struct.

    - [**Example Code Piece - 3**](https://github.com/tberk-s/go-practices/blob/main/src/01-HTTP-STUFF/weather-api-practice-with-struct/main.go)

**TODO:** For the future I plan on creating a better webapp where user will be writing the city name and the application will show the results to the user dynamically.


## 02 Data Types In Go

**@wip**