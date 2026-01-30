package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

// 全域變數
var (
	// DefaultCompressOptions 默認壓縮選項
	DefaultCompressOptions = &ImgCompressorCompressOptions{
		MaxWidth:       1920,
		MaxHeight:      1080,
		Quality:        85,
		Format:         "jpeg",
		PreserveAspect: true,
		MaxFileSize:    2 * 1024 * 1024, // 2MB
		ResizeMethod:   "lanczos",
		Optimize:       true,
		Progressive:    true,
		AutoQuality:    true,
		ImgCompressorQualityRange: ImgCompressorQualityRange{
			Min: 60,
			Max: 95,
		},
	}

	// DefaultThumbnailOptions 默認縮圖選項
	DefaultThumbnailOptions = &ImgCompressorCompressOptions{
		MaxWidth:       200,
		MaxHeight:      200,
		Quality:        80,
		Format:         "jpeg",
		PreserveAspect: true,
		ResizeMethod:   "lanczos",
	}

	// DefaultWebOptions 默認Web優化選項
	DefaultWebOptions = &ImgCompressorCompressOptions{
		MaxWidth:     1920,
		MaxHeight:    1080,
		Quality:      85,
		Format:       "jpeg",
		MaxFileSize:  1 * 1024 * 1024, // 1MB
		AutoQuality:  true,
		Progressive:  true,
		ResizeMethod: "lanczos",
		ImgCompressorQualityRange: ImgCompressorQualityRange{
			Min: 60,
			Max: 90,
		},
	}

	// DefaultKYCOptions 默認KYC處理選項
	DefaultKYCOptions = &ImgCompressorCompressOptions{
		MaxWidth:       1920,
		MaxHeight:      1080,
		Quality:        85,
		Format:         "jpeg",
		PreserveAspect: true,
		MaxFileSize:    2 * 1024 * 1024, // 2MB
		AutoQuality:    true,
		ResizeMethod:   "lanczos",
		ImgCompressorQualityRange: ImgCompressorQualityRange{
			Min: 70,
			Max: 90,
		},
	}
)

// ImgCompressorCompressOptions 圖片壓縮選項
type ImgCompressorCompressOptions struct {
	// 基本設置
	MaxWidth       int    `json:"max_width"`       // 最大寬度
	MaxHeight      int    `json:"max_height"`      // 最大高度
	Quality        int    `json:"quality"`         // JPEG品質 (1-100)
	Format         string `json:"format"`          // 輸出格式 (jpeg, png, gif)
	PreserveAspect bool   `json:"preserve_aspect"` // 保持寬高比

	// 高級設置
	MaxFileSize  int64  `json:"max_file_size"` // 最大文件大小（字節）
	ResizeMethod string `json:"resize_method"` // 縮放算法
	Optimize     bool   `json:"optimize"`      // 優化輸出
	Progressive  bool   `json:"progressive"`   // 漸進式JPEG

	// 自動調整
	AutoQuality               bool                      `json:"auto_quality"`  // 自動調整品質
	ImgCompressorQualityRange ImgCompressorQualityRange `json:"quality_range"` // 品質範圍
}

// ImgCompressorQualityRange 品質範圍
type ImgCompressorQualityRange struct {
	Min int `json:"min"` // 最低品質
	Max int `json:"max"` // 最高品質
}

// ImgCompressorCompressResult 壓縮結果
type ImgCompressorCompressResult struct {
	Data             []byte  `json:"-"`                 // 壓縮後的數據
	OriginalSize     int64   `json:"original_size"`     // 原始文件大小
	CompressedSize   int64   `json:"compressed_size"`   // 壓縮後大小
	CompressionRatio float64 `json:"compression_ratio"` // 壓縮比
	OriginalWidth    int     `json:"original_width"`    // 原始寬度
	OriginalHeight   int     `json:"original_height"`   // 原始高度
	NewWidth         int     `json:"new_width"`         // 新寬度
	NewHeight        int     `json:"new_height"`        // 新高度
	Format           string  `json:"format"`            // 輸出格式
	Quality          int     `json:"quality"`           // 實際使用的品質
	ProcessTime      int64   `json:"process_time_ms"`   // 處理時間（毫秒）
}

