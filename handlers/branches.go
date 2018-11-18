package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aditi23/fyle/db"
)

// BranchesHandler return the details of all the branches of particular given bank and the given city
func BranchesHandler(w http.ResponseWriter, r *http.Request) {

	// Input Validation
	bankName, ok := r.URL.Query()["bank_name"]
	if !ok || len(bankName) != 1 || bankName[0] == "" || len(bankName[0]) > 49 {
		http.Error(w, "Invalid Parameters Received", 400)
		return
	}

	city, ok := r.URL.Query()["city"]
	if !ok || len(city) != 1 || city[0] == "" || len(city[0]) > 50 {
		http.Error(w, "Invalid Parameters Received", 400)
		return
	}

	details := banksDetails{}

	// DB call
	err := queryAllBranches(&details, bankName[0], city[0])

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

// queryAllBranches performs the sql operation and return the data or any error if occurred
func queryAllBranches(branchDetails *banksDetails, bankName, city string) error {
	sqlStatement := `select ifsc, bank_id, branch, address, city, district, state, bank_name
	 from bank_branches where bank_name = $1 and city = $2`
	rows, err := db.DB.Query(sqlStatement, bankName, city)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		branch := branchInformation{}
		err = rows.Scan(
			&branch.Ifsc,
			&branch.BankID,
			&branch.Branch,
			&branch.Address,
			&branch.City,
			&branch.District,
			&branch.State,
			&branch.Name,
		)
		if err != nil {
			return err
		}
		branchDetails.Details = append(branchDetails.Details, branch)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
