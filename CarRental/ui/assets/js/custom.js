/*=========================
	PRE-LOADER 
=========================*/
$(window).load(function(){ 
	$(".preloader").delay(1000).fadeOut("slow"); 
});


/*=========================
	MAIN NAVIGATION 
=========================*/
$(".navbar-nav li a").click(function (event) {
    var toggle = $(".navbar-toggle").is(":visible");
    if (toggle) {
      $(".navbar-collapse").collapse('hide');
    }
  });


$(window).scroll(function(){
	var sticky = $('.main-navigation'),
      scroll = $(window).scrollTop();
  	if (scroll >= 850) {
		sticky.addClass('fixed');
		sticky.addClass('animated slideInDown');
	}
	else {
		sticky.removeClass('fixed');
	}
});

var lastId,
topMenu = $(".main-navigation .navbar-nav"),
topMenuHeight = topMenu.outerHeight()+15,
menuItems = topMenu.find("a"),
scrollItems = menuItems.map(function(){
	var item = $($(this).attr("href"));
	if (item.length) { return item; }
});
	
$(window).scroll(function(){
	var fromTop = $(this).scrollTop()+topMenuHeight;
	var cur = scrollItems.map(function(){
		if ($(this).offset().top < fromTop)
		return this;
	});
	
	cur = cur[cur.length-1];
	var id = cur && cur.length ? cur[0].id : "";
	if (lastId !== id) {
		lastId = id;
		menuItems
			.parent().removeClass("active")
			.end().filter("[href=#"+id+"]").parent().addClass("active");
	}                   
});


/*=========================
	SMOOTH SCROLL
=========================*/
$('.navbar-nav a, .hero-buttons a, .custom-button, .car-details li a').click(function(){
  $.scrollTo( this.hash, 1200, { easing:'swing' });
  return false;
});


/*=========================
	WIDE SLIDER
=========================*/
$(document).ready(function() {
  $("#wide_slider").owlCarousel({
        navigation: false,
		pagination: false,
		singleItem: true,
        slideSpeed: 400,
        autoPlay: 5000,
		rtl: true
		//transitionStyle : 'backSlide',
    });
});


/*=========================
	SMALL SLIDER
=========================*/
$(document).ready(function() {
	var owl = $("#small-slider");
	
   	 owl.owlCarousel({
		autoPlay: false,
		items : 3,
		pagination: true,
		itemsDesktop : [1199,3],
		itemsDesktopSmall : [979,1],
		itemsTablet : [768,1]
  	});
	
    $(".right-arrow").click(function() {
        owl.trigger('owl.next');
    });
	
    $(".left-arrow").click(function() {
        owl.trigger('owl.prev');
    });
});

$(document).ready(function(){
	$(".small-slider .item span").addClass('overlay');
	$(".small-slider .item span").append('<img src="assets/images/icons/play.png" alt="">');
});


$(".small-slider .item iframe").hide();

$(".small-slider .item").click(function() {
	$(this).find("span").hide();
	$(this).find("img").hide();
	if($(this).find("iframe").show()){
		$(this).find("iframe").attr("src", $(this).find("iframe").attr("src").replace("autoplay=0", "autoplay=1"));
	}
});
	

/*=========================
	NUMBERS COUNT
=========================*/
(function($) {
		"use strict";
		function count($this){
		var current = parseInt($this.html(), 10);
		current = current + 1; 
		$this.html(++current);
			if(current > $this.data('count')){
				$this.html($this.data('count'));
			} else {    
				setTimeout(function(){count($this)}, 50);
			}
		}        	
		$(".stat-count").each(function() {
		  $(this).data('count', parseInt($(this).html(), 10));
		  $(this).html('0');
		  count($(this));
		});
})(jQuery);


/*=========================
	CUSTOM SLIDER
=========================*/
$("#tab1 .car-info-container .infos").hide();
$("#tab1 .car-info-container .infos:first").show();
$("#tab1 .custom-nav-container ul li").click(function() {
	  $("#tab1 .car-info-container .infos").hide();
	  var activeTab = $(this).attr("rel"); 
	  $("#"+activeTab).fadeIn();		
	  $("#tab1 .custom-nav-container ul li").removeClass("active-cst");
	  $(this).addClass("active-cst");
});

$("#tab2 .car-info-container .infos").hide();
$("#tab2 .car-info-container .infos:first").show();
$("#tab2 .custom-nav-container ul li").click(function() {
	  $("#tab2 .car-info-container .infos").hide();
	  var activeTab = $(this).attr("rel"); 
	  $("#"+activeTab).fadeIn();		
	  $("#tab2 .custom-nav-container ul li").removeClass("active-cst");
	  $(this).addClass("active-cst");
});


$("#tab3 .car-info-container .infos").hide();
$("#tab3 .car-info-container .infos:first").show();
$("#tab3 .custom-nav-container ul li").click(function() {
	  $("#tab3 .car-info-container .infos").hide();
	  var activeTab = $(this).attr("rel"); 
	  $("#"+activeTab).fadeIn();		
	  $("#tab3 .custom-nav-container ul li").removeClass("active-cst");
	  $(this).addClass("active-cst");
});


$("#tab4 .car-info-container .infos").hide();
$("#tab4 .car-info-container .infos:first").show();
$("#tab4 .custom-nav-container ul li").click(function() {
	  $("#tab4 .car-info-container .infos").hide();
	  var activeTab = $(this).attr("rel"); 
	  $("#"+activeTab).fadeIn();		
	  $("#tab4 .custom-nav-container ul li").removeClass("active-cst");
	  $(this).addClass("active-cst");
});


/*=========================
	TOOLTIP
=========================*/
$(function () {
    $('[data-toggle="tooltip"]').tooltip();
     $('[data-toggle="tooltip"]').hover(function (){
          $(this).next().addClass("animated bounce");               
    });
});


/*=========================
	ANIMATIONS
=========================*/
wow = new WOW(
  {
    mobile: false
  });
wow.init();


/*=========================
	CAR DETAILS
=========================*/
$(document).ready(function(){
	$(".car-details li a").addClass('wow animated bounceIn');
	$(".car-details li a").attr('href', '#section8');
	$(".car-info-container .infos").addClass('animated flipInX');
});



/*=============================
	IE 10 SUPPORT FOR BOOTSTRAP
=============================*/
if (navigator.userAgent.match(/IEMobile\/10\.0/)) {
  var msViewportStyle = document.createElement('style')
  msViewportStyle.appendChild(
    document.createTextNode(
      '@-ms-viewport{width:auto!important}'
    )
  )
  document.querySelector('head').appendChild(msViewportStyle)
}


/*=============================================
	ANDROID STOCK BROWSER SUPPORT FOR BOOTSTRAP
=============================================*/
$(function () {
  var nua = navigator.userAgent
  var isAndroid = (nua.indexOf('Mozilla/5.0') > -1 && nua.indexOf('Android ') > -1 && nua.indexOf('AppleWebKit') > -1 && nua.indexOf('Chrome') === -1)
  if (isAndroid) {
    $('select.form-control').removeClass('form-control').css('width', '100%')
  }
})

