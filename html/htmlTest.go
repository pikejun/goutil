package main

import (
	"html"
	"log"
)

func main(){
	s := html.EscapeString("<div>name</div>") // 逸码
	log.Println(s)
	s = html.UnescapeString(s) // 解码
	log.Println(s)

	log.Println(html.EscapeString(html.UnescapeString("<div>name</div>")))      // 先逸码后解码
	log.Println(html.UnescapeString("&aacute;"), html.UnescapeString("&#225;"))
}
