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

Invariant: The variable `sum` holds the sum of all values of `A` that already have been processed by the loop, i.e. the sum of $A[1;i-1]$.

Initialization: Before the first iteration, the subarray $A[1;0]$ is an empty array with the sum of 0.

Maintenance: The `sum` variable is increased by $A[i]$ in every iteration, holding the sum of $A[1;i]$.

Termination: Once $i=n$, the loop terminates, after $A[n]$ has been added to `sum`.

## 2.1-3

The condition $A[j] > key$ on line 5 needs to replaced with $A[j] < key$.

## 2.1-4

given: $A=<a_1,a_2,\ldots,a_n>, n, x$,

    Linear-Search(A, n, x)
        for i = 1 to n
            if x = A[i]
                return i
        return NIL

Invariant: The subarray $A[1;i]$ does not contain $x$.

Initialization: The subarray $A[1;0]$ is empty.

Maintenance: The loop continues if $A[i]$ is not equal to $x$, or terminates otherwise.

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

## 2.2-1

$f(n) = n^3/1000 + 100n^2 - 100n \Rightarrow \Theta(n^3)$

## 2.2-2

    Selection-Sort(A, n)
        for i = 1 to n-1
            k = i+1
            smallest = A[k]
            for j = k+1 to n
                if A[j] < smallest
                    smallest = A[j]
                    k = j
            if smallest < A[i]
                temp = A[i]
                A[i] = A[k]
                A[k] = temp

Invariant: All items of the subarray $A[1;i]$ are smaller or equal than the items in the subarray $A[i+1;n]$. $\forall x \in A[1;i] \wedge y \in A[i+1;n]: x \leq y$

Only the first $n-1$ items need to be considered, because $A[n]$ must be the largest element after $i=n-1$ elements have been processed. The element ending up at $A[n]$ was considered $n-1$ times to _not_ be the smallest element of $A[i;n]$, so it must be the largest one.

Both best and worst case scenarios need to consider the entire remainder, i.e. the subarray $A[i+1;n]$, for every iteration of $1..n-1$ in order to find the smallest element: $\Theta(n^2)$

## 2.2-3

In the best case, linear search finds the element at the first position: $\Theta(1)$

In the worst case, linear search finds no matching element after processing the entire array: $\Theta(n)$

In the average case, linear search finds the element in the middle of the array after $n/2$ steps: $\Theta(n/2) \Rightarrow \Theta(n)$

## 2.2-4

An array can be checked in $n-1$ steps if it is already in order. If so, the procedure can terminate early, i.e. without doing any sorting at all.

