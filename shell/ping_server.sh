#!/bin/bash
# Author: @kyokomi
# サーバーにpingを送るスクリプト
# Last Edited: 2023/04/23
read -p "Which server should be pinged" server_addr
ping -c3 $server_addr 2>&1 > /dev/null || echo "Server Dead"