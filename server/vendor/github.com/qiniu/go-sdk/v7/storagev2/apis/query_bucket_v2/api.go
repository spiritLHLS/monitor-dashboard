// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 查询存储空间服务域名
package query_bucket_v2

import (
	"encoding/json"
	errors "github.com/qiniu/go-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Bucket    string // 存储空间名称
	AccessKey string // Access Key
}

// 获取 API 所用的响应
type Response struct {
	RegionId     string       // 区域 ID
	TimeToLive   int64        // 查询结果的 TTL
	UpDomains    UpDomains    // 上传域名
	IoDomains    IoDomains    // 下载域名
	IoSrcDomains IoSrcDomains // 源站下载域名
	RsDomains    RsDomains    // 对象管理域名
	RsfDomains   RsfDomains   // 对象列举域名
	ApiDomains   ApiDomains   // API 域名
}

// 空间级别的主加速上传域名列表
type MainBucketAcceleratedUpDomains = []string

// 空间级别的备用加速上传域名列表
type BackupBucketAcceleratedUpDomains = []string

// 空间级别的加速上传域名
type BucketAcceleratedUpDomains struct {
	MainAcceleratedUpDomains   MainBucketAcceleratedUpDomains   // 空间级别的主加速上传域名列表
	BackupAcceleratedUpDomains BackupBucketAcceleratedUpDomains // 空间级别的备用加速上传域名列表
}
type jsonBucketAcceleratedUpDomains struct {
	MainAcceleratedUpDomains   MainBucketAcceleratedUpDomains   `json:"main"`   // 空间级别的主加速上传域名列表
	BackupAcceleratedUpDomains BackupBucketAcceleratedUpDomains `json:"backup"` // 空间级别的备用加速上传域名列表
}

func (j *BucketAcceleratedUpDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonBucketAcceleratedUpDomains{MainAcceleratedUpDomains: j.MainAcceleratedUpDomains, BackupAcceleratedUpDomains: j.BackupAcceleratedUpDomains})
}
func (j *BucketAcceleratedUpDomains) UnmarshalJSON(data []byte) error {
	var nj jsonBucketAcceleratedUpDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainAcceleratedUpDomains = nj.MainAcceleratedUpDomains
	j.BackupAcceleratedUpDomains = nj.BackupAcceleratedUpDomains
	return nil
}
func (j *BucketAcceleratedUpDomains) validate() error {
	if len(j.MainAcceleratedUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainAcceleratedUpDomains"}
	}
	if len(j.BackupAcceleratedUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "BackupAcceleratedUpDomains"}
	}
	return nil
}

// 主加速上传域名列表
type MainAcceleratedUpDomains = []string

// 备用加速上传域名列表
type BackupAcceleratedUpDomains = []string

// 加速上传域名
type AcceleratedUpDomains struct {
	MainAcceleratedUpDomains   MainAcceleratedUpDomains   // 主加速上传域名列表
	BackupAcceleratedUpDomains BackupAcceleratedUpDomains // 备用加速上传域名列表
}
type jsonAcceleratedUpDomains struct {
	MainAcceleratedUpDomains   MainAcceleratedUpDomains   `json:"main"`   // 主加速上传域名列表
	BackupAcceleratedUpDomains BackupAcceleratedUpDomains `json:"backup"` // 备用加速上传域名列表
}

func (j *AcceleratedUpDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonAcceleratedUpDomains{MainAcceleratedUpDomains: j.MainAcceleratedUpDomains, BackupAcceleratedUpDomains: j.BackupAcceleratedUpDomains})
}
func (j *AcceleratedUpDomains) UnmarshalJSON(data []byte) error {
	var nj jsonAcceleratedUpDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainAcceleratedUpDomains = nj.MainAcceleratedUpDomains
	j.BackupAcceleratedUpDomains = nj.BackupAcceleratedUpDomains
	return nil
}
func (j *AcceleratedUpDomains) validate() error {
	if len(j.MainAcceleratedUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainAcceleratedUpDomains"}
	}
	if len(j.BackupAcceleratedUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "BackupAcceleratedUpDomains"}
	}
	return nil
}

// 主源站上传域名列表
type MainSourceUpDomains = []string

