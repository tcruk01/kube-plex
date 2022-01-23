// ffmpeg package contains helpers to unescape and escape FFmpeg strings
package ffmpeg

import "regexp"

// Unescape takes a string that has FFmpeg compatible escape characters and returns the raw string
func Unescape(s string) string {
	re := regexp.MustCompile(`('(?P<s>[^']*)'|\\(?P<c>.))`)
	return string(re.ReplaceAll([]byte(s), []byte("${s}${c}")))
}

// Escape adds FFmpeg compatible escaping to a string
func Escape(s string) string {
	re := regexp.MustCompile(`(\\|')`)
	return string(re.ReplaceAll([]byte(s), []byte("\\$1")))
}

// RemoveOptsCrf removes crf option from given string
//func RemoveOptsCrf(s string) string {
//        re := regexp.MustCompile(`/-crf:[0-9]\S[0-9]\d{1,2}`)
//        return string(re.ReplaceAll([]byte(s), "")
//}
