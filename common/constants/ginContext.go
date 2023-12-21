package constants

const _Prefix = "gocrud_"

var GinContextE = struct {
	UserName      string
	UserObject    string
	PromStartTime string
}{
	UserName:      _Prefix + "userName",
	UserObject:    _Prefix + "userObject",
	PromStartTime: _Prefix + "promStartTime",
}
