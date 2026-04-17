package aliyunoss

import (
	"testing"

	"github.com/zdz1715/appender"
)

func TestClientIsDriver(t *testing.T) {
	var driver appender.Driver

	driver = NewClient(Config{})

	t.Logf("%+v", driver)
}
