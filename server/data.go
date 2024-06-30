package main

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Id               int    `json:"ID"`
	Image            string `json:"Image"`
	ProjectName      string `json:"ProjectName"`
	Description      string `json:"Description"`
	ShortDescription string `json:"Short Description"` // TODO delete in favor of dynamically generated short description on react
}
