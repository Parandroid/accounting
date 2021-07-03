package journals

var idCounter int

type Journaler interface {
	post()
}

type JournalEntry struct {
	id int
}

func (j *JournalEntry) SetId() {
	j.id = idCounter
	idCounter++
}

func (j *JournalEntry) Id() int {
	return j.id
}
