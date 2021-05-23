package appsettings

type AppSettings struct {
	OutputFile struct {
		OutputFilePath     string `envconfig:"optional"`
		AppendToOutputFile bool   `envconfig:"optional"`
	}
	Output struct {
		NumberToDisplay       int  `envconfig:"optional"`
		DisplayUniqueIPs      bool `envconfig:"optional"`
		DisplayUniqueURLs     bool `envconfig:"optional"`
		SortByKeyAdditionally bool `envconfig:"optional"`
	}
	InputFilePath     string `envconfig:"optional"`
	RegularExpression string `envconfig:"optional"`
}
