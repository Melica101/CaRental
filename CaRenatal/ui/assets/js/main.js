(function ($) {
	"use strict";

    jQuery(document).ready(function($){

 
        //active Car Slider
            $('.pro-car-slider').slick({
              centerMode: true,
                autoplay: true,
              centerPadding: '60px',
              slidesToShow: 2,
              responsive: [
                {
                  breakpoint: 768,
                  settings: {
                    arrows: false,
                    centerMode: true,
                    centerPadding: '40px',
                    slidesToShow: 2
                  }
                },
                {
                  breakpoint: 480,
                  settings: {
                    arrows: false,
                    centerMode: true,
                    centerPadding: '40px',
                    slidesToShow: 1
                  }
                }
              ]
            });

        // Active Header Owl carousel
        $('.slider-wrapper').owlCarousel({
            items: 1,
            loop:true,
            autoplay: true,
            nav:false,
            dots: false,
            touchDrag: false,
            mouseDrag: false,
            smartSpeed: 1000,
            lazyLoad: true,
            responsive:{
                0:{
                    items:1
                },
                600:{
                    items:1
                },
                1000:{
                    items:1
                }
            }
        });
        
		// homepage slider animation support
			$(".header-slider").on("translate.owl.carousel", function(){
				$(".slider-content h1, .slider-content p").removeClass("animated fadeInUp").css({'opacity':'0'});
				$(".slider-content a").removeClass("animated fadeInDown").css({'opacity':'0'});
			});
			
			$(".header-slider").on("translated.owl.carousel", function(){
				$(".slider-content h1, .slider-content p").addClass("animated fadeInUp").css({'opacity':'0'});
				$(".slider-content a").addClass("animated fadeInDown").css({'opacity':'0'});
			});        

      
        
        // Wow JS
        new WOW().init();



        



    });


    jQuery(window).load(function(){
                 // -------------------------------------------------------------
    // Animated scrolling / Scroll Up
    // -------------------------------------------------------------

        //var scroll = new SmoothScroll('a[href*="#"]');
        // Add scrollspy to <body>
        $('body').scrollspy({target: ".sh-mainmenu, .single-slider", offset: 50});

        // Add smooth scrolling on all links inside the navbar
        $(".navbar-nav a, .single-slider a").on('click', function(event) {

          // Make sure this.hash has a value before overriding default behavior
          if (this.hash !== "") {

            // Prevent default anchor click behavior
            event.preventDefault();

            // Store hash
            var hash = this.hash;

            // Using jQuery's animate() method to add smooth page scroll
            // The optional number (800) specifies the number of milliseconds it takes to scroll to the specified area
            $('html, body').animate({
              scrollTop: $(hash).offset().top -20
            }, 800, function(){

            // Add hash (#) to URL when done scrolling (default click behavior)
              window.location.hash = hash;
            });

          } // End if

        });

        
    });


}(jQuery));	