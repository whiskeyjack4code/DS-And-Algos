# Linear Search

Liner search is O(n).

In the worst case scenario, where the value to find is located at the end of the array (or not in the array at all),
we need to loop over N elements in an Array of size N.

A computer can jump to any memory location in 1 operation and read a value at that location, but it does
not know the values in the Array without going to each location in turn and reading them.

When declaring an array (or a Slice in go which points back to an underlying array), we get a pointer
to the first element in that array. The computer knows the size of the Array at compile time, so it therefore
knows how much space the array needs, by allocating N\*T blocks of memory, where N is the size of the Array and T is the
data type held in the array.

a = Array[5] {1,2,3,4,5}

1. a is an array that holds Integer values which are generally 4 bytes in size.
2. The computer allocates 5\*4 blocks of contiguous memory and passes back a pointer to it.
3. a now holds a pointer to the first element of the Array: - &Array[0]

Given the above, we can say that a computer can jump to any array index in 1 operation, and read the value
at that location. But it does not know all values in the array without checking each in turn.

Linear search therefore needs to iterate over N elements in a worst case scenario.
