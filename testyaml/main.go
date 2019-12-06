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

type DockerConfig struct {
	UpdateCorn  string       `yaml:"updatecorn"`
	SwarmClient []ClientInfo `yaml:"dockerclient"`
	SingleDockerConfig       `yaml:"singledocker"`
	RouterConfig             `yaml:"router"`
}

type RouterConfig struct {
	CrossNetwork bool
	Routers      map[string]*IntraFromTo `yaml:"innerrouters,omitempty"`
	InnerRouters []FromTo                `yaml:"routers"` //从yaml文件中读取，内部使用上方的Routers
	IptsImage    string                  `yaml:"iptsimage"`
}
type IntraFromTo struct {
	FromIp   string
	FromPort uint32
	ToIp     string
	ToPort   uint32
}

type FromTo struct {
	From string
	To   string
}

type SingleDockerConfig struct {
	BaseDir       string                  `yaml:"basedir"`
	Dockertls     bool                    `yaml:"dockertls"`
	Tlsca         string                  `yaml:"tlsca"`
	Tlscert       string                  `yaml:"tlscert"`
	Tlskey        string                  `yaml:"tlskey"`
}

type ClientInfo struct {
	Address string
	Key     string
}

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
	walkDecode(c.Content[0])

	c.Content[0].Content[1].Value = "custom ec2 value" // <-dangerous

	// TODO: really should walk the yaml.Node tree and find the relevent ec2 field value
	// TODO: all the while staying within slice bounds

	err = writeYaml(c)
	if err != nil {
		panic(err)
	}
	//验证是否可读
	dockerConfig := DockerConfig{}
	yamlFile, err := ioutil.ReadFile(basepath1)
	if err != nil {
		logger.Error("Fail to read router config: ", err.Error())
		return
	}
	err = yaml.Unmarshal(yamlFile, &dockerConfig)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("dockerConfig.RouterConfig.Routers:  ",dockerConfig.RouterConfig.InnerRouters)
}

func walkDecode(node *yaml.Node) {
	for k, v := range node.Content {
		walkDecode(v)
		if v.Value == "dockerclient" {
			tmpnode1 := &yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "address1",
				Line:3,
				Column:10,
			}
			tmpnode2 := &yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "aaaaaaaaaaa",
				Line:3,
				Column:10,
			}
			tmpnode := &yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					tmpnode1, tmpnode2,
				},
			}

			node.Content[k+1].Content = append(node.Content[k+1].Content, tmpnode)
			if len(node.Content[k+1].Content)==1{
				logger.Info("tmpnode:  ", node.Content[k+1])
				node.Content[k+1].Kind = yaml.SequenceNode
				node.Content[k+1].Tag = "!!seq"
			}

			for k1, vv := range node.Content[k+1].Content {
				logger.Info("k1:", k1, vv)
				logger.Info("next v:", vv.Content[0])
				logger.Info("next v:", vv.Content[1])
			}
		}
		if v.Value == "router"{
			logger.Info("node.content: ",node.Content[k+1].Content[0])
		}
		if v.HeadComment != "" {
			res, err := EncodeFromHZGB2312(v.HeadComment)
			if err != nil {
				panic(err)
			}
			v.HeadComment = res
		}
		if v.FootComment != "" {
			res, err := EncodeFromHZGB2312(v.FootComment)
			if err != nil {
				panic(err)
			}
			v.FootComment = res
		}
		if v.LineComment != "" {
			res, err := EncodeFromHZGB2312(v.LineComment)
			if err != nil {
				panic(err)
			}
			v.LineComment = res
		}
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
