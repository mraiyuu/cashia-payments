package handlers

import (
	"net/http"

	"github.com/mraiyuu/cashia-payments/internal/utils"
)

func AuthenticateMerchant(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSONSuccess(w, "endpoint hit and backend running", http.StatusAccepted)
}
