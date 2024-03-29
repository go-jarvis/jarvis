package jarvis

import (
	"os"
	"strings"
	"text/template"

	"github.com/go-jarvis/jarvis/pkg/tmpl"
	"github.com/sirupsen/logrus"
)

func (app *AppCtx) dockerizeCommand() {
	dockerfile, _ := tmpl.GetFile("Dockerfile")
	tmpl, _ := template.New("dockerfile").Parse(string(dockerfile))

	fobj, err := os.OpenFile("Dockerfile.default", os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		logrus.Errorf("create Dockerfile.default failed: %v", err)
	}
	defer fobj.Close()

	data := struct {
		Name    string
		Workdir string
	}{
		Name:    app.name,
		Workdir: workdir(app.rootdir),
	}

	err = tmpl.Execute(fobj, data)
	if err != nil {
		logrus.Errorf("write Dockerfile.default failed: %v", err)
	}
}

func workdir(root string) string {
	workdir, _ := os.Getwd()
	sub := strings.TrimPrefix(workdir, root)

	sub = strings.Trim(sub, "/")
	if sub == "" {
		return "."
	}

	return sub
}
