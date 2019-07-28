const DisjoinSet = require('./DisjointSet');

const ds = new DisjoinSet();
ds.addSet(1);
ds.addSet(2);
ds.addSet(3);
ds.addSet(4);
ds.addSet(5);
ds.addSet(6);
ds.addSet(7);

ds.union(1,2);
ds.union(2,3);
ds.union(4,5);
ds.union(6,7);
ds.union(5,6);
ds.union(3,7);

console.log(ds.findSet(1));
console.log(ds.findSet(2));
console.log(ds.findSet(3));
console.log(ds.findSet(4));
console.log(ds.findSet(5));
console.log(ds.findSet(6));
console.log(ds.findSet(7));