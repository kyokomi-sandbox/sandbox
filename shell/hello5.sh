#!/bin/bash
# ログイン時に歓迎のメッセージをユーザーに表示するスクリプト
if [ "$#" -lt 1 ]; then
  read -p "Enter a name: "
  name=$REPLY
else
  name=$1
fi
echo "Hello $name"
exit 0