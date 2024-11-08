package grpc

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	unzip "tea.gitpark.ru/sast/shpack/files"
)

func LoadLocalFile(filePath string) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("err file info", err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("err buffer", err)
	}

	return buffer, nil
}

func ZipFiles(files []string, baseDir, output string) error {
	newZipFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = addFileToZip(zipWriter, file, baseDir); err != nil {
			return err
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename, baseDir string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	relativePath, err := filepath.Rel(baseDir, filename)
	if err != nil {
		return err
	}
	header.Name = relativePath

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}

func SeparateZipArchive(archive []byte, tempDir string) (goArchivePath, dockerfileArchivePath, jsArchivePath, pyArchivePath string, err error) {

	_, err = unzip.Unzip(archive, tempDir)
	if err != nil {
		return "", "", "", "", err
	}

	var goFiles, dockerFiles, jsFiles, pyFiles []string
	err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), ".go") || info.Name() == "go.mod" {
				goFiles = append(goFiles, path)
			} else if strings.Contains(info.Name(), "Dockerfile") {
				dockerFiles = append(dockerFiles, path)
			} else if strings.HasSuffix(info.Name(), ".js") || strings.HasSuffix(info.Name(), ".ts") || info.Name() == "package.json" || info.Name() == "package-lock.json" {
				jsFiles = append(jsFiles, path)
			} else if strings.Contains(info.Name(), ".py") /*|| info.Name() == "requirements.txt"*/ {
				pyFiles = append(pyFiles, path)
			}
		}

		return nil
	})
	if err != nil {
		return "", "", "", "", err
	}

	goArchivePath = filepath.Join(tempDir, "go_files.zip")
	err = ZipFiles(goFiles, tempDir, goArchivePath)
	if err != nil {
		return "", "", "", "", err
	}

	dockerfileArchivePath = filepath.Join(tempDir, "dockerfile_files.zip")
	err = ZipFiles(dockerFiles, tempDir, dockerfileArchivePath)
	if err != nil {
		return "", "", "", "", err
	}

	jsArchivePath = filepath.Join(tempDir, "js_files.zip")
	err = ZipFiles(jsFiles, tempDir, jsArchivePath)
	if err != nil {
		return "", "", "", "", err
	}

	pyArchivePath = filepath.Join(tempDir, "py_files.zip")
	err = ZipFiles(pyFiles, tempDir, pyArchivePath)
	if err != nil {
		return "", "", "", "", err
	}

	return goArchivePath, dockerfileArchivePath, jsArchivePath, pyArchivePath, nil
}
