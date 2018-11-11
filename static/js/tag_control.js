$(".sidebar_icon").click(function () {
    loadDoc()
});

function loadDoc() {
    console.log("Calling ajax")
    $.ajax({
        url:"/tag",
        type:"GET",
        data: {'tag': 'test'},
        success: function(html){
            console.log("SUCCESS");
        }
    });
    $("#sidebar_icon_list").hide(1000);
    $("#article_search_tag_list").show(500);
    $("#sidebar_close").show(500);
}