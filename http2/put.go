package http2

import (
	"encoding/json"
	"net/http"
	"wmaster/controller"
)

func (h Http) PutOneStokMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateStok
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateStok(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneNameMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateName
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateName(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneModelMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateModel
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateModel(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneTipeMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateTipe
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateTipe(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}

func (h Http) PutOneSubMaster(res http.ResponseWriter, req *http.Request) {
	SetHeader(res)
	col, con := d.MakeConnection()
	var tmp controller.FormatUpdateSub
	json.NewDecoder(req.Body).Decode(&tmp)
	prim := c.FormatUpdateSub(tmp.Set.Data)
	err := d.UpdOne(col, tmp.Find, prim)
	if err != nil {
		bol.Res = false
	} else {
		bol.Res = true
	}
	d.Dis(con)
	json.NewEncoder(res).Encode(bol)
}
