package handlers

import (
	"fmt"
	"net/http"
)

func DataEntryHandler(w http.ResponseWriter, r *http.Request) {

	params, err := ExtractPathParams(r, Params.DATA_ENTRY)

	if err {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	section, id, action := params["section"], params["id"], params["action"]

	fmt.Fprintf(w, "SECTION %s ID %s FILE %s", section, id, action)
}
