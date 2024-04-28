# Binary Search

Binary search is O(logN).

Only works on an ordered array.

In the worst case scenario, where the value to find is located at the end of the array (or not in the array at all),
we reduce the array down by 50% at each iteration (guess).

As the computer knows the start, size, and therefore the end of the array, we can get the midpoint
of the array by adding the lower bound to the upper bound and dividng by 2. Note: be cautious to the off by one error
here, and minus 1 from your array length to get the true upper bound.

1. When we have the middle index of the array, we check the value at that index and compare it to our search value.

2. If it matches, we return it and call it a day.

3. If the search value is lower than the value at the midpoint, we can elimiate all values to the right of the midpoint,
   and set a new 'upper bound' of (mid - 1). We subtract 1 from the current mid becase we already checked the midvalue
   in the first step.

4. If the search value is higher than the value at the midpoint, we can elimiate all values to the left of the midpoint.
   and set a new 'lower bound' of (mid + 1). We add 1 to the current mid becase we already checked the mid value
   in the first step.

5. With each comparison (or guess), we always eliminate half of values i.e. the values it's impossible to be.

This process of elimination is Logarithmic (the inverse of exponents).
If we double the array, we only add 1 extra step (the comparison step), before eliminating half the elements again.

a = Array[11]{2,4,6,10,25,50,100}
mid = (0 (lower bound) + 6 (upper bound)) / 2
mid = 3
midValue[mid] = 10

{2,4,6,10,25,50,100}

searchValue = 50

midValue == searchValue ? No.
searchValue < midValue ? No.
lower = mid + 1

{25,50,100}

lower = 4 + 1 = 5
mid = 5 ((lower bound) + 6 (upper bound)) / 2
mid = 5 (rounded down)
midValue[mid] = 50

midValue == searchValue ? Yes. Return the mid.

return index 5
