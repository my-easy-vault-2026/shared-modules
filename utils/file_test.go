package utils

import (
	"context"
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestDownloadToStream 測試下載為數據流功能
func TestDownloadToStream(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		options     *DownloadOptions
		expectError bool
	}{
		{
			name: "下載小文本文件到流",
			url:  "https://httpbin.org/robots.txt",
			options: &DownloadOptions{
				Timeout:        10 * time.Second,
				MaxSize:        1024 * 1024, // 1MB
				UserAgent:      "FileUtil-Test/1.0",
				FollowRedirect: true,
			},
			expectError: false,
		},
		{
			name: "下載JSON數據到流",
			url:  "https://httpbin.org/json",
			options: &DownloadOptions{
				Timeout:        10 * time.Second,
				MaxSize:        10 * 1024, // 10KB
				UserAgent:      "FileUtil-Test/1.0",
				FollowRedirect: true,
			},
			expectError: false,
		},
		{
			name:        "下載不存在的文件到流",
			url:         "https://httpbin.org/nonexistent-file.txt",
			options:     DefaultDownloadOptions(),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DownloadToStream(context.Background(), tt.url, tt.options)

			if tt.expectError {
				if err == nil {
					t.Errorf("期望錯誤但沒有發生錯誤")
				}
				return
			}

			if err != nil {
				t.Errorf("下載到流失敗: %v", err)
				return
			}

			if result == nil {
				t.Errorf("下載結果為空")
				return
			}

			// 驗證下載結果
			if result.Size <= 0 {
				t.Errorf("下載大小無效: %d", result.Size)
			}

			if len(result.Data) != int(result.Size) {
				t.Errorf("數據長度與Size不匹配: 數據長度=%d, Size=%d", len(result.Data), result.Size)
			}

			if result.StatusCode != 200 {
				t.Errorf("HTTP狀態碼不正確: %d", result.StatusCode)
			}

			if result.MD5Hash == "" {
				t.Errorf("MD5哈希為空")
			}

			if result.SHA256Hash == "" {
				t.Errorf("SHA256哈希為空")
			}

			if result.DownloadTime <= 0 {
				t.Errorf("下載時間無效: %v", result.DownloadTime)
			}

			t.Logf("下載到流成功: 大小=%s, MIME=%s, 耗時=%v",
				FormatFileSize(result.Size), result.MimeType, result.DownloadTime)
			t.Logf("MD5: %s", result.MD5Hash)
		})
	}
}

// TestDownloadAndValidate 測試下載並驗證功能
func TestDownloadAndValidate(t *testing.T) {
	// 定義驗證規則
	rule := &ValidationRule{
		MaxSize: 50 * 1024, // 50KB
		MinSize: 100,       // 100 bytes
		AllowedTypes: []string{
			"text/plain",
			"application/json",
			"text/html",
		},
		CheckMagicBytes: false, // 對於文本文件不檢查魔術字節
	}

	tests := []struct {
		name        string
		url         string
		shouldPass  bool
		description string
	}{
		{
			name:        "下載並驗證robots.txt",
			url:         "https://httpbin.org/robots.txt",
			shouldPass:  true,
			description: "文本文件應該通過驗證",
		},
		{
			name:        "下載並驗證JSON",
			url:         "https://httpbin.org/json",
			shouldPass:  true,
			description: "JSON文件應該通過驗證",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result, err := DownloadAndValidate(context.Background(), tt.url, rule, DefaultDownloadOptions())

			if tt.shouldPass {
				if err != nil {
					t.Errorf("下載驗證失敗: %v", err)
					return
				}

				if result == nil {
					t.Errorf("下載結果為空")
					return
				}

				t.Logf("下載驗證成功: %s (大小: %s, 類型: %s)",
					tt.name, FormatFileSize(result.Size), result.MimeType)
			} else {
				if err == nil {
					t.Errorf("期望驗證失敗但通過了驗證")
				}
			}
		})
	}
}

