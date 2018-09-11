// Get the modal

var modal = $("#contact_form_modal");
modal.addClass("visuallyhidden");
modal.addClass("hidden");

var hireBtn = $("#hire_box");
var contributeBtn = $("#contribute_box");
var otherBtn = $("#other_box");

var span = $(".close").first();

// When the user clicks on the button, open the modal
hireBtn.click(function() {
    modalOpen();
    $("#hidden_subject").val("hire");
    $("#project_select").hide();
    $("#contact_form_company").show();
    $("#msg_to_sender").text("If you like my work and would like to hire me, get in touch and give me" +
        "a brief idea of what you need and introduce yourself. I will be in touch usually within a few days.");
});

contributeBtn.click(function() {
    modalOpen();
    $("#hidden_subject").val("contribute");
    $("#project_select").show();
    $("#contact_form_company").show();
    $("#msg_to_sender").text("I always welcome help on my projects and if you would like to be part of it introduce yourself!" +
        "Let me know as specific or general as you like in what you want to participate in. I will usually respond within" +
        "a few days.");
});

otherBtn.click(function() {
    modalOpen();
    $("#hidden_subject").val("general");
    $("#project_select").hide();
    $("#contact_form_company").hide();
    $("#msg_to_sender").text("If you just want to get in touch for any reason, drop me a message and I will be in touch!");
});

// When the user clicks on <span> (x), close the modal
span.click(function() {
    modalClose();
});

function modalOpen() {
    modal.removeClass("hidden");
    setTimeout(function () {
        modal.removeClass("visuallyhidden")
    }, 10);
    modal.addClass("displayed")
}

function modalClose() {
    modal.addClass('visuallyhidden');
    modal.one('transitionend', function (e) {
        modal.addClass('hidden');
    });
}

// When the user clicks anywhere outside of the modal, close it

$(window).click(function(e) {
    console.log(e);
    if ($(e.target).is(modal)) {
        modalClose();
    }
});

$('#contact_form').submit( function(event) {
    var formId = this.id,
        form = this;

    confirmFormSubmission();

    event.preventDefault();
    setTimeout( function () {
        form.submit();
    }, 2000);
});

function confirmFormSubmission () {
    $('#msg_to_sender').hide();
    $('#form_submission_confirm').show();
    setTimeout(function () {
        $('#msg_to_sender').show();
        $('#form_submission_confirm').hide();
    }, 2500);
}