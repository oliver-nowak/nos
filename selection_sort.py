from numpy import loadtxt

def selectionsort(arrayToSort):

    for i in xrange(0, len(arrayToSort)-1):
        j_min = i

        for j in xrange(i+1, len(arrayToSort)):
            if arrayToSort[j] < arrayToSort[j_min]:
                # update to new minimum
                j_min = j

        if j_min != i:
            # swap
            tmp = arrayToSort[i]
            arrayToSort[i] = arrayToSort[j_min]
            arrayToSort[j_min] = tmp


    return arrayToSort


arrayToSort = loadtxt("/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_worst.txt", dtype='int')

print arrayToSort

arraySorted = selectionsort(arrayToSort)

print arraySorted