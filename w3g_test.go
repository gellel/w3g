package w3g_test

import (
	"fmt"
	"testing"

	"github.com/gellel/w3g"
)

func Test(t *testing.T) {
	fmt.Println(w3g.NewAcceptHeader(w3g.AcceptHTTPHeader{MIMESubType: "png", MIMEType: "image", Q: 0.9}, w3g.AcceptHTTPHeader{MIMESubType: "*", MIMEType: "image", Q: 0.8}))

	fmt.Println(w3g.NewAcceptCHHeader(w3g.AcceptCHHTTPHeader{Width: true}))
}
