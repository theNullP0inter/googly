# Googly

Googly is just a toolkit that has diverse tools that are set to work together to get microservices up & running as with minimal effort

## Tools available
***

Router: https://github.com/gin-gonic/gin

CLI: https://github.com/spf13/cobra

Env management: https://github.com/spf13/viper

Logger: https://github.com/sirupsen/logrus

Dependency Injection: https://github.com/sarulabs/di (v2)

ORM for Rdb: https://gorm.io/

Migrations: https://github.com/golang-migrate/migrate

Mongo DB: https://go.mongodb.org/mongo-driver/mongo


## Key Concepts
***

### Googly

Googly is an application manager. it builds all regietered applications using dependency injection and starts them using Ingress

### App

Generally an application is divided into many sub-apps, `App` represents a single sub-app

in example/rdb_crud, accounts is an `App`

### Ingress

Ingress represents a way to run the application. You can essentially create http_ingress & grpc ingress in the same code base and serve them separately.

Ingress is  nothing but a cobra command that will be registered as a sub command for googly

contrib/rdb implements a migrate ingress which is used for rdb migrations

### Services & Controllers

Service is where you application logic sits. Controllers like(gin, grpc) will interact with these services

Controller acts as a bridge between service and ingress

you can have a single service serving http for client-side applications & grpc for backend services. You need to create 2 different controllers in this case. 


### Model

Model represents a row/document depending on your your database. 
MongoDb & Gorm are available in the toolkit 

### Resource

Resource is what services will be interacting with. This doesnt necessarily need to be a model.

Sometimes when you maintain your resources on a 3rd party database or service, you can create a "virtual" resource and write services to maintain this resource

### ResourceManager

ResourceManager will help you interact with the resource.

This takes away changing your service and controllers for every new type of a database

There are resource managers available for Gorm & MongoDB in the toolkit


### Logger

Googly provides a simple logger interface which is referenced in service and controllers. If logger of your choice doesnt implement the interface, you can write a binding with your logger. There is a binding for logrus in `contrib/` 




## Running Examples

Examples are present in `examples/` directory

Mongo CRUD

```
docker-compose -f example/mongo_crud/docker-compose.yaml up --build
```


RDB CRUD

```
docker-compose -f example/rdb_crud/docker-compose.yaml up --build
```


## Pending tasks
* Unit Tests
* Better documentation
* Add CI/CD pipeline
* Add Grpc example


## More tools to be added
* Kafka
* pushpin integration for websockets