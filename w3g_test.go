package w3g_test

import (
	"fmt"
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
}
