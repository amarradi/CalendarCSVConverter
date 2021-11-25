package config

import (
	"csv2csv/internal/pkg/core"
	"errors"
	"flag"
	"regexp"
	"strconv"
	"strings"
)

const (
	titleColFlag       = "title"
	descriptionColFlag = "description"
	dateFormatFlag     = "date-format"
	defaultDateFormat  = "27.05.2021"
)

type EventParseConfig struct {
	EventCols map[core.EventField]string
	RowRange  *Range
	InputFile string
}

func FromCmdLine() (*EventParseConfig, error) {
	var dateFormat, titleCol string
	rowRangeArg := flag.String("range", "", "The row range for the event fields")
	flag.StringVar(&titleCol, titleColFlag, "", "Column for "+titleColFlag)
	descriptionColArg := flag.String(descriptionColFlag, "", "Column for "+descriptionColFlag)
	flag.StringVar(&dateFormat, dateFormatFlag, "", "Date format for parsing dates")

	flag.Parse()

	args := &EventParseConfig{}

	rowRange, err := parseRange(*rowRangeArg)
	if err != nil {
		return nil, err
	}
	args.RowRange = rowRange

	args.EventCols = map[core.EventField]string{}

	if strings.TrimSpace(titleCol) == "" {
		return nil, errors.New("missing required column for " + titleColFlag)
	}
	args.EventCols[core.Title] = titleCol
	if strings.TrimSpace(*descriptionColArg) == "" {
		return nil, errors.New("missing required column for " + descriptionColFlag)
	}
	args.EventCols[core.Description] = *descriptionColArg

	if strings.TrimSpace(dateFormat) == "" {
		dateFormat = defaultDateFormat
	}

	if flag.NArg() != 1 {
		return nil, errors.New("need exactly one input")
	}
	args.InputFile = flag.Arg(0)

	return args, nil
}

func parseRange(rangeArg string) (*Range, error) {
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

	argRange := &Range{
		StartRow: startRow,
		EndRow:   endRow,
	}
	return argRange, nil
}
