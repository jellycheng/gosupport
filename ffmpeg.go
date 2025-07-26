package gosupport

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// FfmpegMergeFiles([]string{"a[0].mp4", "a[1].m4a"}, "xxx.mp4")
func FfmpegMergeFiles(srcSlice []string, dist string) error {
	cmdArgs := []string{
		"-y",
	}
	for _, path := range srcSlice {
		cmdArgs = append(cmdArgs, "-i", path)
	}
	cmdArgs = append(cmdArgs, "-c:v", "copy", "-c:a", "copy", dist)
	binFile, err := LookPath("ffmpeg")
	if err != nil {
		return err
	}
	bufErr := new(bytes.Buffer)
	cmd := exec.Command(binFile, cmdArgs...)
	cmd.Stderr = bufErr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%s\n%s", err.Error(), bufErr.String())
	}
	return nil
}

func FfmpegMerge2Mp4(srcSlice []string, dist string, txtFile string) error {
	txtFilePath := txtFile + ".txt"
	mergeFile, _ := os.Create(txtFilePath)
	for _, path := range srcSlice {
		_, _ = mergeFile.Write([]byte(fmt.Sprintf("file '%s'\n", path)))
	}
	_ = mergeFile.Close()
	cmdArgs := []string{
		"-y",
		"-f", "concat",
		"-safe", "0",
		"-i", txtFilePath,
		"-c", "copy",
		"-bsf:a", "aac_adtstoasc",
		dist,
	}
	binFile, err := LookPath("ffmpeg")
	if err != nil {
		return err
	}
	bufErr := new(bytes.Buffer)
	cmd := exec.Command(binFile, cmdArgs...)
	cmd.Stderr = bufErr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%s\n%s", err.Error(), bufErr.String())
	}
	return nil
}
