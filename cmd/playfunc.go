package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	colgate "github.com/tcolgate/mp3"
)

func playMP3(args string) error {
	if len(args) == 0 {
		return errors.New("Please insert the .mp3 filename")
	}

	duration, err := Duration(args)
	if err != nil {
		return err
	}

	f, err := os.Open(args)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()

	defer p.Close()

	fmt.Printf("Duration is %v seconds\n", duration)
	fmt.Printf("Length: %d[bytes]\n", d.Length())

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	defer s.Stop() // Start the spinner

	//time.Sleep(time.Duration() * time.Second) // Run for some time to simulate work
	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

func playMP4(args []string) {
	for _, ival := range args {
		fmt.Println("Playing MP4 of", ival, "...")
	}
}

func Duration(filename string) (float64, error) {
	t := 0.0

	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	var dframe colgate.Frame
	dcolgate := colgate.NewDecoder(f)
	skipped := 0

	for {
		if err := dcolgate.Decode(&dframe, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		t = t + dframe.Duration().Seconds()
	}

	return t, nil
}
