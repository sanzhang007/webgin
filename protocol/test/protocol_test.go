package test

import (
	"fmt"
	"testing"

	"github.com/sanzhang007/webgin/protocol"
)

func TestSs(t *testing.T) {
	var ss protocol.Ss
	s := `ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTo5ZDNhYmFiZC02MWIwLTRmMjktYmEyNi0zYmI2ZDg2MDQxNzBAc2cwMy5qaWVkaWFuLmN5b3U6NDMwMjM#%F0%9F%87%B8%F0%9F%87%AC%20%E7%8B%AE%E5%9F%8ES03-M`
	ss.Parse(s)
	fmt.Printf("%+v\n", ss)
}
