package importer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	csrfToken string = "lz4owIFIAf3Ld7zeWju0Rtc6Y8zbLc4HdwQuzs1OaUcqIAAVPm0AyEcbjAJGf8tC"
	csrfData  string = `
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
					<meta property="og:url" content="https://fabtcg.com/gem/158731/add/" />
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
	<script src="https://dgmi4fxzalveh.cloudfront.net/static/7e0fec03/autocomplete_light/i18n/en.js"></script> <p>Note: Players can be added after they have registered on our fabtcg <a style="text-decoration: underline;" href="https://fabtcg.com/register/" target="_blank">register page</a></p><hr><label>Current players</label><ol style='max-height: 300px; overflow-y: scroll; border-top: 1px solid #6e4d33; border-bottom: 1px solid #6e4d33; padding: 10px 0 10px 30px; list-style-type: decimal;'><li>Alexandre, Marc (69434669)</li><li>Player, Test (TEST12)</li><li>Player, Test (TEST11)</li><li>Player, Test (TEST10)</li><li>Player, Test (TEST9)</li><li>Player, Test (TEST8)</li><li>Player, Test (TEST7)</li><li>Player, Test (TEST6)</li><li>Player, Test (TEST5)</li><li>Player, Test (TEST4)</li><li>Player, Test (TEST3)</li><li>Player, Test (TEST2)</li><li>Player, Test (TEST1)</li></ol> <input type="hidden" name="id" value="158731" id="id_id"> <div id="div_id_country_filter" class="form-group"> <label for="id_country_filter" class="">
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
	<br><a href="/gem/158731/run/">Cancel</a>
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
					let path = window.location.origin + "/gem/" + 158731 + "/generate/" + num_players + "/"
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
		{id: "", gemId: "69434669"},
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
			fmt.Println("Here we are")
			assert.Equal(t, serverUrl, req.Header.Get("Origin"))
			assert.Equal(t, csrfToken, req.Header.Get("X-CSRFToken"))
			assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

			req.ParseForm()
			assert.Equal(t, "test_username", req.Form.Get("username"))
			assert.Equal(t, "test_password", req.Form.Get("password"))
			assert.Equal(t, csrfToken, req.Form.Get("csrfmiddlewaretoken"))
		}

		rw.Write([]byte(csrfData))
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

func TestGetCsrfToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/", req.URL.String())
		rw.Write([]byte(csrfData))
	}))
	defer server.Close()

	connector := CookieConnector{client: server.Client(), baseUrl: server.URL}

	got := connector.GetCsrfToken()
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
