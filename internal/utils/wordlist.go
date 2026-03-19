package utils

import (
    "bufio"
    "os"
    "strings"
)

func LoadWordlist(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := strings.TrimSpace(scanner.Text())
        if word != "" && !strings.HasPrefix(word, "#") {
            words = append(words, word)
        }
    }
    return words, scanner.Err()
}
