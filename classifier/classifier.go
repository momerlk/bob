package classifier

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"gocv.io/x/gocv"
	"time"
)

func DisplayDetection(camId int  , cascadeFile string) {
	displayText := "Human";

    cam , err := gocv.VideoCaptureDevice(camId);
    if err != nil {
        fmt.Println(err);
        os.Exit(1); 
    }
    defer cam.Close();

    window := gocv.NewWindow("Face Detection - By M.Omer Ali");
    defer window.Close();

    camImg := gocv.NewMat();
    defer camImg.Close();

    classifier := gocv.NewCascadeClassifier();
    defer classifier.Close();

    if !classifier.Load(cascadeFile) {
        fmt.Println("Error loading Cascade Classifier 'file :" , cascadeFile);
        os.Exit(1);
    }

    for {
        if !cam.Read(&camImg) {
            fmt.Println("Error reading from webcam!");
            os.Exit(1);
        }

        if camImg.Empty() {
            fmt.Println("Image is empty");
            continue;
        }
        
        rects := classifier.DetectMultiScale(camImg);
        
        rectColor := color.RGBA{255 , 0 , 0 , 0};

        fontFace := gocv.FontHersheySimplex;

        for _ , rect := range rects {
            gocv.Rectangle(&camImg , rect , rectColor , 4);
            textSize := gocv.GetTextSize(displayText , fontFace , 1.2 , 2);

            imagePoint := image.Pt(rect.Min.X+(rect.Min.X/2)-(textSize.X/2), rect.Min.Y-2);

            gocv.PutText(&camImg , displayText , imagePoint , fontFace , 1.2 , rectColor , 2);
        }

        window.IMShow(camImg);
        if window.WaitKey(1) >= 0 {
            break;
        }
    }
}

func DisplayDetectionTimeout(camId int , cascadeFile string , timeLimit int){
	displayText := "Human";

    cam , err := gocv.VideoCaptureDevice(camId);
    if err != nil {
        fmt.Println(err);
        os.Exit(1); 
    }
    defer cam.Close();

    window := gocv.NewWindow("Face Detection - By M.Omer Ali");
    defer window.Close();

    camImg := gocv.NewMat();
    defer camImg.Close();

    classifier := gocv.NewCascadeClassifier();
    defer classifier.Close();

    if !classifier.Load(cascadeFile) {
        fmt.Println("Error loading Cascade Classifier 'file :" , cascadeFile);
        os.Exit(1);
    }

    for timeout := time.After(time.Duration(timeLimit) * time.Second); ;{
        if !cam.Read(&camImg) {
            fmt.Println("Error reading from webcam!");
            os.Exit(1);
        }

        if camImg.Empty() {
            continue;
        }
        
        rects := classifier.DetectMultiScale(camImg);
        fmt.Println("No. of people in image :" ,len(rects));
        
        rectColor := color.RGBA{255 , 0 , 0 , 0};

        fontFace := gocv.FontHersheySimplex;

        for _ , rect := range rects {
            gocv.Rectangle(&camImg , rect , rectColor , 4);
            textSize := gocv.GetTextSize(displayText , fontFace , 1.2 , 2);

            imagePoint := image.Pt(rect.Min.X+(rect.Min.X/2)-(textSize.X/2), rect.Min.Y-2);

            gocv.PutText(&camImg , displayText , imagePoint , fontFace , 1.2 , rectColor , 2);
        }

		window.IMShow(camImg);
		if window.WaitKey(1) >= 0 {
            break;
        }

		select {
		case <-timeout:
			return;
		default:
		}
    }
}

func DetectPeople(cameraId int , classifierFile string) (numPeople int) {
    cam , err := gocv.VideoCaptureDevice(cameraId);
    if err != nil {
        fmt.Println(err);
        os.Exit(1); 
    }
    defer cam.Close();

    img := gocv.NewMat();
    defer img.Close();

    classifier := gocv.NewCascadeClassifier();
    defer classifier.Close();

    if !classifier.Load(classifierFile) {
        fmt.Println("Error loading classifier File :" , classifierFile); 
    }

    for i := 0;i < 4;i++{
        if !cam.Read(&img) {
            fmt.Println("error reading image from webcam");
        }
        if img.Empty() {
            continue;
        } 

        rects := classifier.DetectMultiScale(img);
        
        numPeople = len(rects);
    }

    return numPeople;
}

func DetectPeopleTimeout(cameraId int , classifierFile string , timeLimit int) (numPeople int){
	cam , err := gocv.VideoCaptureDevice(cameraId);
    if err != nil {
        fmt.Println(err);
        os.Exit(1); 
    }
    defer cam.Close();

    img := gocv.NewMat();
    defer img.Close();

    classifier := gocv.NewCascadeClassifier();
    defer classifier.Close();

    if !classifier.Load(classifierFile) {
        fmt.Println("Error loading classifier File :" , classifierFile); 
    }

    for timeout := time.After(time.Duration(timeLimit) * time.Second); ;{
        if !cam.Read(&img) {
            fmt.Println("error reading image from webcam");
        }
        if img.Empty() {
            continue;
        } 

        rects := classifier.DetectMultiScale(img);
        
        numPeople = len(rects);

		select{
		case <-timeout:
			return numPeople;
		default:
		}
    }
}