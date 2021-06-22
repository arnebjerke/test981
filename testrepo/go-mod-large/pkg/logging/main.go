package logging

import (
	"fmt"
	"log"

	"github.com/flant/logboek"
)

var (
	imageNameFormat    = "⛵ image %s"
	artifactNameFormat = "🛸 artifact %s"
)

func Init() error {
	if err := logboek.Init(); err != nil {
		return err
	}

	log.SetOutput(logboek.GetOutStream())

	return nil
}

func EnableLogColor() {
	logboek.EnableLogColor()
}

func DisableLogColor() {
	logboek.DisableLogColor()
}

func SetWidth(value int) {
	logboek.SetWidth(value)
}

func DisablePrettyLog() {
	imageNameFormat = "image %s"
	artifactNameFormat = "artifact %s"

	logboek.DisablePrettyLog()
}

func ImageLogName(name string, isArtifact bool) string {
	if !isArtifact {
		if name == "" {
			name = "~"
		}
	}

	return name
}

func ImageLogProcessName(name string, isArtifact bool) string {
	logName := ImageLogName(name, isArtifact)
	if !isArtifact {
		return fmt.Sprintf(imageNameFormat, logName)
	} else {
		return fmt.Sprintf(artifactNameFormat, logName)
	}
}
