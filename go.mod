module github.com/theNullP0inter/googly

go 1.16

require (
	github.com/certifi/gocertifi v0.0.0-20210507211836-431795d63e8d // indirect
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/evalphobia/logrus_sentry v0.8.2
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/getsentry/sentry-go v0.11.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.8.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-migrate/migrate v3.5.4+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/google/uuid v1.3.0
	github.com/jinzhu/copier v0.3.2
	github.com/json-iterator/go v1.1.11
	github.com/kylelemons/go-gypsy v1.0.0 // indirect
	github.com/lib/pq v1.10.2
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/sarulabs/di v2.0.0+incompatible
	github.com/sarulabs/di/v2 v2.4.2
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0 // indirect
	github.com/thedevsaddam/govalidator v1.9.10
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)

replace github.com/theNullP0inter/googly/command => ./command

replace github.com/theNullP0inter/googly/logger => ./logger

replace github.com/theNullP0inter/googly/model => ./model

replace github.com/theNullP0inter/googly/rdb => ./rdb

replace github.com/theNullP0inter/googly/route => ./route

replace github.com/theNullP0inter/googly/service => ./service

replace github.com/theNullP0inter/googly/resource => ./resource

replace github.com/theNullP0inter/googly/controller => ./controller

replace github.com/theNullP0inter/googly/app => ./app

replace github.com/theNullP0inter/googly/app_repository => ./app_repository

replace github.com/theNullP0inter/googly/ingress => ./ingress

// example/rdb_crud

replace github.com/theNullP0inter/googly/example/rdb_crud => ./example/rdb_crud

replace github.com/theNullP0inter/googly/example/rdb_crud/accounts => ./example/rdb_crud/accounts
