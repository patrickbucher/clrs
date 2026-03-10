# Chapter 2

## 2.1-1

given: $A=[31,41,59,26,41,58], n=6$, 1-based indices

1. $31,[41],59,26,41,58$
    - $41$ is already at the right place
2. $31,41,[59],26,41,58$
    - $59$ is already at the right place
3. $31,41,59,[26],41,58$
    - $26$ is moved to index 1
    - $31,41,59$ are shifted to the right
4. $26,31,41,59,[41],58$
    - $41$ (at index 5) is moved to index 4
    - $59$ is shifted to the right
5. $26,31,41,41,59,[58]$
    - $58$ is moved to index 5
    - $59$ is shifted to the right
6. $26,31,41,41,58,59$
    - the array is now sorted

## 2.1-2

Invariant: The variable `sum` holds the sum of all values of `A` that already have been processed by the loop, i.e. the sum of `A[1;i-1]`.

Initialization: Before the first iteration, the subarray `A[1;0]` is an empty array with the sum of 0.

Maintenance: The `sum` variable is increased by `A[i]` in every iteration, holding the sum of `A[1;i]`.

Termination: Once $i=n$, the loop terminates, after `A[n]` has been added to `sum`.

## 2.1-3

The condition `A[j] > key` on line 5 needs to replaced with `A[j] < key`.

## 2.1-4

given: $A=<a_1,a_2,\ldots,a_n>, n, x$,

    Linear-Search(A, n, x)
        for i = 1 to n
            if x = A[i]
                return i
        return NIL

Invariant: The subarray `A[1;i]` does not contain `x`.

Initialization: The subarray `A[1;0]` is empty.

Maintenance: The loop continues if `A[i]` is not equal to `x`, or terminates otherwise.

Termination: If $x \notin A[1;i]$ for $i=n$, the loop terminates without having returned an index. For $x \in A[1;i]$ for $i=n$, the function returned the index $j$ for $A[j] = x$.

## 2.1-5

    Binary-Addition(A, B, n)
        carry = 0
        for i = 0 to n
            bit = A[i] + B[i] + carry
            C[i] = bit mod 2
            if bit > 1
                carry = 1
            else
                carry = 0
        C[i+1] = carry
        return C

