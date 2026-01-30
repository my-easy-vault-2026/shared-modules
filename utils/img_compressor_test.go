package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestImgCompressorBasicUsage 基本使用測試
func TestImgCompressorBasicUsage(t *testing.T) {
	// 創建測試圖片數據（模擬JPEG文件頭）
	testImageData := createMockJPEGData(800, 600, 100*1024) // 800x600, 100KB

	// 使用默認選項壓縮
	result, err := ImgCompressorCompressWithDefaultOptions(testImageData)
	if err != nil {
		t.Fatalf("基本壓縮失敗: %v", err)
	}

	// 驗證結果
	if result.OriginalSize != int64(len(testImageData)) {
		t.Errorf("原始大小不匹配: 期望 %d, 實際 %d", len(testImageData), result.OriginalSize)
	}

	if result.CompressedSize <= 0 {
		t.Errorf("壓縮後大小無效: %d", result.CompressedSize)
	}

	if result.CompressionRatio <= 0 || result.CompressionRatio > 1 {
		t.Errorf("壓縮比無效: %f", result.CompressionRatio)
	}

	t.Logf("基本壓縮成功:")
	t.Logf("- 原始大小: %s (%dx%d)",
		ImgCompressorFormatFileSize(result.OriginalSize), result.OriginalWidth, result.OriginalHeight)
	t.Logf("- 壓縮後大小: %s (%dx%d)",
		ImgCompressorFormatFileSize(result.CompressedSize), result.NewWidth, result.NewHeight)
	t.Logf("- 節省空間: %.1f%%, 處理時間: %dms",
		(1-result.CompressionRatio)*100, result.ProcessTime)
}

// TestImgCompressorKYCProcessing KYC圖片處理測試
func TestImgCompressorKYCProcessing(t *testing.T) {
	tests := []struct {
		name        string
		imageData   []byte
		description string
	}{
		{
			name:        "身份證正面",
			imageData:   createMockJPEGData(1920, 1080, 1.5*1024*1024), // 1.5MB
			description: "高解析度身份證正面照",
		},
		{
			name:        "手持證件照",
			imageData:   createMockJPEGData(1280, 960, 800*1024), // 800KB
			description: "手持身份證照片",
		},
		{
			name:        "護照照片",
			imageData:   createMockPNGData(1600, 1200, 2*1024*1024), // 2.2MB PNG
			description: "高品質護照照片",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 獲取原始圖片信息
			info, err := ImgCompressorGetImageInfo(tt.imageData)
			if err != nil {
				t.Fatalf("獲取圖片信息失敗: %v", err)
			}

			t.Logf("處理 %s (%s):", tt.name, tt.description)
			t.Logf("- 原始信息: %dx%d, %s, %s",
				info.Width, info.Height, info.Format, ImgCompressorFormatFileSize(info.Size))

			// 使用KYC專用處理
			result, err := ImgCompressorProcessKYCImage(tt.imageData)
			if err != nil {
				t.Fatalf("KYC圖片處理失敗: %v", err)
			}

			// 驗證KYC處理結果
			if result.CompressedSize > 2*1024*1024 { // 不應超過2MB
				t.Errorf("KYC處理後文件過大: %s", ImgCompressorFormatFileSize(result.CompressedSize))
			}

			if result.Format != "jpeg" {
				t.Errorf("KYC處理後格式不正確: %s", result.Format)
			}

			t.Logf("- 處理結果: %dx%d, 品質%d, %s",
				result.NewWidth, result.NewHeight, result.Quality, ImgCompressorFormatFileSize(result.CompressedSize))
			t.Logf("- 節省空間: %.1f%%, 處理時間: %dms",
				(1-result.CompressionRatio)*100, result.ProcessTime)

			// 轉換為base64（模擬API返回）
			base64Data := base64.StdEncoding.EncodeToString(result.Data)
			t.Logf("- Base64長度: %d字符", len(base64Data))

			// 驗證base64數據有效性
			if len(base64Data) == 0 {
				t.Errorf("Base64轉換失敗")
			}
		})
	}
}

