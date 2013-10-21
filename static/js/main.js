(function(){
    var fragment = $.deparam.fragment();
    // For debugging!
    console.log("fragment", fragment);
    $.ajax({
        dataType: "json",
        beforeSend: function (xhr) {
            xhr.setRequestHeader("Authorization", "Bearer " + fragment["token"]);
            xhr.setRequestHeader("Accept",        "application/json");
        },
        type: "GET",
        url: "https://localhost:8000/v1/user/12345",
        data: {},
        success: function(response) {
            // For debugging!
            console.log("success", response);
            $("#root").text(response.Name);
        },
        error: function(_, what) {
            // For debugging!
            console.log("error", what);
        }
    });
}())
