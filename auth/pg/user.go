package pg

import (
	"log"
)

type UserRepo struct {
}

func (p *UserRepo) Hi2() {
	log.Println("hi2")
}
