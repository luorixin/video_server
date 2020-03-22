package defs

type Err struct {
  Error string `json:"error"`
  ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
  HttpSC int
  Error Err
}

var (
  ErrorRequestBodyParseFailed = ErrResponse{HttpSC:400, Error:Err{"Request body is not correct", "0001"}}
  ErrorNotAuthUser = ErrResponse{HttpSC:401, Error:Err{"User authentication failed.", "0002"}}
  ErrorDBUser = ErrResponse{HttpSC:500, Error:Err{"DB opts failed.", "0003"}}
  ErrorMarshal = ErrResponse{HttpSC:500, Error:Err{"marshal failed.", "0004"}}
)
