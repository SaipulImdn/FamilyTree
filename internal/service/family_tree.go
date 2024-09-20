package service

import (
	"FamilyTree/internal/model"
	"fmt"
)

type FamilyTree struct {
    people map[string]*model.Person
}

func NewFamilyTree() *FamilyTree {
    return &FamilyTree{
        people: make(map[string]*model.Person),
    }
}

func (ft *FamilyTree) AddPerson(p *model.Person) {
    ft.people[p.ID] = p
}

func (ft *FamilyTree) People() map[string]*model.Person {
    return ft.people
}

func (ft *FamilyTree) DisplayFamilyTree(id string, level int) {
    person, exists := ft.people[id]
    if !exists {
        return
    }

    fmt.Printf("%s%s\n", string('\t'*level), person.Name)

    for _, child := range ft.people {
        if child.ParentID == id {
            ft.DisplayFamilyTree(child.ID, level+1)
        }
    }
}
