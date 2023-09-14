# Integer Ranges

This is a simple library for working with integer ranges.
Given a list of integers, it will return a list of ranges.

## Installation
```shell
$ go install github.com/csgriffis/ranges
```

## Usage
```shell
$ ranges --input=1,2,3,11,12,33,55,56,57
["1-3","11-12","33","55-57"]
```

## Nerdy Things
### Time Complexity
The input list is sorted in place, so the time complexity is O(n*log(n)*log(n)); 
based on the algorithm used in the standard library.
This is worst-case, and assumes that the input list is not already sorted.

This ends up being the overall time complexity, since the output list is
constructed in a single pass over the input list, O(n).

### Space Complexity
The input list is sorted in place, so the space complexity is O(n).

The output list could be, in a worst case, the same size as the input list;
thus the additional space complexity is O(n).

Tracking each range is utilized via a slice with a consistent capacity of 2; and does not materially increase space complexity.





