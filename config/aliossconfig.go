package config

import bmconfig "github.com/alfredyang1986/blackmirror/bmconfighandle"

type AliOssConfig struct {
	Name            string
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
}

func (aoc *AliOssConfig) GenerateConfig() {
	configPath := "resource/alioss.json"
	profileItems := bmconfig.BMGetConfigMap(configPath)

	aoc.Name = profileItems["Name"].(string)
	aoc.EndPoint = profileItems["EndPoint"].(string)
	aoc.AccessKeyID = profileItems["AccessKeyID"].(string)
	aoc.AccessKeySecret = profileItems["AccessKeySecret"].(string)

}
