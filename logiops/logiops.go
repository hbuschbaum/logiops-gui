package logiops

import (
	"os"
	"strings"
)

func deleteWhitespaces(s *string) {
	*s = strings.ReplaceAll(string(*s), " ", "")
	*s = strings.ReplaceAll(string(*s), "\n", "")
	*s = strings.ReplaceAll(string(*s), "\t", "")
}

func toLogiData(s string) LogiData {
	return LogiData{}
}

func parseLogiData(s string) LogiData {
	return LogiData{}
}


func openAndStrip() (string, error) {
	dat, err := os.ReadFile("/etc/logiops.conf")
	if err != nil {
		return "", err
	}
	dat2 := string(dat)
	deleteWhitespaces(&dat2)

	return dat2, nil
}
