package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		TestName      string
		Input         interface{}
		ExpectedInput []string
	}{
		{
			TestName:      "struct with one string field",
			Input:         struct{ Name string }{"Chris"},
			ExpectedInput: []string{"Chris"},
		},
		{
			TestName:      "struct with two string fields",
			Input:         struct{ Name, City string }{"Chris", "London"},
			ExpectedInput: []string{"Chris", "London"},
		},
		{
			TestName: "struct with varying field types",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedInput: []string{"Chris"},
		},
		{
			TestName: "nested fields",
			Input: Person{
				Name:    "Chris",
				Profile: Profile{33, "London"},
			},
			ExpectedInput: []string{"Chris", "London"},
		},
		{
			TestName: "struct with a pionter",
			Input: &Person{
				Name:    "Chris",
				Profile: Profile{33, "London"},
			},
			ExpectedInput: []string{"Chris", "London"},
		},
		{
			TestName: "struct with slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedInput: []string{"London", "Reykjavik"},
		},
		{
			TestName: "struct with arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedInput: []string{"London", "Reykjavik"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			var got []string
			walk(tc.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tc.ExpectedInput) {
				t.Errorf("got: %v, expected: %v", got, tc.ExpectedInput)
			}
		})
	}

	t.Run("struct with maps", func(t *testing.T) {
		mp := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		want := []string{"Bar", "Boz"}

		walk(mp, func(input string) {
			got = append(got, input)
		})

		for _, key := range want {
			assertContains(t, got, key)
		}
	})

	t.Run("struct with channels", func(t *testing.T) {
		mChan := make(chan Profile)

		go func() {
			defer close(mChan)
			mChan <- Profile{33, "Berlin"}
			mChan <- Profile{34, "Katowice"}
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(mChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, expected: %v", got, want)
		}
	})

	t.Run("struct with function", func(t *testing.T) {
		mFunc := func() (Profile, Profile) {
			return Profile{33, "Berling"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berling", "Katowice"}

		walk(mFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, wanted: %v", got, want)
		}
	})
}

func assertContains(t testing.TB, gots []string, want string) {
	t.Helper()
	var contains bool
	for _, got := range gots {
		if got == want {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", gots, want)
	}
}
