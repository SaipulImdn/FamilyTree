package main

import (
	"FamilyTree/internal/model"
	"FamilyTree/internal/service"
	"FamilyTree/internal/traversal"
	"fmt"
)

func main() {
    familyTree := service.NewFamilyTree()

    familyTree.AddPerson(model.NewPerson("1", "Root", ""))
    familyTree.AddPerson(model.NewPerson("2", "Alice", "1"))
    familyTree.AddPerson(model.NewPerson("3", "Bob", "1"))
    familyTree.AddPerson(model.NewPerson("4", "Charlie", "2"))

    fmt.Println("Family Tree (DFS):")
    traversal.DFS(familyTree.People()["1"], familyTree.People(), 0, true)

    fmt.Println("\nFamily Tree (BFS):")
    traversal.BFS(familyTree.People()["1"], familyTree.People())
}
