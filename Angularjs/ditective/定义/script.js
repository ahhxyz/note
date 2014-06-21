// Code goes here

var app = angular.module("directiveExtensionApp", []);


app.controller("foo", function($scope){
  $scope.handleTableClick = function(id){
    alert("This is a function in the page's controller.\n"+
          "I received id: " +id);
  }
});

/* This is a directive which provides rows in a table. */
app.directive("myTableDirective", function() {
  
    return {
      restrict: 'A',
      template: '<tr ng-repeat="thing in things"><td>{{thing.name}}</td></tr>',
      
      controller: function($scope) {
        
        // Some things to display in each row
        $scope.things = [
          { id: "0", name: "Fred" },
          { id: "1", name: "Donna"},
          { id: "2", name: "Goddart"}
        ];
        
      },
      
      compile: function(e){
        // do stuff maybe...
      }
      
    }
  });


  
  /* A simple click binding directive, not used in this example. */
  app.directive("myClick", function() {
  
    return {
      restrict: 'A',
      link: function(scope, element, attrs) {
        element.bind('click', function() { alert('You clicked: ' + attrs.myClick) });
        element.css('cursor', 'pointer')
      }
    }
  });
  
  
  

  
