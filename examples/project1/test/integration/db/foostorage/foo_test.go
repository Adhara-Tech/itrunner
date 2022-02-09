package foostorage_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Adhara-Tech/project1-example/pkg/storage/foostorage"
)

func TestDoFoo(t *testing.T) {
	foostorage.DoFoo()
	fmt.Println(os.Getenv("CUSTOM_KEY"))
}
