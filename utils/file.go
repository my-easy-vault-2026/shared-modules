package utils

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileInfo 文件信息結構
type FileInfo struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	MimeType     string    `json:"mime_type"`
	Extension    string    `json:"extension"`
	MD5Hash      string    `json:"md5_hash"`
	SHA256Hash   string    `json:"sha256_hash"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
	IsValid      bool      `json:"is_valid"`
	ErrorMessage string    `json:"error_message,omitempty"`
}

// ValidationRule 文件驗證規則
type ValidationRule struct {
	MaxSize         int64    `json:"max_size"`          // 最大文件大小（字節）
	MinSize         int64    `json:"min_size"`          // 最小文件大小（字節）
	AllowedTypes    []string `json:"allowed_types"`     // 允許的MIME類型
	AllowedExts     []string `json:"allowed_exts"`      // 允許的文件擴展名
	ForbiddenTypes  []string `json:"forbidden_types"`   // 禁止的MIME類型
	ForbiddenExts   []string `json:"forbidden_exts"`    // 禁止的文件擴展名
	CheckMagicBytes bool     `json:"check_magic_bytes"` // 是否檢查文件魔術字節
}

// DownloadOptions 下載選項
type DownloadOptions struct {
	Timeout        time.Duration     `json:"timeout"`         // 下載超時時間
	MaxSize        int64             `json:"max_size"`        // 最大下載大小
	Headers        map[string]string `json:"headers"`         // 自定義請求頭
	UserAgent      string            `json:"user_agent"`      // User-Agent
	FollowRedirect bool              `json:"follow_redirect"` // 是否跟隨重定向
}

// DefaultValidationRule 返回默認驗證規則
func DefaultValidationRule() *ValidationRule {
	return &ValidationRule{
		MaxSize:         10 * 1024 * 1024, // 10MB
		MinSize:         1,                // 1 byte
		AllowedTypes:    []string{"image/jpeg", "image/png", "image/gif", "application/pdf", "text/plain"},
		AllowedExts:     []string{".jpg", ".jpeg", ".png", ".gif", ".pdf", ".txt"},
		CheckMagicBytes: true,
	}
}

// DefaultDownloadOptions 返回默認下載選項
func DefaultDownloadOptions() *DownloadOptions {
	return &DownloadOptions{
		Timeout:        30 * time.Second,
		MaxSize:        50 * 1024 * 1024, // 50MB
		UserAgent:      "eUSD/FileUtil/1.0",
		FollowRedirect: true,
		Headers:        make(map[string]string),
	}
}

// DownloadFile 從URL下載文件到指定路徑
func DownloadFile(ctx context.Context, url, destPath string, options *DownloadOptions) (*FileInfo, error) {
	if options == nil {
		options = DefaultDownloadOptions()
	}

	// 創建HTTP客戶端
	client := &http.Client{
		Timeout: options.Timeout,
	}

	// 如果不跟隨重定向
	if !options.FollowRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	// 創建請求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("創建請求失败: %w", err)
	}

	// 設置請求頭
	req.Header.Set("User-Agent", options.UserAgent)
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// 發送請求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("下載失败: %w", err)
	}
	defer resp.Body.Close()

	// 檢查HTTP狀態碼
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("下載失败，HTTP狀態碼: %d", resp.StatusCode)
	}

	// 檢查內容長度
	if resp.ContentLength > 0 && resp.ContentLength > options.MaxSize {
		return nil, fmt.Errorf("文件過大: %d bytes, 最大允許: %d bytes", resp.ContentLength, options.MaxSize)
	}

	// 創建目標目錄
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return nil, fmt.Errorf("創建目錄失败: %w", err)
	}

	// 創建目標文件
	file, err := os.Create(destPath)
	if err != nil {
		return nil, fmt.Errorf("創建文件失败: %w", err)
	}
	defer file.Close()

	// 限制讀取大小並複製數據
	limitedReader := io.LimitReader(resp.Body, options.MaxSize+1)
	written, err := io.Copy(file, limitedReader)
	if err != nil {
		os.Remove(destPath) // 清理部分下載的文件
		return nil, fmt.Errorf("保存文件失败: %w", err)
	}

	// 檢查是否超過最大大小
	if written > options.MaxSize {
		os.Remove(destPath)
		return nil, fmt.Errorf("文件過大: %d bytes, 最大允許: %d bytes", written, options.MaxSize)
	}

	// 獲取文件信息
	return GetFileInfo(destPath)
}

// ReadFile 讀取文件內容
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// ReadFileWithLimit 讀取文件內容（帶大小限制）
func ReadFileWithLimit(filePath string, maxSize int64) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 獲取文件信息
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// 檢查文件大小
	if stat.Size() > maxSize {
		return nil, fmt.Errorf("文件過大: %d bytes, 最大允許: %d bytes", stat.Size(), maxSize)
	}

	// 讀取文件
	return io.ReadAll(file)
}

// GetFileInfo 獲取文件詳細信息
func GetFileInfo(filePath string) (*FileInfo, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	info := &FileInfo{
		Name:       stat.Name(),
		Path:       filePath,
		Size:       stat.Size(),
		Extension:  strings.ToLower(filepath.Ext(filePath)),
		CreatedAt:  stat.ModTime(), // 在某些系統上這是修改時間
		ModifiedAt: stat.ModTime(),
		IsValid:    true,
	}

	// 檢測MIME類型
	mimeType, err := DetectMimeType(filePath)
	if err != nil {
		info.ErrorMessage = fmt.Sprintf("檢測MIME類型失败: %v", err)
		info.IsValid = false
	} else {
		info.MimeType = mimeType
	}

	// 計算哈希值
	if info.IsValid {
		md5Hash, sha256Hash, err := CalculateHashes(filePath)
		if err != nil {
			info.ErrorMessage = fmt.Sprintf("計算哈希值失败: %v", err)
			info.IsValid = false
		} else {
			info.MD5Hash = md5Hash
			info.SHA256Hash = sha256Hash
		}
	}

	return info, nil
}

// ValidateFile 驗證文件是否符合規則
func ValidateFile(ctx context.Context, filePath string, rule *ValidationRule) (*FileInfo, error) {
	if rule == nil {
		rule = DefaultValidationRule()
	}

	info, err := GetFileInfo(filePath)
	if err != nil {
		return nil, err
	}

	var errors []string

	// 檢查文件大小
	if info.Size > rule.MaxSize {
		errors = append(errors, fmt.Sprintf("文件過大: %d bytes, 最大允許: %d bytes", info.Size, rule.MaxSize))
	}
	if info.Size < rule.MinSize {
		errors = append(errors, fmt.Sprintf("文件過小: %d bytes, 最小要求: %d bytes", info.Size, rule.MinSize))
	}

	// 檢查允許的MIME類型
	if len(rule.AllowedTypes) > 0 {
		allowed := false
		for _, allowedType := range rule.AllowedTypes {
			if info.MimeType == allowedType {
				allowed = true
				break
			}
		}
		if !allowed {
			errors = append(errors, fmt.Sprintf("不允許的文件類型: %s", info.MimeType))
		}
	}

	// 檢查允許的擴展名
	if len(rule.AllowedExts) > 0 {
		allowed := false
		for _, allowedExt := range rule.AllowedExts {
			if strings.EqualFold(info.Extension, allowedExt) {
				allowed = true
				break
			}
		}
		if !allowed {
			errors = append(errors, fmt.Sprintf("不允許的文件擴展名: %s", info.Extension))
		}
	}

	// 檢查禁止的MIME類型
	for _, forbiddenType := range rule.ForbiddenTypes {
		if info.MimeType == forbiddenType {
			errors = append(errors, fmt.Sprintf("禁止的文件類型: %s", info.MimeType))
			break
		}
	}

	// 檢查禁止的擴展名
	for _, forbiddenExt := range rule.ForbiddenExts {
		if strings.EqualFold(info.Extension, forbiddenExt) {
			errors = append(errors, fmt.Sprintf("禁止的文件擴展名: %s", info.Extension))
			break
		}
	}

	// 檢查魔術字節（如果啟用）
	if rule.CheckMagicBytes {
		if err := ValidateMagicBytes(ctx, filePath, info.MimeType); err != nil {
			errors = append(errors, fmt.Sprintf("文件類型驗證失败: %v", err))
		}
	}

	// 設置驗證結果
	if len(errors) > 0 {
		info.IsValid = false
		info.ErrorMessage = strings.Join(errors, "; ")
	}

	return info, nil
}

// DetectMimeType 檢測文件的MIME類型
func DetectMimeType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 讀取前512字節用於檢測
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	// 使用http.DetectContentType檢測
	mimeType := http.DetectContentType(buffer[:n])

	// 如果檢測結果是application/octet-stream，嘗試從文件擴展名推斷
	if mimeType == "application/octet-stream" {
		ext := strings.ToLower(filepath.Ext(filePath))
		if guessedType := mime.TypeByExtension(ext); guessedType != "" {
			mimeType = guessedType
		}
	}

	return mimeType, nil
}

// ValidateMagicBytes 驗證文件魔術字節
func ValidateMagicBytes(ctx context.Context, filePath, expectedMimeType string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 讀取前幾個字節
	buffer := make([]byte, 16)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}

	header := buffer[:n]

	// 檢查常見文件類型的魔術字節
	switch expectedMimeType {
	case "image/jpeg":
		if len(header) < 2 || header[0] != 0xFF || header[1] != 0xD8 {
			return fmt.Errorf("不是有效的JPEG文件")
		}
	case "image/png":
		expectedPNG := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
		if len(header) < 8 || !bytesEqual(header[:8], expectedPNG) {
			return fmt.Errorf("不是有效的PNG文件")
		}
	case "image/gif":
		if len(header) < 6 {
			return fmt.Errorf("文件太小，不是有效的GIF文件")
		}
		if !bytesEqual(header[:6], []byte("GIF87a")) && !bytesEqual(header[:6], []byte("GIF89a")) {
			return fmt.Errorf("不是有效的GIF文件")
		}
	case "application/pdf":
		if len(header) < 4 || !bytesEqual(header[:4], []byte("%PDF")) {
			return fmt.Errorf("不是有效的PDF文件")
		}
	}

	return nil
}

// CalculateHashes 計算文件的MD5和SHA256哈希值
func CalculateHashes(filePath string) (md5Hash, sha256Hash string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	md5Hasher := md5.New()
	sha256Hasher := sha256.New()

	// 使用MultiWriter同時寫入兩個hasher
	writer := io.MultiWriter(md5Hasher, sha256Hasher)

	if _, err := io.Copy(writer, file); err != nil {
		return "", "", err
	}

	md5Hash = hex.EncodeToString(md5Hasher.Sum(nil))
	sha256Hash = hex.EncodeToString(sha256Hasher.Sum(nil))

	return md5Hash, sha256Hash, nil
}

// FileExists 檢查文件是否存在
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// GetFileSize 獲取文件大小
func GetFileSize(filePath string) (int64, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

// DeleteFile 刪除文件
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// CopyFile 複製文件
func CopyFile(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 創建目標目錄
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}

// MoveFile 移動文件
func MoveFile(srcPath, destPath string) error {
	// 創建目標目錄
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	// 嘗試重命名（如果在同一文件系統上）
	if err := os.Rename(srcPath, destPath); err == nil {
		return nil
	}

	// 如果重命名失敗，則複製然後刪除
	if err := CopyFile(srcPath, destPath); err != nil {
		return err
	}

	return os.Remove(srcPath)
}

// FormatFileSize 格式化文件大小為人類可讀格式
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// bytesEqual 比較兩個字節切片是否相等
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// CreateTempFile 在指定目錄創建臨時文件
func CreateTempFile(dir, pattern string) (*os.File, error) {
	return os.CreateTemp(dir, pattern)
}

// CleanupTempFiles 清理臨時文件
func CleanupTempFiles(pattern string) error {
	tempDir := os.TempDir()
	matches, err := filepath.Glob(filepath.Join(tempDir, pattern))
	if err != nil {
		return err
	}

	for _, match := range matches {
		if err := os.Remove(match); err != nil {
			// 記錄錯誤但繼續清理其他文件
			fmt.Printf("清理臨時文件失败 %s: %v\n", match, err)
		}
	}

	return nil
}

// DownloadResult 下載結果結構
type DownloadResult struct {
	Data         []byte            `json:"-"`             // 文件數據
	Size         int64             `json:"size"`          // 文件大小
	MimeType     string            `json:"mime_type"`     // MIME類型
	Headers      map[string]string `json:"headers"`       // 響應頭
	StatusCode   int               `json:"status_code"`   // HTTP狀態碼
	URL          string            `json:"url"`           // 下載URL
	MD5Hash      string            `json:"md5_hash"`      // MD5哈希
	SHA256Hash   string            `json:"sha256_hash"`   // SHA256哈希
	DownloadTime time.Duration     `json:"download_time"` // 下載耗時
}

// DownloadToStream 從URL下載文件並返回數據流
func DownloadToStream(ctx context.Context, url string, options *DownloadOptions) (*DownloadResult, error) {
	if options == nil {
		options = DefaultDownloadOptions()
	}

	startTime := time.Now()

	// 創建HTTP客戶端
	client := &http.Client{
		Timeout: options.Timeout,
	}

	// 如果不跟隨重定向
	if !options.FollowRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	// 創建請求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("創建請求失敗: %w", err)
	}

	// 設置請求頭
	req.Header.Set("User-Agent", options.UserAgent)
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// 發送請求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("下載失敗: %w", err)
	}
	defer resp.Body.Close()

	// 檢查HTTP狀態碼
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("下載失敗，HTTP狀態碼: %d", resp.StatusCode)
	}

	// 檢查內容長度
	if resp.ContentLength > 0 && resp.ContentLength > options.MaxSize {
		return nil, fmt.Errorf("文件過大: %d bytes, 最大允許: %d bytes", resp.ContentLength, options.MaxSize)
	}

	// 限制讀取大小並讀取數據
	limitedReader := io.LimitReader(resp.Body, options.MaxSize+1)
	data, err := io.ReadAll(limitedReader)
	if err != nil {
		return nil, fmt.Errorf("讀取數據失敗: %w", err)
	}

	// 檢查是否超過最大大小
	if int64(len(data)) > options.MaxSize {
		return nil, fmt.Errorf("文件過大: %d bytes, 最大允許: %d bytes", len(data), options.MaxSize)
	}

	// 檢測MIME類型
	mimeType := http.DetectContentType(data)
	if mimeType == "application/octet-stream" {
		// 嘗試從Content-Type響應頭獲取
		if contentType := resp.Header.Get("Content-Type"); contentType != "" {
			mimeType = strings.Split(contentType, ";")[0] // 移除參數部分
		}
	}

	// 計算哈希值
	md5Hash := fmt.Sprintf("%x", md5.Sum(data))
	sha256Hash := fmt.Sprintf("%x", sha256.Sum256(data))

	// 收集響應頭
	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	return &DownloadResult{
		Data:         data,
		Size:         int64(len(data)),
		MimeType:     mimeType,
		Headers:      headers,
		StatusCode:   resp.StatusCode,
		URL:          url,
		MD5Hash:      md5Hash,
		SHA256Hash:   sha256Hash,
		DownloadTime: time.Since(startTime),
	}, nil
}

// DownloadAndValidate 下載文件並驗證
func DownloadAndValidate(ctx context.Context, url string, rule *ValidationRule, options *DownloadOptions) (*DownloadResult, error) {
	if rule == nil {
		rule = DefaultValidationRule()
	}

	// 下載文件
	result, err := DownloadToStream(ctx, url, options)
	if err != nil {
		return nil, err
	}

	// 驗證文件大小
	if result.Size > rule.MaxSize {
		return nil, fmt.Errorf("文件過大: %d bytes, 最大允許: %d bytes", result.Size, rule.MaxSize)
	}
	if result.Size < rule.MinSize {
		return nil, fmt.Errorf("文件過小: %d bytes, 最小要求: %d bytes", result.Size, rule.MinSize)
	}

	// 驗證MIME類型
	if len(rule.AllowedTypes) > 0 {
		allowed := false
		for _, allowedType := range rule.AllowedTypes {
			if result.MimeType == allowedType {
				allowed = true
				break
			}
		}
		if !allowed {
			return nil, fmt.Errorf("不允許的文件類型: %s", result.MimeType)
		}
	}

	// 檢查禁止的MIME類型
	for _, forbiddenType := range rule.ForbiddenTypes {
		if result.MimeType == forbiddenType {
			return nil, fmt.Errorf("禁止的文件類型: %s", result.MimeType)
		}
	}

	// 驗證魔術字節（如果啟用）
	if rule.CheckMagicBytes {
		if err := validateMagicBytesFromData(ctx, result.Data, result.MimeType); err != nil {
			return nil, fmt.Errorf("文件類型驗證失敗: %w", err)
		}
	}

	return result, nil
}

// validateMagicBytesFromData 從數據驗證魔術字節
func validateMagicBytesFromData(ctx context.Context, data []byte, expectedMimeType string) error {
	if len(data) == 0 {
		return fmt.Errorf("數據為空")
	}

	// 檢查常見文件類型的魔術字節
	switch expectedMimeType {
	case "image/jpeg":
		if len(data) < 2 || data[0] != 0xFF || data[1] != 0xD8 {
			return fmt.Errorf("不是有效的JPEG文件")
		}
	case "image/png":
		expectedPNG := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
		if len(data) < 8 || !bytesEqual(data[:8], expectedPNG) {
			return fmt.Errorf("不是有效的PNG文件")
		}
	case "image/gif":
		if len(data) < 6 {
			return fmt.Errorf("數據太小，不是有效的GIF文件")
		}
		if !bytesEqual(data[:6], []byte("GIF87a")) && !bytesEqual(data[:6], []byte("GIF89a")) {
			return fmt.Errorf("不是有效的GIF文件")
		}
	case "application/pdf":
		if len(data) < 4 || !bytesEqual(data[:4], []byte("%PDF")) {
			return fmt.Errorf("不是有效的PDF文件")
		}
	}

	return nil
}

// SaveStreamToFile 將下載結果保存到文件
func SaveStreamToFile(ctx context.Context, result *DownloadResult, destPath string) error {
	// 創建目標目錄
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("創建目錄失敗: %w", err)
	}

	// 寫入文件
	if err := os.WriteFile(destPath, result.Data, 0644); err != nil {
		return fmt.Errorf("保存文件失敗: %w", err)
	}

	return nil
}

// GetStreamReader 獲取數據流的讀取器
func (r *DownloadResult) GetStreamReader() *strings.Reader {
	return strings.NewReader(string(r.Data))
}

// GetBytesReader 獲取字節流的讀取器
func (r *DownloadResult) GetBytesReader() *bytes.Reader {
	return bytes.NewReader(r.Data)
}

// WriteToResponse 將下載結果寫入HTTP響應
func (r *DownloadResult) WriteToResponse(w http.ResponseWriter, filename string) error {
	// 設置響應頭
	w.Header().Set("Content-Type", r.MimeType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", r.Size))

	if filename != "" {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	}

	// 寫入數據
	_, err := w.Write(r.Data)
	return err
}