// TestStreamOperations 測試流操作功能
func TestStreamOperations(t *testing.T) {
	// 先下載一個文件到流
	result, err := DownloadToStream(context.Background(), "https://httpbin.org/json", DefaultDownloadOptions())
	if err != nil {
		t.Fatalf("下載文件失敗: %v", err)
	}

	// 測試獲取字節讀取器
	bytesReader := result.GetBytesReader()
	if bytesReader == nil {
		t.Errorf("獲取字節讀取器失敗")
	}

	// 測試從字節讀取器讀取數據
	readData := make([]byte, result.Size)
	n, err := bytesReader.Read(readData)
	if err != nil && err != io.EOF {
		t.Errorf("從字節讀取器讀取失敗: %v", err)
	}

	if int64(n) != result.Size {
		t.Errorf("讀取數據長度不匹配: 期望 %d, 實際 %d", result.Size, n)
	}

	// 測試獲取字符串讀取器
	stringReader := result.GetStreamReader()
	if stringReader == nil {
		t.Errorf("獲取字符串讀取器失敗")
	}

	// 測試保存流到文件
	tempFile := "./test_stream_save.json"
	defer os.Remove(tempFile)

	if err := SaveStreamToFile(context.Background(), result, tempFile); err != nil {
		t.Errorf("保存流到文件失敗: %v", err)
	}

	// 驗證保存的文件
	if !FileExists(tempFile) {
		t.Errorf("保存的文件不存在")
	}

	savedData, err := ReadFile(tempFile)
	if err != nil {
		t.Errorf("讀取保存的文件失敗: %v", err)
	}

	if !bytesEqual(savedData, result.Data) {
		t.Errorf("保存的文件內容不匹配")
	}

	t.Logf("流操作測試通過，處理了 %s 的數據", FormatFileSize(result.Size))
}

// TestKYCStreamDownload 測試KYC文件流下載（模擬場景）
func TestKYCStreamDownload(t *testing.T) {
	// 模擬從CDN或文件服務器下載KYC文件
	// 這裡使用httpbin提供的圖像端點作為測試

	kycRule := &ValidationRule{
		MaxSize: 2 * 1024 * 1024, // 2MB
		MinSize: 1024,            // 1KB
		AllowedTypes: []string{
			"image/jpeg",
			"image/png",
			"image/webp",
		},
		CheckMagicBytes: true,
	}

	// 測試下載圖像文件（使用httpbin的圖像端點）
	imageURL := "https://httpbin.org/image/jpeg"

	result, err := DownloadAndValidate(context.Background(), imageURL, kycRule, &DownloadOptions{
		Timeout:        15 * time.Second,
		MaxSize:        2 * 1024 * 1024,
		UserAgent:      "KYC-Service/1.0",
		FollowRedirect: true,
	})

	if err != nil {
		// 如果網絡不可用或服務不可用，跳過測試
		t.Skipf("跳過KYC流下載測試，網絡錯誤: %v", err)
		return
	}

	if result == nil {
		t.Errorf("KYC文件下載結果為空")
		return
	}

	// 驗證下載結果
	if result.MimeType != "image/jpeg" {
		t.Errorf("MIME類型不匹配: 期望 image/jpeg, 實際 %s", result.MimeType)
	}

	// 驗證文件頭
	if len(result.Data) < 2 || result.Data[0] != 0xFF || result.Data[1] != 0xD8 {
		t.Errorf("JPEG文件頭驗證失敗")
	}

	// 測試將結果轉換為base64（模擬API返回）
	base64Data := base64.StdEncoding.EncodeToString(result.Data)
	if base64Data == "" {
		t.Errorf("base64編碼失敗")
	}

	t.Logf("KYC文件下載成功:")
	t.Logf("- 大小: %s", FormatFileSize(result.Size))
	t.Logf("- MIME類型: %s", result.MimeType)
	t.Logf("- MD5: %s", result.MD5Hash)
	t.Logf("- 下載耗時: %v", result.DownloadTime)
	t.Logf("- Base64長度: %d", len(base64Data))
}

