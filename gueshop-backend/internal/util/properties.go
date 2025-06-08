package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type ConfigProperties interface {
	GetProp(name string) string
	GetIntProp(name string) int
	GetInt32Prop(name string) int32
	GetInt64Prop(name string) int64
	GetBoolProp(name string) bool
	GetFloat64Prop(name string) float64
	GetStringArray(name string, splitBy string) []string
}

type GoEnvConfigProperties struct {
}

func NewGoEnvConfigProperties() ConfigProperties {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found!")
	}
	return &GoEnvConfigProperties{}
}

func (util *GoEnvConfigProperties) GetProp(name string) string {
	res := os.Getenv(name)
	if res == "" {
		fmt.Printf("no property %s found! \n", name)
	}
	return res
}

func (util *GoEnvConfigProperties) GetIntProp(name string) int {
	return int(util.GetInt32Prop(name))
}

func (util *GoEnvConfigProperties) GetInt32Prop(name string) int32 {
	result, err := strconv.ParseInt(util.GetProp(name), 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(result)
}

func (util *GoEnvConfigProperties) GetInt64Prop(name string) int64 {
	result, err := strconv.ParseInt(util.GetProp(name), 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func (util *GoEnvConfigProperties) GetBoolProp(name string) bool {
	result, err := strconv.ParseBool(util.GetProp(name))
	if err != nil {
		panic(err)
	}
	return result
}

func (util *GoEnvConfigProperties) GetFloat64Prop(name string) float64 {
	result, err := strconv.ParseFloat(util.GetProp(name), 64)
	if err != nil {
		panic(err)
	}
	return result
}

func (util *GoEnvConfigProperties) GetStringArray(name string, splitBy string) []string {
	text := util.GetProp(name)
	return strings.Split(text, splitBy)
}
