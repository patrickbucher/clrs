# Chapter 1

## 1.1-1

Folded towels are sorted by size before piling them up, so that the biggest ones end up on the bottom, and the smallest ones on the top. This results in a pyramid-shaped heap that is very stable and does not fall over upon transportation to the closet.

A walk from the school building's entrance to a specific class room allows for different routes. Some of those are considerably longer than others. In order to minimize traveling time, the shortest route shall be found. However, unevenly distributed traffic in the hallways is also a relevant factor; not for distance, but for traveling time.

## 1.1-2

Other than speed, memory consumption is also a relevant measure of efficiency. Some algorithms can be sped up using a concurrent implementation that allows for parallel execution when given the right environment, which influences the efficiency as perceived by the end user. Some algorithms scale well for large inputs, but are therefore more complicated, which slows down the computation for very small inputs. So the efficiency of an algorithm needs to be considered for different kinds of inputs, e.g. how a sorting algorithm deals with very small or pre-ordered inputs.

## 1.1-3

A binary search tree, which places smaller items on the left and larger items on the right of every node, can speed up looking for individual elements considerabely. However, this is only the case if the items are not inserted in a pre-ordered sequence. Otherwise, the binary search on such a degenerated tree becomes as slow as a linear search. Re-balancing such a tree is also quite involved and error-prone to implement.

## 1.1-4

Both problems, shortest-path and travelig-salesperson, are about minimizing the distance traveled. However, shortest-path only consideres the path between two given nodes, whereas traveling-salesperson requires visiting all nodes in the most efficient way possible.

## 1.1-5

Ranking the athletes of a competition by a relevant criterium (e.g. time elapsed for running one hundred meters) requires the best solution, i.e. the one correct sorting order.

Scheduling the time table for a school is a problem where the quality of a solution requires compromises, and the outcomes are judged subjectively: The solution that creates the least dissatifaction in sum might be the one that creates considerable dissatifaction for one particular teacher or group of students.

## 1.1-6

Scheduling the time table for a school can take up both shapes: If the scheduling is done while it is still possible for students to register and to drop out, the input is only partially available. If the body of students remains constant, then the scheduling algorithm has the entire input ready.

Finding the optimal route while traveling can also take up both forms. If traffic needs to be considered, the input arrives over time. If traffic is negligible or ignored, only the roads have to be considered, which will not change for the duration of a single trip.

## 1.2-1

The sort function in a spreadsheet application requires a sorting algorithm. Since there are usually only a few rows, maybe thousands, raw computation performance is secondary. However, the sorting algorithm must allow for sorting by multiple criteria or be stable for multiple subsequent sorting operations by different criteria.

## 1.2-2

Insertion Sort: $8n^2$

Merge Sort: $64n \log_{10} n$

Insertion Sort beats Merge Sort:

$$ 8n^2 < 64n \log_{10} n $$

Divide by $8n$:

$$ n < 8 \log_{10} n $$

According to Wolfram Alpha, Insertion Sort beats Merge sort for $2, 3, 4, 5, 6$.

## 1.2-3

$$ 100n^2 < 2^n $$
