/** file: main2.go */

package main

import (
	"encoding/json"
	"fmt"
	"time"
	myStruct "github.com/mezcel/struct-fmt"
	myfunc "github.com/mezcel/struct-fmt"
)

func main() {

	// Open a jsonFile
	var jsonPath string = "json/rosaryJSON-nab.json"
	var byteValue []byte = myfunc.ReturnByteValue(jsonPath)

	// Make a struct DB from a json file

	var rosaryBeads myStruct.RosaryBeads
	var beads myStruct.Beads
	var decades myStruct.Decades
	var mysterys myStruct.Mysterys
	var books myStruct.Books
	var scriptures myStruct.Scriptures
	var messages myStruct.Messages
	var prayers myStruct.Prayers

	json.Unmarshal(byteValue, &rosaryBeads)
	json.Unmarshal(byteValue, &beads)
	json.Unmarshal(byteValue, &decades)
	json.Unmarshal(byteValue, &mysterys)
	json.Unmarshal(byteValue, &books)
	json.Unmarshal(byteValue, &scriptures)
	json.Unmarshal(byteValue, &messages)
	json.Unmarshal(byteValue, &prayers)

	// Main Loop

	var weekdayNo int = int(time.Now().Weekday())
	var accumulator int = myfunc.ReturnStartPosition(weekdayNo)

	for accumulator < 315 {
		myfunc.Cls() // clear terminal screen

		// position progress step increment
		accumulator++

		var decadeIdx int = rosaryBeads.RosaryBeads[accumulator].DecadeIndex
		var mysteryIdx int = rosaryBeads.RosaryBeads[accumulator].MysteryIndex
		var prayerIdx int = rosaryBeads.RosaryBeads[accumulator].PrayerIndex
		var scriptureIdx int = rosaryBeads.RosaryBeads[accumulator].ScriptureIndex
		var messageIdx int = rosaryBeads.RosaryBeads[accumulator].MessageIndex

		var decadeName string = decades.Decades[decadeIdx].DecadeName
		var mysteryName string = mysterys.Mysterys[mysteryIdx].MysteryName
		var scriptureText string = scriptures.Scriptures[scriptureIdx].ScriptureText
		var messageText string = messages.Messages[messageIdx].MesageText
		var prayerText string = prayers.Prayers[prayerIdx].PrayerText

		fmt.Println("Decade:\t\t" + decadeName)
		fmt.Println("Mystery:\t" + mysteryName)
		fmt.Println("Message:\t" + messageText)
		fmt.Println("Scripture:\t" + scriptureText + "\n")
		fmt.Println("Prayer:\t\t" + prayerText)

		fmt.Println("\n---\nPress the Enter to continue / Ctrl+C to Exit")
		fmt.Scanln()
	}
}