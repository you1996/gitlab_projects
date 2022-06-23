# gitlab_projects

## To run the code:

##### 1) make intall

##### 2) make run

###### There is other commands you can use like "make test"... take a look at Makefile

## How to query the API?

### endpoints :

` /projects-names-and-stars` => returns json of projects.

`/projects-total-stars` => returns the number of stars

parameter : `number-of-projects`

## Examples of query:
##### Consider the project is running locally and at port 8082:

`curl -X GET "http://localhost:8082/projects-names-and-stars?number-of-projects=200"`

`curl -X GET "http://localhost:8082/projects-total-stars?number-of-projects=200"`

### Docker :

`docker build -t gitlabProjects .`

##### then

`docker run -d -i gitlabProjects`
