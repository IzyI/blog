$("#send_me_button").click(function (e) {
    e.preventDefault();
    var sendmetext = $('#sendmetext').val();
    var sendmetextarea = $('#sendmetextarea').val();
    var sendmenumber = $('#sendmenumber').val().replace(/\s+/g, ' ').trim();
    var name_bot = 'siuzanna';

    if (isNaN(sendmenumber)) {
        alert("В поле 10+2 должно быть только число");
        return
    }
    var data = {
        "contact": sendmetext,
        "text": sendmetextarea,
        "number": sendmenumber,
        "number_check": '12',
        "name_bot": name_bot,
        "site_name": "kaper.su"
    };
    $.ajax({
        url: "/botik",
        type: "POST",
        data: JSON.stringify(data),
        dataType: "json",
        contentType: 'application/json; charset=utf-8',
        success: function (response) {
            alert(response.text);
        },
        error: function (xhr, status) {
            console.log(status, xhr);
        },

    });
    console.log("!!!!");

});