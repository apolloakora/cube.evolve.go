package tree

import (
	"fmt"
	"testing"
)

type MyFolder struct {
	name    string
	folders []*MyFolder
	files   []string
}

const target = "Lake Albert"

// @todo reduce the visibility of this function
func (f MyFolder) Search() ([]*Item, bool, *Item) {
	for _, lake := range f.files {
		fmt.Println("Searching to see if lake", lake, "matches the target", target)
		if lake == target {
			var item Item = f
			return nil, true, &item
		}
	}
	var itemPointers []*Item
	for _, folderPtr := range f.folders {
		var item Item = folderPtr
		itemPointers = append(itemPointers, &item)
	}
	return itemPointers, false, nil
}

func (f MyFolder) String() string {
	return f.name
}

var uganda = MyFolder{
	name:  "uganda",
	files: []string{"Lake Albert", "Lake Kyoga"},
}

var kenya = MyFolder{
	name:  "kenya",
	files: []string{"Lake Naivasha", "Lake Nakuru", "Lake Turkana"},
}

var tanzania = MyFolder{
	name:  "tanzania",
	files: []string{"Lake Rukwa"},
}

var malawi = MyFolder{
	name:  "malawi",
	files: []string{"Lake Malawi"},
}

var eastAfrica = MyFolder{
	name:    "east africa",
	folders: []*MyFolder{&uganda, &kenya, &tanzania},
	files:   []string{"Lake Victoria", "Lake Tanganyika"},
}

var africa = MyFolder{
	name:    "africa",
	folders: []*MyFolder{&eastAfrica, &malawi},
	files:   []string{"Lake Edward", "Lake Chad", "Superior"},
}

var canada = MyFolder{
	name:  "canada",
	files: []string{"Winnipeg", "Reindeer Lake", "Great Bear Lake"},
}

var america = MyFolder{
	name:  "america",
	files: []string{"Michigan", "Huron"},
}

var northAmerica = MyFolder{
	name:    "north america",
	folders: []*MyFolder{&canada, &america},
	files:   []string{"Ontario", "Erie", "Lake of the Woods", "Superior"},
}

var asia = MyFolder{
	name:  "asia",
	files: []string{"Caspian Sea"},
}

var worldLakes = MyFolder{
	name:    "world lakes",
	folders: []*MyFolder{&northAmerica, &africa, &asia},
}

func TestNewQueue(t *testing.T) {
	var queueItem Item = worldLakes
	queue := NewQueue(&queueItem)
	foundItemPtr := queue.Seek()
	fmt.Println("The queue seek operation returned", (*foundItemPtr).String())
}
