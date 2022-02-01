package varstorage_test

import (
	"github.com/AdharaProjects/project1-example/pkg/storage/varstorage"
	"testing"
)

func TestDoVar(t *testing.T) {

	varstorage.DoVar()
	t.Error("Fail var")
}
