package transcoder

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/swappingio/swapend/pkg/config"
)

var (
	ActiveTranscoders = make(chan bool, 1)
)

type conversionJob struct {
	SourceFilename string
	OutputFilename string

	EncodeAAC  bool
	AACBitrate string

	EncodeOpus  bool
	OpusBitrate string

	EncodeMP3   bool
	MP3VBRLevel string
}

func StartJob(job conversionJob) error {

	ActiveTranscoders <- true

	return ConvertFile(job)
}

func NewConversionJob() conversionJob {
	newJob := conversionJob{}
	newJob.EncodeAAC = true
	newJob.AACBitrate = "128k"
	newJob.EncodeOpus = true
	newJob.OpusBitrate = "128k"
	newJob.EncodeMP3 = true
	newJob.MP3VBRLevel = "4"

	return newJob
}

func prepareJob(job conversionJob, threads int) []string {

	FFmpegArgs := []string{
		"-y",
		"-threads", strconv.Itoa(threads),
		"-i", job.SourceFilename,
	}

	if job.EncodeAAC {
		aacArgs := []string{
			"-acodec", "aac",
			"-b:a", job.AACBitrate,
			job.OutputFilename + ".aac",
		}
		FFmpegArgs = append(FFmpegArgs, aacArgs...)
	}

	if job.EncodeMP3 {
		mp3Args := []string{
			"-acodec", "libmp3lame",
			"-qscale:a", job.MP3VBRLevel,
			job.OutputFilename + ".mp3",
		}
		FFmpegArgs = append(FFmpegArgs, mp3Args...)
	}

	if job.EncodeOpus {
		opusArgs := []string{
			"-acodec", "libopus",
			"-b:a", job.OpusBitrate,
			job.OutputFilename + ".opus",
		}
		FFmpegArgs = append(FFmpegArgs, opusArgs...)
	}

	return FFmpegArgs

}

func ConvertFile(job conversionJob) (err error) {
	defer func() {
		<-ActiveTranscoders
	}()

	c := config.GetConfig()

	cmd := exec.Command("ffmpeg", prepareJob(job, c.Transcoder.Threads)...)

	stderr, _ := cmd.StderrPipe()
	in := bufio.NewScanner(stderr)

	if err = cmd.Start(); err != nil {
		return err
	}

	for in.Scan() {
		if c.Transcoder.Debug {
			fmt.Println(in.Text())
		}
	}

	err = cmd.Wait()

	return err
}
