package models

type Device struct {
	SN       string            `json:"sn"`
	Desc     string            `json:"desc"`
	Protocol map[string]string `json:"protocol"`
}
