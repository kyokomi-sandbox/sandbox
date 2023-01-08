#!/bin/sh

alp json --file logs/nginx/access.log
pt-query-digest logs/mysql/mysql-slow.log

