package folhacerta

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dateRegex = regexp.MustCompile(`/Date\((\d+)\)/`)

func parseDate(s string) time.Time {
	subs := dateRegex.FindStringSubmatch(s)
	if len(subs) > 2 {
		fmt.Fprintf(os.Stderr, "failed to parse date %s\n", s)
		return time.Time{}
	}
	n, err := strconv.Atoi(subs[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse date %s\n", s)
		return time.Time{}
	}
	return time.UnixMilli(int64(n))
}

func parseDuration(s string) time.Duration {
	segments := strings.SplitN(s, ":", 2)
	if len(segments) < 2 {
		fmt.Fprintf(os.Stderr, "failed to parse duration %s\n", s)
		return 0
	}

	h, err := strconv.Atoi(segments[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse duration %s\n", s)
		return 0
	}

	m, err := strconv.Atoi(segments[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse duration %s\n", s)
		return 0
	}

	return time.Duration(h)*time.Hour + time.Duration(m)*time.Minute
}
