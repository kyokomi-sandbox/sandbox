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

D = {"a": 1, "b": 2, "c": 3}
ks = list(D.keys())
print(ks)
ks.sort()
print(ks)

print("--------------")

for key in ks:
    print(key, "=>", D[key])

print("--------------")

for key in sorted(D):
    print(key, "=>", D[key])

squares = [x ** 2 for x in [1, 2, 3, 4, 5]]
print(squares)

if not "d" in D:
    print("d ない")
elif "e" in D:
    print("d ある e もある")
else:
    print("d ある e ない")

f = open("sample.txt", "w")
f.write("Hello\n")
f.write("world\n")
f.close()

# LとMが対応するオブジェクトは別モノなのでisはFALSE

L = [2, 3, 4]
M = [2, 3, 4]

print(L, "==", M, L == M)
print(L, "is", M, L is M)

# XとYはキャッシュが行われるためTRUEになる

X = 2
Y = 2

print(X, "==", Y, X == Y)
print(X, "is", Y, X is Y)

print(int("42"), repr(12), str(23))

S = "hogehogefugahoge"
print(S.replace("hoge", "piyo"))
print("hogehoge %s == %d" % ("fuga", 11111))
print("%(X)d %(Y)d" % {"X": 100, "Y": 200})
print(list(S))