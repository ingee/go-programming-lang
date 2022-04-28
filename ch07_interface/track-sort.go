package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
	"unsafe"
)

// 1: define Track {{{
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2021, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
} // 1: *Track 아이템 초기화 leteral이 낯설다

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
} // }}} 1

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "=====", "======", "=====", "====", "======")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// 4: define byArtist type to sort tracks {{{
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] } // }}} 4

func main() {
	// 2: check size of Track type {{{
	fmt.Println("\n=== Check Track type =======================================")
	aTrack := Track{}
	pTrack := &aTrack
	fmt.Println("sizeof aTrack=", unsafe.Sizeof(aTrack))
	fmt.Println("sizeof aTrack.Title=", unsafe.Sizeof(aTrack.Title))
	fmt.Println("sizeof aTrack.Artist=", unsafe.Sizeof(aTrack.Artist))
	fmt.Println("sizeof aTrack.Album=", unsafe.Sizeof(aTrack.Album))
	fmt.Println("sizeof aTrack.Year=", unsafe.Sizeof(aTrack.Year))
	fmt.Println("sizeof aTrack.Duration=", unsafe.Sizeof(aTrack.Length))
	fmt.Println("sizeof pTrack=", unsafe.Sizeof(pTrack))

	fmt.Println("\n=== Check tracks value =====================================")
	fmt.Printf("tracks= %#v\n", tracks)
	fmt.Printf("tracks[0]= %#v\n", tracks[0])
	fmt.Println("sizeof tracks[0]=", unsafe.Sizeof(tracks[0]))
	fmt.Println("len(tracks)=", len(tracks))
	fmt.Println("sizeof tracks=", unsafe.Sizeof(tracks))
	// }}} 2

	// 3: print tracks {{{
	fmt.Println("\n=== Print tracks ===========================================")
	printTracks(tracks)
	// }}}

	// 4:  sort tracks by Artist {{{
	fmt.Println("\n=== Sort tracks by Artist and print tracks =================")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	// }}}

	// 5:  sort tracks by Artist reversely
	fmt.Println("\n=== Sort tracks by Artist reversely ========================")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	// 5: Reverse()는 reverse-sort를 수행하지 않고
	// 		Less()가 뒤집힌 reverse data type을 리턴
	//
	printTracks(tracks)
	// }}} 5
}
