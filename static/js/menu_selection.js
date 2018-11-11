var isBurgerMenu = false;
var currentSelected = null;
isBurgerMenu = $(window).width() <= 644;
get_current_page_name();

window.addEventListener('resize', function(event){
    isBurgerMenu = $(window).width() <= 644;
    get_current_page_name()
});

function get_current_page_name() {

    var url = window.location.href;
    $(".main_menu_item a").each(function () {
        if (url == (this.href)) {
            currentSelected = $(this).closest("li a");
            $(this).closest("li a").css("color", "honeydew");
        }
    });
}

document.addEventListener("DOMContentLoaded", function (event) {
    var _selector = document.querySelector('input[name=burger_menu_checkbox]');
    _selector.addEventListener('change', function (event) {
        if (_selector.checked) {
            $('#menu').removeClass('hidden');
        } else {
            $('#menu').addClass('hidden');
        }
    });
});