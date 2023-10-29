package build

import (
	"os"
	"fmt"
	"reflect"

	"github.com/microgiantya/asciiart"

	. "github.com/microgiantya/logger"
)

const (
	p = "build"
)

var (
	name        string
	version     string
	date        string
	host        string
	cpuCores    string
	description string

	buildInfo = _buildInfo{
		Name:        name,
		Version:     version,
		Date:        date,
		Host:        host,
		CpuCores:    cpuCores,
		Description: description,
	}
)

type _buildInfo struct {
	Name        string
	Version     string
	Date        string
	Host        string
	CpuCores    string
	Description string
}

func init() {
	info, err := buildinfo.ReadFile(os.Args[0])
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
        fmt.Printf("%+v\n", info.String())

	for _, line := range asciiart.Lib[buildInfo.Name] {
		I_("", line)
	}

	t := reflect.TypeOf(buildInfo)
	v := reflect.ValueOf(buildInfo)
	for i := 0; i < v.NumField(); i++ {
		I_(p, fmt.Sprintf("%s: %+v", t.Field(i).Name, v.Field(i).Interface()))
	}
}
