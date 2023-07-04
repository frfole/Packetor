//go:build generate

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

const (
	version             = "1.20.1"
	versionsManifestUrl = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
)

type (
	block struct {
		Name    string
		ProtoID int32
		States  []int32
	}
	blocks struct {
		Blocks []block
	}
)

//go:generate go run $GOFILE
func main() {
	//log.Println("creating temp directory")
	//tempDir, err := os.MkdirTemp("", "packetor")
	//if err != nil {
	//	log.Fatal("failed to create temp directory", err)
	//	return
	//}
	//log.Printf("Temp directory created at %v\n", tempDir)
	//defer os.RemoveAll(tempDir)
	//
	//log.Printf("Getting version manifest url for version %v\n", version)
	//versionManifestUrl, err := getVersionManifestUrl(version)
	//if err != nil {
	//	log.Fatal("Failed to get version manifest url", err)
	//	return
	//}
	//
	//log.Println("Getting server jar url")
	//serverJarUrl, err := getServerJarUrl(versionManifestUrl)
	//if err != nil {
	//	log.Fatal("Failed to get server jar url", err)
	//	return
	//}
	//
	//log.Printf("Downloading server jar from %v\n", serverJarUrl)
	//err = downloadServerJar(serverJarUrl, tempDir)
	//if err != nil {
	//	log.Fatal("Failed to create temp directory", err)
	//	return
	//}
	//
	//log.Println("Extracting data from server jar")
	//err = extractServerData(tempDir)
	//if err != nil {
	//	log.Fatal("Failed to extract data", err)
	//	return
	//}

	tempDir := "/tmp/packetor2597211569"
	log.Println("Generating items data")
	err := genItems(tempDir)
	if err != nil {
		log.Fatal("Failed to generate items.go", err)
	}
	log.Println("Generating blocks data")
	err = genBlocks(tempDir)
	if err != nil {
		log.Fatal("Failed to generate items.go", err)
	}
	log.Println("Done")
}

func getVersionManifestUrl(version string) (string, error) {
	resp, err := http.Get(versionsManifestUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Versions []struct {
			ID  string `json:"id"`
			Url string `json:"url"`
		} `json:"versions"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	for _, s := range data.Versions {
		if s.ID == version {
			return s.Url, nil
		}
	}
	return "", fmt.Errorf("failed to get manifest for given version")
}

func getServerJarUrl(versionManifestUrl string) (string, error) {
	resp, err := http.Get(versionManifestUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Downloads struct {
			Server struct {
				Url string `json:"url"`
			} `json:"server"`
		} `json:"downloads"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	return data.Downloads.Server.Url, nil
}

func downloadServerJar(serverJarUrl string, tempDir string) error {
	resp, err := http.Get(serverJarUrl)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filepath.Join(tempDir, "server.jar"), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func extractServerData(tempDir string) error {
	cmd := exec.Command("java", "-DbundlerMainClass=net.minecraft.data.Main", "-jar", "server.jar", "--reports")
	cmd.Dir = tempDir
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}

func genItems(tempDir string) error {
	regFile, err := os.Open(filepath.Join(tempDir, "generated", "reports", "registries.json"))
	if err != nil {
		return err
	}
	defer regFile.Close()
	var data struct {
		Items struct {
			Entries map[string]struct {
				ProtoID int32 `json:"protocol_id"`
			} `json:"entries"`
		} `json:"minecraft:item"`
	}
	if err = json.NewDecoder(regFile).Decode(&data); err != nil {
		return err
	}

	items := make([]struct {
		ProtoID int32
		Name    string
	}, len(data.Items.Entries))
	idx := 0
	for key, value := range data.Items.Entries {
		items[idx] = struct {
			ProtoID int32
			Name    string
		}{ProtoID: value.ProtoID, Name: key}
		idx++
	}

	tmpl, err := readFile("items.go.tmpl")
	if err != nil {
		return err
	}

	file, err := os.Create("items.go")
	if err != nil {
		return err
	}
	defer file.Close()
	err = template.Must(template.New("").Parse(tmpl)).Execute(file, items)
	if err != nil {
		return err
	}
	return nil
}

func genBlocks(tempDir string) error {
	regFile, err := os.Open(filepath.Join(tempDir, "generated", "reports", "registries.json"))
	if err != nil {
		return err
	}
	defer regFile.Close()
	var regData struct {
		Blocks struct {
			Entries map[string]struct {
				ProtoID int32 `json:"protocol_id"`
			} `json:"entries"`
		} `json:"minecraft:block"`
	}
	if err = json.NewDecoder(regFile).Decode(&regData); err != nil {
		return err
	}
	blocksFile, err := os.Open(filepath.Join(tempDir, "generated", "reports", "blocks.json"))
	if err != nil {
		return err
	}
	defer blocksFile.Close()
	var blocksData map[string]struct {
		States []struct {
			Id int32 `json:"id"`
		} `json:"states"`
	}
	if err = json.NewDecoder(blocksFile).Decode(&blocksData); err != nil {
		return err
	}

	out := blocks{}

	out.Blocks = make([]block, len(regData.Blocks.Entries))
	idx := 0
	for key, value := range regData.Blocks.Entries {
		states := make([]int32, len(blocksData[key].States))
		for i := range states {
			states[i] = blocksData[key].States[i].Id
		}
		out.Blocks[idx] = block{
			Name:    key,
			ProtoID: value.ProtoID,
			States:  states,
		}
		idx++
	}

	tmpl, err := readFile("blocks.go.tmpl")
	if err != nil {
		return err
	}

	file, err := os.Create("blocks.go")
	if err != nil {
		return err
	}
	defer file.Close()
	err = template.Must(template.New("").Parse(tmpl)).Execute(file, out)
	if err != nil {
		return err
	}
	return nil
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	all, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	file.Close()
	return string(all), nil
}
