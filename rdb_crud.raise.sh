#!/bin/bash

set -euo

. ./rdb_crud.config.sh

# Default values
reset=false
reset_vol=false
daemon=false


while [[ $# -gt 0 ]]
do
    key="$1"

    case $key in
        --reset)
        reset=true
        shift # past argument
        ;;
        --reset_vol)
        reset_vol=true
        shift # past argument
        ;;
        --daemon)
        daemon=true
        shift # past argument
        ;;
        *)    # unknown option
        POSITIONAL+=("$1") # save it in an array for later
        shift # past argument
        ;;
    esac

done


# Reset if needed
if $reset; then
    echo "Will reset the environment..."
    sleep 1
    if $reset_vol; then
        docker-compose down --rmi all -v --remove-orphans
    else
        docker-compose down --rmi all --remove-orphans
    fi

else
    echo "Will down the environment."
    docker-compose down
fi


if $daemon; then
    echo "Will start the environment in daemon mode"
    DOCKER_CONTENT_TRUST=1 docker-compose up -d
else
    echo "Will start the environment"
    DOCKER_CONTENT_TRUST=1 docker-compose up
fi

echo "sleeping for 3 seconds..."
sleep 3