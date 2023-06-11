package test

import (
	"fmt"
	"testing"
	"time"
)

func TestExample(t *testing.T) {
	now := time.Now()
	fmt.Println("Unix Format", now.Format("2006-01-02"))
}
