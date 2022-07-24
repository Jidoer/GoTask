package runer
import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)
func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
func Cmd(cmd string) string{
	err, out, errout := Shellout(cmd)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	/*
	   fmt.Println("--- stdout ---")
	   fmt.Println(out)
	   fmt.Println("--- stderr ---")
	*/
	fmt.Println(out)
	fmt.Println(errout)
	return out

}