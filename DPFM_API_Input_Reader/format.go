package dpfm_api_input_reader

import (
	"data-platform-api-incoterms-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToIncoterms() *requests.Incoterms {
	data := sdc.Incoterms
	return &requests.Incoterms{
		Incoterms: data.Incoterms,
	}
}

func (sdc *SDC) ConvertToIncotermsText() *requests.IncotermsText {
	dataIncoterms := sdc.Incoterms
	data := sdc.Incoterms.IncotermsText
	return &requests.IncotermsText{
		Incoterms:     dataIncoterms.Incoterms,
		Language:      data.Language,
		IncotermsName: data.IncotermsName,
	}
}
