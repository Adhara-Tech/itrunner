package varstorage_test

import (
	"testing"

	"github.com/Adhara-Tech/project1-example/pkg/storage/varstorage"
)

func TestDoVar(t *testing.T) {

	varstorage.DoVar()
	t.Error("Fail var")
}
