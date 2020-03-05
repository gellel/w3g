package w3g_test

import (
	"fmt"
	"testing"

	"github.com/gellel/w3g"
)

func Test(t *testing.T) {

	fmt.Println(w3g.CacheControlHeader{MaxAge: -1, MaxStale: -1}.Value())

}
