$(".contact-bottom-menu-item").each(function () {
    $(this).addClass("hidden")
});

$(".bottom-expansion").mouseover(function () {
    $(".contact-bottom-menu-item").each(function () {
        $(this).removeClass("hidden");
    });

    $(".bottom-expansion").css("flex", "5")
});

$(".bottom-expansion").mouseout(function () {


    $(".bottom-expansion").css("flex", "1")
});
