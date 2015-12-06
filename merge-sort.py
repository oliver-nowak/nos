from numpy import loadtxt


def merge(left, right):
    sorted = []

    # iterate through left and right arrays
    # and compare elements
    while len(left) > 0 and len(right) > 0:

        if left[0] <= right[0]:
            i = left[0]
            sorted.append(i)

            left = left[1:]
        else:
            j = right[0]
            sorted.append(j)

            right = right[1:]

    # iterate through whats left;
    # generally only one of these while statements is executed
    while len(left) > 0:
        i = left[0]
        sorted.append(i)

        left = left[1:]

    while len(right) > 0:
        j = right[0]
        sorted.append(j)

        right = right[1:]

    return sorted



def mergesort(arrayToSort):
    arrayLength = len(arrayToSort)
    middle = arrayLength / 2

    # Base case
    if arrayLength == 1:
        return arrayToSort

    # Split array to sort in half
    left = arrayToSort[0:middle]
    right = arrayToSort[middle:arrayLength]

    # Recurse to continue splitting in half until we hit the base case
    left = mergesort(left)
    right = mergesort(right)

    # Merge the split arrays
    return merge(left, right)


arrayToSort = loadtxt("/Users/nowak/Development/go/src/github.com/oliver-nowak/nos/small_worst.txt", dtype='int')

print arrayToSort

arraySorted = mergesort(arrayToSort)

print arraySorted

