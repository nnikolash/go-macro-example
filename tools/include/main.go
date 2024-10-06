package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	targetFilePath := os.Getenv("GOFILE")
	if targetFilePath == "" {
		log.Fatal("GOFILE must be set")
	}

	headersDir := os.Getenv("GO_INCLUDE_HEADERS_DIR")
	if headersDir == "" {
		workDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("get work dir: %v", err)
		}

		headersDir = workDir
	}

	genDir := os.Getenv("GO_INCLUDE_GEN_DIR")
	if genDir == "" {
		targetFileDir := path.Dir(targetFilePath)
		genDir = path.Join(targetFileDir, "generated")
		if err := os.MkdirAll(genDir, 0755); err != nil {
			log.Fatalf("create dir '%v': %v", genDir, err)
		}
	}

	definesStr := os.Getenv("GO_INCLUDE_DEFINES")
	defines := strings.Split(definesStr, ",")

	fmt.Printf("Including %v into %v, defines = %v\n", os.Args[1:], targetFilePath, defines)

	codeBytes, err := os.ReadFile(targetFilePath)
	if err != nil {
		log.Fatalf("read file: %v", err)
	}

	lines := strings.Split(string(codeBytes), "\n")

	newLines := []string{lines[0]}

	for _, define := range defines {
		if define == "" {
			continue
		}
		newLines = append(newLines, fmt.Sprintf("#define %v", define))
	}

	for _, incl := range os.Args[1:] {
		if incl[0] != '<' && incl[0] != '"' {
			incl = "\"" + incl + "\""
		}
		if incl[0] == '"' {
			incl = "\"" + headersDir + "/" + incl[1:]
		}
		newLines = append(newLines, fmt.Sprintf("#include %v", incl))
	}

	newLines = append(newLines, lines[1:]...)
	newFileContent := strings.Join(newLines, "\n")

	tmpFileWithIncludes := createTempFile()
	fmt.Printf("Writing file with includes to %v\n", tmpFileWithIncludes.Name())
	tmpFileWithIncludes.Write([]byte(newFileContent))
	tmpFileWithIncludes.Close()

	generatedFilePath := genDir + "/" + targetFilePath

	cmd("g++", "-P", "-E", "-x", "c++", tmpFileWithIncludes.Name(), "-o", generatedFilePath)
	cmd("goimports", "-w", generatedFilePath)
}

func cmd(cmd string, args ...string) error {
	fmt.Printf("Running %v %v\n", cmd, "\""+strings.Join(args, "\" \"")+"\"")
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("exec '%v' %v: %v\n%v", cmd, args, err, string(out))
	}
	return nil
}

func createTempFile() *os.File {
	f, err := os.CreateTemp("", "go_include")
	if err != nil {
		log.Fatalf("create temp file: %v", err)
	}

	return f
}
