__author__ = 'kyokomi'

print('hoge')

import math

print(math.pi)
print(math.sqrt(85))

import random

print(random.random())
print(random.choice([1, 10]))

s = "hogehogeあ"
print(len(s))
print(s, s[1], s[-1])  # インデクシング
print(s[1:8])  # スライシング

print(s + 'hoge')  # 連結
print(s * 5)  # 繰り返し

print(s.find("ge"))
print(s.replace("hoge", "fuga"))
# print(dir(s))

# パターンマッチ

import re

print(re.match("Hello[ \t]*(.*)world", "Hello     Python world").group(1))
print(re.match("/(.*)/(.*)/(.*)", "/v2/api/hoge").groups())