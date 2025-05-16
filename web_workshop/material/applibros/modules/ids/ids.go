package ids

import (
	"github.com/cdvelop/unixid"
)

var ID *unixid.UnixID

func init() {
	// Create a new UnixID handler (server-side)
	idHandler, err := unixid.NewUnixID()
	if err != nil {
		panic(err)
	}
	ID = idHandler
}
