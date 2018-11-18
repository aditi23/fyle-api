package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aditi23/fyle/db"
)

// BranchDetailsHandler return the details of the bank branch with the given ifsc_code as input
func BranchDetailsHandler(w http.ResponseWriter, r *http.Request) {

	// Input Validation
	ifscCode, ok := r.URL.Query()["ifsc_code"]
	if !ok || len(ifscCode) != 1 || ifscCode[0] == "" || len(ifscCode[0]) > 11 {
		http.Error(w, "Invalid Parameters Received", 400)
		return
	}

	details := branchInformation{}

	// DB call
	err := queryBranch(&details, ifscCode[0])

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(details)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(out))
}

// queryBranch performs the sql operation and return the data or any error if occurred
func queryBranch(branchDetails *branchInformation, ifscCode string) error {
	sqlStatement := `select ifsc, bank_id, branch, address, city, district, state, bank_name
	 from bank_branches where ifsc = $1`
	row := db.DB.QueryRow(sqlStatement, ifscCode)
	err := row.Scan(
		&branchDetails.Ifsc,
		&branchDetails.BankID,
		&branchDetails.Branch,
		&branchDetails.Address,
		&branchDetails.City,
		&branchDetails.District,
		&branchDetails.State,
		&branchDetails.Name,
	)
	if err != nil {
		return err
	}
	return nil
}
