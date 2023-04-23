#!/bin/bash
# Author: @kyokomi
# Web: kyokomi.dev
# バックアップするファイルと場所を入力するよう促すスクリプト
# ファイルはユーザーの$HOME/binディレクトリから検索され、$HOME内のディレクトリだけにバックアップされます
# Last Edited: 2023/04/23

read -p "What file types do you want to backup " file_suffix
read -p "Which directory do you want to backup to " dir_name

# 指定されたディレクトリが存在しなければ、作成します
test -d ./$dir_name || mkdir -m 700 ./$dir_name

# findコマンドは、検索基準すなわち .shにマッチするファイルをコピーします
# -path, -prune、-oオプションは、バックアップディレクトリをバックアップから除外するためのもの
find ./ -path ./$dir_name -prune -o -name "*$file_suffix" -exec cp {} ./$dir_name/ \;
exit 0