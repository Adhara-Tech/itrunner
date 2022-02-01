package foostorage_test

import (
	"fmt"
	"github.com/AdharaProjects/project1-example/pkg/storage/foostorage"
	"os"
	"testing"
)

func TestDoFoo(t *testing.T) {
	foostorage.DoFoo()
	fmt.Println(os.Getenv("CUSTOM_KEY"))
	//t.Error("Foo not implemented")
}
