# Googly

Googly is just a toolkit that has diverse tools that are set to work together to get microservices up & running as with minimal effort


Tools are just bindings on other libraries. Googly bindings acts like glue between these libraries to make then4m work with each other.

In other words, Googly is just some set of opinions to make your develoment time lesser on go

## Tools / Googly Bindings available
***
https://github.com/theNullP0inter/gly_mongo

https://github.com/theNullP0inter/gly_gorm

https://github.com/theNullP0inter/gly_logrus

https://github.com/theNullP0inter/gly_gin

https://github.com/theNullP0inter/gly_grpc


## Libraries Used
***
CLI: https://github.com/spf13/cobra

Env management: https://github.com/spf13/viper

Dependency Injection: https://github.com/sarulabs/di (v2)


## Examples
***
Examples can be found at https://github.com/theNullP0inter/googly-example



## Key Concepts
***

### Googly

Googly is an application manager. it builds all regietered applications using dependency injection and starts them using Ingress

For more on Dependency Injection, check this blog: https://www.sarulabs.com/post/2/2018-06-12/what-is-a-dependency-injection-container-and-why-use-one.html


### App

Generally an application is divided into many sub-apps, `App` represents a single sub-app

In github.com/theNullP0inter/googly-example/rdb_crud, accounts is an `App`


### Ingress

Ingress represents a way to run the application. You can essentially create http_ingress & grpc ingress in the same code base and serve them separately on different containers.

Ingress is just a cobra command that will be registered as a sub command for googly

github.com/theNullP0inter/gly_rdb implements a migrate ingress which can be used for rdb migrations.


### Services & Controllers

Service is where you application logic sits. Controllers will interact with these services

Controller acts as a bridge between service and ingress

you can have a single service serving http for client-side applications & grpc for backend services.


### Resource

Resource is what services will be interacting with. This doesnt necessarily need to be a model.

Sometimes when you maintain your resources on a 3rd party database or service, you can create a "virtual" resource and write services to maintain this resource

### ResourceManager

ResourceManager will help service interact with the resource.

This takes away changing your service and controllers for every new type of a database

Resource Managers are bindings on orms/ document-drivers.

There are bindings for Gorm & MongoDB available.


### Logger

Googly provides a simple logger interface which is referenced in service and controllers. If logger of your choice doesnt implement the interface, you can write a binding with your logger.

There is a binding for logrus available.



Note: Resource, Service & Controller for a MVC pattern

## Usage
***

In your main package, 


Create a `struct` implementing `GooglyRunner` and implement your main functions to create a new Googly & Run




```
package main

type MyGooglyRunner struct{}

func (i *MyGooglyRunner) Inject(builder *di.Builder) {

    // Inject your global dependencies here

}



func (i *MyGooglyRunner) GetIngressPoints(cnt di.Container) []ingress.Ingress {
	return []ingress.Ingress{
        // Add your ingresses here
	}

}


func main() {
	g := &googly.Googly{
		GooglyRunner: &MyGooglyRunner{},
		InstalledApps:   []googly.App{
            // Register your sub apps here
        },
	}
	googly.Run(g)
}

```



`googly.Run(g)`  will do the following things

1. Register global dependencies
2. Build all the registered Apps
3. Registers ingress points
4. Runs root command for ingress. you can pass the ingress comands as the first argument and googly will serve that ingress. ( https://github.com/spf13/cobra#commands )



## More tools to be added
* Kafka
* pushpin integration for websockets