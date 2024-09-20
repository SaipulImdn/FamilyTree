package traversal

import (
	"FamilyTree/internal/model"
	"fmt"
)

func BFS(root *model.Person, people map[string]*model.Person) {
    if root == nil {
        return
    }

    queue := []*model.Person{root}
    fmt.Println(root.Name)

    for len(queue) > 0 {
        person := queue[0]
        queue = queue[1:]

        for _, child := range people {
            if child.ParentID == person.ID {
                fmt.Printf(" └── %s\n", child.Name)
                queue = append(queue, child)
            }
        }
    }
}

func DFS(root *model.Person, people map[string]*model.Person, level int, isLast bool) {
    if root == nil {
        return
    }

    indent := ""
    if level > 0 {
        for i := 0; i < level-1; i++ {
            indent += "│   "
        }
        if isLast {
            indent += "└── "
        } else {
            indent += "├── "
        }
    }
    fmt.Printf("%s%s\n", indent, root.Name)

    children := []*model.Person{}
    for _, child := range people {
        if child.ParentID == root.ID {
            children = append(children, child)
        }
    }

    for i, child := range children {
        DFS(child, people, level+1, i == len(children)-1)
    }
}
