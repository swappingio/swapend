package song

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func postSong(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("tmp/" + filename + ".png")
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
}
