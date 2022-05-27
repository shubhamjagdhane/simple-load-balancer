# Simple Load Balancer

Implemented simple load balancer.

## About
In computing, load balancing refers to the process of distributing a set of tasks over a set of resources, with the aim of making their overall processing more efficient. Load balancing can optimize the response time and avoid unevenly overloading some compute nodes while other compute nodes are left idle. - from Wiki

### Prerequisites
```
Docker
Docker-Compose
```

### Installation

1. [Docker](https://www.docker.com/):
If you don't have docker install, please install it from [Here](https://docs.docker.com/engine/install/).
2. [Docker-Compose](https://docs.docker.com/compose/):
If you don't have docker-compose install, please install it from [Here](https://docs.docker.com/compose/install/).

## Getting Started

1. Clone the project.
2. cd simple-load-balancer
3. Please execute `docker-compose build && docker-compose up -d` commands to start the containers.
4. Be sure that your machine is available for ports 9000
5. Please wait while all containers get started
6. Use `docker ps` command to see whether all containers are started or not
7. Make sure that `simpleloadbalancer` containers is running
8. If they are working please open `http://localhost:9000/proxy` in any browser of your machine, if you get any response indicate our load balancer is running
9. To create a simple dummy server hit `chmod +x ./example/dummy-server && ./example/dummy-server 8080` where 8080 is port
10. You can start multiple instance of dummy-server by passing difference port
11. Please pass your machine's IP address with this port to register the server
10. Please execute `docker-compose down` command to stop the containers.



## Current Services
Currently our Load Balancer has following running services
<table border="2">
    <thead style="font-weight:bold; text-align:center; background-color:#000; font-size:16px">
        <tr>
            <td>Service Name</td>
            <td>Endpoint</td>
            <td>Description</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Register</td>
            <td style="text-align:center">/urls/register</td>
            <td>Register the server into load balancer</td>
        </tr>               
        <tr>
            <td>Proxy</td>
            <td style="text-align:center">/proxy</td>                   
            <td>Get the Healthy Server in Round Robin manner</td>
        </tr>               
        <tr>
            <td>Get Server</td>
            <td style="text-align:center">/servers</td>                   
            <td>Gives your all information about the registered servers</td>
        </tr>               
    </tbody>
</table>


## API

1. http://localhost:9000/urls/register

    ### Method: POST

    ### Request Body

    ``` json 
    {
        "url":"http://192.168.1.7:8081",
        "name":"server-1",
        "health":false
    }
    ```
    ### Headers (Both headers are mandatory)
    * Content-Type: application/json
    * Accept: application/json

2. http://localhost:9000/proxy
    ### Method: GET
    ### Headers (Both headers are mandatory)
    * Content-Type: application/json
    * Accept: application/json

Please use any http client (curl, POSTMAN) to play with the APIs.


## Functionalities:
As there are lots of functionalities load balancer does but to keep the project as simple I've implemented only few as below

* Distribute incoming traffic to the network by efficiently distributing across multiple servers
* Ease of adding new server
* Proper success & error response in case of unhealthy servers


## Note:
```
As deployment of this project is containerized based, please do not pass URL as localhost as hostname instead alway use the >>>> IP address

Otherwise load balancer won't able to identifies the server address as load balancer is running inside the container
```



## Built With Stack

* [Golang](https://golang.org/) - The language is used for implementation(golang version 1.13)
* [Mux](https://github.com/gorilla/mux) - A powerful HTTP router and URL matcher for building Go web servers with 
* [Go Mod](https://golang.org/ref/mod#introduction) - The package dependency management in golang
* [Viper](https://github.com/spf13/viper) - Used to read configuration variables from .env files
* [Golang-Testing](https://golang.org/pkg/testing/) - The unit testing framework 
* [Dockerfile]() - Used to build custom images, each service has it's own dockerfile
* [Docker-Compose]() - To build project as dockerized application

## Versioning
Version 1.0

## Contact:
If you face any issues please feel free to contact
[me.shubhamjagdhane@gmail.com](mailto:me.shubhamjagdhane@gmail.com)
