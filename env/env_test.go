package env

import (
	"log"
	"os"
	"testing"
)

func TestLookup(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want string
	}{
		{
			name: "missing env value",
			key:  "MISSING_KEY",
			want: "default",
		},
		{
			name: "existing env value",
			key:  "FOO",
			want: "bar",
		},
	}

	err := os.Setenv("FOO", "bar")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := os.Unsetenv("FOO")
		if err != nil {
			log.Fatal(err)
		}
	}()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Lookup(test.key, "default")
			if want := test.want; got != want {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
