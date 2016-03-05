/**
 * 冒泡排序
 * 
    
 * 
 * 
 */

function BubbleSort(data){
    var len = data.length
    for(var i = 0; i < len; i++){   //外循环每次只能完成一个元素的归位
        for(var j = len - 1; j > i; j--){
            if(data[j - 1] > data[j]){
                var tmp = data[j - 1];
                data[j - 1] = data[j];
                data[j] = tmp;
                console.log(data);  //每次的排序结果 
            }
        }
    }
    return data;
}