package parse

import (
	"csv2csv/internal/pkg/mapping"
	"errors"
	"flag"
	"regexp"
	"strconv"
	"strings"
)

const (
	titleColFlag       = "title"
	descriptionColFlag = "description"
)

type EventParseArgs struct {
	EventCols map[mapping.EventField]string
	RowRange  *mapping.Range
}

func ArgsFromCmdLine() (*EventParseArgs, error) {
	rowRangeArg := flag.String("range", "", "The row range for the event fields")
	titleColArg := flag.String(titleColFlag, "", "Column for "+titleColFlag)
	flag.Parse()

	args := &EventParseArgs{}

	rowRange, err := parseRange(*rowRangeArg)
	if err != nil {
		return nil, err
	}
	args.RowRange = rowRange

	args.EventCols = map[mapping.EventField]string{}
	if strings.TrimSpace(*titleColArg) == "" {
		return nil, errors.New("missing required column for " + titleColFlag)
	}
	args.EventCols[mapping.Title] = *titleColArg

	return args, nil
}

func parseRange(rangeArg string) (*mapping.Range, error) {
	if strings.TrimSpace(rangeArg) == "" {
		return nil, errors.New("missing required flag range")
	}
	r := regexp.MustCompile(`(?P<StartRow>\d+):(?P<EndRow>\d+)`)
	matches := r.FindStringSubmatch(rangeArg)
	if len(matches) == 0 {
		return nil, errors.New("invalid format for flag range")
	}
	startRow, _ := strconv.Atoi(matches[r.SubexpIndex("StartRow")])
	endRow, _ := strconv.Atoi(matches[r.SubexpIndex("EndRow")])

	argRange := &mapping.Range{
		StartRow: startRow,
		EndRow:   endRow,
	}
	return argRange, nil
}