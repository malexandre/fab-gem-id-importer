package importer

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	csrfToken           string = "lz4owIFIAf3Ld7zeWju0Rtc6Y8zbLc4HdwQuzs1OaUcqIAAVPm0AyEcbjAJGf8tC"
	csrfDataFromAddPage string = `
	<!DOCTYPE html>
	<html lang="en-us">
			<head>
					<title>Running Tournament | Legend Story Studios</title>


					<meta charset="utf-8" />
					<meta name="description" content="" />
					<meta name="author" content="Legend Story Studios Limited" />
					<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
					<link rel="icon" href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/tournaments/img/favicon.ico" type="image/x-icon" />


					<meta property="og:title" content="

									Retailer Platform | Legend Story Studios

					" />
					<meta property="og:url" content="https://fabtcg.com/gem/67890/run/" />
					<meta property="og:image" content="

									https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/website/img/fb-shared.jpg

					" />
					<meta property="og:image:secure_url" content="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/website/img/fb-shared.jpg" />
					<meta property="og:image:type" content="image/jpeg" />
					<meta property="og:image:width" content="3000" />
					<meta property="og:image:height" content="2250" />


							<link crossorigin="anonymous" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css" integrity="sha384-xOolHFLEh07PJGoPkLv1IbcEPTNtaed2xpHsD9ESMhqIYd0nLMwNLD69Npy4HI+N" rel="stylesheet">
							<script crossorigin="anonymous" integrity="sha384-ZvpUoO/+PpLXR1lu4jmpXWu80pZlYUAfxl5NsBMWOEPSjUn/6Z/hRTt8+pR6L4N2" src="https://code.jquery.com/jquery-3.5.1.min.js"></script>
	<script crossorigin="anonymous" integrity="sha384-Fy6S3B9q64WdZWQUiU+q4/2Lc9npb8tCaSX9FK7E8HnRr0Jz8D6OP9dO5Vg3Q9ct" src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/js/bootstrap.bundle.min.js"></script>

							<link rel="stylesheet" href="https://use.typekit.net/qdt0oox.css" />


							<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/gem/styles/gem_main.css?1689003389" rel="stylesheet" type="text/css" />
							<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fontawesome-4.7/css/font-awesome.min.css" rel="stylesheet" type="text/css" />


		<style>
			.win1 td:nth-child(2) {
				background-color: #add6a6;
			}
			.win2 td:nth-child(3) {
				background-color: #add6a6;
			}
			.code-tag {
			padding: 5px 8px;
			background-color: lightgrey;
			border-radius: 5px;
			margin-top: 5px;
			display: inline-block;
			cursor: pointer;
		}
		.fab-btn.poppin {
			background-color: #873229;
			color: white;
		}
		.fab-btn.poppin:hover {
			background-color: #b44337;
			color: white;
		}
		</style>



							<!-- Global site tag (gtag.js) - Google Analytics -->
							<script async src="https://www.googletagmanager.com/gtag/js?id=UA-108157814-3"></script>
							<script>
									window.dataLayer = window.dataLayer || [];
									function gtag(){dataLayer.push(arguments);}
									gtag('js', new Date());

									gtag('config', 'UA-108157814-3');
							</script>



							<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_US/sdk.js#xfbml=1&version=v5.0"></script>









							<meta name="facebook-domain-verification" content="7ufg2esad4li9030pgtzzvkl716eq2" />

			</head>


			<body class="light">
					<div id="fb-root"></div>



	<style>
			#siteNav {
					box-shadow: 0px 1px 12px 3px rgba(211, 211, 211, 0.75);
					background-color: white;
			}
			#logo {
					max-width: 14px;
			}
	</style>

	<nav id="siteNav" class="navbar navbar-expand-lg navbar-light py-2 sticky-top">
			<div class="container">
					<a class="navbar-brand d-flex align-items-center mr-4" href="/">
							<img id="logo" src=https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/gem_wagtail/img/lss_icon.png class="mr-2">
							Partners
					</a>
					<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
							<span class="navbar-toggler-icon"></span>
					</button>
					<div class="collapse navbar-collapse" id="navbarNavDropdown">
							<ul class="navbar-nav mr-auto">

									<li class="nav-item dropdown">
											<a class="nav-link" href="https://fabtcg.com/retailer-news/">Retailer News</a>
									</li>
									<li class="nav-item dropdown">
											<a class="nav-link dropdown-toggle" href="https://fabtcg.com/" id="navbarDropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" target="_blank">
											FABTCG</a>
											<div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
													<a class="dropdown-item" href="https://fabtcg.com/" target="_blank">Official Website</a>
													<a class="dropdown-item" href="https://fabtcg.com/resources/digital-assets/" target="_blank" >Marketing Assets</a>
													<a class="dropdown-item" href="https://www.facebook.com/groups/FABretailer/" target="_blank" >Retailer Group</a>
											</div>
									</li>
									<li class="nav-item dropdown">
											<a class="nav-link" href="https://b2b.legendstory.com/">Where to Buy?</a>
									</li>
							</ul>
							<div class="form-inline my-2 my-lg-0">

											<a class="fab-btn" href="/accounts/logout/">Logout</a>

							</div>
					</div>
			</div>
	</nav>


					<style>
			.alert {
					border-radius: 0px;
			}
	</style>
	<div class="container text-center alerts">

	</div>


	<div class="container py-5">

		<div class="btn-cluter pb-3">
			<li class="d-inline-block">
				<a href="/profile/" class="d-inline-block fab-btn">Back to Profile</a>
			</li>
		</div>
		<h1>(TEST) Test d&#x27;édition de draft pods</h1>
		<div class="mb-3">
			<h3>Tournament Information</h3>

			<style>
					input[type=number]::-webkit-inner-spin-button,
					input[type=number]::-webkit-outer-spin-button {
							opacity: 1;
							cursor: pointer;
					}
					.switch {
							position: relative;
							display: inline-block;
							width: 90px;
							height: 34px;
							margin: 0 0 -7px 0;
					}
					.switch input {
							display:none;
					}
					.slider {
							position: absolute;
							cursor: pointer;
							top: 0;
							left: 0;
							right: 0;
							bottom: 0;
							background-color: #c16557;
							-webkit-transition: .4s;
							transition: .4s;
							border-radius: 34px;
					}
					.slider:before {
							position: absolute;
							content: "";
							height: 18px;
							width: 18px;
							left: 6px;
							bottom: 8px;
							background-color: #c19a57;
							-webkit-transition: .4s;
							transition: .4s;
							border-radius: 50%;
					}
					.slider:after {
							content:'Hidden';
							color: white;
							display: block;
							position: absolute;
							transform: translate(-50%,-50%);
							top: 50%;
							left: 50%;
							font-size: 10px;
							font-family: Verdana, sans-serif;
					}
					input:checked + .slider {
							background-color: #add6a6;
					}
					input:focus + .slider {
							box-shadow: 0 0 1px #2196F3;
					}
					input:checked + .slider:before {
							-webkit-transform: translateX(60px);
							-ms-transform: translateX(60px);
							transform: translateX(60px);
					}
					input:checked + .slider:after {
							content:'Visible';
							color: #203737;
					}
			</style>

	<table>
			<tr>
					<th>Status</th>
					<td>Upcoming</td>
			</tr>
			<tr>
					<th>Type</th>
					<td>On Demand</td>
			</tr>
			<tr>
					<th>Format</th>
					<td>Booster Draft</td>
			</tr>
			<tr>
					<th>Start Time</th>
					<td>July 4, 2023, 9:13 a.m.</td>
			</tr>
			<tr>
					<th>XP Multiplier</th>
					<td>-</td>
			</tr>
			<tr>
					<th>Default Pairings Visibility</th>
					<td>
							<label class="switch">
									<input type="checkbox" id="togBtn" checked>
									<div class="slider round"></div>
							</label>
					</td>
			</tr>
			<tr>
					<th>Starting Table Number</th>
					<td>
							<input id="starting-table-number" name="starting-table-number" type="number" min="1" max="10000" value="1">
					</td>
			</tr>
			<tr>
					<th>Paper Size</th>
					<td>
							<select id="print_pref" onchange="print_select()">
							<option value="A4">A4</option>
							<option value="legal">Legal</option>
							</select>
							(for printouts)
					</td>
			</tr>
	</table>



		<script type="text/javascript">
			$.ajaxSetup({
				headers: { "X-CSRFToken": 'lz4owIFIAf3Ld7zeWju0Rtc6Y8zbLc4HdwQuzs1OaUcqIAAVPm0AyEcbjAJGf8tC' }
			});
			$(document).ready(function(){
					$("#starting-table-number").change(function() {
							let min = 1;
							let max = 10000;
							let value = $(this).val();
							if (value < min) {
									$(this).val(min)
							} else if (value > max) {
									$(this).val(max)
							}
							updateTournamentTableOffset(67890, $(this).val()-1)
					});
					$("#togBtn").change(function() {
							updateTournamentDefaultPairingsVisibility(67890, $(this).is(":checked"))
					});
			});
			function updateTournamentDefaultPairingsVisibility(tournamentid, visible_pairings) {
					let path = "/gem/tournament/67890/"
					$.ajax({
							type : "PUT",
							url: path,
							data: {
									csrfmiddlewaretoken: "lz4owIFIAf3Ld7zeWju0Rtc6Y8zbLc4HdwQuzs1OaUcqIAAVPm0AyEcbjAJGf8tC",
									visible_pairings: visible_pairings
							}
					})
					.fail(function(request, status, error) {
							console.log("Error");
							location.reload();
					})
					.done(function(data) {
							console.log("Tournament Default Pairings Visibility Updated!");
					});
			}
			function updateTournamentTableOffset(tournamentid, table_offset) {
							let path = "/gem/tournament/67890/"
							$.ajax({
									type : "PUT",
									url: path,
									data: {
											csrfmiddlewaretoken: "lz4owIFIAf3Ld7zeWju0Rtc6Y8zbLc4HdwQuzs1OaUcqIAAVPm0AyEcbjAJGf8tC",
											table_offset: table_offset
									},
									success: function() {
											console.log("Tournament Starting Table Updated!")
									}
							});
					}
		</script>




				<a class="fab-btn mt-3" href="/gem/67890/update/">Edit</a>

		</div>
		<hr>
		<div class="mb-3">
			<p><i>Players can join this tournament by searching for the code below in the store locator, or by using the sharing link (click icon to copy link). You can also add / edit players manually.</i></p>
			<a onclick='setClipboard("https://fabtcg.com/join/emerald-owl/")' data-toggle="tooltip" data-placement="top" title="Click to copy sharing link">
				<h4 class="code-tag mb-3 mt-0">
					emerald-owl
					&nbsp;
					<i class="fa fa-copy" aria-hidden="true"></i>
				</h4>
			</a>


				<h3>Pending Players</h3>
				<p><i>The following players have applied to join your tournament. You should approve pending players before you start your tournament. Non-approved players will not be included in the pairings, and will be dropped from the tournament after the tournament is submitted or cancelled.</i></p>

				<div id="refresh">
			<ul style="max-height: 300px; overflow-y: scroll; border-top: 1px solid #6e4d33; border-bottom: 1px solid #6e4d33; padding: 10px;" id='refreshInner'>

					<li>There are no pending players.</li>

			</ol>
	</div>


			<h3>Registered Players</h3>
			<p>13 registered players in tournament.</p>


				<ol style="max-height: 300px; overflow-y: scroll; border-top: 1px solid #6e4d33; border-bottom: 1px solid #6e4d33; padding: 10px 0 10px 30px; list-style-type: decimal;">

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Alexandre, Marc (314314314314)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1201914/">Assign Hero</a>

							</span>

									<span class="float-right mr-3">
											<a href="/gem/67890/player/424242/remove/" onclick="return confirm('This will DELETE the player Alexandre, Marc (314314314314) from this tournament. Proceed?')">
													<i class="fa fa-window-close mr-1" aria-hidden="true"></i>
											</a>
									</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST12)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179791/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST11)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179790/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST10)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179789/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST9)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179788/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST8)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179787/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST7)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179786/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST6)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179785/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST5)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179784/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST4)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179783/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST3)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179782/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST2)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179781/">Assign Hero</a>

							</span>

					</li>

					<li>
							<span style="min-width:200px; display:inline-block;">
									<span >
											Player, Test (TEST1)
									</span>
							</span>
							<span>

											<a href="/gem/67890/hero/1179780/">Assign Hero</a>

							</span>

					</li>

	</ol>



				<a href="/gem/67890/add/" class="fab-btn mr-2">Add Players</a>

					<a href="/gem/67890/import/" class="fab-btn mr-2">Import Players</a>




				<a href="/gem/67890/print-list/" class="fab-btn mr-2" target="_blank">Player List</a>
				<a href="/gem/67890/coverage/players" class="fab-btn mr-2" target="_blank">Player List.csv</a>



				<a href="/gem/67890/print-seat/" class="fab-btn mr-2" target="_blank">Seat Players Alphabetically</a>
				<a href="/gem/67890/print-seat-random/" class="fab-btn mr-2" target="_blank">Seat Players Randomly</a>



				<a href="/gem/67890/add-myself/" class="fab-btn float-right">Add Yourself</a>



		</div>
		<hr>
		<div class="mb-3">
			<h3>Tournament Rounds</h3>

				<p>Your tournament hasn't started.</p>

					<a href="/gem/67890/next-round/" class="fab-btn mr-2">Pair Round 1</a>


		</div>


			<hr>
			<div class="mb-3">

						<a href="/gem/67890/end/" class="fab-btn poppin mr-2 disabled" onclick="return confirm('This will end the tournament and submit the results. You will be unable to edit this tournament after it has ended. This action is irreversible. Proceed?')">Submit Tournament</a>


					<a class="fab-btn float-right" href="/gem/67890/delete/" onclick="return confirm('WARNING! This will DELETE the tournament permanently. This action is NOT reversible. Are you sure you want to proceed?')">Delete Tournament</a>


			</div>


			<hr>
			<div class="copyright">G.E.M &copy; 2023 <a href="https://legendstory.com">Legend Story Studios</a>. Support available at <a href="mailto:op@fabtcg.com">op@fabtcg.com</a> or via. live chat. Please send bug reports to <a href="mailto:info@legendstory.com">info@legendstory.com</a>.</div>
	</div>



							<!-- begin olark code -->
							<script type="text/javascript" async> ;(function(o,l,a,r,k,y){if(o.olark)return; r="script";y=l.createElement(r);r=l.getElementsByTagName(r)[0]; y.async=1;y.src="//"+a;r.parentNode.insertBefore(y,r); y=o.olark=function(){k.s.push(arguments);k.t.push(+new Date)}; y.extend=function(i,j){y("extend",i,j)}; y.identify=function(i){y("identify",k.i=i)}; y.configure=function(i,j){y("configure",i,j);k.c[i]=j}; k=y._={s:[],t:[+new Date],c:{},l:a}; })(window,document,"static.olark.com/jsclient/loader.js");
							/* custom configuration goes here (www.olark.com/documentation) */
							olark.identify('7373-875-10-8920');</script>
							<!-- end olark code -->



							<script
									src="https://browser.sentry-cdn.com/7.24.2/bundle.min.js"
									integrity="sha384-9uapQGwcpfQ0MhBPF9K0kJQcSl6WlSXljCtI/zMwbz8mDaI+FgInHBwR+MNVz+h+"
									crossorigin="anonymous"
							></script>
							<script>
									if (window.Sentry) {
											Sentry.init({ dsn: 'https://8ba02bad271d46b6bb689ce708df36be@o115950.ingest.sentry.io/4504318134124544' });
									}
							</script>

			</body>


			<footer>



									<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fabsite/jsp/Enquire/enquire.min.js"></script>
									<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fabsite/jsp/Slick/slick.min.js"></script>

									<script>
											$(document).ready(function(){
													// Trigger all bootstrap 4 tooltips
													$(function () {
															$('[data-toggle="tooltip"]').tooltip()
													});
											})
									</script>


		<script src="https://cdn.jsdelivr.net/npm/js-cookie@beta/dist/js.cookie.min.js"></script>
		<script type="text/javascript">
			if(Cookies.get('print_pref') == null){
				Cookies.set('print_pref', $('#print_pref').val(), {expires : 365})
			}
			$('#print_pref').val(Cookies.get('print_pref'));
			function print_select(){
				Cookies.set('print_pref', $('#print_pref').val(), {expires : 365})
			}
			function setClipboard(value) {
				var tempInput = document.createElement("input");
				tempInput.style = "position: absolute; left: -1000px; top: -1000px";
				tempInput.value = value;
				document.body.appendChild(tempInput);
				tempInput.select();
				document.execCommand("copy");
				document.body.removeChild(tempInput);
				alert('Sharing link copied to clipboard!')
			}

			function refreshElement(){
				$('#refresh').load('/gem/67890/run/ #refreshInner');
			}
			$(document).ready(function(){
				setInterval("refreshElement();", 15000);
			})

		</script>


			</footer>
	</html>
	`
	csrfDataFromLoginPage string = `
	<!DOCTYPE html>
	<html lang="en-us">
			<head>
					<title>Retailer Center | Legend Story Studios</title>


					<meta charset="utf-8" />
					<meta name="description" content="" />
					<meta name="author" content="Legend Story Studios Limited" />
					<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
					<link rel="icon" href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/tournaments/img/favicon.ico" type="image/x-icon" />


					<meta property="og:title" content="

									Retailer Platform | Legend Story Studios

					" />
					<meta property="og:url" content="https://fabtcg.com/" />
					<meta property="og:image" content="

									https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/website/img/fb-shared.jpg

					" />
					<meta property="og:image:secure_url" content="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/website/img/fb-shared.jpg" />
					<meta property="og:image:type" content="image/jpeg" />
					<meta property="og:image:width" content="3000" />
					<meta property="og:image:height" content="2250" />


							<link crossorigin="anonymous" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css" integrity="sha384-xOolHFLEh07PJGoPkLv1IbcEPTNtaed2xpHsD9ESMhqIYd0nLMwNLD69Npy4HI+N" rel="stylesheet">
							<script crossorigin="anonymous" integrity="sha384-ZvpUoO/+PpLXR1lu4jmpXWu80pZlYUAfxl5NsBMWOEPSjUn/6Z/hRTt8+pR6L4N2" src="https://code.jquery.com/jquery-3.5.1.min.js"></script>
	<script crossorigin="anonymous" integrity="sha384-Fy6S3B9q64WdZWQUiU+q4/2Lc9npb8tCaSX9FK7E8HnRr0Jz8D6OP9dO5Vg3Q9ct" src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/js/bootstrap.bundle.min.js"></script>

							<link rel="stylesheet" href="https://use.typekit.net/qdt0oox.css" />


							<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/gem/styles/gem_main.css?1688986479" rel="stylesheet" type="text/css" />
							<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fontawesome-4.7/css/font-awesome.min.css" rel="stylesheet" type="text/css" />





							<!-- Global site tag (gtag.js) - Google Analytics -->
							<script async src="https://www.googletagmanager.com/gtag/js?id=UA-108157814-3"></script>
							<script>
									window.dataLayer = window.dataLayer || [];
									function gtag(){dataLayer.push(arguments);}
									gtag('js', new Date());

									gtag('config', 'UA-108157814-3');
							</script>



							<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_US/sdk.js#xfbml=1&version=v5.0"></script>









							<meta name="facebook-domain-verification" content="7ufg2esad4li9030pgtzzvkl716eq2" />

			</head>


			<body class="light">
					<div id="fb-root"></div>



	<style>
			#siteNav {
					box-shadow: 0px 1px 12px 3px rgba(211, 211, 211, 0.75);
					background-color: white;
			}
			#logo {
					max-width: 14px;
			}
	</style>

	<nav id="siteNav" class="navbar navbar-expand-lg navbar-light py-2 sticky-top">
			<div class="container">
					<a class="navbar-brand d-flex align-items-center mr-4" href="/">
							<img id="logo" src=https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/gem_wagtail/img/lss_icon.png class="mr-2">
							Partners
					</a>
					<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
							<span class="navbar-toggler-icon"></span>
					</button>
					<div class="collapse navbar-collapse" id="navbarNavDropdown">
							<ul class="navbar-nav mr-auto">

									<li class="nav-item dropdown">
											<a class="nav-link" href="https://fabtcg.com/retailer-news/">Retailer News</a>
									</li>
									<li class="nav-item dropdown">
											<a class="nav-link dropdown-toggle" href="https://fabtcg.com/" id="navbarDropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" target="_blank">
											FABTCG</a>
											<div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
													<a class="dropdown-item" href="https://fabtcg.com/" target="_blank">Official Website</a>
													<a class="dropdown-item" href="https://fabtcg.com/resources/digital-assets/" target="_blank" >Marketing Assets</a>
													<a class="dropdown-item" href="https://www.facebook.com/groups/FABretailer/" target="_blank" >Retailer Group</a>
											</div>
									</li>
									<li class="nav-item dropdown">
											<a class="nav-link" href="https://b2b.legendstory.com/">Where to Buy?</a>
									</li>
							</ul>
							<div class="form-inline my-2 my-lg-0">

							</div>
					</div>
			</div>
	</nav>


					<style>
			.alert {
					border-radius: 0px;
			}
	</style>
	<div class="container text-center alerts">

	</div>


	<style>
			html {height: 100%;}
			body {height: 100%;}
			.container.container-centered {
					min-height: 80vh;
			}
			.container.container-centered > .row {
					align-items: center;
					height: 100%;
			}
	</style>

	<div class="container container-centered py-5 d-flex justify-content-center align-items-center">
			<div class="col-12 col-lg-8 p-lg-0 m-0">

	<div class="login py-5">
			<h1>G.E.M Login</h1>
			<p>Flesh and Blood TCG Game & Event Management software</p>
			<form method="post" class="form">
					<input type="hidden" name="csrfmiddlewaretoken" value="lz4owIFIAf3Ld7zeWju0Rtc6Y8zbLc4HdwQuzs1OaUcqIAAVPm0AyEcbjAJGf8tC">
					<div class="pt-3">
					<div class="form-group"><label for="id_username">Email address</label><input type="text" name="username" autofocus autocapitalize="none" autocomplete="username" maxlength="254" class="form-control" placeholder="Email address" title="" required id="id_username"></div>
	<div class="form-group"><label for="id_password">Password</label><input type="password" name="password" autocomplete="current-password" class="form-control" placeholder="Password" title="" required id="id_password"></div>
					</div>
					<div class="form-group">
							<button type="submit" id="submit" class="fab-btn my-2">
									Login
							</button>
							<a href="/register/" class="fab-btn my-2 float-right">Register</a>
							<br>
							<a href="/accounts/password_reset/">Forgot password?</a>
					</div>
					<p><i>Note that you may login to GEM using with your Flesh and Blood account, and vice versa. LGS and TO users may find it benefitial to register separate accounts for tournament-management and personal use.</i></p>
			</form>
	</div>

			</div>
	</div>



							<!-- begin olark code -->
							<script type="text/javascript" async> ;(function(o,l,a,r,k,y){if(o.olark)return; r="script";y=l.createElement(r);r=l.getElementsByTagName(r)[0]; y.async=1;y.src="//"+a;r.parentNode.insertBefore(y,r); y=o.olark=function(){k.s.push(arguments);k.t.push(+new Date)}; y.extend=function(i,j){y("extend",i,j)}; y.identify=function(i){y("identify",k.i=i)}; y.configure=function(i,j){y("configure",i,j);k.c[i]=j}; k=y._={s:[],t:[+new Date],c:{},l:a}; })(window,document,"static.olark.com/jsclient/loader.js");
							/* custom configuration goes here (www.olark.com/documentation) */
							olark.identify('7373-875-10-8920');</script>
							<!-- end olark code -->



							<script
									src="https://browser.sentry-cdn.com/7.24.2/bundle.min.js"
									integrity="sha384-9uapQGwcpfQ0MhBPF9K0kJQcSl6WlSXljCtI/zMwbz8mDaI+FgInHBwR+MNVz+h+"
									crossorigin="anonymous"
							></script>
							<script>
									if (window.Sentry) {
											Sentry.init({ dsn: 'https://8ba02bad271d46b6bb689ce708df36be@o115950.ingest.sentry.io/4504318134124544' });
									}
							</script>

			</body>


			<footer>



									<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fabsite/jsp/Enquire/enquire.min.js"></script>
									<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fabsite/jsp/Slick/slick.min.js"></script>

									<script>
											$(document).ready(function(){
													// Trigger all bootstrap 4 tooltips
													$(function () {
															$('[data-toggle="tooltip"]').tooltip()
													});
											})
									</script>




			</footer>
	</html>
	`
	userListData string = `
	<!DOCTYPE html>
	<html lang="en-us">
			<head>
					<title>Add Players | Legend Story Studios</title>


					<meta charset="utf-8" />
					<meta name="description" content="" />
					<meta name="author" content="Legend Story Studios Limited" />
					<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
					<link rel="icon" href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/tournaments/img/favicon.ico" type="image/x-icon" />


					<meta property="og:title" content="

									Retailer Platform | Legend Story Studios

					" />
					<meta property="og:url" content="https://fabtcg.com/gem/67890/add/" />
					<meta property="og:image" content="

									https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/website/img/fb-shared.jpg

					" />
					<meta property="og:image:secure_url" content="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/website/img/fb-shared.jpg" />
					<meta property="og:image:type" content="image/jpeg" />
					<meta property="og:image:width" content="3000" />
					<meta property="og:image:height" content="2250" />


							<link crossorigin="anonymous" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css" integrity="sha384-xOolHFLEh07PJGoPkLv1IbcEPTNtaed2xpHsD9ESMhqIYd0nLMwNLD69Npy4HI+N" rel="stylesheet">
							<script crossorigin="anonymous" integrity="sha384-ZvpUoO/+PpLXR1lu4jmpXWu80pZlYUAfxl5NsBMWOEPSjUn/6Z/hRTt8+pR6L4N2" src="https://code.jquery.com/jquery-3.5.1.min.js"></script>
	<script crossorigin="anonymous" integrity="sha384-Fy6S3B9q64WdZWQUiU+q4/2Lc9npb8tCaSX9FK7E8HnRr0Jz8D6OP9dO5Vg3Q9ct" src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/js/bootstrap.bundle.min.js"></script>

							<link rel="stylesheet" href="https://use.typekit.net/qdt0oox.css" />


							<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/gem/styles/gem_main.css?1688988329" rel="stylesheet" type="text/css" />
							<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fontawesome-4.7/css/font-awesome.min.css" rel="stylesheet" type="text/css" />


	<style>
		.django-ckeditor-widget {
			width: 100%;
		}
		div#cke_id_description {
			width: 100% !important;
			margin: 0px;
		}
		input[type=number]::-webkit-inner-spin-button,
		input[type=number]::-webkit-outer-spin-button {
			opacity: 1;
			cursor: pointer;
		}
		#generate-players-input {
			background-color: white;
			border: 1px solid #c19a57;
			border-radius: 0px;
			color: black;
			padding: 0.2rem 0.5rem;
			display: inline-block;
			font-weight: 400;
			text-align: center;
			white-space: nowrap;
			vertical-align: middle;
			user-select: none;
			font-size: 1rem;
			line-height: 1.5;
			border-radius: 3px;
		}
		#generate-players-input:hover {
			background-color: white;
			cursor: text;
		}
		.add_player__form-control {
			min-height: 42px;
			height: min-content !important;
		}
	</style>



							<!-- Global site tag (gtag.js) - Google Analytics -->
							<script async src="https://www.googletagmanager.com/gtag/js?id=UA-108157814-3"></script>
							<script>
									window.dataLayer = window.dataLayer || [];
									function gtag(){dataLayer.push(arguments);}
									gtag('js', new Date());

									gtag('config', 'UA-108157814-3');
							</script>



							<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_US/sdk.js#xfbml=1&version=v5.0"></script>









							<meta name="facebook-domain-verification" content="7ufg2esad4li9030pgtzzvkl716eq2" />

			</head>


			<body class="light">
					<div id="fb-root"></div>



	<style>
			#siteNav {
					box-shadow: 0px 1px 12px 3px rgba(211, 211, 211, 0.75);
					background-color: white;
			}
			#logo {
					max-width: 14px;
			}
	</style>

	<nav id="siteNav" class="navbar navbar-expand-lg navbar-light py-2 sticky-top">
			<div class="container">
					<a class="navbar-brand d-flex align-items-center mr-4" href="/">
							<img id="logo" src=https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/gem_wagtail/img/lss_icon.png class="mr-2">
							Partners
					</a>
					<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
							<span class="navbar-toggler-icon"></span>
					</button>
					<div class="collapse navbar-collapse" id="navbarNavDropdown">
							<ul class="navbar-nav mr-auto">

									<li class="nav-item dropdown">
											<a class="nav-link" href="https://fabtcg.com/retailer-news/">Retailer News</a>
									</li>
									<li class="nav-item dropdown">
											<a class="nav-link dropdown-toggle" href="https://fabtcg.com/" id="navbarDropdownMenuLink" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" target="_blank">
											FABTCG</a>
											<div class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
													<a class="dropdown-item" href="https://fabtcg.com/" target="_blank">Official Website</a>
													<a class="dropdown-item" href="https://fabtcg.com/resources/digital-assets/" target="_blank" >Marketing Assets</a>
													<a class="dropdown-item" href="https://www.facebook.com/groups/FABretailer/" target="_blank" >Retailer Group</a>
											</div>
									</li>
									<li class="nav-item dropdown">
											<a class="nav-link" href="https://b2b.legendstory.com/">Where to Buy?</a>
									</li>
							</ul>
							<div class="form-inline my-2 my-lg-0">

											<a class="fab-btn" href="/accounts/logout/">Logout</a>

							</div>
					</div>
			</div>
	</nav>


					<style>
			.alert {
					border-radius: 0px;
			}
	</style>
	<div class="container text-center alerts">

	</div>


	<div class="container py-5">

		<div class="btn-cluter pb-3">
			<li class="d-inline-block"><a href="javascript:history.back()" class="fab-btn">Cancel</a></li>
		</div>

			<h1>Generate Players</h1>
			<p>You may quickly generate and add a specified number of non-existent players to your tournament for testing purposes below.</p>
			<input id="generate-players-input" name="generate-players-input" type="number" min="1" max="1000" value="16">
			<span id="generate-players-btn" class="fab-btn">Generate</span>
			<hr/>

		<h1>Add Players</h1>


	<form  method="post" > <input type="hidden" name="csrfmiddlewaretoken" value="VQ8sJ1pdfklpOkR5kSx9ts3jkWyHjAWEjsnl4yBCiiKA4qk4qu0N6L5blnTU071I"> <link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/admin/css/vendor/select2/select2.min.css" media="screen" rel="stylesheet">
	<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/admin/css/autocomplete.css" media="screen" rel="stylesheet">
	<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/select2.css" media="screen" rel="stylesheet">
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/admin/js/vendor/select2/select2.full.js"></script>
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/autocomplete_light.min.js"></script>
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/select2.min.js"></script>
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/i18n/en.js"></script> <p>Note: Players can be added after they have registered on our fabtcg <a style="text-decoration: underline;" href="https://fabtcg.com/register/" target="_blank">register page</a></p><hr><label>Current players</label><ol style='max-height: 300px; overflow-y: scroll; border-top: 1px solid #6e4d33; border-bottom: 1px solid #6e4d33; padding: 10px 0 10px 30px; list-style-type: decimal;'><li>Alexandre, Marc (314314314314)</li><li>Player, Test (TEST12)</li><li>Player, Test (TEST11)</li><li>Player, Test (TEST10)</li><li>Player, Test (TEST9)</li><li>Player, Test (TEST8)</li><li>Player, Test (TEST7)</li><li>Player, Test (TEST6)</li><li>Player, Test (TEST5)</li><li>Player, Test (TEST4)</li><li>Player, Test (TEST3)</li><li>Player, Test (TEST2)</li><li>Player, Test (TEST1)</li></ol> <input type="hidden" name="id" value="67890" id="id_id"> <div id="div_id_country_filter" class="form-group"> <label for="id_country_filter" class="">
									Country filter
							</label> <div> <select name="country_filter" class="select custom-select" id="id_country_filter"> <option value="ALL">All</option> <option value="NZ">New Zealand</option> <option value="AU">Australia</option> <option value="MY">Malaysia</option> <option value="US">United States of America</option>

	</select> </div> </div> <div id="div_id_players_to_add" class="form-group"> <label for="id_players_to_add" class=" requiredField">
									Players to add<span class="asteriskField">*</span> </label> <div> <select name="players_to_add" class="add_player__form-control modelselect2multiple custom-select" placeholder="Players" required id="id_players_to_add" data-autocomplete-light-language="en" data-autocomplete-light-url="/dal-urls/player-autocomplete-add/" data-autocomplete-light-function="select2" multiple>
	</select><div style="display:none" class="dal-forward-conf" id="dal-forward-conf-for_id_players_to_add"><script type="text/dal-forward-conf">[{"type": "field", "src": "id"}, {"type": "field", "src": "country_filter"}]</script></div> </div> </div> <div
			class="buttonHolder"> <input type="submit"
			name="save"
			value="Save"

					class="btn btn-primary fab-btn mb-3"
					id="submit-id-save"


			/>
	<br><a href="/gem/67890/run/">Cancel</a>
	</div> </form>



			<hr>
			<div class="copyright">G.E.M &copy; 2023 <a href="https://legendstory.com">Legend Story Studios</a>. Support available at <a href="mailto:op@fabtcg.com">op@fabtcg.com</a> or via. live chat. Please send bug reports to <a href="mailto:info@legendstory.com">info@legendstory.com</a>.</div>
	</div>



							<!-- begin olark code -->
							<script type="text/javascript" async> ;(function(o,l,a,r,k,y){if(o.olark)return; r="script";y=l.createElement(r);r=l.getElementsByTagName(r)[0]; y.async=1;y.src="//"+a;r.parentNode.insertBefore(y,r); y=o.olark=function(){k.s.push(arguments);k.t.push(+new Date)}; y.extend=function(i,j){y("extend",i,j)}; y.identify=function(i){y("identify",k.i=i)}; y.configure=function(i,j){y("configure",i,j);k.c[i]=j}; k=y._={s:[],t:[+new Date],c:{},l:a}; })(window,document,"static.olark.com/jsclient/loader.js");
							/* custom configuration goes here (www.olark.com/documentation) */
							olark.identify('7373-875-10-8920');</script>
							<!-- end olark code -->



							<script
									src="https://browser.sentry-cdn.com/7.24.2/bundle.min.js"
									integrity="sha384-9uapQGwcpfQ0MhBPF9K0kJQcSl6WlSXljCtI/zMwbz8mDaI+FgInHBwR+MNVz+h+"
									crossorigin="anonymous"
							></script>
							<script>
									if (window.Sentry) {
											Sentry.init({ dsn: 'https://8ba02bad271d46b6bb689ce708df36be@o115950.ingest.sentry.io/4504318134124544' });
									}
							</script>

			</body>


			<footer>



									<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fabsite/jsp/Enquire/enquire.min.js"></script>
									<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/fabsite/jsp/Slick/slick.min.js"></script>

									<script>
											$(document).ready(function(){
													// Trigger all bootstrap 4 tooltips
													$(function () {
															$('[data-toggle="tooltip"]').tooltip()
													});
											})
									</script>


		<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/admin/css/vendor/select2/select2.min.css" media="screen" rel="stylesheet">
	<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/admin/css/autocomplete.css" media="screen" rel="stylesheet">
	<link href="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/select2.css" media="screen" rel="stylesheet">
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/admin/js/vendor/select2/select2.full.js"></script>
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/autocomplete_light.min.js"></script>
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/select2.min.js"></script>
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/i18n/en.js"></script>
		<script type="text/javascript">
			$.ajaxSetup({
				headers: { "X-CSRFToken": 'VQ8sJ1pdfklpOkR5kSx9ts3jkWyHjAWEjsnl4yBCiiKA4qk4qu0N6L5blnTU071I' }
			});
			$(document).ready(function(){
				if ($('#id_tournament_address').val() == "None") {
					$('#id_tournament_address').val('');
				}
				if ($('#id_s_address').val() == "None") {
					$('#id_s_address').val('');
				}
				$("#generate-players-btn").click(function(e) {
					e.preventDefault()
					let min = 1;
					let max = 1000;
					let num_players = $("#generate-players-input").val();
					if (num_players < min) {
						num_players = min;
						$("#generate-players-input").val(num_players)
					} else if (num_players > max) {
						num_players = max;
						$("#generate-players-input").val(num_players)
					}
					let path = window.location.origin + "/gem/" + 67890 + "/generate/" + num_players + "/"
					window.location.href = path
				});
			});
			function autoInit() {
				var input = document.getElementById('id_address_form');
				var autocomplete = new google.maps.places.Autocomplete(input);
			}
		</script>
			</footer>
	</html>
	`
	expectedUsers []User = []User{
		{id: "", gemId: "314314314314"},
		{id: "", gemId: "TEST12"},
		{id: "", gemId: "TEST11"},
		{id: "", gemId: "TEST10"},
		{id: "", gemId: "TEST9"},
		{id: "", gemId: "TEST8"},
		{id: "", gemId: "TEST7"},
		{id: "", gemId: "TEST6"},
		{id: "", gemId: "TEST5"},
		{id: "", gemId: "TEST4"},
		{id: "", gemId: "TEST3"},
		{id: "", gemId: "TEST2"},
		{id: "", gemId: "TEST1"},
	}
)

