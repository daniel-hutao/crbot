package ferry

type Data struct {
	Title     string
	URL       string
	Number    int
	CreatedAt string
}

type Datas []Data

type Message struct {
	Now   string
	Datas []Data
}

type MessageChan chan Message

var GlobalMessageChan = make(MessageChan, 10)
