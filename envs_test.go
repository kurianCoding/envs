package envs

import "testing"
import "os"

func TestReadenv(t *testing.T) {
	os.Setenv("test", "random")
	value := Read()["test"]
	if value != "random" {
		t.Errorf("read value is different than set value %s", value)
	}

}
