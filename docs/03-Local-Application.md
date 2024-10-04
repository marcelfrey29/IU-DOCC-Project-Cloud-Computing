# Development of the Local Version of the Application

## Requirements

To develop the inital, local version of the web app, requirements have been defined as User Stories as described in the [Requirements Documentation](02-Requirements.md).

Tasks:

- #15
- #1
- #2
- #3
- #4
- #5

User Stories:

- #6
- #7
- #11
- #10
- #9
- #8
- #12
- #13
- #14
- #16

## Architecture

The Web Application is a Single Page Application (SPA). 
[React](https://react.dev/) and the [NextUI](https://nextui.org/) Framework are used to build the SPA.
The static `.html`, `.css` and `.js` files of the SPA are hosted on a web server.
For the local application, [nginx](https://nginx.org/en/) is used as web server. 

The backend service provides a REST API which allows Create, Read, Update and Delete (CRUD) operations on resources.
All API endpoints will be handled by a single service.
[Golang](https://go.dev/) and the [Fiber](https://gofiber.io/) Web Framework are used to build the backend service.

Application data like Travel Guides and Activities are stored in a database. 
The [DynamoDB Local](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html) NoSQL database is used. 

HTTP Request are performed by the SPA from the browser to the backend service. 
Request and response data transfered in JSON format. 
The backend service reads and writes data to the database.

![Architecture Diagram of the Local Application](assets/architecture-local.svg)

The web application, backend, and database are running in seperate Docker Containers.
[Docker Compose](https://docs.docker.com/compose/) is used to launch the entire application stack with a single command.
Persistent storage is only needed for the database container.
Locally, all applications are running on the same Docker Host.
There is no redundancy or load balancing.
