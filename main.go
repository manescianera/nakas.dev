package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port     int
	crt, key string
)

func init() {
	flag.IntVar(&port, "port", 0, "server port")
	flag.StringVar(&crt, "crt", "", "define tls certificate path")
	flag.StringVar(&key, "key", "", "define tls public key path")
	flag.Parse()
}

func main() {
	const (
		colorBackground          = "#bcc5d8"
		colorText                = "#1b1d28"
		colorLink                = "#bf616a"
		colorLinkBackgroundHover = "#b3bbcc"
	)
	style := `
		* {
			font-family: 'JetBrains Mono', monospace;
			box-sizing: border-box;
			margin: 0;
			padding: 0;
		}
		html, body {
			min-height: 100vh;
		}
		body {
			background: ` + colorBackground + `;
			display: flex;
			justify-content: center;
			align-items: center;
		}
		p {
			color: ` + colorText + `;
			margin: 1.5rem 0;
		}
		a {
			color: ` + colorLink + `;
			text-decoration: none;
		}
		a:hover {
			text-decoration: underline;
			background: ` + colorLinkBackgroundHover + `;
		}
		div {
			text-align: justify;
			word-wrap: break-word;
			width: 24rem;
		}
		div#srsly {
			text-align: center;
		}
	`
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset='utf-8'>
			<meta name='viewport' content='width=device-width,initial-scale=1'>
			<style>
				` + style + `
			</style>
			<title>nakas.dev</title>
		</head>

		<body>
			<div>
				<p>
 	 				i'm a data engineer at
					<a href="https://labanoras.io" target="_blank">labanoras tech</a>
					where we are harnessing modern data
					solutions to accelerate business growth
				</p>
				<p>
					currently working on a fully-managed
					data-as-a-service platform
					<a href="https://tables.dev" target="_blank">external tables</a>
					delivering ready-to-use datasets
				</p>
				<p>
					any interactions preferable at
					<a href="https://linkedin.com/in/žygimantas-nakas-a55baa1a9" target="_blank">linkedin</a>
					or <a href="https://t.me/manescianera" target="_blank">telegram</a>,
					or of course you just can go all sneaky - and spy
					on me on <a href="https://github.com/manescianera" target="_blank">github</a>
				</p>
				<p>
					i care about <a href="https://go.dev/" target="_blank">go</a>,
					<a href="https://www.snowflake.com/en/" target="_blank">snowflake</a>,
					<a href="https://aws.amazon.com/">aws</a>, <a href="https://www.getdbt.com/">dbt</a>,
					<a href="/sql">sql</a>, <a href="https://www.docker.com/">docker</a>
					and some other stuff i've probably forgot about while writing this
				</p>
			</div>
		</body>
		</html>
	`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html)
	})

	http.HandleFunc("/sql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset='utf-8'>
					<meta name='viewport' content='width=device-width,initial-scale=1'>
					<style>
						`+style+`
					</style>
					<title>nakas.dev</title>
				</head>
				<div id="srsly"><p>seriously?</p></div>
			</html>
		`)
	})

	if crt != "" && key != "" {
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%v", port), crt, key, nil))
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	}
}