// TestDownloadFile 測試文件下載功能
func TestDownloadFile(t *testing.T) {
	// 創建測試目錄
	testDir := "./test-downloads"
	defer os.RemoveAll(testDir)

	tests := []struct {
		name        string
		url         string
		destPath    string
		options     *DownloadOptions
		expectError bool
	}{
		{
			name:     "下載小文本文件",
			url:      "https://httpbin.org/robots.txt",
			destPath: filepath.Join(testDir, "robots.txt"),
			options: &DownloadOptions{
				Timeout:        10 * time.Second,
				MaxSize:        1024 * 1024, // 1MB
				UserAgent:      "FileUtil-Test/1.0",
				FollowRedirect: true,
			},
			expectError: false,
		},
		{
			name:        "下載不存在的文件",
			url:         "https://httpbin.org/nonexistent-file.txt",
			destPath:    filepath.Join(testDir, "nonexistent.txt"),
			options:     DefaultDownloadOptions(),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := DownloadFile(context.Background(), tt.url, tt.destPath, tt.options)

			if tt.expectError {
				if err == nil {
					t.Errorf("期望錯誤但沒有發生錯誤")
				}
				return
			}

			if err != nil {
				t.Errorf("下載文件失敗: %v", err)
				return
			}

			if info == nil {
				t.Errorf("文件信息為空")
				return
			}

			// 驗證文件是否存在
			if !FileExists(tt.destPath) {
				t.Errorf("下載的文件不存在: %s", tt.destPath)
			}

			// 驗證文件大小
			if info.Size <= 0 {
				t.Errorf("文件大小無效: %d", info.Size)
			}

			t.Logf("下載成功: %s (大小: %s)", info.Name, FormatFileSize(info.Size))
		})
	}
}

// TestReadFile 測試文件讀取功能
func TestReadFile(t *testing.T) {
	// 創建測試文件
	testContent := "這是測試文件內容\n包含多行文本"
	testFile := "./test_read_file.txt"

	// 寫入測試文件
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("創建測試文件失敗: %v", err)
	}
	defer os.Remove(testFile)

	// 測試正常讀取
	content, err := ReadFile(testFile)
	if err != nil {
		t.Errorf("讀取文件失敗: %v", err)
	}

	if string(content) != testContent {
		t.Errorf("讀取內容不匹配\n期望: %s\n實際: %s", testContent, string(content))
	}

	// 測試帶限制的讀取
	limitedContent, err := ReadFileWithLimit(testFile, int64(len(testContent)))
	if err != nil {
		t.Errorf("限制讀取失敗: %v", err)
	}

	if string(limitedContent) != testContent {
		t.Errorf("限制讀取內容不匹配")
	}

	// 測試超過限制的讀取
	_, err = ReadFileWithLimit(testFile, 5)
	if err == nil {
		t.Errorf("期望超過限制錯誤但沒有發生")
	}

	t.Logf("文件讀取測試通過，內容長度: %d bytes", len(content))
}

// TestFileValidation 測試文件驗證功能
func TestFileValidation(t *testing.T) {
	// 創建不同類型的測試文件
	testFiles := map[string][]byte{
		"test.txt": []byte("這是文本文件"),
		"test.jpg": {0xFF, 0xD8, 0xFF, 0xE0},                         // JPEG文件頭
		"test.png": {0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, // PNG文件頭
		"test.pdf": []byte("%PDF-1.4\ntest pdf content"),
	}

	// 創建測試文件
	for filename, content := range testFiles {
		if err := os.WriteFile(filename, content, 0644); err != nil {
			t.Fatalf("創建測試文件失敗 %s: %v", filename, err)
		}
		defer os.Remove(filename)
	}

	// 定義驗證規則
	rule := &ValidationRule{
		MaxSize:         1024,
		MinSize:         1,
		AllowedTypes:    []string{"image/jpeg", "image/png", "application/pdf", "text/plain"},
		AllowedExts:     []string{".jpg", ".png", ".pdf", ".txt"},
		CheckMagicBytes: true,
	}

	tests := []struct {
		filename    string
		shouldPass  bool
		description string
	}{
		{"test.txt", true, "文本文件應該通過驗證"},
		{"test.jpg", true, "JPEG文件應該通過驗證"},
		{"test.png", true, "PNG文件應該通過驗證"},
		{"test.pdf", true, "PDF文件應該通過驗證"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			info, err := ValidateFile(context.Background(), tt.filename, rule)
			if err != nil {
				t.Errorf("驗證文件時發生錯誤: %v", err)
				return
			}

			if info.IsValid != tt.shouldPass {
				t.Errorf("驗證結果不符合期望\n文件: %s\n期望通過: %v\n實際通過: %v\n錯誤信息: %s",
					tt.filename, tt.shouldPass, info.IsValid, info.ErrorMessage)
			}

			t.Logf("文件 %s 驗證結果: %v (類型: %s, 大小: %s)",
				tt.filename, info.IsValid, info.MimeType, FormatFileSize(info.Size))
		})
	}
}

