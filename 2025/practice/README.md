## Problem List

-----

sources : 

1. stack questions (Aditya Verma): https://www.youtube.com/watch?v=P1bAPZg5uaE&list=PL_z_8CaSLPWdeOezg68SKkeLN4-T_jNHd
2. binary search (Striver) : https://www.youtube.com/playlist?list=PLgUwDviBIf0pMFMWuuvDNMAkoQFi-h0ZF
3. 
-----

1. stack
   1. Implement MinStack (push, pop, min APIs)
      1. Order-n space
      2. Order-1 space : IMP https://www.youtube.com/watch?v=ZvaRHYYI0-4&list=PL_z_8CaSLPWdeOezg68SKkeLN4-T_jNHd&index=11
   2. PATTERN:: nearest greater to right
      1. nearest greater to right
      2. nearest greater to left
      3. nearest smaller to right
      4. nearest smaller to left
   3. stock span : nearest smaller to left
   4. maximum area of histogram
   5. max area in binary matrix
   6. Rain Water Trapping : IMP
2. binary search
   1. PATTERN:: binary search on 1d array 
      1. simple binary search
      2. lower_bound : find element just lower than given
      3. upper_bound : find element just bigger than given
      4. peak element in rotated array : https://www.geeksforgeeks.org/dsa/find-a-peak-in-a-given-array/
      5. rotated array search : https://www.geeksforgeeks.org/dsa/search-an-element-in-a-sorted-and-pivoted-array/#expected-approach-1using-two-binary-searches-olog-n-time-and-o1-space
         1. find pivot
         2. find certain element
         3. how many times array is rotated
      6. bitonic array search 
   2. PATTERN:: binary search on 2d matrix
      1. find row with max 1s (all row are ascending, contain 0 or 1 only)
      2. row-wise and column-wise sorted matrix search
      3. peak element in a matrix
      4. matrix median
   3.  PATTERN::binary search on answer : IMP
       1.  find square root on N
       2.  find nth root of N
       3.  koko eating bananas
       4.  book allocation problem
       5.  split array - largest sum
       6.  median of two sorted arrays : IMP
3. graph : (know the space, time complexity)
   1. defining graph
   2. BFS traversal
      1. directed graph
      2. undirected graph
   3. DFS traversal
      1. directed graph
      2. undirected graph
   4. detect cycle in undirected graph
      1. using dfs traversal
   5. detect cycle in directed graph
      1. using dfs
      2. using bfs
   6. topological sorting (can only be in directed graph)
      1. using bfs
      2. using dfs
   7. minimum spanning tree
      1. kruskal tree
      2. prim algorithm
   8. minimum distance from source
      1. 0-1 BFS
      2. NO NEGTIVE CYCLE => dijkstra
      3. POSSIBLE NEGTIVE CYCLE => Bellman Ford
   9.  minimum distance between any two nodes
      1.  FLOYD-WARSHALL 
4. dynamic programming : 
   1. 1D DP
      1. frong jump, can jump from 1 or 2 position behind.
   2. 2D DP
   3. 3D DP
   4. 
5. sliding window
6. linklist
7. array::PATTERN::interval related
8. array::PATTERN::stock buy-sell related
9.  array::PATTERN::2_sum_and_4_sum
10. binary trees
11. binary search tree
12. heap
13. backtracking
14. implementation
    1.  sort a stack using stack APIs
    2.  min stack, Order(1) space : https://www.youtube.com/watch?v=ZvaRHYYI0-4&list=PL_z_8CaSLPWdeOezg68SKkeLN4-T_jNHd&index=11