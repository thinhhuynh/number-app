# number-app

## Environment setup

You need to have 
* [Go](https://golang.org/),
* [Node.js](https://nodejs.org/),
* [Yarn](https://www.npmjs.com/package/yarn)
* [Docker](https://www.docker.com/), and
* [Docker Compose](https://docs.docker.com/compose/)
(comes pre-installed with Docker on Mac and Windows)
installed on your computer.

Verify the tools by running the following commands:

```sh
go version
npm --version
docker --version
docker-compose --version
```

If you are using Windows you will also need
[gcc](https://gcc.gnu.org/). It comes installed
on Mac and almost all Linux distributions.

## Start in development mode

In the project directory run the command (you might
need to prepend it with `sudo` depending on your setup):

### * Start database (optional)
```sh
docker-compose -f docker-compose-dev.yml up
```

This starts a local MongoDB on `localhost:27017`.
The database will be populated with test records
from the [init-db.js](init-db.js) file.


### 1. Start API
Navigate to the `server` folder and start the back end:

```sh
cd server

# start API
go run server.go

# Start test
go test ./... --cover -v
```
The back end will serve on http://localhost:8080.


### 2. Start Web UI
Navigate to the `webapp` folder, install dependencies,
and start the front end development server by running:

```sh
cd webapp
yarn
yarn start
```
The application will be available on http://localhost:3000.
 
## Start in production mode

Perform:
```sh
docker-compose up
```
This will build the application and start it together with
its database. Access the application on http://localhost:8080.
