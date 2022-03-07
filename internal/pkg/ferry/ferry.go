package ferry

type Reviewer string

type Data struct {
	Title              string
	User               string
	URL                string
	Number             int
	CreatedAt          string
	RequestedReviewers []Reviewer
}

type Datas []Data

type Message struct {
	Now   string
	Datas []Data
}

type MessageChan chan Message

var GlobalMessageChan = make(MessageChan, 10)
