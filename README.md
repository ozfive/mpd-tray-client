# mpd-tray-client

---

mpd-tray-client is a Go program that provides a system tray menu for playing internet radio stations using the MPD protocol.

## Description
The program uses the [getlantern/systray](https://github.com/getlantern/systray) package to create the system tray icon and menu, and the [fhs/gompd/mpd](https://github.com/fhs/gompd/tree/master/mpd) package to communicate with the MPD server.

mpd-tray-client loads station information and icon images from files and creates menu items for each station with an associated URL. When a station is selected from the menu, mpd-tray-client clears the MPD playlist, adds the station URL to the playlist, and starts playing.

In addition, mpd-tray-client provides menu items for toggling play/pause, skipping to the previous or next track, and quitting the program.

## Installation and Usage

To use mpd-tray-client, you will need to have an MPD server running on your local machine. You can download and install MPD from [musicpd.org](https://www.musicpd.org/).

Follow these steps to install and run mpd-tray-client:

    1. Clone the mpd-tray-client repo: git clone [https://github.com/ozfive/mpd-tray-client.git](https://github.com/ozfive/mpd-tray-client.git)
    2. Change to the mpd-tray-client directory: cd mpd-tray-client
    3. Build the program: go build
    4. Run the program: ./mpd-tray-client

mpd-tray-client will display a system tray icon. Right/Left-click on the icon to open the menu and select a station to play. Use the other menu items to control playback or quit the program. 

Currently the only functionality that works fully is the Play/Pause menu item, the Quit menu item and The Stations drop-down menu item. Further functionality will be added shortly.

## Configuration

mpd-tray-client loads station information from a list defined in the main function. You can modify this list to add, remove, or edit station information.

Icon images for the system tray menu items are loaded from files located in the ./icons/ directory. You can replace these images with your own or modify the code in mpd-tray-client to load images from a different directory.

mpd-tray-client connects to the MPD server using the default address localhost:6600. If your MPD server is running on a different host or port, you can modify the mpd.Dial function call in mpd-tray-client to specify the correct address.

## Image Attribution

The icons used for mpd-tray-client were aquired from [Flaticon](https://www.flaticon.com)
app icon (icon.png): [Music icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/music")

last next icons(right.png, left.png): [Arrows icons created by vectaicon - Flaticon](https://www.flaticon.com/free-icons/arrows)

quit icon (cross.png): [Quit icons created by Dixit Lakhani_02 - Flaticon](https://www.flaticon.com/free-icons/quit)

play icon (play-squared-button): [Start icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/start)

pause icon (pause): [Ui icons created by adiobae - Flaticon](https://www.flaticon.com/free-icons/ui)
## License
This program is licensed under the [MIT License](https://opensource.org/license/mit/). See the LICENSE file for details.
