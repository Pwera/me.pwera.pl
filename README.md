# Visit my web page
https://me.pwera.pl/

Tech stack
- [go](https://go.dev)
- [templ](https://github.com/a-h/templ)
- [webslides](https://github.com/webslides/WebSlides)
  


How to run:
1) `go install github.com/a-h/templ/cmd/templ@latest`
2) `templ generate -watch`
3) `go run main.go`
4) `docker build . -t webpagecontainer`
5) `docker run --rm -p3000:3000 webpagecontainer`


