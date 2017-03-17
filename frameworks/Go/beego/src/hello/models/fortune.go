package models

import (
	"log"
	"sort"

	"github.com/astaxie/beego/orm"
)

const FortuneSelect = "SELECT id,message FROM Fortune"

type Fortune struct {
	Id      uint16 `json:"id" orm:"pk"`
	Message string `json:"message"`
}

func GetFortunes() []*Fortune {
	o := orm.NewOrm()
	fortunes := make([]*Fortune, 0, 16)
	_, err := o.Raw(FortuneSelect).QueryRows(&fortunes)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	fortunes = append(fortunes, &Fortune{Message: "Additional fortune added at request time."})

	sort.Sort(ByMessage(fortunes))
	return fortunes
}

type ByMessage []*Fortune

func (s ByMessage) Len() int           { return len(s) }
func (s ByMessage) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByMessage) Less(i, j int) bool { return s[i].Message < s[j].Message }