func TestConnect(t *testing.T) {
	serverUrl := ""

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/", req.URL.String())

		if req.Method == http.MethodPost {
			assert.Equal(t, serverUrl, req.Header.Get("Origin"))
			assert.Equal(t, csrfToken, req.Header.Get("X-CSRFToken"))
			assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

			req.ParseForm()
			assert.Equal(t, "test_username", req.Form.Get("username"))
			assert.Equal(t, "test_password", req.Form.Get("password"))
			assert.Equal(t, csrfToken, req.Form.Get("csrfmiddlewaretoken"))
		}

		rw.Write([]byte(csrfDataFromLoginPage))
	}))
	defer server.Close()

	serverUrl = server.URL
	connector := CookieConnector{
		client:   server.Client(),
		baseUrl:  serverUrl,
		username: "test_username",
		password: "test_password",
	}

	got := connector.Connect()
	assert.Equal(t, true, got)
}

func TestGetUserByGemId(t *testing.T) {
	gemId := "12345"
	eventId := "67890"

	params := url.Values{}
	params.Add("q", gemId)
	params.Add("forward", `{"id":"`+eventId+`","country_filter":"ALL"}`)

	t.Run("Valid user", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/dal-urls/player-autocomplete-add/?"+params.Encode(), req.URL.String())
			rw.Write([]byte(`{
				"results": [
					{
						"id": "42314",
						"text": "Test User (12345)",
						"selected_text": "Test User (12345)"
					}
				],
				"pagination": {
					"more": false
				}
			}`))
		}))
		defer server.Close()

		connector := CookieConnector{client: server.Client(), baseUrl: server.URL}
		got, err := connector.GetUserByGemId(eventId, gemId)
		assert.Equal(t, nil, err)
		assert.Equal(t, User{id: "42314"}, got)
	})

	t.Run("No user", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/dal-urls/player-autocomplete-add/?"+params.Encode(), req.URL.String())
			rw.Write([]byte(`{
				"results": [],
				"pagination": {
					"more": false
				}
			}`))
		}))
		defer server.Close()

		connector := CookieConnector{client: server.Client(), baseUrl: server.URL}
		got, err := connector.GetUserByGemId(eventId, gemId)
		assert.Equal(t, ErrNotFound, err)
		assert.Equal(t, User{}, got)
	})

	t.Run("Invalid user", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/dal-urls/player-autocomplete-add/?"+params.Encode(), req.URL.String())
			rw.Write([]byte(`{
				"results": [
					{
						"id": null,
						"text": "Test User (12345)",
						"selected_text": "Test User (12345)"
					}
				],
				"pagination": {
					"more": false
				}
			}`))
		}))
		defer server.Close()

		connector := CookieConnector{client: server.Client(), baseUrl: server.URL}
		got, err := connector.GetUserByGemId(eventId, gemId)
		assert.Equal(t, ErrInvalidData, err)
		assert.Equal(t, User{}, got)
	})
}

