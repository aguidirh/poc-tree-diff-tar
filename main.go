package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	repos := []string{
		"first-mirroring-albo-only",
		"second-mirroring-albo-and-noo",
	}

	for _, repo := range repos {
		argument := "./" + repo + "/docker/registry/v2/repositories"
		cmd := exec.Command("tree", argument)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("tree of "+repo+" %s\n", output)
		err = os.WriteFile("."+repo+"-tree", output, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	argument1 := "." + repos[0] + "-tree"
	argument2 := "." + repos[1] + "-tree"
	cmd := exec.Command("diff", argument1, argument2)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 1 {
				log.Printf("diff between trees %s\n", output)
				err = os.WriteFile(".diff-tree", output, 0644)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

}
