#!/bin/bash

set -euo

PATH_REPO=${PWD}
PATH_ENV="${PATH_REPO}/env"

mkdir -p ${PATH_ENV}

export PATH_ENV_APP_RDB="${PATH_ENV}/app_rdb.env"
export PATH_ENV_APP="${PATH_ENV}/app.env"



# Setting up DB Config

if [ -f "${PATH_ENV_APP_RDB}" ]; then
    echo "Local DB configuration exists."
    export $(grep -v '^#' "${PATH_ENV_APP_RDB}" | xargs)
else
    echo "Creating local DB configuration..."
    MYSQL_DATABASE=app
    MYSQL_USER=app
    MYSQL_PASSWORD=$(head -c 18 /dev/urandom | base64 | tr -dc 'a-zA-Z0-9' | head -c 12)
    MYSQL_ROOT_PASSWORD=$(head -c 18 /dev/urandom | base64 | tr -dc 'a-zA-Z0-9' | head -c 12)
    APP_RDB_HOST=app_rdb

    echo "APP_RDB_HOST=${APP_RDB_HOST}
MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
MYSQL_DATABASE=${MYSQL_DATABASE}
MYSQL_USER=${MYSQL_USER}
MYSQL_PASSWORD=${MYSQL_PASSWORD}" > "${PATH_ENV_APP_RDB}"

    echo "DB configuration created."
fi




# Setting up App env 

if [ -f "${PATH_ENV_APP}" ]; then
    echo "Local App configuration exists."
    export $(grep -v '^#' "${PATH_ENV_APP}" | xargs)
else
    echo "Creating local auth configuration..."
    
    
    AUTH_USER=admin
    AUTH_PASSWORD=$(head -c 18 /dev/urandom | base64 | tr -dc 'a-zA-Z0-9' | head -c 12)

    echo "RDB_URL=${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${APP_RDB_HOST}:3306)/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local" > "${PATH_ENV_APP}"

    echo "APP configuration created."
fi