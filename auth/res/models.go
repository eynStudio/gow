package res

import . "github.com/eynstudio/gobreak"

type AuthRes struct {
	Id   GUID
	Mc   string
	Bz   string
	Qz   int
	Icon string
	Ns   string
	Args []KeyValue
	Opts []Opt
}

func NewRes() *AuthRes          { return &AuthRes{Id: Guid()} }
func (p AuthRes) GetMc() string { return p.Mc }
func (p AuthRes) GetNs() string { return p.Ns }
func (p AuthRes) GetQz() int    { return p.Qz }

func (p *AuthRes) ReplaceOpt(opt Entity) { Slice(&p.Opts).ReplaceEntity(opt) }
func (p *AuthRes) DelOpt(id GUID)        { Slice(&p.Opts).RemoveEntity(id) }

type Opt struct {
	Id   GUID
	Mc   string
	Bz   string
	Icon string
	Qz   int
	View bool
	Act  string
	Args []KeyValue
}
