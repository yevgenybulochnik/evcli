package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
			return errors.New("Please specify a single image name\n\n" + core.GetAvailableImageNames())
		}
		if _, ok := core.ImageDict[args[0]]; !ok {
			return errors.New("Image name not found\n\n" + core.GetAvailableImageNames())
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		core.CheckOrCreateConfigDir()
		if core.ProfileExists(args[0]) {
			fmt.Printf("Profile name %v already exists\n", args[0])
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		pool_name, _ := cmd.Flags().GetString("pool")
		downloadImage(args[0], pool_name)
	},
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

	imageFileName := filepath.Base(core.ImageDict[imgName])

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
			core.AddProfile(imgName, imageFileName)
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("pool", "p", "images", "specify storage pool to save image")
}
