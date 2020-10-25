package cmd

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "option: play [.mp3, .mp4] file will initialize the music",
	RunE: func(cmd *cobra.Command, args []string) error {

		statusmp3, err := cmd.Flags().GetBool("mp3")
		if err != nil {
			return err
		}

		statusmp4, err := cmd.Flags().GetBool("mp4")
		if err != nil {
			return err
		}

		if statusmp3 {
			log.Println("Loading ....")
			if err := playMP3(args[0]); err != nil {
				log.Fatal(err)
			}
		}

		if statusmp4 {
			playMP4(args)
		}

		if !statusmp3 && !statusmp4 {
			prompt := promptui.Select{
				Label: "Select Day",
				Items: []string{"aor.mp3", "testing.mp3"},
			}

			_, result, err := prompt.Run()

			prompt = promptui.Select{
				Label: "Select Loop",
				Items: []string{"yes", "no"},
			}
			_, loop, err := prompt.Run()
			if err != nil {
				return err
			}

			if err != nil {
				return err
			}

			loopController(result, loop)

		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().BoolP("mp3", "m", false, "Play MP3")
	playCmd.Flags().BoolP("mp4", "c", false, "Play MP4")
	//playCmd.Flags().String("filename", "f", "Name of the .mp3 file")
}
