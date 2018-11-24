$(".contact-bottom-menu").mouseover(function () {
    $("#contact-page-container").css("grid-template-rows", "30% auto auto 15%");
    $(".contact-bottom-menu-item").each(function () {
        $(this).removeClass("hidden");
    })
});

$(".contact-bottom-menu").mouseout(function () {
    $("#contact-page-container").css("grid-template-rows", "30% auto auto 5%");
    $(".contact-bottom-menu-item").each(function () {
        $(this).addClass("hidden")
    });
});
