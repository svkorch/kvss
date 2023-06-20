package kvss

import (
	"testing"
)

var testCase1 = map[string]string{
	"a":      "A",
	"b":      "B",
	"aa":     "AAAAAA",
	"bb":     "BBBB",
	"123":    "1234567890",
	"StrStr": "qwerty000",
}

var testCase2 = map[string]string{
	"a":      "mMm",
	"b":      "MmMmM",
	"aa":     "qqq",
	"bb":     "qwerty",
	"123":    "aabbccddEEE",
	"StrStr": "101112",
}

func TestStorage(t *testing.T) {
	s := New()

	{ // Test an empty storage
		for key := range testCase1 {
			gotValue, ok := s.Get(key)

			if ok {
				t.Fatalf("An unset key [%s] exists.", key)
			}

			if gotValue != "" {
				t.Fatalf("A value of the unset key [%s] exists: [%s].", key, gotValue)
			}
		}
	}

	{ // Test of addings to an empty storage
		for k, v := range testCase1 {
			s.Add(k, v)
		}

		for key, value := range testCase1 {
			gotValue, ok := s.Get(key)

			if !ok {
				t.Fatalf("A set key [%s] does not exist.", key)
			}

			if gotValue == "" {
				t.Fatalf("A value of the set key [%s] is an empty string.", key)
			}

			if gotValue != value {
				t.Fatalf("A value of the key [%s] is [%s]. Expected [%s].", key, gotValue, value)
			}
		}
	}

	{ // Test of addings to non empty storage
		for k, v := range testCase2 {
			s.Add(k, v)
		}

		for key, value := range testCase2 {
			gotValue, ok := s.Get(key)

			if !ok {
				t.Fatalf("A set key [%s] does not exist.", key)
			}

			if gotValue == "" {
				t.Fatalf("A value of the set key [%s] is an empty string.", key)
			}

			if gotValue != value {
				t.Fatalf("A value of the key [%s] is [%s]. Expected [%s].", key, gotValue, value)
			}
		}
	}

}
