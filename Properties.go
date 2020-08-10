package properties

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var propertie *Propertie
var path string

type Propertie struct {
	Pairs []PropertiePair
}

type PropertiePair struct {
	Key   string
	Value string
}

func SetPropertiePath(s string) {
	path = s
}

func (p *Propertie) GetProperty(key string) string {
	for _, s := range p.Pairs {
		if s.Key == key {
			return s.Value
		}
	}
	return ""
}

func GetPropertyFile() *Propertie {
	return propertie
}

func LoadProperty() (*Propertie, error) {
	if propertie == nil {
		var p Propertie
		file, err := os.Open(path)
		if err != nil {
			return nil, errors.New("Not able to open Properties")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s := strings.Split(scanner.Text(), ": ")
			p.Pairs = append(p.Pairs, PropertiePair{Key: s[0], Value: s[1]})
		}
		propertie = &p
		if err := scanner.Err(); err != nil {
			return nil, errors.New("Propertie file not well formed")
		}
	}
	return propertie, nil
}