// ImgCompressorImageInfo 圖片信息
type ImgCompressorImageInfo struct {
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Format      string  `json:"format"`
	Size        int64   `json:"size"`
	AspectRatio float64 `json:"aspect_ratio"`
}

// ImgCompressorCompressFromBytes 從字節數據壓縮圖片
func ImgCompressorCompressFromBytes(data []byte, options *ImgCompressorCompressOptions) (*ImgCompressorCompressResult, error) {
	if options == nil {
		options = DefaultCompressOptions
	}

	startTime := time.Now()

	// 解碼圖片
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("解碼圖片失敗: %w", err)
	}

	originalBounds := img.Bounds()
	originalWidth := originalBounds.Dx()
	originalHeight := originalBounds.Dy()

	// 計算新尺寸
	newWidth, newHeight := imgCompressorCalculateNewSize(originalWidth, originalHeight, options)

	// 縮放圖片
	var resizedImg image.Image
	if newWidth != originalWidth || newHeight != originalHeight {
		resizedImg = imgCompressorResizeImage(img, newWidth, newHeight, options.ResizeMethod)
	} else {
		resizedImg = img
	}

	// 設置輸出格式
	outputFormat := options.Format
	if outputFormat == "" {
		outputFormat = format // 保持原格式
	}

	// 壓縮圖片
	compressedData, actualQuality, err := imgCompressorCompressImage(resizedImg, outputFormat, options)
	if err != nil {
		return nil, fmt.Errorf("壓縮圖片失敗: %w", err)
	}

	// 如果啟用自動品質調整且文件過大，則降低品質重新壓縮
	if options.AutoQuality && options.MaxFileSize > 0 && int64(len(compressedData)) > options.MaxFileSize {
		compressedData, actualQuality, err = imgCompressorAutoAdjustQuality(resizedImg, outputFormat, options)
		if err != nil {
			return nil, fmt.Errorf("自動品質調整失敗: %w", err)
		}
	}

	processTime := time.Since(startTime).Milliseconds()

	return &ImgCompressorCompressResult{
		Data:             compressedData,
		OriginalSize:     int64(len(data)),
		CompressedSize:   int64(len(compressedData)),
		CompressionRatio: float64(len(compressedData)) / float64(len(data)),
		OriginalWidth:    originalWidth,
		OriginalHeight:   originalHeight,
		NewWidth:         newWidth,
		NewHeight:        newHeight,
		Format:           outputFormat,
		Quality:          actualQuality,
		ProcessTime:      processTime,
	}, nil
}

// CompressFromFile 從文件壓縮圖片
func ImgCompressorCompressFromFile(inputPath string, options *ImgCompressorCompressOptions) (*ImgCompressorCompressResult, error) {
	// 讀取文件
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("讀取文件失敗: %w", err)
	}

	return ImgCompressorCompressFromBytes(data, options)
}

// CompressAndSave 壓縮圖片並保存到文件
func ImgCompressorCompressAndSave(inputPath, outputPath string, options *ImgCompressorCompressOptions) (*ImgCompressorCompressResult, error) {
	result, err := ImgCompressorCompressFromFile(inputPath, options)
	if err != nil {
		return nil, err
	}

	// 保存到文件
	if err := os.WriteFile(outputPath, result.Data, 0644); err != nil {
		return nil, fmt.Errorf("保存文件失敗: %w", err)
	}

	return result, nil
}

// BatchCompress 批量壓縮圖片
func ImgCompressorBatchCompress(files []string, options *ImgCompressorCompressOptions) ([]*ImgCompressorCompressResult, error) {
	results := make([]*ImgCompressorCompressResult, 0, len(files))

	for _, file := range files {
		result, err := ImgCompressorCompressFromFile(file, options)
		if err != nil {
			return nil, fmt.Errorf("壓縮文件 %s 失敗: %w", file, err)
		}
		results = append(results, result)
	}

	return results, nil
}

// CreateThumbnail 創建縮圖
func ImgCompressorCreateThumbnail(data []byte, size int) (*ImgCompressorCompressResult, error) {
	options := &ImgCompressorCompressOptions{
		MaxWidth:       size,
		MaxHeight:      size,
		Quality:        80,
		Format:         "jpeg",
		PreserveAspect: true,
		ResizeMethod:   "lanczos",
	}

	return ImgCompressorCompressFromBytes(data, options)
}

