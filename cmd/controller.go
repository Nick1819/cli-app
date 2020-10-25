package cmd

func loopController(result string, loop string) {
	if loop == "yes" {
		for {
			playMP3(result)
		}
	}
	if loop == "no" {
		playMP3(result)
	}
}
