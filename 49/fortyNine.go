package main

import (
	"os"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"fmt"
	"net/http"
	"io/ioutil"
		"strings"
	"regexp"
	"github.com/therecipe/qt/gui"
)

const UrlToService= "http://api.flickr.com/services/feeds/photos_public.gne?tags="

var searchTextEdit *widgets.QLineEdit
var imageLabels = make([]*widgets.QLabel , 0)
var imageLabelLayout  *widgets.QVBoxLayout
var tagDescriptionLabel *widgets.QLabel


//var imageLinkRegExp = regexp.MustCompile(`.*img src=&quot;http(s)*://([^/]+)(.*\\.jpg)&quot;.*`)
//var imageLinkRegExp = regexp.MustCompile(`.*(.*\.jpg).*`)
var imageLinkRegExp = regexp.MustCompile(`.*img src=&quot;(http://[^/]+.*\.jpg)&quot;.*`)
//&lt;p&gt;&lt;a href=&quot;http://www.flickr.com/photos/10924140@N02/30071663698/&quot; title=&quot;Alligator&quot;&gt;&lt;img src=&quot;http://farm1.staticflickr.com/859/30071663698_8785c7f1d3_m.jpg&quot; width=&quot;240&quot; height=&quot;160&quot; alt=&quot;Alligator&quot; /&gt;&lt;/a&gt;&lt;/p&gt;
func main() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := widgets.NewQMainWindow(nil, 0)
	window.Resize2(400, 400)
	window.SetWindowTitle("49")
	centerWindow(window)

	// create a regular widget

	// Create central widget.
	widget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(widget)

	// Create the widgets.
	searchStringLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	searchStringLabel.SetText("Search string")
	searchTextEdit = widgets.NewQLineEdit(nil)
	searchButton := widgets.NewQPushButton2("&Search", nil)
	tagDescriptionLabel = widgets.NewQLabel(nil, core.Qt__Widget)

	searchLayout := widgets.NewQHBoxLayout()
	searchLayout.AddWidget(searchStringLabel, 0, core.Qt__AlignLeft)
	searchLayout.AddWidget(searchTextEdit, 1, core.Qt__AlignHCenter)
	searchLayout.AddWidget(searchButton, 0, core.Qt__AlignRight)

	imageLabelLayout = widgets.NewQVBoxLayout()

	// Create scroll area to put the input widgets and the images in.
	groupBox := widgets.NewQGroupBox(nil)
	groupBox.SetLayout(imageLabelLayout)
	scrollArea := widgets.NewQScrollArea(nil)
	scrollArea.SetWidget(groupBox)
	scrollArea.SetWidgetResizable(true)
	scrollArea.SetMinimumSize(core.NewQSize2(300, 300)) // This is not very pretty...

	// Add the scroll area to a layout...
	mainLayout := widgets.NewQVBoxLayout()
	mainLayout.AddLayout(searchLayout, 0)
	mainLayout.AddWidget(tagDescriptionLabel, 0, core.Qt__AlignLeft)
	mainLayout.AddWidget(scrollArea, 1, core.Qt__AlignCenter)

	// A set layout of the central widget.
	widget.SetLayout(mainLayout)

	// Connect the button to a slot.
	searchButton.ConnectClicked(searchButtonClicked)


	// make the window visible
	window.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}

func centerWindow(window *widgets.QMainWindow) {
	qr := window.FrameGeometry()
	cp := widgets.QApplication_Desktop().AvailableGeometry(window).Center()
	qr.MoveCenter(cp)
	window.Move(qr.TopLeft())
}

func searchButtonClicked(_ bool) {
	searchText := searchTextEdit.Text()
	fetchDataFromFlickrAndDisplayIt(searchText)
}

func fetchDataFromFlickrAndDisplayIt(searchText string) {
	removeAllImageLabels()
	tagDescriptionLabel.SetText("Photos about " + searchText)

	// Load feed and extract URLs of the images to fetch from it.
	feed, _ := loadFeedFromFlickr(searchText)
	imageUrls := extractImageUrls(feed)

	// Load each image concurrently.
	imageChan := make(chan []byte, len(imageUrls))
	for _, imageUrl := range imageUrls {
		go func(urlOfImageToFetch string, outdataChan chan<- []byte) {
			fmt.Println("Loading:", urlOfImageToFetch)
			image, _ := loadImage(urlOfImageToFetch)
			outdataChan <- image

		}(imageUrl, imageChan)
	}

	// Display the desired number of images. Display each image as soon as it is loaded.
	// The update of the UI must be made in the main thread.
	for i:= 0; i < len(imageUrls); i++ {
		displayPictureInLabel(<-imageChan)
	}
}


func displayPictureInLabel(image []byte) {
	imageLabel := createImageLabel(image)
	imageLabelLayout.AddWidget(imageLabel, 0, core.Qt__AlignHCenter)
	imageLabels = append(imageLabels, imageLabel)
}

func removeAllImageLabels() {
	for _, imageLabel := range imageLabels {
		imageLabelLayout.RemoveWidget(imageLabel)
	}

	imageLabels = make([]*widgets.QLabel , 0)
}

func loadFeedFromFlickr(searchText string) (feed string, errorToReturn error) {
	url := UrlToService + searchText
	fmt.Printf("Fetching data from %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data from URL! Error: %v", err)
		return
	} else {
		// Make sure the connection is closed.
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(contents)
			fmt.Fprintf(os.Stderr,"Could not read from response! Error: %v", err)
			errorToReturn = err
			return
		}

		feed = string(contents)
	}
	return

}
func extractImageUrls(feed string) (imageUrls []string){
	imageUrls = make([]string, 0)
	lines := strings.Split(feed, "\n")
	for _, line := range lines {
		if imageLinkRegExp.MatchString(line) {
			groups := imageLinkRegExp.FindAllStringSubmatch(line, -1)
			if len(groups) > 0 && len(groups[0]) > 1 {
				imageUrl := groups[0][1]
				imageUrls = append(imageUrls, imageUrl)
			}
		}
	}

	return
}

func loadImage(imageUrl string) (image []byte, errorToReturn error){
	fmt.Printf("Fetching data from %s\n", imageUrl)
	resp, err := http.Get(imageUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read data from URL! Error: %v", err)
		return
	}
	// Make sure the connection is closed.
	defer resp.Body.Close()

	image, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Could not read from response! Error: %v", err)
		errorToReturn = err
	}

	return
}

func createImageLabel(imageData []byte) (imageLabel *widgets.QLabel){
	pixmap := gui.NewQPixmap()
	pixmap.LoadFromData(string(imageData), uint(len(imageData)), "jpg", core.Qt__AutoColor)

	imageLabel = widgets.NewQLabel(nil, core.Qt__Widget)
	imageLabel.SetPixmap(pixmap)

	return
}




