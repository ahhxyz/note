/*
 * 链表：链式存储的线性表
 * 实现链表需要定义2个数据结构，分别类似于PHP内核的Bucket和HashTable
 */
function Node(data){
    this.data = data;
    this.next = null;
    this.previous = null;
}
function ChainList(head){
    this.head = new Node(head);
}
ChainList.prototype.find = function(item){
    var currentNode = this.head;
    while(currentNode.data != item){
        console.log(currentNode)
        currentNode = currentNode.next;
    }
    return currentNode; //引用类型的数据，作为函数返回值返回的依然是一个引用，即地址的副本，同传参时的道理是一样的
}

ChainList.prototype.insert = function(data, item){
    var newNode = new Node(data);
    var targetNode = this.find(item);
    newNode.next = targetNode.next;
    newNode.previous = targetNode;
    targetNode.next = newNode;
}

ChainList.prototype.display = function(){
    var currentNode = this.head;
    console.log(currentNode);
    while(currentNode.next !== null){
        console.log(currentNode.next);
        currentNode = currentNode.next;
    }
}

ChainList.prototype.findPrevious = function(item){
    var node = this.head;
    while(node.next && node.next.data == item){
        return node;
    }
}

ChainList.prototype.remove = function(item){
    var node = this.find(item);
    node.previous.next = node.next;
    node.next.previous = node.previous;
    node.next = null;
    node.previous = null;
}

var chainList = new ChainList('张三');
//console.log(chainList); // { head: Node { data: '张三', next: null } }
chainList.insert('李四', '张三');
chainList.insert('王麻子', '李四');
//var node = chainList.find('李四');
//node.data = 'lisi';
//console.log(chainList.findPrevious('lisi'))
chainList.display();
//chainList.remove('李四')
//chainList.display();
