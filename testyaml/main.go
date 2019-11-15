package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"github.com/op/go-logging"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("test")
}

var basepath = "testyaml/app.yaml"
var basepath1 = "testyaml/config/app.yaml"

func writeYaml(c *yaml.Node) error {

	bytes, err := yaml.Marshal(c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(basepath1, bytes, 0644)
	if err != nil {
		logger.Error(err)
		return err
	}
	return err
}

func readYamlSchema(b []byte) (*yaml.Node, error) {
	var c = yaml.Node{}
	err := yaml.Unmarshal(b, &c)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &c, nil
}

func main() {

	all, err := ioutil.ReadFile(basepath)
	if err != nil {
		panic(err)
	}
	logger.Info(string(all))
	encoded, err := DecodeToHZGB2312(string(all))
	if err != nil {
		panic(err)
	}
	c, err := readYamlSchema([]byte(encoded))
	if err != nil {
		panic(err)
	}
	//logger.Info("c:",c.Content[0].Content)
	logger.Info("------将中文注释decode回utf8--------")
	for _, v := range c.Content[0].Content {
		if v.HeadComment != "" {
			res, err := EncodeFromHZGB2312(v.HeadComment)
			if err != nil {
				panic(err)
			}
			v.HeadComment = res
		}
	}

	c.Content[0].Content[1].Value = "custom ec2 value" // <-dangerous

	// TODO: really should walk the yaml.Node tree and find the relevent ec2 field value
	// TODO: all the while staying within slice bounds

	err = writeYaml(c)
	if err != nil {
		panic(err)
	}
}

//utf8 -> HZGB2312
func DecodeToHZGB2312(utf8Str string) (dst string, err error) {
	var trans transform.Transformer = simplifiedchinese.HZGB2312.NewEncoder()
	var reader *strings.Reader = strings.NewReader(utf8Str)
	var transReader *transform.Reader = transform.NewReader(reader, trans)
	bytes, err := ioutil.ReadAll(transReader)
	if err != nil {
		return
	}
	dst = string(bytes)
	return
}

//HZGB2312 -> utf8
func EncodeFromHZGB2312(gbkStr string) (utf8Str string, err error) {
	var trans transform.Transformer = simplifiedchinese.HZGB2312.NewDecoder()
	var reader *strings.Reader = strings.NewReader(gbkStr)
	var transReader *transform.Reader = transform.NewReader(reader, trans)
	bytes, err := ioutil.ReadAll(transReader)
	if err != nil {
		return
	}
	utf8Str = string(bytes)
	return
}
