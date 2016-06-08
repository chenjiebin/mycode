/**=========================================================
 * Module: config.js
 * App routes and resources configuration
 =========================================================*/

angular.module('angle').config(['$stateProvider', '$locationProvider', '$urlRouterProvider', 'RouteHelpersProvider',
function ($stateProvider, $locationProvider, $urlRouterProvider, helper) {
  'use strict';

  // Set the following to true to enable the HTML5 Mode
  // You may have to set <base> tag in index and a routing configuration in your server
  $locationProvider.html5Mode(true);

  // default route
  $urlRouterProvider.otherwise('/singleview');

  // 
  // Application Routes
  // -----------------------------------   
  $stateProvider
    .state('app', {
        // url: '/app',
        abstract: true,
        templateUrl: helper.basepath('app.ng.html'),
        controller: 'AppController',
        resolve: helper.resolveFor('modernizr', 'icons')
    })
    .state('app.singleview', {
        url: '/singleview',
        title: 'Single View',
        templateUrl: helper.basepath('singleview.ng.html')
    })
    .state('app.submenu', {
        url: '/submenu',
        title: 'Submenu',
        templateUrl: helper.basepath('submenu.ng.html')
    })
    // 
    // CUSTOM RESOLVES
    //   Add your own resolves properties
    //   following this object extend
    //   method
    // ----------------------------------- 
    // .state('app.someroute', {
    //   url: '/some_url',
    //   templateUrl: 'path_to_template.html',
    //   controller: 'someController',
    //   resolve: angular.extend(
    //     helper.resolveFor(), {
    //     // YOUR RESOLVES GO HERE
    //     }
    //   )
    // })
    ;


}]).config(['$ocLazyLoadProvider', 'APP_REQUIRES', function ($ocLazyLoadProvider, APP_REQUIRES) {
    'use strict';

    // Lazy Load modules configuration
    $ocLazyLoadProvider.config({
      debug: false,
      events: true,
      modules: APP_REQUIRES.modules
    });

}]).config(['$controllerProvider', '$compileProvider', '$filterProvider', '$provide',
    function ( $controllerProvider, $compileProvider, $filterProvider, $provide) {
      'use strict';
      // registering components after bootstrap
      angular.module('angle').controller = $controllerProvider.register;
      angular.module('angle').directive  = $compileProvider.directive;
      angular.module('angle').filter     = $filterProvider.register;
      angular.module('angle').factory    = $provide.factory;
      angular.module('angle').service    = $provide.service;
      angular.module('angle').constant   = $provide.constant;
      angular.module('angle').value      = $provide.value;

}]).config(['$translateProvider', function ($translateProvider) {

    $translateProvider.useStaticFilesLoader({
        prefix : 'translate/',
        suffix : '.json'
    });
    $translateProvider.preferredLanguage('en');
    $translateProvider.useLocalStorage();
    $translateProvider.usePostCompiling(true);

}]).config(['cfpLoadingBarProvider', function(cfpLoadingBarProvider) {
    cfpLoadingBarProvider.includeBar = true;
    cfpLoadingBarProvider.includeSpinner = false;
    cfpLoadingBarProvider.latencyThreshold = 500;
    cfpLoadingBarProvider.parentSelector = '.wrapper > section';
}]).config(['$tooltipProvider', function ($tooltipProvider) {

    $tooltipProvider.options({appendToBody: true});

}])
;
