gatherTags();

function gatherTags() {
    var tag_list = $('#article_search_tag_list');
    $('.article_tag').each(function (index) {
        tag_list.append("<li>" + this.innerText +"</li>")
    });
}

/* Set the width of the side navigation to 250px */
function openNav() {
    document.getElementById("mySidenav").style.width = "250px";
}

/* Set the width of the side navigation to 0 */
function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
}