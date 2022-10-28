package main

type (
	Main struct {
		Version string `yaml:"openapi"`
		Info    `yaml:"info"`
		Servers []Server               `yaml:"servers"`
		Paths   map[string]interface{} `yaml:"paths"`
	}

	Info struct {
		Title   string `yaml:"title"`
		Desc    string `yaml:"description"`
		Version string `yaml:"version"`
	}

	Server struct {
		Url string `yaml:"url"`
	}

	Parameter struct {
		Item map[string]string
	}
)
