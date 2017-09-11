package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"os/exec"
	"io/ioutil"
	"bytes"
)

func RunCommand(name string, arg ...string) (bytes.Buffer, bytes.Buffer, error) {
	cmd := exec.Command(name, arg...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout, stderr, err
}


const PackageName = "summon-file"
func TestPackage(t *testing.T) {

	Convey("Given a compiled summon-file package", t, func() {

		Convey("Given summon-file is run with no arguments", func() {
			_, stderr, err := RunCommand(PackageName)

			Convey("Returns with error", func() {
				So(err, ShouldNotBeNil)
				So(stderr.String(), ShouldEqual, "A variable identifier must be given as the first and only argument!")
			})
		})

		Convey("Given summon-file is run with too many arguments", func() {
			_, stderr, err := RunCommand(PackageName, "one", "two")

			Convey("Returns with error", func() {
				So(err, ShouldNotBeNil)
				So(stderr.String(), ShouldEqual, "A variable identifier must be given as the first and only argument!")
			})
		})

		Convey("Given summon-file is run with an existent file path", func() {
			content := []byte("existent-file-content")
			tmpfile, err := ioutil.TempFile("", "existent-file")
			tmpfileName := tmpfile.Name()
			defer os.Remove(tmpfileName)

			tmpfile.Write(content)
			tmpfile.Close()

			stdout, _, err := RunCommand(PackageName, tmpfileName)

			Convey("Prints the contents of the file on stdout", func() {
				So(err, ShouldBeNil)
				So(stdout.String(), ShouldEqual, "existent-file-content")
			})
		})

		Convey("Given summon-file is run with a non-existent file path", func() {
			_, stderr, err := RunCommand(PackageName, "non-existent-file")

			Convey("Returns with error", func() {
				So(err, ShouldNotBeNil)
				So(stderr.String(), ShouldContainSubstring, "no such file or directory")
			})
		})

	})
}

