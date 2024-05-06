package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDataSourceName(t *testing.T) {
	port := "1111"
	host := "localhoost"
	name := ""
	user := "user01"
	password := "pass02!"
	sslmode := "disable"
	config := ConfigDatabase{
		Port: port, Host: host, Name: name, User: user, Password: password, SSLMode: sslmode,
	}
	assert.Equal(
		t,
		"host="+host+" port="+port+" dbname="+name+" user="+user+" password="+password+" sslmode="+sslmode,
		config.GetDataSourceName(),
	)
}

func TestDefaultConfigInfo(t *testing.T) {
	assert.Equal(
		t,
		&ConfigInfo{
			Name:      "config",
			Extension: "yaml",
			Paths:     []string{".", "config", "pkg/config"},
		},
		DefaultConfigInfo(),
	)
}

func TestInitConfigFile(t *testing.T) {
	assert.Panics(
		t,
		func() {
			InitConfig(ConfigInfo{Name: "^^^", Extension: "not right5%", Paths: []string{"...///"}})
		},
	)
}

// func TestInitConfigUnmarshal(t *testing.T) {
// 	configTest := ConfigApp{
// 		Database: ConfigDatabase{
// 			Host:     "lcalhost",
// 			Port:     "532",
// 			Name:     "pstgres",
// 			User:     "pstgres",
// 			Password: "",
// 			SSLMode:  "isable",
// 		},
// 	}

// 	yamlData, err := yaml.Marshal(&configTest)
// 	if err != nil {
// 		fmt.Println("Error marshaling YAML:", err)
// 		return
// 	}

// 	// Write the YAML data to a file
// 	err = os.WriteFile("config_test.yaml", yamlData, 0777)
// 	if err != nil {
// 		fmt.Println("Error writing YAML file:", err)
// 		return
// 	}

// 	configName := "config_test"
// 	configExtension := "yaml"
// 	configPaths := []string{".", "config", "pkg/config"}
// 	assert.PanicsWithError(
// 		t,
// 		"a",
// 		func() { InitConfig(ConfigInfo{Name: configName, Extension: configExtension, Paths: configPaths}) },
// 	)

// 	fmt.Println("Config written to config_test.yaml")

// 	// Delete the file
// 	err = os.Remove("config_test.yaml")
// 	if err != nil {
// 		fmt.Println("Error deleting file:", err)
// 		return
// 	}
// }

// func TestInitConfigCorrect(t *testing.T) {
// 	configTest := ConfigApp{
// 		Database: ConfigDatabase{
// 			Host:     "localhost",
// 			Port:     "5432",
// 			Name:     "postgres",
// 			User:     "postgres",
// 			Password: "",
// 			SSLMode:  "disable",
// 		},
// 		Server: ConfigServer{
// 			Host: "localhost",
// 			Port: "8080",
// 		},
// 	}

// 	yamlData, err := yaml.Marshal(&configTest)
// 	if err != nil {
// 		fmt.Println("Error marshaling YAML:", err)
// 		return
// 	}

// 	// Write the YAML data to a file
// 	err = os.WriteFile("config_test.yaml", yamlData, 0777)
// 	if err != nil {
// 		fmt.Println("Error writing YAML file:", err)
// 		return
// 	}

// 	configName := "config_test"
// 	configExtension := "yaml"
// 	configPaths := []string{".", "config", "pkg/config"}
// 	assert.Equal(t, configTest, InitConfig(ConfigInfo{Name: configName, Extension: configExtension, Paths: configPaths}))

// 	// Delete the file
// 	err = os.Remove("config_test.yaml")
// 	if err != nil {
// 		fmt.Println("Error deleting file:", err)
// 		return
// 	}
// }