// OptimizeForWeb 針對Web優化圖片
func ImgCompressorOptimizeForWeb(data []byte, targetSize int64) (*ImgCompressorCompressResult, error) {
	options := *DefaultWebOptions // 複製默認選項
	options.MaxFileSize = targetSize

	return ImgCompressorCompressFromBytes(data, &options)
}

// GetImageInfo 獲取圖片信息
func ImgCompressorGetImageInfo(data []byte) (*ImgCompressorImageInfo, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("解碼圖片失敗: %w", err)
	}

	bounds := img.Bounds()
	return &ImgCompressorImageInfo{
		Width:       bounds.Dx(),
		Height:      bounds.Dy(),
		Format:      format,
		Size:        int64(len(data)),
		AspectRatio: float64(bounds.Dx()) / float64(bounds.Dy()),
	}, nil
}

// ProcessKYCImage KYC圖片專用處理
func ImgCompressorProcessKYCImage(data []byte) (*ImgCompressorCompressResult, error) {
	return ImgCompressorCompressFromBytes(data, DefaultKYCOptions)
}

// imgCompressorCalculateNewSize 計算新的圖片尺寸
func imgCompressorCalculateNewSize(originalWidth, originalHeight int, options *ImgCompressorCompressOptions) (int, int) {
	if options.MaxWidth <= 0 && options.MaxHeight <= 0 {
		return originalWidth, originalHeight
	}

	// 如果圖片已經小於限制，不需要縮放
	if (options.MaxWidth <= 0 || originalWidth <= options.MaxWidth) &&
		(options.MaxHeight <= 0 || originalHeight <= options.MaxHeight) {
		return originalWidth, originalHeight
	}

	if !options.PreserveAspect {
		// 不保持寬高比，直接使用最大值
		width := options.MaxWidth
		height := options.MaxHeight
		if width <= 0 {
			width = originalWidth
		}
		if height <= 0 {
			height = originalHeight
		}
		return width, height
	}

	// 保持寬高比計算
	ratio := float64(originalWidth) / float64(originalHeight)

	newWidth := originalWidth
	newHeight := originalHeight

	if options.MaxWidth > 0 && newWidth > options.MaxWidth {
		newWidth = options.MaxWidth
		newHeight = int(float64(newWidth) / ratio)
	}

	if options.MaxHeight > 0 && newHeight > options.MaxHeight {
		newHeight = options.MaxHeight
		newWidth = int(float64(newHeight) * ratio)
	}

	return newWidth, newHeight
}

// imgCompressorResizeImage 縮放圖片
func imgCompressorResizeImage(img image.Image, width, height int, method string) image.Image {
	switch method {
	case "nearest":
		return resize.Resize(uint(width), uint(height), img, resize.NearestNeighbor)
	case "bilinear":
		return resize.Resize(uint(width), uint(height), img, resize.Bilinear)
	case "bicubic":
		return resize.Resize(uint(width), uint(height), img, resize.Bicubic)
	case "mitchell":
		return resize.Resize(uint(width), uint(height), img, resize.MitchellNetravali)
	case "lanczos2":
		return resize.Resize(uint(width), uint(height), img, resize.Lanczos2)
	case "lanczos3":
		return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	case "imaging_lanczos":
		return imaging.Resize(img, width, height, imaging.Lanczos)
	case "imaging_linear":
		return imaging.Resize(img, width, height, imaging.Linear)
	case "imaging_cubic":
		return imaging.Resize(img, width, height, imaging.CatmullRom)
	default: // "lanczos"
		return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	}
}

// imgCompressorCompressImage 壓縮圖片到指定格式
func imgCompressorCompressImage(img image.Image, format string, options *ImgCompressorCompressOptions) ([]byte, int, error) {
	var buf bytes.Buffer
	var actualQuality int

	switch format {
	case "jpeg", "jpg":
		quality := options.Quality
		if quality <= 0 || quality > 100 {
			quality = 85
		}
		actualQuality = quality

		opts := &jpeg.Options{Quality: quality}
		if err := jpeg.Encode(&buf, img, opts); err != nil {
			return nil, 0, err
		}

	case "png":
		actualQuality = 100 // PNG是無損的
		encoder := &png.Encoder{
			CompressionLevel: png.BestCompression,
		}
		if err := encoder.Encode(&buf, img); err != nil {
			return nil, 0, err
		}

	case "gif":
		actualQuality = 100 // GIF品質固定
		if err := gif.Encode(&buf, img, nil); err != nil {
			return nil, 0, err
		}

	default:
		return nil, 0, fmt.Errorf("不支持的格式: %s", format)
	}

	return buf.Bytes(), actualQuality, nil
}

