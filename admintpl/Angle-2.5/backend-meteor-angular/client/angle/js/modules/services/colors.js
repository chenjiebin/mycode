/**=========================================================
 * Module: colors.js
 * Services to retrieve global colors
 =========================================================*/
 
angular.module('angle').factory('colors', ['APP_COLORS', function(colors) {
  
  return {
    byName: function(name) {
      return (colors[name] || '#fff');
    }
  };

}]);
