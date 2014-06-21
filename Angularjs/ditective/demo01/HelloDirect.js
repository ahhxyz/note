var appModule = angular.module('app', []);
appModule.directive('hello', function() {
    return {
        restrict: 'E',
        template: '<div ng-repeat="i in [0, 1, 2, 3, 4, 5, 6, 7]" >{{i}}:Hi there</div>',
        replace: true
    }; 
});