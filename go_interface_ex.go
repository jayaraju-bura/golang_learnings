package main
import "fmt"

type Book struct {
    author, title string
}

type Magazine struct {
    title string
    issue int
}

func (b *Book) Assign(a, name string) {
    b.author = a
    b.title = name
    
}

func (m *Magazine) Assign(name string, iss int) {
    m.title  = name
    m.issue = iss
}

func (b *Book) print() {
    fmt.Println("contents of the Book instance", b.author, " -> ",  b.title)
}

func (m *Magazine) print() {
    fmt.Println("contents of the Magazine instance ", m.title, " -> ",  m.issue)
    
}

type printer interface {
    print()
}



func main() {
    var b Book
    var m Magazine
    b.Assign("Richard Thaler", "Misbehaving")
    m.Assign("Zero to one", 556)
    
    var ptr printer
    ptr = &b
    ptr.print()
    ptr = &m 
    ptr.print()
    
}

// contents of the Book instance Richard Thaler  ->  Misbehaving
// contents of the Magazine instance  Zero to one  ->  556
