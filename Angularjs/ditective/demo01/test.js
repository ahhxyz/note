var appModule = angular.module('app', []);
appModule.controller('TestCtrl', function($scope, $compile){
var link = $compile('<p>{{ text }}</p>');
  //var node = link($scope);
  console.log(link);  });