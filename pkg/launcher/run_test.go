package launcher

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type Runner2 struct {
}

func (r *Runner2) Appname() string {
	return "runner2s"
}

func (r *Runner2) Shutdown() error {
	return nil
}

func (r *Runner2) Run() error {
	time.Sleep(3 * time.Second)
	return fmt.Errorf("sleep 5 s")
}

func Test_Run(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	la := &Launcher{}
	la.Launch(ctx, &Runner2{})
}
