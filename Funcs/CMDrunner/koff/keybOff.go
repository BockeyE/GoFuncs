package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func kbOff() {
	cmd := exec.Command("cmd", "/C", "sc", "config", "i8042prt", "start=", "disabled")
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
	fmt.Println(">>>>>>>>>>>>>>>>>  end  >>>>>>>>>>>>>>>>")
	fmt.Println("请输入任意退出")
	//读键盘
	reader := bufio.NewReader(os.Stdin)
	//以换行符结束
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
}
func main() {
	kbOff()
}
