package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yourusername/backup-restore-tool/backup"
)

func main() {
	fmt.Println("=== Backup and Restore Tool ===")
	fmt.Println("1. Backup a folder")
	fmt.Println("2. Restore from backup")

	var choice int
	fmt.Print("Enter choice (1 or 2): ")
	fmt.Scanln(&choice)

	reader := bufio.NewReader(os.Stdin)

	if choice == 1 {
		fmt.Print("Enter folder path to backup: ")
		source, _ := reader.ReadString('\n')
		source = sanitize(source)

		fmt.Print("Enter destination zip file (e.g., backup.zip): ")
		dest, _ := reader.ReadString('\n')
		dest = sanitize(dest)

		err := backup.BackupFolder(source, dest)
		if err != nil {
			fmt.Println("Backup failed:", err)
		} else {
			fmt.Println("Backup completed successfully!")
		}
	} else if choice == 2 {
		fmt.Print("Enter zip file to restore from: ")
		zipPath, _ := reader.ReadString('\n')
		zipPath = sanitize(zipPath)

		fmt.Print("Enter destination folder to restore to: ")
		dest, _ := reader.ReadString('\n')
		dest = sanitize(dest)

		err := backup.RestoreBackup(zipPath, dest)
		if err != nil {
			fmt.Println("Restore failed:", err)
		} else {
			fmt.Println("Restore completed successfully!")
		}
	} else {
		fmt.Println("Invalid choice")
	}
}

func sanitize(input string) string {
	return input[:len(input)-1] // remove newline
}
