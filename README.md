# Cloud Resume Challenge

Purpose to build a website to host my resume.


## Architecture

Web
|
API
|
DB


### Web layer

Simple webpage using HTML5, CSS, Javascript.

To run locally, it is using nginx as defined in the docker-compose.yml.


### API layer

RESTful APIs built using Go. Intended to deployed as containers.

To run this layer locally, the Go app needs to be containerized and then run as defined in the docker-compose.yml.


### DB layer

Using Firebase on GCP, since the solution is intended to be ran on GCP.



# Web server

Using NGINX in a docker container.


to run environment

`docker-compose up --detach`