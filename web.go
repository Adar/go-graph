/*
go-graph - Graphs from shell STDIN.

Copyright (c) 2016 Christian Senkowski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func startWeb() {
	tmpl = template.Must(template.New(".").Parse(html))
	http.HandleFunc("/", webServer)
	listenOn := fmt.Sprintf("0.0.0.0:%d", *portFlag)
	log.Fatal(http.ListenAndServe(listenOn, nil))
}

func webServer(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, allStats); err != nil {
		log.Print(err)
	}
}

const html = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>rtop-viz</title>
    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet">
	<link href='https://fonts.googleapis.com/css?family=Source+Sans+Pro' rel='stylesheet' type='text/css'>
	{{$all := .}}
	<style type="text/css">
	body { background: #E8F5E9; font-family: "Source Sans Pro", sans-serif; font-size: 12px; }
	.chart {
		border-radius: 3px; background-color: #fff;
		box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
		margin: 5px;
	}
	.chartc {
		display: flex; display: -webkit-flex; flex-wrap: wrap; -webkit-flex-wrap: wrap;
	}
	.dygraph-legend { font-size: 12px !important; }
	.dygraph-title { font-size: 14px; font-weight: 400; }
	h2 { text-align: center; font-size: 24px; padding: 1.1em; }
	.footer { margin: 5em 0 2em 0; color: #aaa; text-align: center; font-size: 14px }
	a { color: #777; }
	</style>
	<style type="text/css">
	#div_g {
		position: absolute;
		left: 10px;
		right: 10px;
		top: 40px;
		bottom: 10px;
	}
	</style>

  </head>
  <body>
	<div class="container-fluid">
	  <div class="row">
	  <div class="col-sm-12 chartc">
		  <div id="id-1" class="chart" style="width: 100%;"></div>
		</div>
	  </div>
	</div>

    <script src="https://code.jquery.com/jquery-2.1.4.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/dygraph/1.1.1/dygraph-combined.js"></script>
	<script type="text/javascript">
	$(function() {
		{{$sr := ($all.GetStats)}}
		new Dygraph(
			document.getElementById("id-1"),
			[
				{{range $sr.Entries}}
				[ new Date( {{.At.Unix}} * 1000 ), {{.Value}} ],
				{{end}}
			],
			{
				title: "{{$all.GetTitle}}",
				axisLabelFontSize: 10,
				axes: { y: { axisLabelWidth: 70 } },
				labels: [ "X", "Value" ],
				maxNumberWidth: 100,
          		gridLineColor: 'rgb(200,200,200)',
			});
			var search = location.search || '';
			if (search.indexOf('noreload') === -1) {
				window.setTimeout(function() {
				window.location.reload();
				}, 5000);
			}
		});
	</script>
  </body>
</html>
`
