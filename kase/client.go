package kase

import (
	"fmt"
	"github.com/joncalhoun/desk"
)

type Client struct {
	*desk.Backend
	Subdomain, Username, Password string
}

func getC() Client {
	return Client{desk.GetBackend(), desk.Subdomain, desk.Username, desk.Password}
}

func List(params *desk.CaseListParams) ([]desk.Case, error) {
	return getC().List(params)
}

func (c Client) List(params *desk.CaseListParams) ([]desk.Case, error) {
	rawCases := desk.RawCases{}
	path := "/cases"
	err := c.Call("GET", path, c.Subdomain, c.Username, c.Password, params.UrlValues(), &rawCases)
	if err != nil {
		return nil, err
	}
	return desk.ParseRawCases(&rawCases), nil
}

func Search(params *desk.CaseSearchParams) ([]desk.Case, error) {
	return getC().Search(params)
}

func (c Client) Search(params *desk.CaseSearchParams) ([]desk.Case, error) {
	rawCases := desk.RawCases{}
	path := "/cases/search"
	err := c.Call("GET", path, c.Subdomain, c.Username, c.Password, params.UrlValues(), &rawCases)
	if err != nil {
		return nil, err
	}
	return desk.ParseRawCases(&rawCases), nil
}

func Get(id int) (*desk.Case, error) {
	return getC().Get(id)
}

func (c Client) Get(id int) (*desk.Case, error) {
	kase := desk.Case{}
	path := fmt.Sprintf("/cases/%d", id)
	err := c.Call("GET", path, c.Subdomain, c.Username, c.Password, nil, &kase)
	return &kase, err
}

func GetExternal(externalId string) (*desk.Case, error) {
	return getC().GetExternal(externalId)
}

func (c Client) GetExternal(externalId string) (*desk.Case, error) {
	kase := desk.Case{}
	path := fmt.Sprintf("/cases/%d", externalId)
	err := c.Call("GET", path, c.Subdomain, c.Username, c.Password, nil, &kase)
	return &kase, err
}

func GetMessage(kase *desk.Case) (*desk.Message, error) {
	return getC().GetMessage(kase)
}

func (c Client) GetMessage(kase *desk.Case) (*desk.Message, error) {
	kase.Message = &desk.Message{}
	path := fmt.Sprintf("/cases/%d/message", kase.Id)
	err := c.Call("GET", path, c.Subdomain, c.Username, c.Password, nil, kase.Message)
	return kase.Message, err
}

func GetNotes(kase *desk.Case) ([]desk.Note, error) {
	return getC().GetNotes(kase)
}

func (c Client) GetNotes(kase *desk.Case) ([]desk.Note, error) {
	// TODO(joncalhoun): Make pagination on notes work - currently only the first page is ever grabbed.

	rawNotes := desk.RawNotes{}
	path := fmt.Sprintf("/cases/%d/notes", kase.Id)
	err := c.Call("GET", path, c.Subdomain, c.Username, c.Password, nil, &rawNotes)
	kase.Notes = desk.ParseRawNotes(&rawNotes)
	return kase.Notes, err
}

func GetReplies(kase *desk.Case) ([]desk.Message, error) {
	return getC().GetReplies(kase)
}

func (c Client) GetReplies(kase *desk.Case) ([]desk.Message, error) {
	// TODO(joncalhoun): Make pagination on notes work - currently only the first page is ever grabbed.

	rawReplies := desk.RawReplies{}
	path := fmt.Sprintf("/cases/%d/replies", kase.Id)
	err := c.Call("GET", path, c.Subdomain, c.Username, c.Password, nil, &rawReplies)
	kase.Replies = desk.ParseRawReplies(&rawReplies)
	return kase.Replies, err
}
