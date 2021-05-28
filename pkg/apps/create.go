package apps

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
)

func Create(appNames []string, flaggedPlatforms []string, shell string, appsDirPath string, continueOnError bool) {
	platforms := ResolvePlatforms(flaggedPlatforms)

	if _, err := os.Stat(appsDirPath); os.IsNotExist(err) {
		log.Fatalln("Apps folder at '" + appsDirPath + "' doesn't exist.\nPlease run 'eac init' first.")
	} else if err != nil {
		log.Fatalln("There was an error while accessing the appsDir at '" + appsDirPath + "':\n" + err.Error())
	}

	for _, appName := range appNames {
		appPath := path.Join(appsDirPath, appName)
		if _, err := os.Stat(appPath); os.IsNotExist(err) { // if folder doesn't exist yet
			err := os.Mkdir(appPath, os.ModePerm)
			if err != nil {
				log.Fatalln("There was an error while creating the app at '" + appPath + "':\n" + err.Error())
			}
			log.Println("Created '" + appPath + "' folder.")
		} else if err == nil { // if folder does exist
			fmt.Println("App '" + appName + "' does already exist.")
		} else {
			log.Fatalln("There was an error while accessing the app at '" + appPath + "':\n" + err.Error())
		}

		for _, platform := range platforms {
			platformPath := path.Join(appPath, platform)
			if _, err := os.Stat(platformPath); os.IsNotExist(err) {
				err := os.Mkdir(platformPath, os.ModePerm) // ignore errors
				if err != nil {
					log.Fatalln("There was an error while creating the platform '" + platform + "' for app '" + appPath + "' at '" + platformPath + "':\n" + err.Error())
				}
				log.Println("Created '" + platformPath + "' folder.")
			} else if err == nil {
				fmt.Println("Platform '" + platform + "' for app '" + appName + "' does already exist.")
			} else {
				log.Fatalln("There was an error while accessing the platform '" + platform + "' for app '" + appPath + "' at '" + platformPath + "':\n" + err.Error())
			}

			platformDemoFiles := demoFiles[platform]

			for filename, fileContent := range platformDemoFiles {
				f, err := os.Create(path.Join(platformPath, filename))
				if err != nil {
					log.Fatalln("There was an error while creating the file '" + filename + "' at '" + platformPath + "':\n" + err.Error())
				}
				defer f.Close()

				w := bufio.NewWriter(f)

				_, err = fmt.Fprintf(w, fileContent, appName)
				if err != nil {
					log.Fatalln("There was an error while writing to the file '" + filename + "' at '" + platformPath + "':\n" + err.Error())
				}

				w.Flush()
				log.Println("Created '" + path.Join(platformPath, filename) + "' file.")
			}
		}
	}
}
