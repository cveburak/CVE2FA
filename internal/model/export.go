package model

type ExportData struct {
	Version  int             `json:"version"`
	Accounts []ExportAccount `json:"accounts"`
}

type ExportAccount struct {
	Issuer       string `json:"issuer"`
	AccountName  string `json:"accountName"`
	SecretBase32 string `json:"secretBase32"`
}

type ExportRequest struct {
	MasterPassword string `json:"masterPassword"`
}

type ImportRequest struct {
	MasterPassword string `json:"masterPassword"`
}

type ImportPayload struct {
	EncryptedBlob  string `json:"encryptedBlob"`
	MasterPassword string `json:"masterPassword"`
}
