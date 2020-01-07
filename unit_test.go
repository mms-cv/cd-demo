package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestUnit(t *testing.T) {
	cases := []struct{ question, reply string }{
		{"ping", "pong"},
		{"hi", "holla"}, 
	}

	for _, x := range cases {
		r := httptest.NewRequest("GET", "http://dummy/"+x.question, nil)
		w := httptest.NewRecorder()
		index(w, r)

		if body, err := ioutil.ReadAll(w.Body); err != nil {
			t.Error(err)
		} else if string(body) != x.reply {
			t.Error("oops we have a problem: expected reply - ", x.reply, ", but got - ", string(body))
		}
	}
}
func TestSum(t *testing.T) {
	tables := []struct {
		x int 
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
