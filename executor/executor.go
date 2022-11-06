package executor

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
)

type languageDetails struct {
	extension string
	name      string
	rootPath  string
	srcPath   string
	isBin     bool
	binPath   string
}

type projectDetails struct {
	projectName     string
	projectId       string
	defaultFileName string
	languageDetails languageDetails
}

var languagesBinConfig = make(map[string]bool)

func init() {
	languagesBinConfig["c"] = true
	languagesBinConfig["py"] = false
	languagesBinConfig["java"] = true
	languagesBinConfig["rb"] = false
	languagesBinConfig["go"] = true
}

func ExecuteRequest(extension string, projectName string, code string) string {
	var result string = ""
	fileCreated, filePath, projPath := checkAndCreateFile(extension, projectName, code)
	if fileCreated {
		switch extension {
		case "java":
			result = buildAndExecuteJava(filePath, projPath)
		case "py":
			result = buildAndExecutePython(filePath, projPath)
		case "rb":
			result = buildAndExecuteRuby(filePath, projPath)

		case "c":
			result = buildAndExecuteC(filePath, projPath)

		}
	}
	return result
}

func checkAndCreateFile(extension string, projectName string, code string) (fileCreated bool, filaPath string, projPath string) {
	srcPath, defaultFileName, projPath := checkAndCreateDir(extension, projectName)
	// sampleProgram := samplePrograms(extension)
	filePath := fmt.Sprintf("%s/%s.%s", srcPath, defaultFileName, extension)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	io.WriteString(f, code)

	return true, filePath, projPath
}

func checkAndCreateDir(extension string, projectName string) (src string, defaultFileName string, projPath string) {
	binPathValue := ""
	projectConfig := projectConfigurationDetails(extension, projectName)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	projectPath := filepath.Join(cwd, projectConfig.projectId, projectConfig.projectName)
	srcPathValue := filepath.Join(projectPath, projectConfig.languageDetails.rootPath, projectConfig.languageDetails.srcPath)
	if _, err := os.Stat(srcPathValue); os.IsNotExist(err) {
		err := os.MkdirAll(srcPathValue, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	if projectConfig.languageDetails.isBin {
		binPathValue = filepath.Join(projectPath, projectConfig.languageDetails.binPath)
		if _, err := os.Stat(binPathValue); os.IsNotExist(err) {
			err := os.MkdirAll(binPathValue, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return srcPathValue, projectConfig.defaultFileName, projectPath
}

func projectConfigurationDetails(extension string, projectName string) projectDetails {
	id := uuid.New()
	switch extension {
	case "c":
		return projectDetails{
			projectName:     projectName,
			projectId:       id.String(),
			defaultFileName: "main",
			languageDetails: languageDetails{
				name:      "C",
				extension: extension,
				rootPath:  "src",
				srcPath:   "main",
				isBin:     languagesBinConfig[extension],
				binPath:   "bin",
			},
		}
	case "java":
		return projectDetails{
			projectName:     projectName,
			projectId:       id.String(),
			defaultFileName: "Main",
			languageDetails: languageDetails{
				name:      "JAVA",
				extension: extension,
				rootPath:  "src",
				srcPath:   "main",
				isBin:     languagesBinConfig[extension],
				binPath:   "bin",
			},
		}
	case "py":
		return projectDetails{
			projectName:     projectName,
			projectId:       id.String(),
			defaultFileName: "main",
			languageDetails: languageDetails{
				name:      "PYTHON",
				extension: extension,
				rootPath:  "src",
				srcPath:   "main",
				isBin:     languagesBinConfig[extension],
				binPath:   "bin",
			},
		}
	case "rb":
		return projectDetails{
			projectName:     projectName,
			projectId:       id.String(),
			defaultFileName: "main",
			languageDetails: languageDetails{
				name:      "RUBY",
				extension: extension,
				rootPath:  "src",
				srcPath:   "main",
				isBin:     languagesBinConfig[extension],
				binPath:   "bin",
			},
		}
	case "go":
		return projectDetails{
			projectName:     projectName,
			projectId:       id.String(),
			defaultFileName: "main",
			languageDetails: languageDetails{
				name:      "GO",
				extension: extension,
				rootPath:  "src",
				srcPath:   "main",
				isBin:     languagesBinConfig[extension],
				binPath:   "bin",
			},
		}
	}
	return projectDetails{}
}

func buildAndExecuteJava(filePath string, projectPath string) string {
	binPath := fmt.Sprintf("%s/bin", projectPath)
	cmd := exec.Command("javac", "-d", binPath, filePath)
	var errb bytes.Buffer
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return string(errb.String())
	}
	buildExecuteString := fmt.Sprintf("%s/bin", projectPath)

	out := exec.Command("java", "-classpath", buildExecuteString, "Main")
	output, err := out.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	return string(output)
}

func buildAndExecutePython(filePath string, projectPath string) string {
	cmd := exec.Command("python", filePath)
	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return string(stderr.String())
	}
	return string(stdout.String())
}

func buildAndExecuteRuby(filePath string, projectPath string) string {
	cmd := exec.Command("ruby", filePath)
	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return string(stderr.String())
	}
	return string(stdout.String())
}

func buildAndExecuteC(filePath string, projectPath string) string {
	binPath := fmt.Sprintf("%s/bin", projectPath)
	cmd := exec.Command("gcc", filePath, "-o", binPath+"/main")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return string(stderr.String())
	}
	buildExecuteString := fmt.Sprintf("%s/bin", projectPath)

	out := exec.Command(buildExecuteString + "/main")
	output, err := out.Output()
	if err != nil {
		log.Println(err)
	}
	return string(output)
}
