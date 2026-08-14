package main

import (
	"log"
	"os"

	"go.yaml.in/yaml/v3"
)

type ProxyConfig struct {
	Services map[string]Servers `yaml:"services"`
	
}
type Servers struct{
	Url string `yaml:"url,omitempty"`
	Connections struct{
		Rest struct {
			Paths map[string]RestSettings `yaml:"paths,omitempty"`
		} `yaml:"rest,omitempty"`
		Grpc struct {
			ProtoFilePath string `yaml:"proto_file_path,omitempty"`
			Services map[string]GrpcSettings `yaml:"services,omitempty"`
		} `yaml:"grpc,omitempty"`
	} `yaml:"connections,omitempty"`
	Assets map[string]StaticSettings `yaml:"assets,omitempty"`
}

type RestSettings struct{
	MappedPath string `yaml:"mapped_path,omitempty"`
	RequestType string `yaml:"request_type,omitempty"`
}

type GrpcSettings struct{
	Rpc map[string]RpcSettings `yaml:"rpc,omitempty"`
}

type StaticSettings struct {
	MappedPath string `yaml:"mapped_path,omitempty"`
}

type RpcSettings struct{
	MappedPath string `yaml:"mapped_path,omitempty"`
}

func LoadYamlConfigs(filename string) (*ProxyConfig,error){
	var proxyConfig *ProxyConfig
	yamlFile,err:=os.ReadFile(filename)
	if err!=nil{
		return nil,err
	}
	
	err=yaml.Unmarshal(yamlFile,&proxyConfig)
	if err!=nil{
		return nil,err
	}
	return proxyConfig,nil



}


func main(){
	proxyConfig,err:=LoadYamlConfigs("../.config/proxy-config.yaml")
	if err!=nil{
		log.Fatal(err)
	}
	log.Print(proxyConfig)



}
