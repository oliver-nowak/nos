import random

order = 2500

i = 0
max_num = order

with open("/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_best.txt", "wb") as fin:
    while i < max_num:
        i = i+1
        fin.write(str(i))
        fin.write("\n")

i = order
max_num = 0

with open("/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_worst.txt", "wb") as fin:
    while i > max_num:
        i = i-1
        fin.write(str(i))
        fin.write("\n")


i = order
max_num = 0

arry = []

while i > max_num:
    i = i-1
    arry.append(i)

random.shuffle(arry)
# print arry

i = 0
max_num = order
print len(arry)
with open("/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_avg.txt", "wb") as fin:
    while i < max_num-1:
        i = i+1
        # print i
        fin.write(str(arry[i]))
        fin.write("\n")