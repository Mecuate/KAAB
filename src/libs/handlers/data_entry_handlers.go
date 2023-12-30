package handlers

import (
	"fmt"
	"net/http"
)

func DataEntryHandler(w http.ResponseWriter, r *http.Request) {

	params, err := ExtractPathParams(r, Params.DATA_ENTRY)

	if err != nil {
		FailReq(w, 1)
		return
	}

	section, id, action := params["section"], params["subject_id"], params["action"]

	fmt.Fprintf(w, "SECTION %s ID %s FILE %s", section, id, action)
}
