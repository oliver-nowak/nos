from numpy import loadtxt

def insertionsort(arrayToSort):
    for i in xrange(1, len(arrayToSort)):

        j = i

        # iterate through the rest of the list
        while j > 0 and arrayToSort[j-1] > arrayToSort[j]:

            # swap
            tmp = arrayToSort[j]
            arrayToSort[j] = arrayToSort[j-1]
            arrayToSort[j-1] = tmp

            # decrement
            j = j - 1

    return arrayToSort


arrayToSort = loadtxt("/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_worst.txt", dtype='int')

print arrayToSort

arraySorted = insertionsort(arrayToSort)

print arraySorted