// 备用源站上传域名列表
type BackupSourceUpDomains = []string

// 源站上传域名
type SourceUpDomains struct {
	MainSourceUpDomains   MainSourceUpDomains   // 主源站上传域名列表
	BackupSourceUpDomains BackupSourceUpDomains // 备用源站上传域名列表
}
type jsonSourceUpDomains struct {
	MainSourceUpDomains   MainSourceUpDomains   `json:"main"`   // 主源站上传域名列表
	BackupSourceUpDomains BackupSourceUpDomains `json:"backup"` // 备用源站上传域名列表
}

func (j *SourceUpDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonSourceUpDomains{MainSourceUpDomains: j.MainSourceUpDomains, BackupSourceUpDomains: j.BackupSourceUpDomains})
}
func (j *SourceUpDomains) UnmarshalJSON(data []byte) error {
	var nj jsonSourceUpDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainSourceUpDomains = nj.MainSourceUpDomains
	j.BackupSourceUpDomains = nj.BackupSourceUpDomains
	return nil
}
func (j *SourceUpDomains) validate() error {
	if len(j.MainSourceUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainSourceUpDomains"}
	}
	if len(j.BackupSourceUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "BackupSourceUpDomains"}
	}
	return nil
}

// 已经过时的主加速上传域名列表
type OldMainAcceleratedUpDomains = []string

// 已经过时的加速上传域名
type OldAcceleratedDomains struct {
	OldMainAcceleratedUpDomains OldMainAcceleratedUpDomains // 主加速上传域名列表
	Info                        string                      // 描述信息
}

// 已经过时的加速上传域名
type OldAcceleratedUpDomains = OldAcceleratedDomains
type jsonOldAcceleratedDomains struct {
	OldMainAcceleratedUpDomains OldMainAcceleratedUpDomains `json:"main"` // 主加速上传域名列表
	Info                        string                      `json:"info"` // 描述信息
}

func (j *OldAcceleratedDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonOldAcceleratedDomains{OldMainAcceleratedUpDomains: j.OldMainAcceleratedUpDomains, Info: j.Info})
}
func (j *OldAcceleratedDomains) UnmarshalJSON(data []byte) error {
	var nj jsonOldAcceleratedDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.OldMainAcceleratedUpDomains = nj.OldMainAcceleratedUpDomains
	j.Info = nj.Info
	return nil
}
func (j *OldAcceleratedDomains) validate() error {
	if len(j.OldMainAcceleratedUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "OldMainAcceleratedUpDomains"}
	}
	if j.Info == "" {
		return errors.MissingRequiredFieldError{Name: "Info"}
	}
	return nil
}

// 已经过时的主源站上传域名列表
type OldMainSourceUpDomains = []string

// 已经过时的源站上传域名
type OldSourceDomains struct {
	OldMainSourceUpDomains OldMainSourceUpDomains // 主源站上传域名列表
	Info                   string                 // 描述信息
}

// 已经过时的源站上传域名
type OldSourceUpDomains = OldSourceDomains
type jsonOldSourceDomains struct {
	OldMainSourceUpDomains OldMainSourceUpDomains `json:"main"` // 主源站上传域名列表
	Info                   string                 `json:"info"` // 描述信息
}

func (j *OldSourceDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonOldSourceDomains{OldMainSourceUpDomains: j.OldMainSourceUpDomains, Info: j.Info})
}
func (j *OldSourceDomains) UnmarshalJSON(data []byte) error {
	var nj jsonOldSourceDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.OldMainSourceUpDomains = nj.OldMainSourceUpDomains
	j.Info = nj.Info
	return nil
}
func (j *OldSourceDomains) validate() error {
	if len(j.OldMainSourceUpDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "OldMainSourceUpDomains"}
	}
	if j.Info == "" {
		return errors.MissingRequiredFieldError{Name: "Info"}
	}
	return nil
}

