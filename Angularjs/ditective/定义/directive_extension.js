// Obtain the module that the directive we want to extend is in.
var module = angular.module("directiveExtensionApp");

  /* Extending the myTableDirective directive */
  module.directive("myTableDirective", function() {
  
    return {
      restrict: 'A',
      
      compile: function(e){
        
        // Find the tr element
        e.find('tr')
        
          // Add the ng-click element
          .attr('ng-click','handleTableClick({{thing.id}})')
          
          // Make the cursor a pointer when hovering.
          .css('cursor', 'pointer');
        
      
      }
      
    }
  });