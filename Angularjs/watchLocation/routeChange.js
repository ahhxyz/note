var appModule = angular.module('app', []);
function LocationController($scope, $location) {
$scope.$on('$routeChangeStart', function(next, current) { 
alert("a")
 });
}