// TestImgCompressorBatchProcessing 批量處理測試
func TestImgCompressorBatchProcessing(t *testing.T) {
	// 創建測試目錄和文件
	testDir := "./test-batch-images"
	outputDir := "./test-output"

	// 清理測試目錄
	defer func() {
		os.RemoveAll(testDir)
		os.RemoveAll(outputDir)
	}()

	// 創建測試目錄
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("創建測試目錄失敗: %v", err)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		t.Fatalf("創建輸出目錄失敗: %v", err)
	}

	// 創建測試文件
	testFiles := []struct {
		filename string
		data     []byte
	}{
		{"photo1.jpg", createMockJPEGData(1920, 1080, 2*1024*1024)},
		{"photo2.jpg", createMockJPEGData(1280, 720, 1*1024*1024)},
		{"photo3.jpg", createMockJPEGData(800, 600, 500*1024)},
	}

	var filePaths []string
	for _, file := range testFiles {
		filePath := filepath.Join(testDir, file.filename)
		if err := os.WriteFile(filePath, file.data, 0644); err != nil {
			t.Fatalf("創建測試文件失敗: %v", err)
		}
		filePaths = append(filePaths, filePath)
	}

	// 自定義批量處理選項
	options := &ImgCompressorCompressOptions{
		MaxWidth:     1200,
		MaxHeight:    800,
		Quality:      80,
		Format:       "jpeg",
		MaxFileSize:  500 * 1024, // 500KB
		AutoQuality:  true,
		ResizeMethod: "lanczos",
		ImgCompressorQualityRange: ImgCompressorQualityRange{
			Min: 60,
			Max: 85,
		},
	}

	// 執行批量處理
	results, err := ImgCompressorBatchCompress(filePaths, options)
	if err != nil {
		t.Fatalf("批量處理失敗: %v", err)
	}

	// 驗證批量處理結果
	if len(results) != len(filePaths) {
		t.Errorf("處理結果數量不匹配: 期望 %d, 實際 %d", len(filePaths), len(results))
	}

	totalOriginalSize := int64(0)
	totalCompressedSize := int64(0)

	for i, result := range results {
		t.Logf("文件 %s 處理結果:", filepath.Base(filePaths[i]))
		t.Logf("- 尺寸變化: %dx%d -> %dx%d",
			result.OriginalWidth, result.OriginalHeight, result.NewWidth, result.NewHeight)
		t.Logf("- 大小變化: %s -> %s",
			ImgCompressorFormatFileSize(result.OriginalSize),
			ImgCompressorFormatFileSize(result.CompressedSize))
		t.Logf("- 品質: %d, 處理時間: %dms", result.Quality, result.ProcessTime)

		totalOriginalSize += result.OriginalSize
		totalCompressedSize += result.CompressedSize

		// 保存處理後的文件
		outputPath := filepath.Join(outputDir, filepath.Base(filePaths[i]))
		if err := os.WriteFile(outputPath, result.Data, 0644); err != nil {
			t.Errorf("保存處理後文件失敗: %v", err)
		}
	}

	// 統計總體效果
	totalCompressionRatio := float64(totalCompressedSize) / float64(totalOriginalSize)
	t.Logf("批量處理統計:")
	t.Logf("- 總原始大小: %s", ImgCompressorFormatFileSize(totalOriginalSize))
	t.Logf("- 總壓縮後大小: %s", ImgCompressorFormatFileSize(totalCompressedSize))
	t.Logf("- 總節省空間: %.1f%%", (1-totalCompressionRatio)*100)
}

// TestImgCompressorAdvancedFeatures 高級功能測試
func TestImgCompressorAdvancedFeatures(t *testing.T) {
	testImageData := createMockJPEGData(1600, 1200, 1*1024*1024)

	// 測試高級壓縮選項
	advancedOptions := &ImgCompressorAdvancedCompressOptions{
		ImgCompressorCompressOptions: &ImgCompressorCompressOptions{
			MaxWidth:     800,
			MaxHeight:    600,
			Quality:      80,
			Format:       "jpeg",
			MaxFileSize:  200 * 1024, // 200KB
			AutoQuality:  true,
			ResizeMethod: "imaging_lanczos",
		},
		ResizeMode: ImgCompressorResizeModeFit,
		Sharpen:    0.5,
		Brightness: 5,
		Contrast:   3,
		Saturation: 2,
	}

	result, err := ImgCompressorAdvancedCompress(testImageData, advancedOptions)
	if err != nil {
		t.Fatalf("高級壓縮失敗: %v", err)
	}

	t.Logf("高級壓縮結果:")
	t.Logf("- 尺寸: %dx%d -> %dx%d",
		result.OriginalWidth, result.OriginalHeight, result.NewWidth, result.NewHeight)
	t.Logf("- 大小: %s -> %s",
		ImgCompressorFormatFileSize(result.OriginalSize),
		ImgCompressorFormatFileSize(result.CompressedSize))
	t.Logf("- 品質: %d, 處理時間: %dms", result.Quality, result.ProcessTime)

	// 驗證文件大小是否符合要求
	if result.CompressedSize > advancedOptions.MaxFileSize {
		t.Errorf("高級壓縮後文件仍過大: %s > %s",
			ImgCompressorFormatFileSize(result.CompressedSize),
			ImgCompressorFormatFileSize(advancedOptions.MaxFileSize))
	}
}