// 上传域名
type UpDomains struct {
	BucketAcceleratedUpDomains BucketAcceleratedUpDomains // 空间级别的加速上传域名
	AcceleratedUpDomains       AcceleratedUpDomains       // 加速上传域名
	SourceUpDomains            SourceUpDomains            // 源站上传域名
	OldAcceleratedDomains      OldAcceleratedUpDomains    // 已经过时的加速上传域名
	OldSourceDomains           OldSourceUpDomains         // 已经过时的源站上传域名
}
type jsonUpDomains struct {
	BucketAcceleratedUpDomains BucketAcceleratedUpDomains `json:"bucket_acc"` // 空间级别的加速上传域名
	AcceleratedUpDomains       AcceleratedUpDomains       `json:"acc"`        // 加速上传域名
	SourceUpDomains            SourceUpDomains            `json:"src"`        // 源站上传域名
	OldAcceleratedDomains      OldAcceleratedUpDomains    `json:"old_acc"`    // 已经过时的加速上传域名
	OldSourceDomains           OldSourceUpDomains         `json:"old_src"`    // 已经过时的源站上传域名
}

func (j *UpDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonUpDomains{BucketAcceleratedUpDomains: j.BucketAcceleratedUpDomains, AcceleratedUpDomains: j.AcceleratedUpDomains, SourceUpDomains: j.SourceUpDomains, OldAcceleratedDomains: j.OldAcceleratedDomains, OldSourceDomains: j.OldSourceDomains})
}
func (j *UpDomains) UnmarshalJSON(data []byte) error {
	var nj jsonUpDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.BucketAcceleratedUpDomains = nj.BucketAcceleratedUpDomains
	j.AcceleratedUpDomains = nj.AcceleratedUpDomains
	j.SourceUpDomains = nj.SourceUpDomains
	j.OldAcceleratedDomains = nj.OldAcceleratedDomains
	j.OldSourceDomains = nj.OldSourceDomains
	return nil
}
func (j *UpDomains) validate() error {
	if err := j.BucketAcceleratedUpDomains.validate(); err != nil {
		return err
	}
	if err := j.AcceleratedUpDomains.validate(); err != nil {
		return err
	}
	if err := j.SourceUpDomains.validate(); err != nil {
		return err
	}
	if err := j.OldAcceleratedDomains.validate(); err != nil {
		return err
	}
	if err := j.OldSourceDomains.validate(); err != nil {
		return err
	}
	return nil
}

// 主源站下载域名列表
type MainSourceIoDomains = []string

// 源站下载域名
type SourceIoDomains struct {
	MainSourceIoDomains MainSourceIoDomains // 主源站下载域名列表
}
type jsonSourceIoDomains struct {
	MainSourceIoDomains MainSourceIoDomains `json:"main"` // 主源站下载域名列表
}

func (j *SourceIoDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonSourceIoDomains{MainSourceIoDomains: j.MainSourceIoDomains})
}
func (j *SourceIoDomains) UnmarshalJSON(data []byte) error {
	var nj jsonSourceIoDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainSourceIoDomains = nj.MainSourceIoDomains
	return nil
}
func (j *SourceIoDomains) validate() error {
	if len(j.MainSourceIoDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainSourceIoDomains"}
	}
	return nil
}

// 下载域名
type IoDomains struct {
	SourceIoDomains SourceIoDomains // 源站下载域名
}
type jsonIoDomains struct {
	SourceIoDomains SourceIoDomains `json:"src"` // 源站下载域名
}

func (j *IoDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonIoDomains{SourceIoDomains: j.SourceIoDomains})
}
func (j *IoDomains) UnmarshalJSON(data []byte) error {
	var nj jsonIoDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.SourceIoDomains = nj.SourceIoDomains
	return nil
}
func (j *IoDomains) validate() error {
	if err := j.SourceIoDomains.validate(); err != nil {
		return err
	}
	return nil
}

// 主源站下载域名列表
type MainSourceIoSrcDomains = []string

// 源站下载域名
type SourceIoSrcDomains struct {
	MainSourceIoSrcDomains MainSourceIoSrcDomains // 主源站下载域名列表
}
type jsonSourceIoSrcDomains struct {
	MainSourceIoSrcDomains MainSourceIoSrcDomains `json:"main"` // 主源站下载域名列表
}

func (j *SourceIoSrcDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonSourceIoSrcDomains{MainSourceIoSrcDomains: j.MainSourceIoSrcDomains})
}
func (j *SourceIoSrcDomains) UnmarshalJSON(data []byte) error {
	var nj jsonSourceIoSrcDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainSourceIoSrcDomains = nj.MainSourceIoSrcDomains
	return nil
}
func (j *SourceIoSrcDomains) validate() error {
	if len(j.MainSourceIoSrcDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainSourceIoSrcDomains"}
	}
	return nil
}

