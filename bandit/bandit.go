package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	filesPath := []string{"1.js"} // Список директорий проектов
	dir := "eslint/"
	var wg sync.WaitGroup

	t1 := time.Now()

	for _, path := range filesPath {
		wg.Add(1)
		path = "eslint/" + path
		go BanditAnalyzeFile(path, &wg)
	}

	wg.Wait()
	log.Println("time files check cost:", time.Since(t1))

	t2 := time.Now()
	BanditAnalyzeDir(dir, filesPath)
	log.Println("time dir check cost:", time.Since(t2))
}

func BanditAnalyzeFile(filePath string, wg *sync.WaitGroup) ([]byte, []byte) {
	defer wg.Done()

	cmd := exec.Command("npx", "eslint", filePath, "--format", "json")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	_ = cmd.Run()

	//log.Println("Linting completed successfully")
	//fmt.Println(out.String())
	return out.Bytes(), stderr.Bytes()
}

func BanditAnalyzeDir(dir string, filesPath []string) ([]byte, []byte) {

	diffTempDir := "tempDir"
	err := CopyFilesToTempDir("eslint/", diffTempDir, filesPath)
	if err != nil {
		log.Println("CopyFilesToTempDir:", err)
	}
	defer os.RemoveAll(diffTempDir)

	cmd := exec.Command("npx", "eslint", diffTempDir, "--format", "json")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	_ = cmd.Run()

	//log.Println("Linting completed successfully")
	fmt.Println(out.String())
	return out.Bytes(), stderr.Bytes()
}

func CopyFilesToTempDir(src, dest string, files []string) error {
	// Создаем временную директорию
	err := os.Mkdir(dest, os.ModePerm)
	if err != nil {
		log.Printf("Failed to create temp dir: %v", err)
		return err
	}

	// Копируем файлы из списка во временную директорию
	for _, file := range files {
		filePath := filepath.Join(src, file)

		fileName := filepath.Base(file)
		destPath := filepath.Join(dest, fileName)

		err := CopyFile(filePath, destPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func CopyFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}
