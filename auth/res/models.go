package res

import . "github.com/eynstudio/gobreak"

type AuthRes struct {
	Id        GUID
	Name      string
	SortOrder int
	Icon      string
	Ns        string
	Args      Params
	Opts      []Opt
}

func NewRes() *AuthRes              { return &AuthRes{Id: Guid()} }
func (p AuthRes) GetName() string   { return p.Name }
func (p AuthRes) GetNs() string     { return p.Ns }
func (p AuthRes) GetSortOrder() int { return p.SortOrder }
func (p AuthRes) GetId() GUID       { return p.Id }

func (p *AuthRes) ReplaceOpt(opt Entity) { Slice(&p.Opts).ReplaceEntity(opt) }
func (p *AuthRes) DelOpt(id GUID)        { Slice(&p.Opts).RemoveEntity(id) }

type Opt struct {
	Id        GUID
	Name      string
	Icon      string
	SortOrder int
	View      bool
	Act       string
	Args      Params
}
