function InsertSort(data){
    var len = data.length;
    for(var i = 1; i < len; i++){
        for(var j = i -1; j >= 0; j--){
            if(data[j] > data[i]){
                var tmp = data[j];
                data[j] = data[i];
                data[i] = tmp;
                break;
            }
        }
    }
    return data;
}