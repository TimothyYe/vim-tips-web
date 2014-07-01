$(document).ready(function(){
	setInterval("refreshTips()", 10000);
});

function refreshTips()
{
	$("#tip-content").hide();

	$.get("http://localhost:3000", function(data){
		var result = $(data);
		$(".content").text(result.find('.content').text());
		$(".comment").text(result.find('.comment').text());
		$("#tip_id").text(result.find("#tip_id").text());
		$("#tip-content a").attr("href", "/" + result.find("#tip_id").text());
	});

	$("#tip-content").delay(1000).fadeIn(2500);
}