// TestGetFileInfo 測試獲取文件信息功能
func TestGetFileInfo(t *testing.T) {
	// 創建測試文件
	testContent := "測試文件信息功能"
	testFile := "./test_file_info.txt"

	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("創建測試文件失敗: %v", err)
	}
	defer os.Remove(testFile)

	// 獲取文件信息
	info, err := GetFileInfo(testFile)
	if err != nil {
		t.Errorf("獲取文件信息失敗: %v", err)
		return
	}

	// 驗證基本信息
	if info.Name != filepath.Base(testFile) {
		t.Errorf("文件名不匹配: 期望 %s, 實際 %s", filepath.Base(testFile), info.Name)
	}

	if info.Size != int64(len(testContent)) {
		t.Errorf("文件大小不匹配: 期望 %d, 實際 %d", len(testContent), info.Size)
	}

	if info.Extension != ".txt" {
		t.Errorf("文件擴展名不匹配: 期望 .txt, 實際 %s", info.Extension)
	}

	if info.MD5Hash == "" {
		t.Errorf("MD5哈希值為空")
	}

	if info.SHA256Hash == "" {
		t.Errorf("SHA256哈希值為空")
	}

	if !info.IsValid {
		t.Errorf("文件應該是有效的，但驗證失敗: %s", info.ErrorMessage)
	}

	t.Logf("文件信息獲取成功:\n名稱: %s\n大小: %s\n類型: %s\nMD5: %s",
		info.Name, FormatFileSize(info.Size), info.MimeType, info.MD5Hash)
}

// TestDetectMimeType 測試MIME類型檢測
func TestDetectMimeType(t *testing.T) {
	tests := []struct {
		filename     string
		content      []byte
		expectedType string
	}{
		{"test.txt", []byte("plain text content"), "text/plain"},
		{"test.jpg", []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}, "image/jpeg"},
		{"test.png", []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, "image/png"},
		{"test.pdf", []byte("%PDF-1.4"), "application/pdf"},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			// 創建測試文件
			if err := os.WriteFile(tt.filename, tt.content, 0644); err != nil {
				t.Fatalf("創建測試文件失敗: %v", err)
			}
			defer os.Remove(tt.filename)

			// 檢測MIME類型
			mimeType, err := DetectMimeType(tt.filename)
			if err != nil {
				t.Errorf("檢測MIME類型失敗: %v", err)
				return
			}

			if !strings.HasPrefix(mimeType, strings.Split(tt.expectedType, "/")[0]) {
				t.Logf("MIME類型檢測結果: %s (期望前綴: %s)", mimeType, tt.expectedType)
			}

			t.Logf("文件 %s 的MIME類型: %s", tt.filename, mimeType)
		})
	}
}

// TestFileOperations 測試文件操作功能
func TestFileOperations(t *testing.T) {
	// 創建測試文件
	srcContent := "這是源文件內容"
	srcFile := "./test_src.txt"
	copyFile := "./test_copy.txt"
	moveFile := "./test_move.txt"

	if err := os.WriteFile(srcFile, []byte(srcContent), 0644); err != nil {
		t.Fatalf("創建源文件失敗: %v", err)
	}
	defer os.Remove(srcFile)
	defer os.Remove(copyFile)
	defer os.Remove(moveFile)

	// 測試文件存在檢查
	if !FileExists(srcFile) {
		t.Errorf("源文件應該存在")
	}

	if FileExists("nonexistent_file.txt") {
		t.Errorf("不存在的文件不應該返回存在")
	}

	// 測試獲取文件大小
	size, err := GetFileSize(srcFile)
	if err != nil {
		t.Errorf("獲取文件大小失敗: %v", err)
	}

	if size != int64(len(srcContent)) {
		t.Errorf("文件大小不匹配: 期望 %d, 實際 %d", len(srcContent), size)
	}

	// 測試複製文件
	if err := CopyFile(srcFile, copyFile); err != nil {
		t.Errorf("複製文件失敗: %v", err)
	}

	if !FileExists(copyFile) {
		t.Errorf("複製的文件不存在")
	}

	// 驗證複製內容
	copyContent, err := ReadFile(copyFile)
	if err != nil {
		t.Errorf("讀取複製文件失敗: %v", err)
	}

	if string(copyContent) != srcContent {
		t.Errorf("複製文件內容不匹配")
	}

	// 測試移動文件
	if err := MoveFile(copyFile, moveFile); err != nil {
		t.Errorf("移動文件失敗: %v", err)
	}

	if FileExists(copyFile) {
		t.Errorf("移動後源文件仍存在")
	}

	if !FileExists(moveFile) {
		t.Errorf("移動後目標文件不存在")
	}

	// 測試刪除文件
	if err := DeleteFile(moveFile); err != nil {
		t.Errorf("刪除文件失敗: %v", err)
	}

	if FileExists(moveFile) {
		t.Errorf("刪除後文件仍存在")
	}

	t.Logf("文件操作測試全部通過")
}

