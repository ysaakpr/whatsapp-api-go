package whatsapp

import (
	"encoding/json"
	"fmt"
)

func (api *API) NewText(body string, previewUrl bool) *Text {
	return &Text{api: api, Type: "text", Body: body, PreviewUrl: previewUrl}
}

func (api *API) NewMediaId(file, _type string) *Media {
	return &Media{api: api, File: file, IsItAnID: true, Type: _type}
}

func (api *API) NewMediaLink(file, _type string) *Media {
	return &Media{api: api, File: file, IsItAnID: false, Type: _type}
}

func (api *API) NewLocation(lng, ltd, name, address string) *Location {
	return &Location{api: api, Type: "location", Longitude: lng, Latitude: ltd, Name: name, Address: address}
}

func (api *API) NewContacts(c []Contacts) *ContactsReq {
	return &ContactsReq{api: api, Type: "contacts", Contacts: c}
}

func (api *API) NewInteractive() *Interactive {
	return &Interactive{api: api, Type: "interactive"}
}
func (api *API) NewInteractiveBtnReq() *InteractiveBtnReq {
	return &InteractiveBtnReq{api: api, Type: "interactive"}
}

func (api *API) NewTextBasedTemplate() *TextBasedTemplate {
	return &TextBasedTemplate{api: api, Type: "template"}
}

func (api *API) NewMultiBasedTemplate() *MultiBasedTemplate {
	return &MultiBasedTemplate{api: api, Type: "template"}
}

func (api *API) send(phoneId, to, _type string, obj interface{}) (*MessageResponse, error) {
	endpoint := fmt.Sprintf("/%s/messages", phoneId)

	body := map[string]interface{}{}
	body[_type] = obj
	body["type"] = _type
	body["messaging_product"] = "whatsapp"
	body["recipient_type"] = "individual"
	body["to"] = to

	res, status, err := api.request(endpoint, "POST", nil, body)
	if err != nil {
		return nil, err
	}

	if status >= 400 {
		e := ErrorResponse{}
		json.Unmarshal(res, &e)
		return nil, &e
	}

	r := MessageResponse{}
	json.Unmarshal(res, &r)
	
	return &r, nil
}

func (obj *Media) Send(phoneId, to string) (*MessageResponse, error) {

	var rObj interface{}
	if obj.IsItAnID {
		rObj = obj.ToId()
	} else {
		rObj = obj.ToLink()
	}

	r, err := obj.api.send(phoneId, to, obj.Type, rObj)
	if err != nil {
		return nil, err
	}

	return r, err
}

func (obj *Text) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "text", obj)
	if err != nil {
		return nil, err
	}
	return r, err
}

func (obj *Location) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "location", obj)
	if err != nil {
		return nil, err
	}
	return r, err
}

func (obj *ContactsReq) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "contacts", obj)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (obj *Interactive) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "interactive", obj)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (obj *InteractiveBtnReq) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "interactive", obj)
	if err != nil {
		return nil, err
	}
	return r, err
}

func (obj *TextBasedTemplate) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "template", obj)
	if err != nil {
		return nil, err
	}
	return r, err
}

func (obj *MultiBasedTemplate) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "template", obj)
	if err != nil {
		return nil, err
	}
	return r, err
}

type MessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaId  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		Id string `json:"id"`
	} `json:"messages"`
}
