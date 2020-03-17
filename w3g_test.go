package w3g_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/gellel/w3g"
)

func Test(t *testing.T) {

	/*fmt.Println(w3g.AcceptHeader{MIMESubType: "gif", MIMEType: "image", Q: 0.8}.String())

	fmt.Println(w3g.AcceptCHHeader{AcceptCH: true, AcceptCHLifetime: true, ContentDPR: true, DeviceMemory: true}.String())

	fmt.Println(w3g.AcceptCHLifetimeHeader{}.String())

	fmt.Println(w3g.CacheControlHeader{Public: true}.String())

	fmt.Println(w3g.ContentDispositionHeader{Attachment: true, FileName: "hello.html"}.String())

	fmt.Println(w3g.ContentRangeHeader{}.String())

	fmt.Println(w3g.ContentTypeHeader{Boundary: "something", Charset: "utf-8", MIMESubType: "plain", MIMEType: "text"}.String())

	fmt.Println(w3g.CookieHeader{Cookies: []*http.Cookie{&http.Cookie{Name: "hello", Value: "world"}}}.String())

	fmt.Println(w3g.DPRHeader{DPR: 1}.String())

	fmt.Println(w3g.DeviceMemoryHeader{Memory: 1}.String())

	fmt.Println(w3g.ETagHeader{Value: "1224", W: true}.String())*/

	fmt.Println(w3g.FeaturePolicyHeader{Accelerometer: "*", Camera: "none"}.String())

	fmt.Println(w3g.ForwardedHeader{Identifier: net.ParseIP("2001:db8:cafe::17")}.String())
}
