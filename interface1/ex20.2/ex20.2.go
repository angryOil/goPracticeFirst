package main

import (
	"awesomeProject/interface1/fedex"
	"awesomeProject/interface1/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	korSender := &koreaPost.PostSender{}
	fexSender := &fedex.FedexSender{}
	SendBook("어린왕자", korSender)
	SendBook("꼬북좌 화보집", fexSender)

}
