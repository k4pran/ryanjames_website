$(".article_wrapper").each(function () {
    var wrapper = $(this);
    $(this).find(".collapse_article").on("click", function () {
        wrapper.find(".collapsible").addClass("hidden");
        wrapper.find(".code-toolbar").addClass("hidden");
        wrapper.find("br").addClass("hidden");
        wrapper.find(".collapse_article").addClass("hidden");
        wrapper.find(".expand_article").removeClass("hidden");
        wrapper.find(".article_summary").removeClass("hidden");
    });

    $(this).find(".expand_article").on("click", function () {
        wrapper.find(".collapsible").removeClass("hidden");
        wrapper.find(".code-toolbar").removeClass("hidden");
        wrapper.find("br").removeClass("hidden");
        wrapper.find(".collapse_article").removeClass("hidden");
        wrapper.find(".expand_article").addClass("hidden");
        wrapper.find(".article_summary").addClass("hidden");
    });
});

$(".project_wrapper").each(function () {
    var wrapper = $(this);
    $(this).find(".collapse_project").on("click", function () {
        wrapper.find(".collapsible").addClass("hidden");
        wrapper.find(".code-toolbar").addClass("hidden");
        wrapper.find("br").addClass("hidden");
        wrapper.find(".collapse_project").addClass("hidden");
        wrapper.find(".expand_project").removeClass("hidden");
        wrapper.find("iframe").addClass("hidden");
    });

    $(this).find(".expand_project").on("click", function () {
        wrapper.find(".collapsible").removeClass("hidden");
        wrapper.find(".code-toolbar").removeClass("hidden");
        wrapper.find("br").removeClass("hidden");
        wrapper.find(".collapse_project").removeClass("hidden");
        wrapper.find(".expand_project").addClass("hidden");
        wrapper.find("iframe").removeClass("hidden");
    });
});