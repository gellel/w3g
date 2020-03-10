package w3g_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gellel/w3g"
)

func Test(t *testing.T) {

	fmt.Println(w3g.AcceptHeader{MIMESubType: "gif", MIMEType: "image", Q: 0.8}.Value())

	fmt.Println(w3g.AcceptCHHeader{AcceptCH: true, AcceptCHLifetime: true, ContentDPR: true, DeviceMemory: true}.Value())

	fmt.Println(w3g.AcceptCHLifetimeHeader{}.Value())

	fmt.Println(w3g.CacheControlHeader{Public: true}.Value())

	fmt.Println(w3g.ContentDispositionHeader{Attachment: true, FileName: "hello.html"}.Value())

	fmt.Println(w3g.ContentRangeHeader{}.Value())

	fmt.Println(w3g.ContentTypeHeader{Boundary: "something", Charset: "utf-8", MIMESubType: "plain", MIMEType: "text"}.Value())

	fmt.Println(w3g.CookieHeader{Cookies: []*http.Cookie{&http.Cookie{Name: "hello", Value: "world"}}}.Value())

	fmt.Println(w3g.DPRHeader{DPR: 4.223}.Value())
}
