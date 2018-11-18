package handlers

type banksDetails struct {
	Details []branchInformation `json:"bank_details"`
}

type branchInformation struct {
	Name     string `json:"name,omitempty"`
	Ifsc     string `json:"ifsc,omitempty"`
	BankID   int64  `json:"bank_id,omitempty"`
	Branch   string `json:"branch,omitempty"`
	Address  string `json:"address,omitempty"`
	City     string `json:"city,omitempty"`
	District string `json:"district,omitempty"`
	State    string `json:"state,omitempty"`
}
