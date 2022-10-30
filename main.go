package main

import (
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

func main() {
	excecute_request("rb", "firstrubyProject")
}

func excecute_request(extension string, projectName string) {
	fileCreated, filePath, projPath := checkAndCreateFile(extension, projectName)
	if fileCreated {
		switch extension {
		case "java":
			buildAndExecuteJava(filePath, projPath)
		case "py":
			buildAndExecutePython(filePath, projPath)
		case "rb":
			buildAndExecuteRuby(filePath, projPath)

		case "c":
			buildAndExecuteC(filePath, projPath)

		}
	}
}

func samplePrograms(extension string) string {
	switch extension {
	case "c":
		return `#include <stdio.h>
int main()
{
	printf("Hello World! by Pradeek\n");
	return 0;
}`
	case "java":
		return `public class Main {
	public static void main(String[] args) {
		System.out.println("Hello World!");
	}
}`
	case "py":
		return `print("Hello World!")`
	case "rb":
		return `print("Hello World!")`
	case "go":
		return `package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
`
	}
	return ""
}

func checkAndCreateFile(extension string, projectName string) (fileCreated bool, filaPath string, projPath string) {
	srcPath, defaultFileName, projPath := checkAndCreateDir(extension, projectName)
	sampleProgram := samplePrograms(extension)
	filePath := fmt.Sprintf("%s/%s.%s", srcPath, defaultFileName, extension)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	io.WriteString(f, sampleProgram)

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

func buildAndExecuteJava(filePath string, projectPath string) {
	binPath := fmt.Sprintf("%s/bin", projectPath)
	cmd := exec.Command("javac", "-d", binPath, filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	buildExecuteString := fmt.Sprintf("%s/bin", projectPath)

	out := exec.Command("java", "-classpath", buildExecuteString, "Main")
	// out.Path = "java"
	output, err := out.CombinedOutput()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(output))
}

func buildAndExecutePython(filePath string, projectPath string) {
	cmd := exec.Command("python", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func buildAndExecuteRuby(filePath string, projectPath string) {
	cmd := exec.Command("ruby", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func buildAndExecuteC(filePath string, projectPath string) {
	binPath := fmt.Sprintf("%s/bin", projectPath)
	cmd := exec.Command("gcc", filePath, "-o", binPath+"/main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	buildExecuteString := fmt.Sprintf("%s/bin", projectPath)

	out := exec.Command(buildExecuteString + "/main")
	// out.Path = "java"
	output, err := out.Output()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(output))
}
