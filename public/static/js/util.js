var endTime = undefined;
var timer = undefined;
function createCountdownTimer(endTime, textDiv, iframe) {
	if (endTime == -1) {
		$('#time_remaining').text("Not Started");
	}
	else if (!timer) {
		clearInterval(timer);
		timer = setInterval(function() {
			var distance = endTime - Date.now();
			var countDownText;
			if (endTime == -1) {
				textDiv.text("Not Started");
			}
			if (distance > 0) {
				// Time calculations for minutes and seconds
				var minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
				var seconds = Math.floor((distance % (1000 * 60)) / 1000);
				if (seconds < 10 && minutes >= 1) {
					seconds = "0" + seconds;
				}
				
				if (minutes >= 1) {
					// Format "mm:ss minutes"
					countDownText = minutes + ":" + seconds + " minutes";
				}
				else if (seconds > 1) {
					// Format "ss seconds"
					countDownText = seconds + " seconds";
				}
				else {
					// Format "1 second"
					countDownText = "1 second";
				}
			}
			else {
				// When time's up stop the timer, clear the frame, and alert the user
				clearInterval(timer);
				timer = undefined;
				if (iframe) {
					iframe.html('');
				}
				countDownText = "EXPIRED";
				alert("Time's up!");
			}
			textDiv.text(countDownText);
		}, 1500);
	};
};

function stopCountDown() {
	$('#time_remaining').text("");
	clearInterval(timer);
	timer = undefined;
}

function getUrlParameter(name) {
    var url = decodeURIComponent(window.location.search.substring(1));
    var urlVariables = url.split('&');
    
    for (var i = 0; i < urlVariables.length; i++) {
        var param = urlVariables[i].split('=');
        if (name == param[0]) {
            return param[1] === undefined ? true : param[1];
        }
    }
    return undefined;
};