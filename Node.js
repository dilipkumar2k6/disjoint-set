class Node {
    /**
     * Node representation of tree.
     * @param {Object} options
     * @param {*} data 
     * @param {Node} parent 
     * @param {number} rank 
     */
    constructor({data, parent, rank}){
        this.data = data; 
        this.parent = parent;
        this.rank = rank; 
    }
}

module.exports = Node;