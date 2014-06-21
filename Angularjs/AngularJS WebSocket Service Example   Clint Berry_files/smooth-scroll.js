// JavaScript Document
jQuery(document).ready(function() {

  var device = navigator.userAgent.toLowerCase();
  var ios = device.match(/(iphone|ipod|ipad)/);

 //function that returns element's y position
    
    jQuery("a[href*=#]").on('click', function() {
      var scrollTarget = jQuery(this.hash).offset().top;
        if(parseInt(scrollTarget) !== parseInt(jQuery(window).scrollTop())) {
          var intro = jQuery("#intro"),
            nav2 = jQuery(".nav2");
        if (ios) nav2.hide();
          jQuery('html,body').animate({scrollTop:scrollTarget}, 1200, "swing", function(evt) {
          if (ios) {
            nav2.css({position:'absolute', top:scrollTarget + 34});
            var nav2clone = jQuery(".nav2")
            nav2clone.show();
          }
      });
    }
    });

    if (ios) {
        jQuery(document).bind('touchmove',function(){
          var intro = jQuery("#intro"),
            nav2 = jQuery(".nav2");
        if(intro.height() <= nav2.position().top)
        {
            nav2.css({position:'fixed', top:'34px'});
            nav2.show();
          }
          else
            nav2.hide();
      });   
    }
});