var myapp = angular.module('myapp', ['ui.state']);
myapp.config(function($stateProvider, $routeProvider){
$stateProvider
    .state('route1', {
        url: "/route1",
        views: {
            "viewA": {
                template: 'route1.viewA.html',
            },
            "viewB": {
                templateUrl: "route1.viewB.html"
            }
        }
    })
    .state('route2', {
        url: "/route2",
        views: {
            "viewA": {
                templateUrl: "route2.viewA.html"
            },
            "viewB": {
                templateUrl: "route2.viewB.html"
            }
        }
    })
})