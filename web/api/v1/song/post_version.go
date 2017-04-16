package song

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/coral/swapend/db"
	"github.com/gin-gonic/gin"
)

func postVersion(c *gin.Context) {
	song := c.Request.FormValue("song")
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := header.Filename
	fmt.Println(header.Filename)
	tempFilename := "tmp/" + filename + ".png"
	out, err := os.Create(tempFilename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	newFilename := "store/" + db.CreateVersion(song)
	os.Rename(tempFilename, newFilename)
}
