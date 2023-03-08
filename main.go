package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/fhs/gompd/mpd"
	"github.com/getlantern/systray"
)

type Station struct {
	Name string
	URL  string
}

func main() {
	iconFile, err := os.Open("./icons/icon.png")
	if err != nil {
		panic(err)
	}
	defer iconFile.Close()
	iconImage, _, err := image.Decode(iconFile)
	if err != nil {
		panic(err)
	}
	var iconBuf bytes.Buffer
	err = png.Encode(&iconBuf, iconImage)
	if err != nil {
		panic(err)
	}
	iconBytes := iconBuf.Bytes()

	systray.SetIcon(iconBytes)
	// Connect to MPD server
	client, err := mpd.Dial("tcp", "localhost:6600")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Load icons from files
	playIconFile, err := os.Open("icons/play.png")
	if err != nil {
		panic(err)
	}
	defer playIconFile.Close()
	playIconImage, _, err := image.Decode(playIconFile)
	if err != nil {
		panic(err)
	}

	pauseIconFile, err := os.Open("icons/pause.png")
	if err != nil {
		panic(err)
	}
	defer pauseIconFile.Close()
	pauseIconImage, _, err := image.Decode(pauseIconFile)
	if err != nil {
		panic(err)
	}

	previousIconFile, err := os.Open("icons/left.png")
	if err != nil {
		panic(err)
	}
	defer previousIconFile.Close()
	previousIconImage, _, err := image.Decode(previousIconFile)
	if err != nil {
		panic(err)
	}

	nextIconFile, err := os.Open("icons/right.png")
	if err != nil {
		panic(err)
	}
	defer nextIconFile.Close()
	nextIconImage, _, err := image.Decode(nextIconFile)
	if err != nil {
		panic(err)
	}

	quitIconFile, err := os.Open("icons/quit.png")
	if err != nil {
		panic(err)
	}
	defer quitIconFile.Close()
	quitIconImage, _, err := image.Decode(quitIconFile)
	if err != nil {
		panic(err)
	}

	// Convert icon images to []byte
	var playIconBuf bytes.Buffer
	err = png.Encode(&playIconBuf, playIconImage)
	if err != nil {
		panic(err)
	}
	playIconBytes := playIconBuf.Bytes()

	var pauseIconBuf bytes.Buffer
	err = png.Encode(&pauseIconBuf, pauseIconImage)
	if err != nil {
		panic(err)
	}
	pauseIconBytes := pauseIconBuf.Bytes()

	var nextIconBuf bytes.Buffer
	err = png.Encode(&nextIconBuf, nextIconImage)
	if err != nil {
		panic(err)
	}
	nextIconBytes := nextIconBuf.Bytes()

	var previousIconBuf bytes.Buffer
	err = png.Encode(&nextIconBuf, previousIconImage)
	if err != nil {
		panic(err)
	}
	previousIconBytes := previousIconBuf.Bytes()

	var quitIconBuf bytes.Buffer
	err = png.Encode(&quitIconBuf, quitIconImage)
	if err != nil {
		panic(err)
	}
	quitIconBytes := quitIconBuf.Bytes()

	// Create systray menu items with icons
	playPause := systray.AddMenuItem("Play/Pause", "Toggle play/pause")
	playPause.SetIcon(playIconBytes)

	previous := systray.AddMenuItem("Previous", "Return to the previous track")
	previous.SetIcon(previousIconBytes)

	next := systray.AddMenuItem("Next", "Skip to the next track")
	next.SetIcon(nextIconBytes)

	quit := systray.AddMenuItem("Quit", "Quit the program")
	quit.SetIcon(quitIconBytes)

	// Create stations menu item
	stations := systray.AddMenuItem("Stations", "Select a station to play")

	// Define list of stations
	stationList := []Station{
		{Name: "Mother Earth Radio", URL: "http://server9.streamserver24.com:18900/motherearth.mp3"},
		{Name: "Dance Wave!", URL: "http://dancewave.online/dance.mp3"},
		{Name: "Radio Mast", URL: "http://ingest-mia.radiomast.io/8a760c25-fb95-4fcb-9c0e-ca0f269a7360"},
	}

	// Create menu items for each station
	for _, station := range stationList {
		menuItemTitle := station.Name
		menuItem := stations.AddSubMenuItem(menuItemTitle, "")
		fmt.Println("Creating menu item:", menuItemTitle)
		// Define closure for the menu item click handler
		func(url string) {
			menuItem.ClickedCh = make(chan struct{})
			go func() {
				defer close(menuItem.ClickedCh)
				for range menuItem.ClickedCh {
					// Play the selected station
					err := client.Clear()
					if err != nil {
						// handle error
						continue
					}
					err = client.Add(url)
					if err != nil {
						// handle error
						continue
					}
					err = client.Play(-1)
					if err != nil {
						// handle error
						continue
					}
					fmt.Println("Playing station:", menuItemTitle)
				}
			}()
		}(station.URL)
	}
	fmt.Println("Stations menu items created")

	// Handle systray menu item clicks
	go func() {
		for {
			select {
			case <-playPause.ClickedCh:
				status, err := client.Status()
				if err != nil {
					// handle error
					continue
				}
				state, ok := status["state"]
				if !ok {
					// handle error
					continue
				}
				if state == "play" {
					client.Pause(true)
					playPause.SetIcon(pauseIconBytes)
				} else {
					client.Pause(false)
					playPause.SetIcon(playIconBytes)
				}
			case <-previous.ClickedCh:
				client.Previous()
			case <-next.ClickedCh:
				client.Next()
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
	// Start systray
	systray.Run(nil, nil)
}
