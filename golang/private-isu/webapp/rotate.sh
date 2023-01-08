#!/bin/sh

mv ./logs/nginx/access.log "./logs/nginx/access.log.$(date +%Y%m%d-%H%M%S)"
docker-compose restart nginx
mv ./logs/mysql/mysql-slow.log "./logs/mysql/mysql-slow.log.$(date +%Y%m%d-%H%M%S)"
docker-compose restart mysql
