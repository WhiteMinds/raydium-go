package logs

import (
	"fmt"
	"regexp"
	"strconv"
)

type SolanaProgramLog struct {
	Index                int // Current index in array
	CallIndex            int // Index of call in array
	Address              string
	Status               string
	LogMessages          []string
	ConsumedComputeUnits int
	MaxComputeUnits      int
	ProgramData          string
}

func ParseProgramLogs(logs []string) []SolanaProgramLog {
	programs := make([]SolanaProgramLog, 0)

	// Regex patterns
	invokePattern := regexp.MustCompile(`Program (.*?) invoke \[(\d+)]`)
	consumePattern := regexp.MustCompile(`Program .*? consumed (\d+) of (\d+) compute units`)
	successPattern := regexp.MustCompile(`Program (.*?) (success|failed)`)
	dataPattern := regexp.MustCompile(`Program data: (.*)`)

	// Stack stores indices instead of pointers
	stack := 0
	tmpProgramStack := make([]SolanaProgramLog, 0)

	var tmpLog []string
	// left bracket: Program invoke
	// right bracket: Program success|failed

	// other part:
	// Program Data:
	// Program log:
	// Program consumed X of Y compute units

	for idx, line := range logs {
		// for each line, check if it's a left bracket
		if invokePattern.MatchString(line) {
			tmpProgram := SolanaProgramLog{}
			tmpProgram.Address = invokePattern.FindStringSubmatch(line)[1]
			tmpProgram.Index = idx
			tmpProgram.CallIndex, _ = strconv.Atoi(invokePattern.FindStringSubmatch(line)[2])
			tmpProgramStack = append(tmpProgramStack, tmpProgram)
			stack = stack + 1
		} else if successPattern.MatchString(line) {
			if len(tmpProgramStack) == 0 {
				fmt.Printf("Error: unmatched successPattern")
				continue
			}
			tmpProgram := tmpProgramStack[len(tmpProgramStack)-1]
			tmpProgram.Address = successPattern.FindStringSubmatch(line)[1]
			tmpProgram.Status = successPattern.FindStringSubmatch(line)[2]
			// pop from tmpLog
			if len(tmpLog) > 0 {
				tmpLog = tmpLog[:len(tmpLog)-1]
				tmpProgram.LogMessages = tmpLog
				tmpLog = make([]string, 0)
			}
			for _, log := range tmpProgram.LogMessages {
				if consumePattern.MatchString(log) {
					tmpProgram.ConsumedComputeUnits, _ = strconv.Atoi(consumePattern.FindStringSubmatch(log)[1])
					tmpProgram.MaxComputeUnits, _ = strconv.Atoi(consumePattern.FindStringSubmatch(log)[2])
				} else if dataPattern.MatchString(log) {
					tmpProgram.ProgramData = dataPattern.FindStringSubmatch(log)[1]
				}
			}
			programs = append(programs, tmpProgram)
			tmpProgramStack = tmpProgramStack[:len(tmpProgramStack)-1]
		} else {
			tmpLog = append(tmpLog, line)
		}
	}

	return programs
}
