const Node = require('./Node');

class DisjointSet {
    constructor() {
        this.map = new Map();
    }
    /**
     * Add data as set if doesn't exist and update map.
     * @param {*} data 
     */
    addSet(data){
        const node = new Node({
            data,
            rank: 0,
            parent: null
        });
        node.parent = node; // keep self as parent as default
        // update map
        if(!this.map.has(data)) {
            this.map.set(data, node);
        }
        return this.map.get(data);
    }
    findSet(data) {
        const node = this.map.get(data); 
        // if not found then simply return null
        if(!node) {
            return null;
        } 
        return this._findSet(node);       
    }
    _findSet(node){
        // if node parent is same as parent then return 
        const { parent } = node;
        if(parent === node) {
            return parent;
        } 
        // find parent recursively and also update parent for path compression for quick search next time
        node.parent = this._findSet(parent);
        return node.parent; // Return optimized path compression based parent
    }
    union(data1, data2) {
        // Get parent 
        const parent1 = this.findSet(data1);
        const parent2 = this.findSet(data2);

        // If both nodes are part of same set then do nothing
        if(parent1 === parent2) {
            return;
        }
        // Whosover has higher rank become parent 
        if(parent1.rank === parent2.rank) {
            // increment rank
            parent1.rank = parent1.rank + 1;

            // choose parent1 as new parent
            parent2.parent = parent1;            
        }
        else if(parent1.rank > parent2.rank) {
            parent2.parent = parent1;
        } else {
            parent1.parent = parent2; 
        }
    }

}

module.exports = DisjointSet;