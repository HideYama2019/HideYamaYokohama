$(function(){
    $("h1").click(function() {
        $(this).css("color", "red");
    });
    
    $("h2").hover(function() {
        $(this).css("color", "red");
    });    

    $("li").hover(function() {
        $(this).text("momonga");
    }).mouseout(function() {
        $(this).html("<h1>pepe</h1>").css("color", "pink");
    });
})
