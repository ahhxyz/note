function Bucket(){
    
}

function HashTable(){
    this.table = new Array(137);
    this.buildChain();
}

HashTable.prototype.hash = function(data){
    var total = 0;
    for(var i = 0; i < data.length; i++){
        total += 37 * data.charCodeAt(i);   //乘以一个质数，以更好 解决哈希碰撞的问题
    }
    return parseInt(total % this.table.length);
}

HashTable.prototype.put = function(key, data){
    var hashKey = this.hash(key);
    var v = key || data;
    var pos = this.table[hashKey];
    if(pos.length){
        pos[pos.length + 1] = v;
    }else{
        pos[0] = v;
    }
}

HashTable.prototype.get = function(key){
    return this.table[this.hash(key)];
}

HashTable.prototype.buildChain = function(){
    for(var i = 0; i < this.table.length; i++){
        this.table[i] = new Array();
    }
}

HashTable.prototype.show = function(){
    for(var i = 0; i < this.table.length; i++){
        if(this.table[i].length){
            console.log(i + ' : ', this.table[i])
        }
    }
}

var list = ['David', 'Jennifer', 'Doonie', 'Raymond', 'Cynthia', 'Mike', 'Clayton', 'Danny', 'Jonathan'];
var hTable = new HashTable;
for(var i = 0; i < list.length; i++){
    hTable.put(list[i]);
}
hTable.show();