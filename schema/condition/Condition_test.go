package condition

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestCondition(t *testing.T) {
	// test all test
	validUsage := []Condition{All, New, Used, Collectible, Refurbished}

	for _, c := range validUsage {
		condition, err := xml.Marshal(c)
		if err != nil {
			t.Error(err)
		}

		err = xml.Unmarshal(condition, &c)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestInvalidConditionValue(t *testing.T) {
	var c Condition = 6
	condition, err := xml.Marshal(c)
	if err != nil {
		t.Error(err)
	}

	err = xml.Unmarshal(condition, &c)
	if err != nil {
		t.Error(err)
	}

	if c.String() == "All" {
		t.Error("Default value is New if condition is out of range")
	}
}

func TestFailedConditionUnmarshal(t *testing.T) {
	condition, err := xml.Marshal(All)
	if err != nil {
		t.Error(err)
	}

	s := strings.Replace(string(condition), "Condition", "Cond", -1)
	s = strings.Replace(string(s), "</Cond", "</Conda", -1)

	all := All
	err = xml.Unmarshal([]byte(s), &all)
	if err == nil {
		t.Error(err)
	}

}