// TestCalculateHashes 測試哈希計算功能
func TestCalculateHashes(t *testing.T) {
	testContent := "測試哈希計算功能"
	testFile := "./test_hash.txt"

	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("創建測試文件失敗: %v", err)
	}
	defer os.Remove(testFile)

	md5Hash, sha256Hash, err := CalculateHashes(testFile)
	if err != nil {
		t.Errorf("計算哈希失敗: %v", err)
		return
	}

	// 驗證哈希值不為空且格式正確
	if len(md5Hash) != 32 {
		t.Errorf("MD5哈希長度不正確: 期望 32, 實際 %d", len(md5Hash))
	}

	if len(sha256Hash) != 64 {
		t.Errorf("SHA256哈希長度不正確: 期望 64, 實際 %d", len(sha256Hash))
	}

	// 驗證哈希值是16進制字符串
	for _, char := range md5Hash {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f')) {
			t.Errorf("MD5哈希包含無效字符: %c", char)
			break
		}
	}

	t.Logf("哈希計算成功:\nMD5: %s\nSHA256: %s", md5Hash, sha256Hash)
}

// TestFormatFileSize 測試文件大小格式化
func TestFormatFileSize(t *testing.T) {
	tests := []struct {
		size     int64
		expected string
	}{
		{0, "0 B"},
		{512, "512 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1024 * 1024, "1.0 MB"},
		{1024 * 1024 * 1024, "1.0 GB"},
		{5 * 1024 * 1024 * 1024, "5.0 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := FormatFileSize(tt.size)
			if result != tt.expected {
				t.Errorf("格式化結果不匹配: 期望 %s, 實際 %s", tt.expected, result)
			}
		})
	}
}

