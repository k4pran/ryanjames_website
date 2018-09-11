test = [
    "<pre><code class=\"language-treeview\">\n" +
    "root_folder/\n" +
    "|-- a first folder/\n" +
    "|   |-- holidays.mov\n" +
    "|   |-- javascript-file.js\n" +
    "|   `-- some_picture.jpg\n" +
    "|-- documents/\n" +
    "|-- .gitignore\n" +
    "|-- .htaccess\n" +
    "|-- .npmignore\n" +
    "|-- archive 1.zip\n" +
    "|-- archive 2.tar.gz\n" +
    "`-- logo.svg\n" +
    "</code></pre>"
]

function f () {
    console.log(document.body)
    document.body.innerText = test
    console.log("yest")
}