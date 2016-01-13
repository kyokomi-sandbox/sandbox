package main

import (
	"fmt"
	"strconv"
)

type DataType string

const (
	UserData  DataType = "user"
	EventData DataType = "event"
)

type Stream struct {
	DataType DataType
	Data     interface{}
}

type User struct {
	Name string
}

type Event struct {
	ID int
}

func (s Stream) User() *User {
	if s.DataType != UserData {
		return nil
	}

	return s.Data.(*User)
}

func (s Stream) Event() *Event {
	if s.DataType != EventData {
		return nil
	}

	return s.Data.(*Event)
}

func main() {
	
	case1()
	case2()
}

func case2() {
	uch, ech, quit := GetUserStream2()

	go func () {
		for u := range uch {
			fmt.Println(u)
		}
	}()

	go func () {
		for e := range ech {
			fmt.Println(e)
		}
	}()

	<-quit
}

func GetUserStream2() (chan User, chan Event, chan int) {
	
	uch := make(chan User)
	ech := make(chan Event)
	quit := make(chan int)
	
	go func() {
		for i := 0; i < 5; i++ {
			if i%2 == 0 {
				uch <- User{Name: "test"}
			} else {
				ech <- Event{ID: 1000}
			}
		}

		quit <- 0 // 終了判定用なので値は何でもいい		
	}()

	return uch, ech, quit
}

func case1() {
	
	ch := make(chan Stream, 10)
	GetUserStream1(ch)
	for c := range ch {
		
		switch c.DataType {
			case UserData:
			fmt.Println(c.User())
			case EventData:
			fmt.Println(c.Event())
		}
	}
}

func GetUserStream1(ch chan Stream) {

	for i := 0; i < 5; i++ {
		var s Stream
		if i%2 == 0 {
			s = Stream{
				DataType: UserData,
				Data:     &User{Name: "test:" + strconv.Itoa(i)},
			}
		} else {
			s = Stream{
				DataType: EventData,
				Data:     &Event{ID: i},
			}
		}

		ch <- s
	}

	close(ch)
}
