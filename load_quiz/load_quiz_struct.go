package load_quiz

type Question struct {
	Id         int32           `bson:"id" json:"id"`
	En         QuestionDetails `bson:"en" json:"en"`
	De         QuestionDetails `bson:"de" json:"de"`
	Hu         QuestionDetails `bson:"hu" json:"hu"`
	Difficulty string          `bson:"difficulty" json:"difficulty"`
	Topic      string          `bson:"topic" json:"topic"`
}

type QuestionDetails struct {
	Text         string   `bson:"text" json:"text"`
	Answers      []string `bson:"answers" json:"answers"`
	CorrectOrder string   `bson:"correct_order" json:"correct_order"`
}
