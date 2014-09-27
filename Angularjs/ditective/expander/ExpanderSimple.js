var expanderModule=angular.module('expanderModule', [])

expanderModule.controller('SomeController',function($scope,$rootScope) {
   
    $rootScope.style="a";
    $scope.title = '点击展开';
    $scope.text = '这里是内部的内容。';
});
expanderModule.directive('expander', function($rootScope,$location) {
    
    return {
        restrict : 'EA',
        replace : true,
        transclude : true,
        scope : {
            title : '=expanderTitle'
        },
        templateUrl : 'tpl.html',
        link : function(scope, element, attrs) {
            scope.showMe = false;
            scope.toggle = function toggle() {
                scope.showMe = !scope.showMe;
            }
        }
    }
});
