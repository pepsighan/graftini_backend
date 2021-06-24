package templates

import (
	"encoding/json"
	"log"

	"github.com/valyala/quicktemplate"
)

// streamstringify converts the object into a json.
func streamstringify(writer *quicktemplate.Writer, obj interface{}) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Panicf("could not stringify obj: %v", err)
	}

	writer.N().S(string(bytes[:]))
}
