package rc

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const (
	fileModeUserReadWrite = 0600
)

type PivnetRC struct {
	Profiles []PivnetProfile `yaml:"profiles"`
}

type RCHandler struct {
	configFilepath string
}

func NewRCHandler(configFilepath string) *RCHandler {
	return &RCHandler{
		configFilepath: configFilepath,
	}
}

func (h *RCHandler) SaveProfile(
	profileName string,
	apiToken string,
	host string,
) error {
	pivnetRC, err := h.loadPivnetRC()
	if err != nil {
		return err
	}

	if pivnetRC == nil {
		pivnetRC = &PivnetRC{}
	}

	var newInfo *PivnetProfile
	var index int
	for i, p := range pivnetRC.Profiles {
		if p.Name == profileName {
			newInfo = &p
			index = i
			break
		}
	}

	if newInfo == nil {
		newInfo = &PivnetProfile{
			Name:     profileName,
			APIToken: apiToken,
			Host:     host,
		}
		index = len(pivnetRC.Profiles)
		pivnetRC.Profiles = append(pivnetRC.Profiles, PivnetProfile{})
	}

	newInfo.APIToken = apiToken
	newInfo.Host = host

	pivnetRC.Profiles[index] = *newInfo

	return h.writePivnetRC(h.configFilepath, *pivnetRC)
}

// ProfileForName will return (nil,nil) if the file does not exist,
// or if the profile does not exist,
// but will return (nil,err) for other reasons e.g. the file cannot be read.
func (h *RCHandler) ProfileForName(profileName string) (*PivnetProfile, error) {
	pivnetRC, err := h.loadPivnetRC()
	if err != nil {
		return nil, err
	}

	if pivnetRC == nil {
		return nil, nil
	}

	for _, p := range pivnetRC.Profiles {
		if p.Name == profileName {
			return &p, nil
		}
	}

	return nil, nil
}

// RemoveProfileWithName will return error for all errors
func (h *RCHandler) RemoveProfileWithName(profileName string) error {
	pivnetRC, err := h.loadPivnetRC()
	if err != nil {
		return err
	}

	if pivnetRC == nil {
		return nil
	}

	var foundIndex int
	for i, p := range pivnetRC.Profiles {
		if p.Name == profileName {
			foundIndex = i
		}
	}

	pivnetRC.Profiles = append(pivnetRC.Profiles[:foundIndex], pivnetRC.Profiles[foundIndex+1:]...)

	return h.writePivnetRC(h.configFilepath, *pivnetRC)
}

// loadPivnetRC does not return an error if the file does not exist
// but will return an error for other reasons e.g. the file cannot be read.
func (h *RCHandler) loadPivnetRC() (*PivnetRC, error) {
	_, err := os.Stat(h.configFilepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	pivnetRCBytes, err := ioutil.ReadFile(h.configFilepath)
	if err != nil {
		return nil, err
	}

	var pivnetRC PivnetRC
	err = yaml.Unmarshal(pivnetRCBytes, &pivnetRC)
	if err != nil {
		return nil, err
	}

	return &pivnetRC, nil
}

func (h *RCHandler) writePivnetRC(configFileLocation string, contents PivnetRC) error {
	yamlBytes, err := yaml.Marshal(contents)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFileLocation, yamlBytes, fileModeUserReadWrite)
	if err != nil {
		return err
	}

	err = os.Chmod(configFileLocation, fileModeUserReadWrite)
	if err != nil {
		return err
	}

	return nil
}
