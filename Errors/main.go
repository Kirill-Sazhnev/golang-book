package main

import (
	"fmt"
	"net/url"
	"os"
	//"errors"
	//"strings"
)

func main() {
	//u, err := url.Parse("https://a b.com/")
	u, err := url.Parse("https://4p a.to/")

	if err != nil {
		fmt.Println(err)         // Выводит: parse https://a b.com/: invalid character “ ” in host name
		fmt.Printf("%#v\n", err) // Выводит: &url.Error{Op:“parse”, URL:“https://a b.com/”, Err:” “}

		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)   // Выводит: Op: parse
			fmt.Println("URL:", e.URL) // Выводит: URL: https://a b.com/
			fmt.Println("Err:", e.Err) // Выводит: Err: invalid character “ ” in host name
		}
		os.Exit(1)
	}
	fmt.Printf("This is u: %#v", u) //  &url.URL{Scheme:"https", Opaque:"", User:(*url.Userinfo)(nil), Host:"4pda.to", Path:"/", RawPath:"", OmitHost:false, ForceQuery:false, RawQuery:"", Fragment:"", RawFragment:""}
}

/*
type myURL struct {
	url *url.URL
}

func (u *myURL) Error() string {

	return fmt.Sprintf("%#v\n", u.url)

}

func main() {

	var myURL myURL

	url, err := url.Parse("https://a b.com/")
	if err != nil {
		fmt.Println(err)
	}

	myURL.url = url

	defer fmt.Println(myURL.Error())

}
*/
