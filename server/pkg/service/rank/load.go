package rank

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/urlvet/urlvet/pkg/constants"
)

var domainRankMap map[string]int

// topDomains holds the most popular domains in rank order (see TopDomains).
var topDomains []string

// topKept is how many of the most popular domains are kept in order.
const topKept = 10000

func findRankFile() string {
	candidates := []string{
		constants.DOMAIN_RANK_FILE_PATH,
		"assets/top-1m.csv",
		"./assets/top-1m.csv",
		"../assets/top-1m.csv",
		"server/assets/top-1m.csv",
	}
	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return constants.DOMAIN_RANK_FILE_PATH
}

func LoadDomainRanks() error {
	filePath := findRankFile()
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}

	domainRankMap = make(map[string]int, len(records)-1)
	top := make([]string, topKept+1)

	for _, record := range records {
		if len(record) < 2 {
			continue
		}
		rankStr := record[0]
		domain := record[1]

		rank, err := strconv.Atoi(rankStr)
		if err != nil {
			continue // skip malformed rank
		}

		domainRankMap[domain] = rank
		if rank >= 1 && rank <= topKept {
			top[rank] = domain
		}
	}
	topDomains = topDomains[:0]
	for _, d := range top[1:] {
		if d != "" {
			topDomains = append(topDomains, d)
		}
	}

	return nil
}
