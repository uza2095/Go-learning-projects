package main
import "fmt"
func validity(score int) bool{
	check:=score>=0 && score<=100gi
	return check
}
func AvgCalculator(grades map[string]int) float64{
	if len(grades)==0{
		return 0
	}
	var total int
	for _,score:=range grades{
		total=total+score
	}
	avg:=float64(total)/float64(len(grades))
	return avg
}
func main(){
	var name string
	fmt.Println("please enter ur name")
	fmt.Scanln(&name)
	var numofsubjs int
	fmt.Println("Please enter number of courses you have taken:")
	fmt.Scanln(&numofsubjs)

	studentInfo:=make(map[string]int)
		for i := 0; i < numofsubjs; i++ {
			var course string
			fmt.Println("Please enter the course name:")
			fmt.Scanln(&course)

			var score int
			fmt.Println("Please enter the score you got for", course)
			fmt.Scanln(&score)

			if !validity(score) {
				fmt.Println("Invalid score! Score must be between 0 and 100.")
				valid := false
				for attempts := 0; attempts < 5; attempts++ {
					fmt.Println("Please enter a valid score between 0 and 100:")
					fmt.Scanln(&score)
					if validity(score) {
						valid = true
						break
					} else {
						fmt.Println("Invalid input. Try again.")
					}
				}
				if !valid {
					fmt.Println("Too many invalid attempts. Skipping this course.")
					continue
				}
			}

			studentInfo[course] = score
	}

	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Grades:")
	for subject, score := range studentInfo {
		fmt.Printf("  %s: %d\n", subject, score)
	}

	avg := AvgCalculator(studentInfo)
	fmt.Printf("Average grade: %.2f\n", avg)
}
	




