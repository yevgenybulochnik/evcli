package cmd

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/cheggaaa/pb"
	"github.com/spf13/cobra"
	"github.com/yevgenybulochnik/evcli/core"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download image",
	Long:  `Download specific image given image name, (ex: ubuntu1804)`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Please specify a single image name\n\n" + getAvailableImageNames())
		}
		if _, ok := core.ImageDict[args[0]]; !ok {
			return errors.New("Image name not found\n\n" + getAvailableImageNames())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pool_name, _ := cmd.Flags().GetString("pool")
		downloadImage(args[0], pool_name)
	},
}

func getAvailableImageNames() string {
	var imgNames []string

	availableImagesText := "Available Images:\n"

	for imgKey := range core.ImageDict {
		imgNames = append(imgNames, imgKey)
	}

	sort.Strings(imgNames)

	for _, imgName := range imgNames {
		availableImagesText += "\t" + imgName + "\n"
	}

	return availableImagesText
}

func downloadImage(imgName string, pool_name string) {
	conn := core.Connect()

	pool, err := conn.LookupStoragePoolByName(pool_name)
	if err != nil {
		panic(err)
	}

	_, poolPath := core.GetPoolInfo(pool)

	client := grab.NewClient()

	req, _ := grab.NewRequest(poolPath, core.ImageDict[imgName])

	fmt.Printf("Downloading %v..\n", req.URL())
	resp := client.Do(req)

	bar := pb.New(int(resp.Size))
	bar.Start()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			bar.SetCurrent(resp.BytesComplete())
		case <-resp.Done:
			fmt.Println("Completed")
			return
		}

	}
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("pool", "p", "images", "specify storage pool to save image")
}
