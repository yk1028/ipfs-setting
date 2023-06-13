package util

import (
	"os"
	"sync"

	"github.com/Moonyongjung/tx-stress-test/types"
	"gopkg.in/yaml.v3"
)

var once sync.Once
var tester sync.Once
var instance *configManager
var testerInstance *testerManager

type configManager struct {
	conf types.ConfigType
}

type testerManager struct {
	tester types.TesterType
}

// Read and get config.yaml
func Config() *configManager {
	once.Do(func() {
		instance = &configManager{}
	})
	return instance
}

func (s *configManager) Get() types.ConfigType {
	return s.conf
}

func (s *configManager) Read(filePath string) error {
	ConfigType, err := readConfigFile(filePath)
	if err != nil {
		return err
	}
	s.conf = *ConfigType

	return nil
}

func readConfigFile(filePath string) (*types.ConfigType, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var ConfigType types.ConfigType

	err = yaml.Unmarshal(yamlFile, &ConfigType)
	if err != nil {
		return nil, err
	}

	return &ConfigType, nil
}

func Tester() *testerManager {
	tester.Do(func() {
		testerInstance = &testerManager{}
	})
	return testerInstance
}

func (t *testerManager) Get() types.TesterType {
	return t.tester
}

func (t *testerManager) Read(filePath string) error {
	testerType, err := readTesterFile(filePath)
	if err != nil {
		return err
	}

	t.tester = *testerType

	return nil
}

func readTesterFile(filePath string) (*types.TesterType, error) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var testerType types.TesterType

	err = yaml.Unmarshal(yamlFile, &testerType)
	if err != nil {
		return nil, err
	}

	return &testerType, nil
}
