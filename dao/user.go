package dao

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var database = make(map[string]string, 10000)

func DownLoad() {
	file, err := os.Open("dao/name")
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		k := strings.TrimRight(line, "\n")
		k = strings.TrimSpace(line)
		line2, err := reader.ReadString('\n')
		v := strings.TrimRight(line2, "\n")
		v = strings.TrimSpace(line2)
		database[k] = v
	}
}

func AddUser(username, password string) {
	file, err := os.OpenFile("dao/name", os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(username + "\n" + password + "\n")
	writer.Flush()
	database[username] = password
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	DownLoad()
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

func ChangePassword(username string, newpassword string) {
	file, err := os.OpenFile("dao/name", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	i := 0
	var k int
	for {
		i++
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		k += len([]byte(line))
		if i%2 == 1 {
			line = strings.TrimRight(line, "\n")
			if strings.Compare(line, username) == 0 {
				line2, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				k2 := len(line2)
				for y := 0; y < k2; y++ {
					file.WriteAt([]byte(" "), int64(k+y))
				}
				file.WriteAt([]byte(newpassword+"\n"), int64(k))
				return
			}
		}
	}
}

func FindPassword(username string) string {
	DownLoad()
	v := database[username]
	v = strings.TrimSpace(v)
	return v
}

func Board(username string, say string) {
	file, err := os.OpenFile("dao/board", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(username + ":" + say + "\n")
	writer.Flush()
}
