

function generateArticle(title, date, tags, image, content) {
    var sectionElement = document.createElement("section", {"class": "article"});
    sectionElement.setAttribute("class", "article_single");
    var titleElement   = document.createElement("h3", {"class": "article_title"});
    titleElement.setAttribute("class", "article_title");
    var dateElement    = document.createElement("p", {"class": "article_date"});
    dateElement.setAttribute("class", "article_date");
    var imageElement   = document.createElement("img", {"class": "article_image"});
    imageElement.setAttribute("class", "article_image_main");


    titleElement.innerText = title;
    dateElement.innerText  = date;
    imageElement.innerText = image;

    sectionElement.appendChild(titleElement);
    sectionElement.appendChild(dateElement);
    sectionElement.appendChild(setTags(tags));
    sectionElement.appendChild(imageElement);
    sectionElement.appendChild(setContent(content));

    sectionElement.appendChild(document.createElement("hr"));

    articles_block = document.getElementById("articles_list");
    articles_block.appendChild(sectionElement)
}

function setTags(tags) {
    var tagList = document.createElement("ul");
    tagList.setAttribute("class", "article_tags_list");
    tags.sort();

    tags.forEach(function (tagText) {
        var tag = document.createElement("li");
        tag.setAttribute("class", "article_tag");
        tag.innerText = tagText;
        tagList.appendChild(tag);
    });
    return tagList;
}

function setContent(content) {
    var contentParts = [];
    content.forEach(function(contentPart) {
        contentParts.push(renderContentType(contentPart));
    });

    contentSectionElement = document.createElement("div");
    contentSectionElement.setAttribute("class", "article_content");
    contentParts.forEach(function (element) {
        contentSectionElement.appendChild(element);
        contentSectionElement.appendChild(document.createElement("br"))
    });


    return contentSectionElement;
}

function renderContentType(contentPart) {
    switch (contentPart["type"][0]) {
        case "text":
            var contentBlock = document.createElement("div");
            contentBlock.setAttribute("class", "article_content_block");
            for (i = 0; i < contentPart["lines"].length; i++) {
                var element = document.createElement("p");
                element.setAttribute("class", "article_text_block_part");
                element.innerText = contentPart["lines"][i];
                contentBlock.appendChild(element);
            }
            return contentBlock;
        case "image":
            var element = document.createElement("img");
            element.setAttribute("src", contentPart["lines"]);
            element.setAttribute("alt", "article image");
            element.setAttribute("class", "article_content_block");
            element.setAttribute("class", "article_image_block");
            return element;
        case "code":
            var wrapper = document.createElement("pre");
            var lines = document.createElement("code");
            lines.setAttribute("class", "language-" + contentPart["language"][0]);
            lines.innerText = contentPart["lines"].join("\n");
            wrapper.appendChild(lines);
            return wrapper;

        default:
            // console.log("Invalid content part block - Must be 'text', 'image' or 'code. Instead it is '" +
            //     contentPart["type"] + "'");
            return document.createElement("empty")
    }
}