func TestAddUsersToEvent(t *testing.T) {
	serverUrl := ""
	eventId := "67890"

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			assert.Equal(t, "/gem/"+eventId+"/add/", req.URL.String())
			assert.Equal(t, serverUrl, req.Header.Get("Origin"))
			assert.Equal(t, csrfToken, req.Header.Get("X-CSRFToken"))
			assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

			req.ParseForm()
			assert.Equal(t, eventId, req.Form.Get("id"))
			assert.Equal(t, "ALL", req.Form.Get("country_filter"))
			assert.Equal(t, "Save", req.Form.Get("save"))
			assert.Equal(t, csrfToken, req.Form.Get("csrfmiddlewaretoken"))
			assert.Contains(t, req.Form["players_to_add"], "123")
			assert.Contains(t, req.Form["players_to_add"], "345")
			assert.Contains(t, req.Form["players_to_add"], "1179791")
		}

		rw.Write([]byte(csrfDataFromAddPage))
	}))
	defer server.Close()

	serverUrl = server.URL
	connector := CookieConnector{
		client:  server.Client(),
		baseUrl: serverUrl,
	}

	got := connector.AddUsersToEvent(eventId, []User{{id: "123"}, {id: "345"}, {id: "1179791"}})
	assert.Equal(t, true, got)
}

func TestGetCsrfTokenFromLoginPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/", req.URL.String())
		rw.Write([]byte(csrfDataFromLoginPage))
	}))
	defer server.Close()

	connector := CookieConnector{client: server.Client(), baseUrl: server.URL}

	got := connector.GetCsrfTokenFromLoginPage()
	assert.Equal(t, csrfToken, got)
}

func TestGetCsrfTokenFromAddPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/gem/42314/run/", req.URL.String())
		rw.Write([]byte(csrfDataFromAddPage))
	}))
	defer server.Close()

	connector := CookieConnector{client: server.Client(), baseUrl: server.URL}

	got := connector.GetCsrfTokenFromAddPage("42314")
	assert.Equal(t, csrfToken, got)
}

func TestGetUsersFromEvent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/gem/42314/add/", req.URL.String())
		rw.Write([]byte(userListData))
	}))
	defer server.Close()
	connector := CookieConnector{client: server.Client(), baseUrl: server.URL}

	players := connector.GetUsersFromEvent("42314")
	assert.ElementsMatch(t, expectedUsers, players)
}
