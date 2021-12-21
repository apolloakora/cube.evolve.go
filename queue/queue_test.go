package queue

import (
	"fmt"
	"testing"
)

type folder struct {
	name    string
	folders []*folder
	files   []string
}

const target = "Lake Albert"

func (f folder) Search() ([]*Item, bool, *Item) {
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

func (f folder) String() string {
	return f.name
}

var uganda = folder{
	name:  "uganda",
	files: []string{"Lake Albert", "Lake Kyoga"},
}

var kenya = folder{
	name:  "kenya",
	files: []string{"Lake Naivasha", "Lake Nakuru", "Lake Turkana"},
}

var tanzania = folder{
	name:  "tanzania",
	files: []string{"Lake Rukwa"},
}

var malawi = folder{
	name:  "malawi",
	files: []string{"Lake Malawi"},
}

var eastAfrica = folder{
	name:    "east africa",
	folders: []*folder{&uganda, &kenya, &tanzania},
	files:   []string{"Lake Victoria", "Lake Tanganyika"},
}

var africa = folder{
	name:    "africa",
	folders: []*folder{&eastAfrica, &malawi},
	files:   []string{"Lake Edward", "Lake Chad", "Superior"},
}

var canada = folder{
	name:  "canada",
	files: []string{"Winnipeg", "Reindeer Lake", "Great Bear Lake"},
}

var america = folder{
	name:  "america",
	files: []string{"Michigan", "Huron"},
}

var northAmerica = folder{
	name:    "north america",
	folders: []*folder{&canada, &america},
	files:   []string{"Ontario", "Erie", "Lake of the Woods", "Superior"},
}

var asia = folder{
	name:  "asia",
	files: []string{"Caspian Sea"},
}

var worldLakes = folder{
	name:    "world lakes",
	folders: []*folder{&northAmerica, &africa, &asia},
}

func TestNewQueue(t *testing.T) {
	var queueItem Item = worldLakes
	queue := NewQueue(&queueItem)
	foundItemPtr := queue.Seek()
	fmt.Println("The queue seek operation returned", (*foundItemPtr).String())
}
