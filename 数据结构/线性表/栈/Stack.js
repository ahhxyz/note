/**
 * Stack类
 */
function Stack(){
    this.data = []; //该数组不是栈，它只是用来存储栈的数据，从数组的第一个元素到栈顶的元素集合才是栈
    this.top = -1;   //栈顶，能进行操作的一端，即数组的最后一个元素
}

Stack.prototype.isEmpty = function(){
    return this.top === -1;
}

/**
 * 入栈：向栈顶添加元素
 * @param fixed item 值
 * @returns void
 */
Stack.prototype.push = function(item){
    this.data[++this.top] = item;
}
/**
 * 出栈：弹出栈顶元素
 * @returns void
 */
Stack.prototype.pop = function(){
    return this.data[this.top--];
}

Stack.prototype.peak = function(){
    return this.data[this.top];
}

Stack.prototype.shift = function(){
    
}

Stack.prototype.unshift = function(){
    
}

Stack.prototype.clear = function(){
    this.top = 0;
}

Stack.prototype.free = function(){
    this.data = null;
}



