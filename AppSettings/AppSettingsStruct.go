package appsettings

type AppSettings struct {
	// Setting for Output File
	OutputFile struct {
		// If this is set to a non-empty string and a path with filename.txt is provided it will output the information to a file and the console, otherwise just to console
		OutputFilePath string `envconfig:"optional"`
		// false will create a new file every time, and true will append to the same file
		AppendToOutputFile bool `envconfig:"optional"`
	}
	// These option allows one to vary how the output is displayed
	Output struct {
		// This will allow you to increase the number of HITS and Occurences to display on output, if this exceeds the number available it will only display what it can
		NumberToDisplay int `envconfig:"optional"`
		//  Include an ouput for Number of Unique IP Addresses
		DisplayUniqueIPs bool `envconfig:"optional"`
		// Include an output for Number of Unique URLS
		DisplayUniqueURLs bool `envconfig:"optional"`
		// true, KeyValuePair<string,int> will be sorted first by int (number of occurrences) and then by string (IP / URL) will provide more consistent output when the number of occurrences are the same
		// false, KeyValuePair<string,int> will only be sorted first by int (number of occurrences)
		SortByKeyAdditionally bool `envconfig:"optional"`
	}
	// This is the path to the log file that will be parsed containing a IP and URL per line
	InputFilePath string `envconfig:"optional"`
	// Provided Regular Expression pattern can not be empty, it must contain two groups the first identifying the IP Address, and the second identifying the URL
	RegularExpression string `envconfig:"optional"`
}
