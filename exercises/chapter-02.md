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
