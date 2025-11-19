package main
import ("fmt"
		"time")


type message interface{
	getMessage() string
}

type BirthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}

func (bm BirthdayMessage) getMessage() string {
		return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type SenderReport struct {
	reportName string
	numberOfSends int
}

func (sr SenderReport) getMessage() string {
		return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func sendMessage(msg message) (string, int) {
	text := msg.getMessage()
	return text, len(text) * 3
}