var A={} ///为应该程序创建命名空间
A.C={};//整个应用程序的控制器的命名空间
A.M={};//整个应用程序的模型的命名空间

A.C.Order = function($scope){
    $scope.data = [
      {name: 'B', age: 4},  
      {name: 'A', age: 1},  
      {name: 'D', age: 3},  
      {name: 'C', age: 2},    
    ];
    var div=angular.element("<div>甘延寿</div>");

    $('table').append(div);

  }