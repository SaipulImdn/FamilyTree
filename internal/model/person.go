package model

type Person struct {
    ID       string
    Name     string
    ParentID string
}

func NewPerson(id, name, parentID string) *Person {
    return &Person{
        ID:       id,
        Name:     name,
        ParentID: parentID,
    }
}