// TestImgCompressorWebOptimization Web優化測試
func TestImgCompressorWebOptimization(t *testing.T) {
	tests := []struct {
		name        string
		imageData   []byte
		targetSize  int64
		description string
	}{
		{
			name:        "大圖Web優化",
			imageData:   createMockJPEGData(3000, 2000, 5*1024*1024), // 5MB大圖
			targetSize:  500 * 1024,                                  // 目標500KB
			description: "將大圖優化為Web適用大小",
		},
		{
			name:        "中圖Web優化",
			imageData:   createMockJPEGData(1920, 1080, 2*1024*1024), // 2MB
			targetSize:  300 * 1024,                                  // 目標300KB
			description: "標準Web圖片優化",
		},
		{
			name:        "小圖Web優化",
			imageData:   createMockJPEGData(800, 600, 200*1024), // 200KB
			targetSize:  100 * 1024,                             // 目標100KB
			description: "小圖進一步優化",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ImgCompressorOptimizeForWeb(tt.imageData, tt.targetSize)
			if err != nil {
				t.Fatalf("Web優化失敗: %v", err)
			}

			t.Logf("%s (%s):", tt.name, tt.description)
			t.Logf("- 尺寸優化: %dx%d -> %dx%d",
				result.OriginalWidth, result.OriginalHeight, result.NewWidth, result.NewHeight)
			t.Logf("- 大小優化: %s -> %s (目標: %s)",
				ImgCompressorFormatFileSize(result.OriginalSize),
				ImgCompressorFormatFileSize(result.CompressedSize),
				ImgCompressorFormatFileSize(tt.targetSize))
			t.Logf("- 品質: %d, 壓縮比: %.1f%%, 處理時間: %dms",
				result.Quality, result.CompressionRatio*100, result.ProcessTime)

			// 驗證是否達到目標大小
			if result.CompressedSize > tt.targetSize*120/100 { // 允許20%誤差
				t.Logf("警告: 優化後大小超出目標較多 (%s > %s)",
					ImgCompressorFormatFileSize(result.CompressedSize),
					ImgCompressorFormatFileSize(tt.targetSize))
			}
		})
	}
}

// TestImgCompressorThumbnailGeneration 縮圖生成測試
func TestImgCompressorThumbnailGeneration(t *testing.T) {
	originalImage := createMockJPEGData(2000, 1500, 3*1024*1024) // 2000x1500, 3MB

	thumbnailSizes := []int{50, 100, 200, 300}

	for _, size := range thumbnailSizes {
		t.Run(fmt.Sprintf("縮圖_%dx%d", size, size), func(t *testing.T) {
			result, err := ImgCompressorCreateThumbnail(originalImage, size)
			if err != nil {
				t.Fatalf("生成%dx%d縮圖失敗: %v", size, size, err)
			}

			t.Logf("生成 %dx%d 縮圖:", size, size)
			t.Logf("- 原始: %dx%d (%s)",
				result.OriginalWidth, result.OriginalHeight,
				ImgCompressorFormatFileSize(result.OriginalSize))
			t.Logf("- 縮圖: %dx%d (%s)",
				result.NewWidth, result.NewHeight,
				ImgCompressorFormatFileSize(result.CompressedSize))
			t.Logf("- 處理時間: %dms", result.ProcessTime)

			// 驗證縮圖尺寸
			if result.NewWidth > size || result.NewHeight > size {
				t.Errorf("縮圖尺寸超出限制: %dx%d", result.NewWidth, result.NewHeight)
			}

			// 驗證縮圖大小顯著減小
			if result.CompressedSize >= result.OriginalSize {
				t.Errorf("縮圖大小沒有減小")
			}
		})
	}
}

// TestImgCompressorConvenienceFunctions 便利函數測試
func TestImgCompressorConvenienceFunctions(t *testing.T) {
	testImage := createMockJPEGData(1024, 768, 500*1024)

	t.Run("快速壓縮", func(t *testing.T) {
		result, err := ImgCompressorQuickCompress(testImage, 70)
		if err != nil {
			t.Fatalf("快速壓縮失敗: %v", err)
		}

		if result.Quality != 70 {
			t.Errorf("品質設置不正確: 期望 70, 實際 %d", result.Quality)
		}

		t.Logf("快速壓縮 (品質70): %s -> %s",
			ImgCompressorFormatFileSize(result.OriginalSize),
			ImgCompressorFormatFileSize(result.CompressedSize))
	})

	t.Run("僅調整尺寸", func(t *testing.T) {
		result, err := ImgCompressorResizeOnly(testImage, 800, 600)
		if err != nil {
			t.Fatalf("調整尺寸失敗: %v", err)
		}

		if result.NewWidth != 800 {
			t.Errorf("寬度調整不正確: 期望 800, 實際 %d", result.NewWidth)
		}

		t.Logf("尺寸調整: %dx%d -> %dx%d",
			result.OriginalWidth, result.OriginalHeight, result.NewWidth, result.NewHeight)
	})

	t.Run("格式轉換", func(t *testing.T) {
		result, err := ImgCompressorConvertFormat(testImage, "png")
		if err != nil {
			t.Fatalf("格式轉換失敗: %v", err)
		}

		if result.Format != "png" {
			t.Errorf("格式轉換不正確: 期望 png, 實際 %s", result.Format)
		}

		t.Logf("格式轉換 JPEG->PNG: %s -> %s",
			ImgCompressorFormatFileSize(result.OriginalSize),
			ImgCompressorFormatFileSize(result.CompressedSize))
	})
}

