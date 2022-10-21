package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	contourX, contourY := GetIntArrays("contour.txt")
	Draw(contourX, contourY)

	leftEarX, leftEarY := GetIntArrays("left_ear.txt")
	Draw(leftEarX, leftEarY)

	leftEyeX, leftEyeY := GetIntArrays("left_eye.txt")
	Draw(leftEyeX, leftEyeY)

	leftPupilX, leftPupilY := GetIntArrays("left_pupil.txt")
	Draw(leftPupilX, leftPupilY)

	rightEarX, rightEarY := GetIntArrays("right_ear.txt")
	Draw(rightEarX, rightEarY)

	rightEyeX, rightEyeY := GetIntArrays("right_eye.txt")
	Draw(rightEyeX, rightEyeY)

	rightPupilX, rightPupilY := GetIntArrays("right_pupil.txt")
	Draw(rightPupilX, rightPupilY)

	snoutTopX, snoutTopY := GetIntArrays("snout_top.txt")
	Draw(snoutTopX, snoutTopY)

	snoutBottomX, snoutBottomY := GetIntArrays("snout_bottom.txt")
	Draw(snoutBottomX, snoutBottomY)

	toothCountourX, toothCountourY := GetIntArrays("tooth_contour.txt")
	Draw(toothCountourX, toothCountourY)

	toothMiddleX, toothMiddleY := GetIntArrays("tooth_middle.txt")
	Draw(toothMiddleX, toothMiddleY)
}

func Draw(x []int, y []int) {

	robotgo.Move(x[0], y[0])

	for i := 0; i < len(x); i++ {
		robotgo.DragSmooth(x[i], y[i])
		fmt.Println(x[i], y[i])
	}

}

func CreateLogBasedOnMousePosition() {
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	for {
		mouseLocationX, mouseLocationY := robotgo.GetMousePos()
		log.Println(mouseLocationX, mouseLocationY)
		log.Println(mouseLocationX, mouseLocationY)
		time.Sleep(1 * time.Millisecond)
	}
}

func GetIntArrays(archive string) ([]int, []int) {

	file, err := os.Open(archive)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lineCompare string
	var finalResult []string

	for scanner.Scan() {

		line := scanner.Text()

		//remove duplicated lines
		if line != lineCompare {
			lineCompare = line
			finalResult = append(finalResult, line)
		} else {
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var return3first []int
	var return3last []int

	for i := 0; i < len(finalResult); i++ {

		listIndex := finalResult[i]

		//-----String tratative
		//Get the last 7 charachters
		last7 := listIndex[len(listIndex)-7:]

		//Get the first 3 characters of the 7 characthers
		first3 := last7[0:3]

		//Get the last 3 first of the 7 characthers
		last3 := last7[len(last7)-3:]
		//-----

		iFirst3, err := strconv.Atoi(first3)
		if err != nil {
			fmt.Println("Conversion error:", err)
		}

		iLast3, err := strconv.Atoi(last3)
		if err != nil {
			fmt.Println("Conversion error:", err)
		}

		return3first = append(return3first, iFirst3)
		return3last = append(return3last, iLast3)
	}
	return return3first, return3last
}