// imgCompressorAutoAdjustQuality 自動調整品質以滿足文件大小要求
func imgCompressorAutoAdjustQuality(img image.Image, format string, options *ImgCompressorCompressOptions) ([]byte, int, error) {
	if format != "jpeg" && format != "jpg" {
		// 非JPEG格式無法調整品質
		return imgCompressorCompressImage(img, format, options)
	}

	minQuality := options.ImgCompressorQualityRange.Min
	maxQuality := options.ImgCompressorQualityRange.Max

	if minQuality <= 0 {
		minQuality = 60
	}
	if maxQuality <= 0 || maxQuality > 100 {
		maxQuality = 95
	}

	// 二分搜索最佳品質
	bestQuality := maxQuality
	var bestData []byte

	for low, high := minQuality, maxQuality; low <= high; {
		mid := (low + high) / 2

		testOptions := *options
		testOptions.Quality = mid

		data, _, err := imgCompressorCompressImage(img, format, &testOptions)
		if err != nil {
			return nil, 0, err
		}

		if int64(len(data)) <= options.MaxFileSize {
			// 文件大小符合要求，嘗試提高品質
			bestQuality = mid
			bestData = data
			low = mid + 1
		} else {
			// 文件過大，降低品質
			high = mid - 1
		}
	}

	if bestData == nil {
		// 即使最低品質也無法滿足大小要求
		testOptions := *options
		testOptions.Quality = minQuality
		data, _, err := imgCompressorCompressImage(img, format, &testOptions)
		return data, minQuality, err
	}

	return bestData, bestQuality, nil
}

// ResizeMode 縮放模式
type ImgCompressorResizeMode int

const (
	ImgCompressorResizeModeStretch ImgCompressorResizeMode = iota // 拉伸
	ImgCompressorResizeModeFit                                    // 適應（保持比例）
	ImgCompressorResizeModeFill                                   // 填充（可能裁剪）
	ImgCompressorResizeModeCenter                                 // 居中
)

// ImgCompressorAdvancedCompressOptions 高級壓縮選項
type ImgCompressorAdvancedCompressOptions struct {
	*ImgCompressorCompressOptions
	ResizeMode      ImgCompressorResizeMode `json:"resize_mode"`
	BackgroundColor string                  `json:"background_color"` // 背景色（十六進制）
	Sharpen         float64                 `json:"sharpen"`          // 銳化強度 (0-2)
	Blur            float64                 `json:"blur"`             // 模糊半徑
	Brightness      float64                 `json:"brightness"`       // 亮度調整 (-100 to 100)
	Contrast        float64                 `json:"contrast"`         // 對比度調整 (-100 to 100)
	Saturation      float64                 `json:"saturation"`       // 飽和度調整 (-100 to 100)
}

// AdvancedCompress 高級圖片壓縮（使用imaging庫的高級功能）
func ImgCompressorAdvancedCompress(data []byte, options *ImgCompressorAdvancedCompressOptions) (*ImgCompressorCompressResult, error) {
	if options == nil || options.ImgCompressorCompressOptions == nil {
		return ImgCompressorCompressFromBytes(data, DefaultCompressOptions)
	}

	startTime := time.Now()

	// 解碼圖片
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("解碼圖片失敗: %w", err)
	}

	originalBounds := img.Bounds()
	originalWidth := originalBounds.Dx()
	originalHeight := originalBounds.Dy()

	// 應用圖像增強
	enhancedImg := imgCompressorApplyImageEnhancements(img, options)

	// 計算新尺寸
	newWidth, newHeight := imgCompressorCalculateNewSize(originalWidth, originalHeight, options.ImgCompressorCompressOptions)

	// 高級縮放
	resizedImg := imgCompressorAdvancedResize(enhancedImg, newWidth, newHeight, options)

	// 設置輸出格式
	outputFormat := options.Format
	if outputFormat == "" {
		outputFormat = format
	}

	// 壓縮圖片
	compressedData, actualQuality, err := imgCompressorCompressImage(resizedImg, outputFormat, options.ImgCompressorCompressOptions)
	if err != nil {
		return nil, fmt.Errorf("壓縮圖片失敗: %w", err)
	}

	// 自動品質調整
	if options.AutoQuality && options.MaxFileSize > 0 && int64(len(compressedData)) > options.MaxFileSize {
		compressedData, actualQuality, err = imgCompressorAutoAdjustQuality(resizedImg, outputFormat, options.ImgCompressorCompressOptions)
		if err != nil {
			return nil, fmt.Errorf("自動品質調整失敗: %w", err)
		}
	}

	processTime := time.Since(startTime).Milliseconds()

	return &ImgCompressorCompressResult{
		Data:             compressedData,
		OriginalSize:     int64(len(data)),
		CompressedSize:   int64(len(compressedData)),
		CompressionRatio: float64(len(compressedData)) / float64(len(data)),
		OriginalWidth:    originalWidth,
		OriginalHeight:   originalHeight,
		NewWidth:         newWidth,
		NewHeight:        newHeight,
		Format:           outputFormat,
		Quality:          actualQuality,
		ProcessTime:      processTime,
	}, nil
}

