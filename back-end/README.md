# README for backend

## Overview

This is a generic go server with no framework. Currently, the main responsibly of the server is to interact with mongo.

## SetUp

Download Postman here https://www.postman.com/downloads/. Postman is a tool for making API calls and can be used to simulate api calls to servers running on your local machine.

Download Docker Desktop here https://www.docker.com/products/docker-desktop/. Docker is a "tool" that lets you run different applications on your computer or any server regardless of the hardware without needing to install all dependencies like languages, libraries, databases.

Download Studio3T here https://studio3t.com/download/. Studio3T is a tool that lets you visualize mongoDB database and the contents inside of them.

## Configuration

There is a Dockerfile in this folder indicating that the go server has been containerized. There is also two docker compose yaml files in the root directory. Compose files let you spin up multiple containers based on some image and they can communicate to each other on a local network using the container names. The image that we use for the go server is the one built by the Dockerfile here.

### Docker compose

To run the application using compose run the following commands:

- Before running the next command make sure your docker desktop is running. After the command you will be able to see the containers running in your docker desktop.
- `docker compose -f compose.be.yaml up --build` This command will use a specific compose file with the -f flag and rebuild the images everytime. The compose.be.yaml file is to start only the go and the mongo server for development with the backend API. The other yaml file will start the Front end server as well for full stack development.

The go server will be available on localhost:8080 and the mongoDB will be available on localhost:27017.
In Postman the api calls will start with http://localhost:8080 and append the api path as needed.

### studio3T

In studio3T you will need to set up a configuration for the first time. The following times you will just connect to your existing configuration.

- In studio3T click on connect in the upper left hand corner.
- Click on new connection. You will be in the first tab named "Server". In the Server field put "localhost" and in the port put "27017". Select connection type to be Standalone. Then click on the "Authentication" tab. Here put "root" as the user name and "example" as the password. Put "admin" in the Authentication DB. At the very bottom put "admin,local". For the "Authentication mode" field at the top select "Legacy". Finally click save.
- Now you can connect to this configuration and see the UI populate with the data in the DB.
- Note: connection is only possible when docker compose is running.

To populate dummy data into the mongo DB run the following python file:

- TODO

## Development

Now when you run the docker compose command all containers will start on your local machine. Use postman to make API calls and studio3T to visualize the data.

The go image is built with a library that restarts the server when it detects that go files have been changed. Hence, when compose is running any changes made to the code will restart the server and will be applied. That way you can develop and test the server without needing to rerun any commands.

To see the output logs from the go server head over to your docker desktop.

- Click on the containers tab in the upper left hand corner. Here there will be a collabsource entry. Expand it and there will be a container with the name "backend-1".
- Click on the three dots in the right hand side of this container and click on "open in terminal".
- Click on the logs tab and you will see logs here.

## Running without Docker

### starting the app

- run: `$ go run back-end/main.go`
  - this will install necessary dependencies

### tests (back end)

- `$ cd tests/testName && go test`
  - for contract tests, start server first
