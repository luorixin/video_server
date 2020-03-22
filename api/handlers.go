package main

import (
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "io"
  "io/ioutil"
  "net/http"
  "video_server/api/dbops"
  "video_server/api/defs"
  "video_server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  res, _ := ioutil.ReadAll(r.Body)
  ubody := &defs.UserCredential{}
  if err := json.Unmarshal(res, ubody); err !=nil {
    sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
    return
  }
  if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err!=nil{
    sendErrorResponse(w, defs.ErrorDBUser)
    return
  }
  id := session.GenerateNewSessionId(ubody.Username)
  su := &defs.SignedUp{
    Success:   true,
    SessionId: id,
  }
  if resp, err := json.Marshal(su); err != nil {
    sendErrorResponse(w, defs.ErrorMarshal)
    return
  }else {
    sendNormalResponse(w, string(resp), 200)
  }
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  uname := p.ByName("user_name")
  io.WriteString(w, uname)
}
