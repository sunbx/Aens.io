
(function ($) {
    'use strict';


    $(function () {

        $.extend({
            StandardPost: function (url, args) {
                var body = $(document.body),
                    form = $("<form method='post'></form>"),
                    input;
                form.attr({"action": url});
                $.each(args, function (key, value) {
                    input = $("<input type='hidden'>");
                    input.attr({"name": key});
                    input.val(value);
                    form.append(input);
                });

                form.appendTo(document.body);
                form.submit();
                document.body.removeChild(form[0]);
            }
        });

     



        var dt = new Date();
        var m = new Array("Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Spt", "Oct", "Nov", "Dec");
        var w = new Array("Monday", "Tuseday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday");
        var d = new Array("st", "nd", "rd", "th");
        let mn = dt.getMonth();
        let wn = dt.getDay();
        let dn = dt.getDate();
        var dns;
        if (((dn) < 1) || ((dn) > 3)) {
            dns = d[3];
        } else {
            dns = d[(dn) - 1];
            if ((dn == 11) || (dn == 12)) {
                dns = d[3];
            }
        }
        // alert(m[mn] + " " + dn + dns + " " + w[wn - 1] + " " + dt.getFullYear());

        $('#date_time').text((m[mn] + " " + dn + dns + " " + w[wn - 1] + " " + dt.getFullYear()));

        var body = $('body');
        var contentWrapper = $('.content-wrapper');
        var scroller = $('.container-scroller');
        var footer = $('.footer');
        var sidebar = $('.sidebar');
        var navbar = $('.navbar').not('.top-navbar');


        //Add active class to nav-link based on url dynamically
        //Active class can be hard coded directly in html file also as required


        function addActiveClass(element) {
            if (current === "") {
                //for root url
                if (element.attr('href').indexOf("index.html") !== -1) {
                    element.parents('.nav-item').last().addClass('active');
                    if (element.parents('.sub-menu').length) {
                        element.closest('.collapse').addClass('show');
                        element.addClass('active');
                    }
                }
            } else {
                //for other url
                if (element.attr('href').indexOf(current) !== -1) {
                    element.parents('.nav-item').last().addClass('active');
                    if (element.parents('.sub-menu').length) {
                        element.closest('.collapse').addClass('show');
                        element.addClass('active');
                    }
                    if (element.parents('.submenu-item').length) {
                        element.addClass('active');
                    }
                }
            }
        }

        var current = location.pathname.split("/").slice(-1)[0].replace(/^\/|\/$/g, '');
        $('.nav li a', sidebar).each(function () {
            var $this = $(this);
            addActiveClass($this);
        })

        //Close other submenu in sidebar on opening any

        sidebar.on('show.bs.collapse', '.collapse', function () {
            sidebar.find('.collapse.show').collapse('hide');
        });


        //Change sidebar and content-wrapper height
        applyStyles();

        function applyStyles() {
            //Applying perfect scrollbar
        }

        $('[data-toggle="minimize"]').on("click", function () {
            if (body.hasClass('sidebar-toggle-display')) {
                body.toggleClass('sidebar-hidden');
            } else {
                body.toggleClass('sidebar-icon-only');
            }
        });

        //checkbox and radios
        $(".form-check label,.form-radio label").append('<i class="input-helper"></i>');


        // fixed navbar on scroll
        $(window).scroll(function () {
            if (window.matchMedia('(min-width: 991px)').matches) {
                if ($(window).scrollTop() >= 197) {
                    $(navbar).addClass('navbar-mini fixed-top');
                    $(body).addClass('navbar-fixed-top');
                } else {
                    $(navbar).removeClass('navbar-mini fixed-top');
                    $(body).removeClass('navbar-fixed-top');
                }
            }
            if (window.matchMedia('(max-width: 991px)').matches) {
                $(navbar).addClass('navbar-mini fixed-top');
                $(body).addClass('navbar-fixed-top');
            }
        });
    });
})(jQuery);