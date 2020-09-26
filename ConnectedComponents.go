func quickFind(n int, edges [][]int) int {
    totalComponents := n;
    componentIds :=  make([] int, n);
    // Initialize components
    for i, _ := range componentIds {
       componentIds[i] = i; 
    }
    // process each edge
    for _, edge := range edges {
        u := edge[0];
        v := edge[1];
        // Update component id if not same
        if componentIds[u] != componentIds[v] {
            uComponentId := componentIds[u];
            vComponentId := componentIds[v];
            for i, val := range componentIds {
                if val == vComponentId {
                    componentIds[i] = uComponentId;
                }
            }
            // decrement total
            totalComponents--;
        }
    }
    
    return totalComponents;
}
func quickUnion (n int, edges [][]int) int {
    total := n;
    componentIds := make([]int, n);
    // Initialize componentIds
    for i, _ := range componentIds {
        componentIds[i] = i;
    }
    find := func (current int) int {
        for componentIds[current] != current {
            current = componentIds[current];
        }
        return componentIds[current];
    }
    for _, edge := range edges {
        u := edge[0];
        v := edge[1];
        // find component of u and v
        uComponent := find(u);
        vComponent := find(v);
        
        // union components if required
        if (uComponent != vComponent) {
            componentIds[vComponent] = uComponent; // use uComponent
            total--;
        }
    }
    return total;
}
func weightedQuickUnion (n int, edges [][]int) int {
    total := n;
    parent := make([]int, n);
    size := make([]int, n);
    for i, _:= range parent {
        parent[i] = i;
        size[i] = 1;
    }
    find := func (node int) int {        
        for parent[node] != node {
            node = parent[node];
        }
        return node;
    }
    for _, edge := range edges {
        u := edge[0];
        v := edge[1];
        uRoot := find(u);
        vRoot := find(v);
        if uRoot != vRoot {
            total--;
            // Use root of larger tree
            if size[vRoot] > size [uRoot] {
                // make v as parent of u
                parent[uRoot] = vRoot;
                size[vRoot] += size[uRoot];
            } else {
                // make u as parent of v
                parent[vRoot] = uRoot;
                size[uRoot] += size[vRoot];
            }
        }
    }
    return total;
}
func findAndCompressed (parent [] int, node int) int {  
    // Base case for root node
    if parent[node] == node {
        return node;
    }
    // ask worker node to find out root
    root := findAndCompressed(parent, parent[node]);
    // compress path
    parent[node] = root;
    return root;
}
func pathCompressionUnionAndFind (n int, edges [][]int) int {
    total := n;
    parent := make([]int, n);
    size := make([]int, n);
    for i, _:= range parent {
        parent[i] = i;
        size[i] = 1;
    }

    for _, edge := range edges {
        u := edge[0];
        v := edge[1];
        uRoot := findAndCompressed(parent, u);
        vRoot := findAndCompressed(parent, v);
        if uRoot != vRoot {
            total--;
            // Use root of larger tree
            if size[vRoot] > size [uRoot] {
                // make v as parent of u
                parent[uRoot] = vRoot;
                size[vRoot] += size[uRoot];
            } else {
                // make u as parent of v
                parent[vRoot] = uRoot;
                size[uRoot] += size[vRoot];
            }
        }
    }
    return total;
}
func countComponents(n int, edges [][]int) int {
    return pathCompressionUnionAndFind(n, edges);
}