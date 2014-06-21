var appModule = angular.module('app', []);
appModule.directive('hello', function() {
    return {
        restrict: 'E',
        template: '<span>Hi there</span>',
        replace: true
    };
});
appModule.controller('MyController',function($scope) {
    $scope.things = [1,2,3,4,5,6];
});