// imgCompressorApplyImageEnhancements 應用圖像增強
func imgCompressorApplyImageEnhancements(img image.Image, options *ImgCompressorAdvancedCompressOptions) image.Image {
	result := img

	// 銳化
	if options.Sharpen > 0 {
		result = imaging.Sharpen(result, options.Sharpen)
	}

	// 模糊
	if options.Blur > 0 {
		result = imaging.Blur(result, options.Blur)
	}

	// 亮度調整
	if options.Brightness != 0 {
		result = imaging.AdjustBrightness(result, options.Brightness)
	}

	// 對比度調整
	if options.Contrast != 0 {
		result = imaging.AdjustContrast(result, options.Contrast)
	}

	// 飽和度調整
	if options.Saturation != 0 {
		result = imaging.AdjustSaturation(result, options.Saturation)
	}

	return result
}

// imgCompressorAdvancedResize 高級縮放
func imgCompressorAdvancedResize(img image.Image, width, height int, options *ImgCompressorAdvancedCompressOptions) image.Image {
	switch options.ResizeMode {
	case ImgCompressorResizeModeStretch:
		return imaging.Resize(img, width, height, imaging.Lanczos)
	case ImgCompressorResizeModeFit:
		return imaging.Fit(img, width, height, imaging.Lanczos)
	case ImgCompressorResizeModeFill:
		return imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)
	case ImgCompressorResizeModeCenter:
		return imaging.Thumbnail(img, width, height, imaging.Lanczos)
	default:
		return imaging.Fit(img, width, height, imaging.Lanczos)
	}
}

// 便利函數

// ImgCompressorCompressWithDefaultOptions 使用默認選項壓縮圖片
func ImgCompressorCompressWithDefaultOptions(data []byte) (*ImgCompressorCompressResult, error) {
	return ImgCompressorCompressFromBytes(data, DefaultCompressOptions)
}

// ImgCompressorQuickCompress 快速壓縮（保持原尺寸，僅降低品質）
func ImgCompressorQuickCompress(data []byte, quality int) (*ImgCompressorCompressResult, error) {
	options := &ImgCompressorCompressOptions{
		Quality: quality,
		Format:  "jpeg",
	}
	return ImgCompressorCompressFromBytes(data, options)
}

// ImgCompressorResizeOnly 僅調整尺寸（不壓縮品質）
func ImgCompressorResizeOnly(data []byte, width, height int) (*ImgCompressorCompressResult, error) {
	options := &ImgCompressorCompressOptions{
		MaxWidth:       width,
		MaxHeight:      height,
		Quality:        95,
		Format:         "jpeg",
		PreserveAspect: true,
	}
	return ImgCompressorCompressFromBytes(data, options)
}

// ImgCompressorConvertFormat 轉換圖片格式
func ImgCompressorConvertFormat(data []byte, targetFormat string) (*ImgCompressorCompressResult, error) {
	options := &ImgCompressorCompressOptions{
		Format:  targetFormat,
		Quality: 95,
	}
	return ImgCompressorCompressFromBytes(data, options)
}

// ImgCompressorFormatFileSize 格式化文件大小
func ImgCompressorFormatFileSize(size int64) string {
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
