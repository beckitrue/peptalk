package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func generate() (string, error) {

	var a1 = []string{
		"Champ,", "Fact:", "Everybody says", "Dang...", "Check it:",
		"Just saying...", "Superstar,", "Tiger,", "Self,", "Know this:",
		"News alert:", "Girl,", "Ace,", "Excuse me but", "Experts agree:",
		"In my opinion,", "Hear ye, hear ye:", "Ok, listen up:"}

	var a2 = []string{
		"the mere idea of you", "your soul", "your hair today",
		"everything you do", "your personal style", "every thought you have",
		"that sparkle in your eye", "your presense here", "what you got going on",
		"the essential you", "your life's journey", "that saucy personality",
		"your DNA", "that brain of yours", "your choice of attire", "the way your roll",
		"whatever your secret is", "all of y'all"}
	var a3 = []string{
		"has serious game,", "rains magic,", "deserves the Nobel Prize,", "raises the roof,",
		"breeds miracles,", "is paying off big time,", "shows mad skills,", "just shimmers,",
		"is a national treasure,", "gets the party hopping,", "is the next big thing,",
		"roars like a lion", "is a rainbow factory,", "is made of diamonds,", "makes birds sing,",
		"should be taught in school,", "makes my world go 'round,", "is 100% legit,"}
	var a4 = []string{
		"24/7.", "can I get an amen?", "and that's a fact.", "so treat yourself.", "you feel me?",
		"that's just science.", "would I lie?", "for reals.", "mic drop.", "you hidden gem.",
		"snuggle bear.", "period.", "can I get an amen?", "now let's dance.", "high five.",
		"say it again!", "according to CNN.", "so get used to it."}

	return fmt.Sprintf(concat(randomPick(a1), randomPick(a2), randomPick(a3), randomPick(a4))), nil

}

func randomPick(arr []string) string {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pick := r1.Intn(len(arr))
	return arr[pick]

}

func concat(s1, s2, s3, s4 string) string {
	return s1 + " " + s2 + " " + s3 + " " + s4
}

func main() {
	lambda.Start(generate)
}
