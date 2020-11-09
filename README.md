# Matching-app
Matching projects with the best participant by industry, job title, and location.

## Introduction
This was a really fun project! The first time I read it, my first impression was it will be an easy one. But at the moment of start coding, I identify some complications about the data and how to match participants with projects. That makes me start to think sharper and focused on the business logic first and then on how I need to manipulate/sanitize the data in order to use it.

## Main Challenges
I like to share the main challenges I face and what decisions I made to achieve them.
* **Participants data in CSV file:** As I mentioned before, I detect that not all participants in the file had the same format for **city** field. That makes me difficult to filter users for locations in the project. The way I solve this is by calling Google Geocoding API for each participant when I load the data, using latitude and longitude, to obtain a normalized **formatedAddress**. This also allows me to detect one row (390) with inconsistent information `40.8516701	-93.2599318`.
* **Scoring Participants:** My implementation is simple, one point for every matching industry + one point if the job title is a full match with the professional titles specified in the project + 0.5 points if I detect an expertise indicator in the job title. But I'm not sure about false positives.
* **Calculating Distances:** I implement the Haversine formula as is suggested in the project description. But at the moment of testing against the `project.json` I found a lot of zero distances as result. At the moment of code, I have tested this case but makes me doubt it. I have 495 results over 500 cases where distance is less than 100km but a strange number of cases with 0 distance.

## Improvements Opportunities
This is a spoiler alert of what I think I could improve. 
* **Dockerization:** Create the docker image for the project that starts a web server to test the application.
* **Gender filter:** I couldn’t take the time to add a gender filter for participants.
* **Hardcoded Maps Key:** The google maps API key is hardcoded in code. This must be an environment variable.
* **CI/CD pipeline:** Add the pipeline definition for the project.
* **CSV loading time:** Takes too much, around 30 seconds. I believe I could improve this using `go functions`.

## Language
I use Golang to develop the project. I have been using Golang for the last 8 months and, as the challenge description says, is the language I feel strongest now.

## Structure
```
.
├── README.md
├── cmd
│   └── http
│       └── main.go
├── go.mod
├── go.sum
├── pkg
│   ├── delivery
│   │   └── http
│   │       ├── handler.go
│   │       └── matching_participants_handler.go
│   ├── infrastructure
│   │   └── storage
│   │       ├── csv_participants_repository.go
│   │       ├── csv_participants_repository_test.go
│   │       └── respondents_data_test.csv
│   └── matching
│       ├── action.go
│       ├── action_test.go
│       ├── distance_service.go
│       ├── distance_service_test.go
│       ├── matching-coverage.html
│       ├── model.go
│       ├── repository.go
│       ├── score_service.go
│       └── score_service_test.go
└── respondents_data_test.csv
```
I develop the project by the Domain-Driven Design approach and Out Side In. That drives me to an isolated domain `matching` with an `Action` that represents the matching project and participant’s action that is requested to the application. `Action` uses a `ParticipantRepository` interface to retrieve participants by address and filter them by distance with the help of the `DistanceService`. Then the participants are evaluated for the project using the `ScoreService`.

In the `infrastructure` folder are located all implementations for external dependencies not related to the domain, like the `CsvParticipantsRepository`. This is a `ParticipantRepository` that load participants from a CSV file.

`delivery` is the layer where are implemented the delivery mechanisms for the domain. In this case, I only implement an Http Rest API. But if a Command Line Interface is needed too, it could implement in this layer.

`cmd` folder stores all different ways to initiate the application. In this case, as it happens in `delivery`, we only have a function that starts the HTTP server to allows the use of the API.

## How to test it
First, clone the project
```
> git clone https://github.com/carlos-rodrigo/matching-app.git
```
Move to the main.go file location
```
> cd cmd/http/
> go run main.go
```
Now we must see the following on the console
```
2020/11/09 11:39:30 Loading Participants Storage...
2020/11/09 11:40:04 FormattedAddress can't be obtained "Not found address for location"
2020/11/09 11:40:14 Participants loaded....
```
And to test it we can just make this call that contains the content of `project.json` file

