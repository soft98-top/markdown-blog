package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"html/template"
	"regexp"
	"runtime"
	"strings"
)

// FormatAppVersion 格式化应用版本信息
func FormatAppVersion(appVersion, GitCommit, BuildDate string) (string, error) {
	content := `
   Version: {{.Version}}
Go Version: {{.GoVersion}}
Git Commit: {{.GitCommit}}
     Built: {{.BuildDate}}
   OS/ARCH: {{.GOOS}}/{{.GOARCH}}
`
	tpl, err := template.New("version").Parse(content)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, map[string]string{
		"Version":   appVersion,
		"GoVersion": runtime.Version(),
		"GitCommit": GitCommit,
		"BuildDate": BuildDate,
		"GOOS":      runtime.GOOS,
		"GOARCH":    runtime.GOARCH,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), err
}

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Sha1(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// IsInSlice 判断目标字符串是否是在切片中
func IsInSlice(slice []string, s string) bool {
	if len(slice) == 0 {
		return false
	}

	isIn := false
	for _, f := range slice {
		// 增加正则支持, r'正则表达式'
		if strings.HasPrefix(f, "r'") && strings.HasSuffix(f, "'") {
			regexStr := strings.TrimPrefix(f, "r'")
			regexStr = strings.TrimSuffix(regexStr, "'")
			if ok, _ := regexp.Match(regexStr, []byte(s)); ok {
				isIn = true
				break
			}
		} else if f == s {
			isIn = true
			break
		}
	}

	return isIn
}
