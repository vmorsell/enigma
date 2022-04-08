package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

var messages = []string{
	"Yes",
	"No",
	"Food please",
	"Wood please.",
	"Gold please.",
	"Stone please.",
	"Ahh!",
	"All hail, king of the losers!",
	"Ooh!",
	"I'll beat you back to Age of Empires.",
	"Ah! being rushed.",
	"Sure, blame it on your ISP.",
	"Start the game already!",
	"Don't point that thing at me!",
	"Enemy sighted!",
	"It is good to be the king.",
	"Monk! I need a monk!",
	"Long time, no siege.",
	"My granny could scrap better than that.",
	"Nice town, I'll take it.",
	"Quit touching me!",
	"Raiding party!	",
	"Dadgum.",
	"Eh, smite me.",
	"The wonder, the wonder, the... no!",
	"You played two hours to die like this.",
	"Yeah, well, you should see the other guy.",
	"Roggan.",
	"Wololo.",
	"Attack an enemy now.",
	"Cease creating extra villagers.",
	"Create extra villagers.",
	"Build a navy.",
	"Stop building a navy.",
	"Wait for my signal to attack.",
	"Build a wonder.",
	"Give me your extra resources",
	"What age are you in",
}

// randomMessage returns a random message from the messages list.
func randomMessage() (string, error) {
	msg := messages[rand.Intn(len(messages)-1)]
	return format(msg)
}

// format prepares a string for encryption by removing all characters
// except for A-Z.
func format(s string) (string, error) {
	s = strings.ToUpper(s)

	p, err := regexp.Compile("[^A-Z]+")
	if err != nil {
		return "", fmt.Errorf("compile: %w", err)
	}
	s = p.ReplaceAllString(s, "")

	return s, nil
}
