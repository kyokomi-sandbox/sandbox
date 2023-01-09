#!/bin/sh

docker-compose stop app
docker-compose build app
docker-compose start app
