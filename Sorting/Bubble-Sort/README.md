# Bubble Sort

Bubble sort is O(n^2).

The Bubble Sort algorithm has two main parts contributing to the O(n2) score.

1. Comparisons - Two numbers are compared against eachother in each iteration of the inner loop
2. Swaps - A swap may or may not be made, in a worst case scenario we say the number of Swaps equals the number of Comparisons

For an Array of size N, the number of Comparisons will be recurssive:

(N-1) + (N-2) + (N-3) + (N-4) ... len(N)

We never compare the last element of an array with anything, which is why the first Comparison will always be (N-1).

We subtract 1 at each iteration, as the highest number will always bubble up to the i'th index of the array i.e We don't need to check the i'th
element after its iteration has completed.

Array(6,5,4,3,2,1)

Comparisons (Worst Case Scenario):

(6-1) + (6-2) + (6-3) + (6-4) + (6-5) + (6-6)
= 15 Comparisons

Swaps (Worst Case Scenario):

(6-1) + (6-2) + (6-3) + (6-4) + (6-5) + (6-6)
= 15 Steps

Total Operations = 15 Comparisons + 15 Swaps
= 30 Operations

This is Quadratic time. As N input doubles, the number of steps for Bubble sort to complete becomes ~ N^2
