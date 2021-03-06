# Minimum Requirements
- Docker version 18.03.1-ce, build
- docker-compose version 1.19.0
- docker-py version: 2.7.0
- CPython version: 2.7.13
- GCC

## Setup
- git clone git@github.com:yoesoff/Foreign-Currency-BE-Exercise.git
- docker-compose up
- open http://localhost:9000/exchanges on your browser

## Available APIs
- List of exchange http://localhost:9000/api/exchanges [get]
- Delete item http://localhost:9000/api/exchanges/:id/delete [delete]
- Create item http://localhost:9000/api/exchanges [post]
- Update item http://localhost:9000/api/exchanges/:id/update [put]
- Get last 7 days http://localhost:9000/api/exchanges/pairlast7days/:currency1/:currency2 [get] //Average taken/calculated from this API
- FYI, Both Create and Update are using normal form-data

## Simple GUI/Explorer to Browse your Data
- open http://localhost:9000/exchanges on your browser

# Used Programming Technologies (Inside container)
- go version go1.10.3 linux/amd64 
- Revel Framework v0.19.0 (2018-02-06)
- Gorm ORM for Go

# Snapshot and Images
- https://drive.google.com/drive/folders/1kzsCh5BKZHvjVVYlfANwdfWuCysaDG88?usp=sharing




# Go Revel Basic Information

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

