package main

import (
	"fmt"
	"time"
	"unsafe"
)

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
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	fmt.Println("=== Check Track type =======================================")
	aTrack := Track{}
	pTrack := &aTrack
	fmt.Println("sizeof aTrack=", unsafe.Sizeof(aTrack))
	fmt.Println("sizeof aTrack.Title=", unsafe.Sizeof(aTrack.Title))
	fmt.Println("sizeof aTrack.Artist=", unsafe.Sizeof(aTrack.Artist))
	fmt.Println("sizeof aTrack.Album=", unsafe.Sizeof(aTrack.Album))
	fmt.Println("sizeof aTrack.Year=", unsafe.Sizeof(aTrack.Year))
	fmt.Println("sizeof aTrack.Duration=", unsafe.Sizeof(aTrack.Length))
	fmt.Println("sizeof pTrack=", unsafe.Sizeof(pTrack))

	fmt.Println("=== Check tracks value =====================================")
	fmt.Printf("tracks= %#v\n", tracks)
	fmt.Printf("tracks[0]= %#v\n", tracks[0])
	fmt.Println("sizeof tracks[0]=", unsafe.Sizeof(tracks[0]))
	fmt.Println("len(tracks)=", len(tracks))
	fmt.Println("sizeof tracks=", unsafe.Sizeof(tracks))
}
