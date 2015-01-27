$(document).ready(function(){
	connectServer();
});

function connectServer()
{
	var sock = null;
	var wsuri = "ws://vim-tips.com:3000/ws";

	try
	{
		sock = new WebSocket(wsuri);
	}catch (e) {
	}

	sock.onopen = function() {
		console.log("connected to " + wsuri);
	};

	sock.onerror = function(e) {
		console.log(" error from connect " + e);
	};

	sock.onclose = function(e) {
		console.log("connection closed (" + e.code + ")");
	};

	sock.onmessage = function(e) {
		console.log("message received: " + e.data);

		var data = $.parseJSON(e.data);

		if(data.Type == "txt")
		{
			$("#txt-counter").text(data.Count);
		}
		else if (data.Type == "json")
		{
			$("#json-counter").text(data.Count);	
		}
	};
}