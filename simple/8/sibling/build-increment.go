package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

const filename = "/app/build-increment.yml"

type Build struct {
	Data []Data `yaml:"data"`
}

type Data struct {
	Project string `yaml:"project"`
	Counter int64  `yaml:"counter"`
}

func (b *Build) LoadData(filename string) {
	if fileExists(filename) {
		file, _ := ioutil.ReadFile(filename)
		_ = yaml.Unmarshal(file, b)
	} else {
		return
	}
}

func (b *Build) SearchProject(project_name string) *Data {
	for i, row := range b.Data {
		if project_name == row.Project {
			b.Data[i].Counter += 1
			return &row
		}
	}
	return nil
}

func (b *Build) SaveData(filename string) {
	file, _ := yaml.Marshal(b)
	_ = ioutil.WriteFile(filename, file, 0644)
}

func BuildProject(c *gin.Context) {
	project_name := c.PostForm("project")
	fmt.Println(project_name)
	var b Build
	b.LoadData(filename)
	data := b.SearchProject(project_name)
	b.SaveData(filename)
	if data != nil {
		c.JSON(http.StatusOK, gin.H{"build_number": data.Counter})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"build_number": "Not Found"})

}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