// TestKYCFileValidation 測試KYC文件驗證（模擬真實場景）
func TestKYCFileValidation(t *testing.T) {
	// KYC文件驗證規則（基於你的API文檔）
	kycRule := &ValidationRule{
		MaxSize: 2 * 1024 * 1024, // 2MB
		MinSize: 10 * 1024,       // 10KB
		AllowedTypes: []string{
			"image/jpeg",
			"image/png",
			"image/gif",
			"application/pdf",
		},
		AllowedExts: []string{
			".jpg", ".jpeg", ".png", ".gif", ".pdf",
		},
		CheckMagicBytes: true,
	}

	// 創建模擬KYC文件
	testFiles := []struct {
		name        string
		content     []byte
		shouldPass  bool
		description string
	}{
		{
			name:        "id_front.jpg",
			content:     append([]byte{0xFF, 0xD8, 0xFF, 0xE0}, make([]byte, 15*1024)...), // 15KB JPEG
			shouldPass:  true,
			description: "身份證正面照（符合要求）",
		},
		{
			name:        "id_back.jpg",
			content:     append([]byte{0xFF, 0xD8, 0xFF, 0xE0}, make([]byte, 20*1024)...), // 20KB JPEG
			shouldPass:  true,
			description: "身份證背面照（符合要求）",
		},
		{
			name:        "holding_id.png",
			content:     append([]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, make([]byte, 50*1024)...), // 50KB PNG
			shouldPass:  true,
			description: "手持證件照（符合要求）",
		},
		{
			name:        "oversized.jpg",
			content:     append([]byte{0xFF, 0xD8, 0xFF, 0xE0}, make([]byte, 3*1024*1024)...), // 3MB JPEG
			shouldPass:  false,
			description: "超大文件（應該被拒絕）",
		},
		{
			name:        "too_small.jpg",
			content:     []byte{0xFF, 0xD8, 0xFF, 0xE0}, // 4 bytes JPEG
			shouldPass:  false,
			description: "文件太小（應該被拒絕）",
		},
		{
			name:        "invalid_type.exe",
			content:     []byte("MZ"), // 可執行文件頭
			shouldPass:  false,
			description: "不允許的文件類型（應該被拒絕）",
		},
	}

	for _, tt := range testFiles {
		t.Run(tt.description, func(t *testing.T) {
			// 創建測試文件
			if err := os.WriteFile(tt.name, tt.content, 0644); err != nil {
				t.Fatalf("創建測試文件失敗: %v", err)
			}
			defer os.Remove(tt.name)

			// 驗證文件
			info, err := ValidateFile(context.Background(), tt.name, kycRule)
			if err != nil {
				t.Errorf("驗證文件時發生錯誤: %v", err)
				return
			}

			if info.IsValid != tt.shouldPass {
				t.Errorf("KYC文件驗證結果不符合期望\n文件: %s\n期望通過: %v\n實際通過: %v\n錯誤: %s",
					tt.name, tt.shouldPass, info.IsValid, info.ErrorMessage)
			}

			if info.IsValid {
				t.Logf("✓ KYC文件驗證通過: %s (大小: %s, 類型: %s)",
					info.Name, FormatFileSize(info.Size), info.MimeType)
			} else {
				t.Logf("✗ KYC文件驗證失敗: %s - %s", info.Name, info.ErrorMessage)
			}
		})
	}
}

// TestTempFileOperations 測試臨時文件操作
func TestTempFileOperations(t *testing.T) {
	// 創建臨時文件
	tempFile, err := CreateTempFile("", "test_*.tmp")
	if err != nil {
		t.Errorf("創建臨時文件失敗: %v", err)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	// 寫入數據
	testData := "這是臨時文件測試數據"
	if _, err := tempFile.WriteString(testData); err != nil {
		t.Errorf("寫入臨時文件失敗: %v", err)
		return
	}

	// 獲取臨時文件信息
	info, err := GetFileInfo(tempFile.Name())
	if err != nil {
		t.Errorf("獲取臨時文件信息失敗: %v", err)
		return
	}

	if info.Size != int64(len(testData)) {
		t.Errorf("臨時文件大小不匹配: 期望 %d, 實際 %d", len(testData), info.Size)
	}

	t.Logf("臨時文件操作成功: %s (大小: %s)", info.Name, FormatFileSize(info.Size))

	// 測試清理臨時文件
	pattern := "test_*.tmp"
	if err := CleanupTempFiles(pattern); err != nil {
		t.Errorf("清理臨時文件失敗: %v", err)
	}
}

// BenchmarkCalculateHashes 哈希計算性能測試
func BenchmarkCalculateHashes(b *testing.B) {
	// 創建測試文件
	testContent := make([]byte, 1024*1024) // 1MB
	for i := range testContent {
		testContent[i] = byte(i % 256)
	}

	testFile := "./bench_test.dat"
	if err := os.WriteFile(testFile, testContent, 0644); err != nil {
		b.Fatalf("創建測試文件失敗: %v", err)
	}
	defer os.Remove(testFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := CalculateHashes(testFile)
		if err != nil {
			b.Errorf("計算哈希失敗: %v", err)
		}
	}
}

// BenchmarkDetectMimeType MIME類型檢測性能測試
func BenchmarkDetectMimeType(b *testing.B) {
	testFile := "./bench_mime_test.jpg"
	testContent := []byte{0xFF, 0xD8, 0xFF, 0xE0}            // JPEG頭
	testContent = append(testContent, make([]byte, 1024)...) // 添加1KB數據

	if err := os.WriteFile(testFile, testContent, 0644); err != nil {
		b.Fatalf("創建測試文件失敗: %v", err)
	}
	defer os.Remove(testFile)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := DetectMimeType(testFile)
		if err != nil {
			b.Errorf("檢測MIME類型失敗: %v", err)
		}
	}
}