// 源站下载域名
type IoSrcDomains struct {
	SourceIoSrcDomains SourceIoSrcDomains // 源站下载域名
}
type jsonIoSrcDomains struct {
	SourceIoSrcDomains SourceIoSrcDomains `json:"src"` // 源站下载域名
}

func (j *IoSrcDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonIoSrcDomains{SourceIoSrcDomains: j.SourceIoSrcDomains})
}
func (j *IoSrcDomains) UnmarshalJSON(data []byte) error {
	var nj jsonIoSrcDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.SourceIoSrcDomains = nj.SourceIoSrcDomains
	return nil
}
func (j *IoSrcDomains) validate() error {
	if err := j.SourceIoSrcDomains.validate(); err != nil {
		return err
	}
	return nil
}

// 主加速对象管理域名列表
type MainAcceleratedRsDomains = []string

// 加速对象管理域名
type AcceleratedRsDomains struct {
	MainAcceleratedRsDomains MainAcceleratedRsDomains // 主加速对象管理域名列表
}
type jsonAcceleratedRsDomains struct {
	MainAcceleratedRsDomains MainAcceleratedRsDomains `json:"main"` // 主加速对象管理域名列表
}

func (j *AcceleratedRsDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonAcceleratedRsDomains{MainAcceleratedRsDomains: j.MainAcceleratedRsDomains})
}
func (j *AcceleratedRsDomains) UnmarshalJSON(data []byte) error {
	var nj jsonAcceleratedRsDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainAcceleratedRsDomains = nj.MainAcceleratedRsDomains
	return nil
}
func (j *AcceleratedRsDomains) validate() error {
	if len(j.MainAcceleratedRsDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainAcceleratedRsDomains"}
	}
	return nil
}

// 对象管理域名
type RsDomains struct {
	AcceleratedRsDomains AcceleratedRsDomains // 加速对象管理域名
}
type jsonRsDomains struct {
	AcceleratedRsDomains AcceleratedRsDomains `json:"acc"` // 加速对象管理域名
}

func (j *RsDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRsDomains{AcceleratedRsDomains: j.AcceleratedRsDomains})
}
func (j *RsDomains) UnmarshalJSON(data []byte) error {
	var nj jsonRsDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.AcceleratedRsDomains = nj.AcceleratedRsDomains
	return nil
}
func (j *RsDomains) validate() error {
	if err := j.AcceleratedRsDomains.validate(); err != nil {
		return err
	}
	return nil
}

// 主加速对象列举域名列表
type MainAcceleratedRsfDomains = []string

// 加速对象列举域名
type AcceleratedRsfDomains struct {
	MainAcceleratedRsfDomains MainAcceleratedRsfDomains // 主加速对象列举域名列表
}
type jsonAcceleratedRsfDomains struct {
	MainAcceleratedRsfDomains MainAcceleratedRsfDomains `json:"main"` // 主加速对象列举域名列表
}

func (j *AcceleratedRsfDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonAcceleratedRsfDomains{MainAcceleratedRsfDomains: j.MainAcceleratedRsfDomains})
}
func (j *AcceleratedRsfDomains) UnmarshalJSON(data []byte) error {
	var nj jsonAcceleratedRsfDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainAcceleratedRsfDomains = nj.MainAcceleratedRsfDomains
	return nil
}
func (j *AcceleratedRsfDomains) validate() error {
	if len(j.MainAcceleratedRsfDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainAcceleratedRsfDomains"}
	}
	return nil
}

// 对象列举域名
type RsfDomains struct {
	AcceleratedRsfDomains AcceleratedRsfDomains // 加速对象列举域名
}
type jsonRsfDomains struct {
	AcceleratedRsfDomains AcceleratedRsfDomains `json:"acc"` // 加速对象列举域名
}

