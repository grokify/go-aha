package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"syscall"
)

func fmt2() {
	fw, err := os.OpenFile("", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	wb := bufio.NewWriter(fw)
	wb.WriteString("hello world\n")
	wb.Flush()
}

type Gofmt struct {
	gofmtPath string
}

func NewGofmt() (Gofmt, error) {
	gofmt := Gofmt{}
	err := gofmt.Inflate()
	return gofmt, err
}

func (gofmt *Gofmt) Inflate() error {
	binary, err := exec.LookPath("gofmt")
	if err != nil {
		return err
	}
	gofmt.gofmtPath = binary
	return nil
}

func (gofmt *Gofmt) Format(filepath string) error {
	//panic("A0")
	if gofmt.gofmtPath == "" {
		err := gofmt.Inflate()
		if err != nil {
			return err
		}
	}
	//panic("A01")
	args := []string{"gofmt", "-w", filepath}

	if 1 == 0 {
		err := syscall.Exec(gofmt.gofmtPath, args, os.Environ())
		panic("AZ")
		if err != nil {
			return err
		}
		panic("AZ")
	}

	//panic("A1")
	cmd := exec.Command(gofmt.gofmtPath, fmt.Sprintf("-w %v", filepath)) //.Output()
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("OUT> %s\n", out.String())
	if err.Error() == "exit status 2" {
		err = nil
	}
	if err != nil {
		panic(err)
	}
	//panic("A2")

	return err
}

type LineInfo struct {
	SrcPattern string
	SrcRegexp  *regexp.Regexp
	OutString  string
}

func (li *LineInfo) Inflate() error {
	rx, err := regexp.Compile(li.SrcPattern)
	if err != nil {
		return err
	}
	li.SrcRegexp = rx
	return nil
}

type FileModifier struct {
	FilepathSrc string
	FilepathOut string
	Lines       [][]string
	LineInfos   []LineInfo
}

func (fm *FileModifier) Inflate() error {
	if len(fm.Lines) > 0 {
		lineInfos := []LineInfo{}
		for _, lineData := range fm.Lines {
			if len(lineData) != 2 {
				return fmt.Errorf("Line data array is not length of 2.")
			}
			li := LineInfo{
				SrcPattern: lineData[0],
				OutString:  lineData[1],
			}
			err := li.Inflate()
			if err != nil {
				return err
			}
			lineInfos = append(lineInfos, li)
		}
		fm.LineInfos = lineInfos
	}
	return nil
}

func (fm *FileModifier) Modify() error {
	err := fm.Inflate()
	if err != nil {
		return err
	}
	fhSrc, err := os.Open(fm.FilepathSrc)
	if err != nil {
		return err
	}

	fhOut, err := os.OpenFile(fm.FilepathOut, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer fhOut.Close()
	wb := bufio.NewWriter(fhOut)

	scanSrc := bufio.NewScanner(fhSrc)
LINE:
	for scanSrc.Scan() {
		fmt.Println(scanSrc.Text())

		tryText := scanSrc.Text()

		matchFound := false

		for _, li := range fm.LineInfos {
			fmt.Printf("TRY: %v %v\n", fm.FilepathOut, tryText)
			if li.SrcRegexp.MatchString(tryText) {
				wb.WriteString(li.OutString)
				matchFound = true
				continue LINE
			}
		}
		if !matchFound {
			wb.WriteString(fmt.Sprintf("%v\n", tryText))
		}
	}
	wb.Flush()
	fhOut.Close()
	return fm.Format()
}

func (fm *FileModifier) Format() error {
	//return nil
	gofmt, err := NewGofmt()
	if err != nil {
		return err
	}
	return gofmt.Format(fm.FilepathOut)
}

func main() {
	if 1 == 1 {
		fm2 := FileModifier{
			FilepathSrc: "./api_client.go",
			FilepathOut: "api_client.mod.go",
			Lines: [][]string{
				{`^\t+"gopkg\.in/go-resty/resty\.v0"`, "\t\"gopkg.in/go-resty/resty.v1\"\n"},
			},
		}
		err := fm2.Modify()
		if err != nil {
			panic(err)
		}
	}

	if 1 == 1 {
		fm1 := FileModifier{
			FilepathSrc: "./configuration.go",
			FilepathOut: "configuration.mod.go",
			Lines: [][]string{
				{`^\t+Transport\s+\*http\.Transport`, "\tTransport\thttp.RoundTripper\n"},
			},
		}
		err := fm1.Modify()
		if err != nil {
			panic(err)
		}
	}

	if 1 == 0 {
		outFile := "configuration.mod.go"
		if 1 == 0 {
			fhSrc, _ := os.Open("./configuration.go")

			//fhOut, err := os.OpenFile("configuration.go.out", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			fhOut, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer fhOut.Close()
			wb := bufio.NewWriter(fhOut)

			scanSrc := bufio.NewScanner(fhSrc)
			rx1 := regexp.MustCompile(`^\t+Transport\s+\*http\.Transport`)
			rx2 := regexp.MustCompile(`^\t+"gopkg\.in/go-resty/resty\.v0"`)
			for scanSrc.Scan() {
				fmt.Println(scanSrc.Text())

				tryText := scanSrc.Text()
				fmt.Sprintf("TRY: %v\n", tryText)
				if rx1.MatchString(tryText) {
					wb.WriteString("\tTransport\thttp.RoundTripper\n")
				} else if rx2.MatchString(tryText) {
					wb.WriteString("\t\"gopkg.in/go-resty/resty.v1\"\n")
				} else {
					wb.WriteString(fmt.Sprintf("%v\n", tryText))
				}
			}
			wb.Flush()
			fhOut.Close()

		}
		/*
			gofmt, err := NewGofmt()
			if err != nil {
				panic(err)
			}
			err = gofmt.Format(outFile)
			if err != nil {
				panic(err)
			}
		*/
	}

	fmt.Println("DONE")
}
