# Summary
- Calculate the number of disconnected component in given graph
- Edges are coming in stream i.e. we don't have predefined list of edges
- can't apply dfs/bfs optimally
# Basic Union and Find i.e. Quick Find
![](assets/components-problem.png)
![](assets/quick-find-step1.png)
![](assets/quick-find-complexity.png)
![](assets/quick-find-complexity-2.png)
# Improved Union and Find i.e. Quick Union
- Religion of king will be followed by rest in that component
![](assets/quick-union-approach.png)
![](assets/quick-union-step1.png)
![](assets/quick-union-complexity.png)
# Weighted Quick Union and Find
- Goal is to minimize the height of the tree
![](assets/union-find-approach.png)
![](assets/union-find-approach-2.png)
![](assets/union-find-approach-3.png)
# Optimzed Union and Find using path compression
![](assets/compressed-path.png)