// TestImgCompressorFileOperations 文件操作測試
func TestImgCompressorFileOperations(t *testing.T) {
	// 創建測試目錄
	testDir := "./test-file-ops"
	defer os.RemoveAll(testDir)

	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("創建測試目錄失敗: %v", err)
	}

	// 創建測試文件
	testImageData := createMockJPEGData(1200, 800, 800*1024)
	inputFile := filepath.Join(testDir, "input.jpg")
	outputFile := filepath.Join(testDir, "output.jpg")

	if err := os.WriteFile(inputFile, testImageData, 0644); err != nil {
		t.Fatalf("創建測試文件失敗: %v", err)
	}

	// 測試從文件壓縮
	result1, err := ImgCompressorCompressFromFile(inputFile, DefaultCompressOptions)
	if err != nil {
		t.Fatalf("從文件壓縮失敗: %v", err)
	}

	t.Logf("從文件壓縮: %s -> %s",
		ImgCompressorFormatFileSize(result1.OriginalSize),
		ImgCompressorFormatFileSize(result1.CompressedSize))

	// 測試壓縮並保存
	result2, err := ImgCompressorCompressAndSave(inputFile, outputFile, &ImgCompressorCompressOptions{
		MaxWidth:    800,
		MaxHeight:   600,
		Quality:     75,
		Format:      "jpeg",
		MaxFileSize: 300 * 1024, // 300KB
		AutoQuality: true,
	})
	if err != nil {
		t.Fatalf("壓縮並保存失敗: %v", err)
	}

	// 驗證輸出文件存在
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("輸出文件不存在: %s", outputFile)
	}

	// 驗證輸出文件大小
	outputData, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("讀取輸出文件失敗: %v", err)
	}

	if int64(len(outputData)) != result2.CompressedSize {
		t.Errorf("輸出文件大小不匹配: 文件 %d, 結果 %d",
			len(outputData), result2.CompressedSize)
	}

	t.Logf("壓縮並保存: %s -> %s",
		ImgCompressorFormatFileSize(result2.OriginalSize),
		ImgCompressorFormatFileSize(result2.CompressedSize))
}

// 輔助函數：創建模擬JPEG數據
func createMockJPEGData(width, height int, targetSize int) []byte {
	// JPEG文件頭
	header := []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}

	// 填充數據以達到目標大小
	data := make([]byte, targetSize)
	copy(data, header)

	// 填充隨機數據
	for i := len(header); i < targetSize-2; i++ {
		data[i] = byte(i % 256)
	}

	// JPEG文件尾
	data[targetSize-2] = 0xFF
	data[targetSize-1] = 0xD9

	return data
}

// 輔助函數：創建模擬PNG數據
func createMockPNGData(width, height int, targetSize int) []byte {
	// PNG文件頭
	header := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

	// 填充數據以達到目標大小
	data := make([]byte, targetSize)
	copy(data, header)

	// 填充隨機數據
	for i := len(header); i < targetSize; i++ {
		data[i] = byte(i % 256)
	}

	return data
}

// BenchmarkImgCompressorBasicCompress 基本壓縮性能測試
func BenchmarkImgCompressorBasicCompress(b *testing.B) {
	testData := createMockJPEGData(1920, 1080, 1*1024*1024) // 1MB圖片

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ImgCompressorCompressWithDefaultOptions(testData)
		if err != nil {
			b.Errorf("壓縮失敗: %v", err)
		}
	}
}

// BenchmarkImgCompressorKYCProcess KYC處理性能測試
func BenchmarkImgCompressorKYCProcess(b *testing.B) {
	testData := createMockJPEGData(1920, 1080, 2*1024*1024) // 2MB KYC圖片

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ImgCompressorProcessKYCImage(testData)
		if err != nil {
			b.Errorf("KYC處理失敗: %v", err)
		}
	}
}

// BenchmarkImgCompressorThumbnail 縮圖生成性能測試
func BenchmarkImgCompressorThumbnail(b *testing.B) {
	testData := createMockJPEGData(2000, 1500, 3*1024*1024) // 3MB大圖

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ImgCompressorCreateThumbnail(testData, 200)
		if err != nil {
			b.Errorf("縮圖生成失敗: %v", err)
		}
	}
}
