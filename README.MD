# Boilerplate code for Golang application



## Libs Used

Router: https://github.com/gin-gonic/gin

CLI: https://github.com/spf13/cobra

env management: https://github.com/spf13/viper

Logger: https://github.com/sirupsen/logrus

Dependency Injection: https://github.com/sarulabs/di (v2)

ORM: https://gorm.io/

Migrations: https://github.com/golang-migrate/migrate

Error reporting: https://github.com/getsentry/sentry-go


## Instructions to run

To Run, use `raise.sh`.
* Use `--reset` to delete all the images and pull/recreate
* Use `--reset --reset_vol` to delete all the images & volumes
* Use `--daemon`  to run in daemon mode


## ENV VARIABLES

| Variable           | Description                               |
| ------------------ | ------------------------------------------|
| RDB_URL*           | URL for database connection               |
| RDB_MAX_CONNECTIONS| Max Number pf connections to DB           |
| SENTRY_DSN         | DSN to sentry project                     |
| GIN_MODE           | `debug, test, release` debug by default   |


## Available commands

| Command            | Description                                    |
| ------------------ | ---------------------------------------------- |
| serve              | Runs Server on port 8080                       |
| migrate up         | Migrates up to the latest version of migration |
| migrate down       | Migrates down all migrations                   |