```
curl --location --request GET 'http://localhost:8080/matching/' \
--header 'Content-Type: application/json' \
--data-raw '{

    "numberOfParticipants" : 8,
    "timezone" : "America/New_York",
    "cities" : [ 
        {
            "location" : {
                "id" : "ChIJOwg_06VPwokRYv534QaPC8g",
                "city" : "New York",
                "state" : "NY",
                "country" : "US",
                "formattedAddress" : "New York, NY, USA",
                "location" : {
                    "latitude" : 40.7127753,
                    "longitude" : -74.0059728
                }
            }
        }, 
        {
            "location" : {
                "id" : "ChIJ60u11Ni3xokRwVg-jNgU9Yk",
                "city" : "Philadelphia",
                "state" : "PA",
                "country" : "US",
                "formattedAddress" : "Philadelphia, PA, USA",
                "location" : {
                    "latitude" : 39.9525839,
                    "longitude" : -75.1652215
                }
            }
        }, 
        {
            "location" : {
                "id" : "ChIJEcHIDqKw2YgRZU-t3XHylv8",
                "city" : "Miami",
                "state" : "FL",
                "country" : "US",
                "formattedAddress" : "Miami, FL, USA",
                "location" : {
                    "latitude" : 25.7616798,
                    "longitude" : -80.1917902
                }
            }
        }, 
        {
            "location" : {
                "id" : "ChIJS5dFe_cZTIYRj2dH9qSb7Lk",
                "city" : "Dallas",
                "state" : "TX",
                "country" : "US",
                "formattedAddress" : "Dallas, TX, USA",
                "location" : {
                    "latitude" : 32.7766642,
                    "longitude" : -96.7969879
                }
            }
        }, 
        {
            "location" : {
                "id" : "ChIJE9on3F3HwoAR9AhGJW_fL-I",
                "city" : "Los Angeles",
                "state" : "CA",
                "country" : "US",
                "formattedAddress" : "Los Angeles, CA, USA",
                "location" : {
                    "latitude" : 34.0522342,
                    "longitude" : -118.2436849
                }
            }
        }, 
        {
            "location" : {
                "id" : "ChIJIQBpAG2ahYAR_6128GcTUEo",
                "city" : "San Francisco",
                "state" : "CA",
                "country" : "US",
                "formattedAddress" : "San Francisco, CA, USA",
                "location" : {
                    "latitude" : 37.7749295,
                    "longitude" : -122.4194155
                }
            }
        }, 
        {
            "location" : {
                "id" : "ChIJGzE9DS1l44kRoOhiASS_fHg",
                "city" : "Boston",
                "state" : "MA",
                "country" : "US",
                "formattedAddress" : "Boston, MA, USA",
                "location" : {
                    "latitude" : 42.3600825,
                    "longitude" : -71.0588801
                }
            }
        }, 
        {
            "location" : {
                "id" : "ChIJW-T2Wt7Gt4kRKl2I1CJFUsI",
                "city" : "Washington",
                "state" : "DC",
                "country" : "US",
                "formattedAddress" : "Washington, DC, USA",
                "location" : {
                    "latitude" : 38.9071923,
                    "longitude" : -77.0368707
                }
            }
        }
    ],
    "genders" : "N/A",
    "country" : "US",
    "incentive" : 120,
    "name" : "Looking for software engineers experienced with Kafka",
    "professionalJobTitles" : [ 
        "Developer", 
        "Software Engineer", 
        "Software Developer", 
        "Programmer", 
        "Java Developer", 
        "Java/J2EE Developer", 
        "Java Full Stack Developer", 
        "Java Software Engineer", 
        "Java Software Developer", 
        "Application Architect", 
        "Application Developer"
    ],
    "professionalIndustry" : [ 
        "Banking", 
        "Financial Services", 
        "Government Administration", 
        "Insurance", 
        "Retail", 
        "Supermarkets", 
        "Automotive",
        "Computer Software"
    ],
    "education":[
        "bachelordegree",
        "masterdegree"
    ]
}'
```
