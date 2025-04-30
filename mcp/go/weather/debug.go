package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 3 {
		fmt.Println("Usage: debug <command> <output_error_file>")
		return
	}

	// 解析命令行参数
	command := os.Args[1]      // 目标程序命令
	errorLogFile := os.Args[2] // 错误日志文件路径

	// 打开错误日志文件
	errorFile, err := os.OpenFile(errorLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to open error log file: %v\n", err)
		return
	}
	defer errorFile.Close()

	// 创建目标程序的执行命令
	cmd := exec.Command(command)

	// 设置标准输入、输出和错误管道
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Failed to create stdin pipe: %v\n", err)
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Failed to create stdout pipe: %v\n", err)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Failed to create stderr pipe: %v\n", err)
		return
	}

	// 启动目标程序
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start command: %v\n", err)
		return
	}

	// 使用 goroutine 收集标准输出
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text()) // 打印到终端
		}
	}()

	// 使用 goroutine 收集标准错误并写入日志文件
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			errorLine := scanner.Text()
			fmt.Fprintln(errorFile, errorLine) // 写入错误日志文件
			fmt.Println("Error:", errorLine)   // 同时打印到终端
		}
	}()

	// 从标准输入读取用户输入并发送到目标程序
	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Enter input for the program (type 'exit' to quit):")
	for scanner.Scan() {
		input := scanner.Text()
		if strings.TrimSpace(input) == "exit" {
			break
		}
		fmt.Fprintln(stdin, input) // 发送到目标程序的标准输入
	}

	// 关闭标准输入以通知目标程序结束输入
	stdin.Close()

	// 等待目标程序退出
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Command exited with error: %v\n", err)
	}
}
