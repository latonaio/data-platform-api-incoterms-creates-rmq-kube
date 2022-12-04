package dpfm_api_output_formatter

type Incoterms struct {
	Incoterms     string        `json:"Incoterms"`
	IncotermsText IncotermsText `json:"IncotermsText"`
}

type IncotermsText struct {
	Incoterms     string `json:"Incoterms"`
	Language      string `json:"Language"`
	IncotermsName string `json:"IncotermsName"`
}
