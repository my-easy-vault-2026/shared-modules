package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/image/webp"
)

// ImgConverterOptions 圖片轉換選項
type ImgConverterOptions struct {
	Quality      int    // JPEG 品質 (1-100)
	OutputFormat string // 輸出格式 "jpeg", "jpg"
}

// ImgConverterResult 圖片轉換結果
type ImgConverterResult struct {
	Data            []byte `json:"-"`                // 轉換後的圖片數據
	OriginalSize    int64  `json:"original_size"`    // 原始文件大小
	ConvertedSize   int64  `json:"converted_size"`   // 轉換後文件大小
	OriginalFormat  string `json:"original_format"`  // 原始格式
	ConvertedFormat string `json:"converted_format"` // 轉換後格式
	Width           int    `json:"width"`            // 圖片寬度
	Height          int    `json:"height"`           // 圖片高度
	ProcessTime     int64  `json:"process_time"`     // 處理時間(毫秒)
}

// ImgConverterImageInfo 圖片資訊
type ImgConverterImageInfo struct {
	Format string `json:"format"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int    `json:"size"`
}

// ImgConverterDefaultOptions 預設轉換選項
var ImgConverterDefaultOptions = &ImgConverterOptions{
	Quality:      85,
	OutputFormat: "jpeg",
}

// ImgConverterFromBytes 從字節數據轉換圖片
func ImgConverterFromBytes(data []byte, options *ImgConverterOptions) (*ImgConverterResult, error) {
	if options == nil {
		options = ImgConverterDefaultOptions
	}

	startTime := time.Now()

	// 檢測原始格式
	originalFormat, err := ImgConverterDetectFormat(data)
	if err != nil {
		return nil, fmt.Errorf("檢測圖片格式失敗: %w", err)
	}

	// 如果已經是目標格式，直接返回
	if strings.ToLower(originalFormat) == "jpeg" || strings.ToLower(originalFormat) == "jpg" {
		if options.OutputFormat == "jpeg" || options.OutputFormat == "jpg" {
			img, err := imgConverterDecodeImage(data, originalFormat)
			if err != nil {
				return nil, fmt.Errorf("解碼圖片失敗: %w", err)
			}

			bounds := img.Bounds()
			processTime := time.Since(startTime).Milliseconds()

			return &ImgConverterResult{
				Data:            data,
				OriginalSize:    int64(len(data)),
				ConvertedSize:   int64(len(data)),
				OriginalFormat:  originalFormat,
				ConvertedFormat: "jpeg",
				Width:           bounds.Dx(),
				Height:          bounds.Dy(),
				ProcessTime:     processTime,
			}, nil
		}
	}

	// 解碼圖片
	img, err := imgConverterDecodeImage(data, originalFormat)
	if err != nil {
		return nil, fmt.Errorf("解碼圖片失敗: %w", err)
	}

	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	// 轉換為目標格式
	convertedData, err := imgConverterEncodeToJPEG(img, options.Quality)
	if err != nil {
		return nil, fmt.Errorf("編碼圖片失敗: %w", err)
	}

	processTime := time.Since(startTime).Milliseconds()

	return &ImgConverterResult{
		Data:            convertedData,
		OriginalSize:    int64(len(data)),
		ConvertedSize:   int64(len(convertedData)),
		OriginalFormat:  originalFormat,
		ConvertedFormat: options.OutputFormat,
		Width:           originalWidth,
		Height:          originalHeight,
		ProcessTime:     processTime,
	}, nil
}

// ImgConverterFromFile 從文件轉換圖片
func ImgConverterFromFile(inputPath string, options *ImgConverterOptions) (*ImgConverterResult, error) {
	// 讀取文件
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("讀取文件失敗: %w", err)
	}

	return ImgConverterFromBytes(data, options)
}

// ImgConverterAndSave 轉換圖片並保存到文件
func ImgConverterAndSave(inputPath, outputPath string, options *ImgConverterOptions) (*ImgConverterResult, error) {
	result, err := ImgConverterFromFile(inputPath, options)
	if err != nil {
		return nil, err
	}

	// 確保輸出目錄存在
	outputDir := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return nil, fmt.Errorf("創建輸出目錄失敗: %w", err)
	}

	// 保存到文件
	if err := os.WriteFile(outputPath, result.Data, 0644); err != nil {
		return nil, fmt.Errorf("保存文件失敗: %w", err)
	}

	return result, nil
}

// ImgConverterFromReader 從 io.Reader 轉換圖片
func ImgConverterFromReader(reader io.Reader, options *ImgConverterOptions) (*ImgConverterResult, error) {
	// 讀取所有資料
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("讀取數據失敗: %w", err)
	}

	return ImgConverterFromBytes(data, options)
}

// ImgConverterFromBase64 從 Base64 轉換圖片
func ImgConverterFromBase64(base64Data string, options *ImgConverterOptions) (*ImgConverterResult, error) {
	// 解析 Base64
	imageData, err := imgConverterParseBase64Image(base64Data)
	if err != nil {
		return nil, fmt.Errorf("解析Base64失敗: %w", err)
	}

	return ImgConverterFromBytes(imageData, options)
}

// ImgConverterToBase64 轉換圖片並返回 Base64
func ImgConverterToBase64(data []byte, options *ImgConverterOptions) (string, *ImgConverterResult, error) {
	result, err := ImgConverterFromBytes(data, options)
	if err != nil {
		return "", nil, err
	}

	base64String := base64.StdEncoding.EncodeToString(result.Data)
	return base64String, result, nil
}

// ImgConverterBatchFromBytes 批量轉換圖片
func ImgConverterBatchFromBytes(inputs [][]byte, options *ImgConverterOptions) ([]*ImgConverterResult, error) {
	results := make([]*ImgConverterResult, len(inputs))
	var errors []error

	for i, input := range inputs {
		result, err := ImgConverterFromBytes(input, options)
		if err != nil {
			errors = append(errors, fmt.Errorf("轉換圖片 %d 失敗: %w", i, err))
			continue
		}
		results[i] = result
	}

	if len(errors) > 0 {
		return results, fmt.Errorf("批量轉換有 %d 個錯誤: %v", len(errors), errors)
	}

	return results, nil
}

// ImgConverterGetImageInfo 獲取圖片資訊
func ImgConverterGetImageInfo(data []byte) (*ImgConverterImageInfo, error) {
	format, err := ImgConverterDetectFormat(data)
	if err != nil {
		return nil, err
	}

	img, err := imgConverterDecodeImage(data, format)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	return &ImgConverterImageInfo{
		Format: format,
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
		Size:   len(data),
	}, nil
}

// 內部輔助函數

// ImgConverterDetectFormat 檢測圖片格式
func ImgConverterDetectFormat(data []byte) (string, error) {
	if len(data) < 12 {
		return "", fmt.Errorf("無效的圖片數據")
	}

	// PNG 格式檢測
	if bytes.HasPrefix(data, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "png", nil
	}

	// JPEG 格式檢測
	if bytes.HasPrefix(data, []byte{0xFF, 0xD8, 0xFF}) {
		return "jpeg", nil
	}

	// WebP 格式檢測
	if bytes.HasPrefix(data, []byte("RIFF")) && bytes.Contains(data[8:12], []byte("WEBP")) {
		return "webp", nil
	}

	// 使用 http.DetectContentType 作為後備
	contentType := http.DetectContentType(data)
	switch contentType {
	case "image/png":
		return "png", nil
	case "image/jpeg":
		return "jpeg", nil
	case "image/webp":
		return "webp", nil
	default:
		return "", fmt.Errorf("不支援的圖片格式: %s", contentType)
	}
}

// imgConverterDecodeImage 根據格式解碼圖片
func imgConverterDecodeImage(data []byte, format string) (image.Image, error) {
	reader := bytes.NewReader(data)

	switch strings.ToLower(format) {
	case "png":
		return png.Decode(reader)
	case "jpeg", "jpg":
		return jpeg.Decode(reader)
	case "webp":
		return webp.Decode(reader)
	default:
		return nil, fmt.Errorf("不支援的格式: %s", format)
	}
}

// imgConverterEncodeToJPEG 將圖片編碼為 JPEG
func imgConverterEncodeToJPEG(img image.Image, quality int) ([]byte, error) {
	var buf bytes.Buffer

	options := &jpeg.Options{
		Quality: quality,
	}

	err := jpeg.Encode(&buf, img, options)
	if err != nil {
		return nil, fmt.Errorf("JPEG編碼失敗: %w", err)
	}

	return buf.Bytes(), nil
}

// imgConverterParseBase64Image 解析 Base64 圖片資料
func imgConverterParseBase64Image(base64Data string) ([]byte, error) {
	// 移除 data URL 前綴 (如果有的話)
	if strings.Contains(base64Data, ",") {
		parts := strings.Split(base64Data, ",")
		if len(parts) == 2 {
			base64Data = parts[1]
		}
	}

	// 解碼 Base64
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, fmt.Errorf("Base64解碼失敗: %w", err)
	}

	return imageData, nil
}

// 便利函數

// ImgConverterQuickConvert 快速轉換為 JPEG (使用預設選項)
func ImgConverterQuickConvert(data []byte) ([]byte, error) {
	result, err := ImgConverterFromBytes(data, ImgConverterDefaultOptions)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// ImgConverterQuickConvertBase64 快速轉換 Base64 為 JPEG Base64
func ImgConverterQuickConvertBase64(base64Data string) (string, error) {
	imageData, err := imgConverterParseBase64Image(base64Data)
	if err != nil {
		return "", err
	}

	result, err := ImgConverterFromBytes(imageData, ImgConverterDefaultOptions)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(result.Data), nil
}
