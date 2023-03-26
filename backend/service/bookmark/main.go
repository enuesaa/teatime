package bookmark
 
import ()

type Bookmark struct {
	Id string
}

type BookmarkService struct {}

func (b *BookmarkService) List() []Bookmark {
	return []Bookmark{ Bookmark { Id: "" } }
}

func (b *BookmarkService) Get(id string) Bookmark {
	return Bookmark { Id: "" }
}


func (b *BookmarkService) Create() string {
	return "" // id
}

func (b *BookmarkService) Update(id string) string {
	return id
}

func (b *BookmarkService) Delete(id string) {}