func (j *RsfDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRsfDomains{AcceleratedRsfDomains: j.AcceleratedRsfDomains})
}
func (j *RsfDomains) UnmarshalJSON(data []byte) error {
	var nj jsonRsfDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.AcceleratedRsfDomains = nj.AcceleratedRsfDomains
	return nil
}
func (j *RsfDomains) validate() error {
	if err := j.AcceleratedRsfDomains.validate(); err != nil {
		return err
	}
	return nil
}

// 主加速 API 域名列表
type MainAcceleratedApiDomains = []string

// 加速 API 域名
type AcceleratedApiDomains struct {
	MainAcceleratedApiDomains MainAcceleratedApiDomains // 主加速 API 域名列表
}
type jsonAcceleratedApiDomains struct {
	MainAcceleratedApiDomains MainAcceleratedApiDomains `json:"main"` // 主加速 API 域名列表
}

func (j *AcceleratedApiDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonAcceleratedApiDomains{MainAcceleratedApiDomains: j.MainAcceleratedApiDomains})
}
func (j *AcceleratedApiDomains) UnmarshalJSON(data []byte) error {
	var nj jsonAcceleratedApiDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.MainAcceleratedApiDomains = nj.MainAcceleratedApiDomains
	return nil
}
func (j *AcceleratedApiDomains) validate() error {
	if len(j.MainAcceleratedApiDomains) == 0 {
		return errors.MissingRequiredFieldError{Name: "MainAcceleratedApiDomains"}
	}
	return nil
}

// API 域名
type ApiDomains struct {
	AcceleratedApiDomains AcceleratedApiDomains // 加速 API 域名
}
type jsonApiDomains struct {
	AcceleratedApiDomains AcceleratedApiDomains `json:"acc"` // 加速 API 域名
}

func (j *ApiDomains) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonApiDomains{AcceleratedApiDomains: j.AcceleratedApiDomains})
}
func (j *ApiDomains) UnmarshalJSON(data []byte) error {
	var nj jsonApiDomains
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.AcceleratedApiDomains = nj.AcceleratedApiDomains
	return nil
}
func (j *ApiDomains) validate() error {
	if err := j.AcceleratedApiDomains.validate(); err != nil {
		return err
	}
	return nil
}

// 存储空间服务域名查询结果
type BucketQueryResult = Response
type jsonResponse struct {
	RegionId     string       `json:"region"`           // 区域 ID
	TimeToLive   int64        `json:"ttl"`              // 查询结果的 TTL
	UpDomains    UpDomains    `json:"up"`               // 上传域名
	IoDomains    IoDomains    `json:"io"`               // 下载域名
	IoSrcDomains IoSrcDomains `json:"io_src,omitempty"` // 源站下载域名
	RsDomains    RsDomains    `json:"rs"`               // 对象管理域名
	RsfDomains   RsfDomains   `json:"rsf"`              // 对象列举域名
	ApiDomains   ApiDomains   `json:"api"`              // API 域名
}

func (j *Response) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonResponse{RegionId: j.RegionId, TimeToLive: j.TimeToLive, UpDomains: j.UpDomains, IoDomains: j.IoDomains, IoSrcDomains: j.IoSrcDomains, RsDomains: j.RsDomains, RsfDomains: j.RsfDomains, ApiDomains: j.ApiDomains})
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var nj jsonResponse
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.RegionId = nj.RegionId
	j.TimeToLive = nj.TimeToLive
	j.UpDomains = nj.UpDomains
	j.IoDomains = nj.IoDomains
	j.IoSrcDomains = nj.IoSrcDomains
	j.RsDomains = nj.RsDomains
	j.RsfDomains = nj.RsfDomains
	j.ApiDomains = nj.ApiDomains
	return nil
}
func (j *Response) validate() error {
	if j.RegionId == "" {
		return errors.MissingRequiredFieldError{Name: "RegionId"}
	}
	if j.TimeToLive == 0 {
		return errors.MissingRequiredFieldError{Name: "TimeToLive"}
	}
	if err := j.UpDomains.validate(); err != nil {
		return err
	}
	if err := j.IoDomains.validate(); err != nil {
		return err
	}
	if err := j.RsDomains.validate(); err != nil {
		return err
	}
	if err := j.RsfDomains.validate(); err != nil {
		return err
	}
	if err := j.ApiDomains.validate(); err != nil {
		return err
	}
	return nil
}
