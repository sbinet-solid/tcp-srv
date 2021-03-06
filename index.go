// Copyright 2017 The solid-mon-rpi Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const indexTmpl = `
<html>
	<head>
		<title>SoLiD sensors monitoring</title>
		<script type="text/javascript">
		var sock = null;

		function update(data) {
			var p = null;
			
			p = document.getElementById("update-message");
			p.innerHTML = "Last Update: <code>"+data.update+"</code>";

			p = document.getElementById("sensor-plot");
			p.innerHTML = data.plot;

			p = document.getElementById("fast-data");
			p.innerHTML = "<pre>"+data.data+"</pre>";

			p = document.getElementById("sensor-plot-trends");
			p.innerHTML = data.trends;
		};

		window.onload = function() {
			sock = new WebSocket("ws://"+location.host+"/data");
			sock.onmessage = function(event) {
				var data = JSON.parse(event.data);
				update(data);
			};
		};
		</script>

		<style>
		.solid-plot-style {
			font-size: 14px;
			line-height: 1.2em;
		}
		</style>
	</head>

	<body>
		<h2>SoLiD sensors monitoring plots ({{.Freq}} Hz)</h2>

		<div id="fast-plots">
			<div id="sensor-plot" class="solid-plot-style"></div>
		</div>

		<br>
		<div>Server Version: {{.Version}}</div>
		<div id="update-message">Last Update: N/A
		</div>
		<div id="fast-data"></div>

		<h2>SoLiD sensors monitoring plots (trends)</h2>

		<div id="trend-plots">
			<div id="sensor-plot-trends" class="solid-plot-style"></div>
		</div>
	</body>
</html>
`
