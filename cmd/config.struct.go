package cmd

type Config struct {
	SopreQasPath     string `yaml:"sopre-qas-path"`
	SopreQasBinary   string `yaml:"sopre-qas-binary"`
	SopreInstallPath string `yaml:"sopre-install-path"`
	SopreGuiPath     struct {
		D struct {
			Ap struct {
				Designer string `yaml:"designer"`
				Client   string `yaml:"client"`
				Editor   string `yaml:"editor"`
			} `yaml:"ap"`
			Ep struct {
				Designer string `yaml:"designer"`
				Client   string `yaml:"client"`
				Editor   string `yaml:"editor"`
			} `yaml:"ep"`
		} `yaml:"d"`
		Bc struct {
			Ep struct {
				Designer string `yaml:"designer"`
				Client   string `yaml:"client"`
				Editor   string `yaml:"editor"`
			} `yaml:"ep"`
			Global struct {
				Designer string `yaml:"designer"`
				Client   string `yaml:"client"`
				Editor   string `yaml:"editor"`
			} `yaml:"global"`
			Le struct {
				Designer string `yaml:"designer"`
				Client   string `yaml:"client"`
				Editor   string `yaml:"editor"`
			} `yaml:"le"`
			Plst struct {
				Designer string `yaml:"designer"`
				Client   string `yaml:"client"`
				Editor   string `yaml:"editor"`
			} `yaml:"plst"`
		} `yaml:"bc"`
	} `yaml:"sopre-gui-path"`
}
