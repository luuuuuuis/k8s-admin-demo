package service

import (
	"sort"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	nwv1 "k8s.io/api/networking/v1"
)

// DataSelector 封装了对数据进行排序、过滤和分页的功能
type DataSelector struct {
	GenericDataList []DataCell // 存储数据的列表
	FilterQuery     *FilterQuery // 过滤条件，具体属性看FilterQuery结构体{Name string}
	PaginateQuery   *PaginateQuery // 分页条件，具体属性看PaginateQuery结构体{Limit int,Page  int}
}

// DataCell 是数据元素的接口，用于各种资源列表的类型转换
type DataCell interface {
	GetCreation() time.Time // 获取数据元素的创建时间
	GetName() string // 获取数据元素的名称
}

// FilterQuery 定义了过滤条件，这里只有一个名称过滤条件
type FilterQuery struct {
	Name string // 名称过滤条件
}

// PaginateQuery 定义了分页条件，包括每页数据条数和页数
type PaginateQuery struct {
	Limit int // 每页数据条数
	Page  int // 页数
}

// Len 返回数据列表的长度，用于排序
func (d *DataSelector) Len() int {
	return len(d.GenericDataList)
}

// Swap 交换数据列表中的元素位置，用于排序
func (d *DataSelector) Swap(i, j int) {
	d.GenericDataList[i], d.GenericDataList[j] = d.GenericDataList[j], d.GenericDataList[i]
}

// Less 比较两个数据元素的创建时间，用于排序
func (d *DataSelector) Less(i, j int) bool {
	return d.GenericDataList[j].GetCreation().Before(d.GenericDataList[i].GetCreation())
}

// Sort 对数据列表进行排序
func (d *DataSelector) Sort() *DataSelector {
	sort.Sort(d)
	return d
}

// Filter 根据过滤条件过滤数据列表中的元素
func (d *DataSelector) Filter() *DataSelector {
	if d.FilterQuery.Name == "" {
		return d
	}

	filteredList := []DataCell{}
	for _, value := range d.GenericDataList {
		if strings.Contains(value.GetName(), d.FilterQuery.Name) {
			filteredList = append(filteredList, value)
		}
	}

	d.GenericDataList = filteredList
	return d
}

// Paginate 对数据列表进行分页
func (d *DataSelector) Paginate() *DataSelector {
	limit := d.PaginateQuery.Limit
	page := d.PaginateQuery.Page

	if limit <= 0 || page <= 0 {
		return d
	}

	startIndex := limit * (page - 1)
	endIndex := limit * page

	if len(d.GenericDataList) < endIndex {
		endIndex = len(d.GenericDataList)
	}

	d.GenericDataList = d.GenericDataList[startIndex:endIndex]
	return d
}

// podCell 是 corev1.Pod 类型的数据元素，实现了 DataCell 接口
type podCell corev1.Pod

func (p podCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func (p podCell) GetName() string {
	return p.Name
}

// deploymentCell 是 appsv1.Deployment 类型的数据元素，实现了 DataCell 接口
type deploymentCell appsv1.Deployment

func (d deploymentCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func (d deploymentCell) GetName() string {
	return d.Name
}

// 其他类型的 DataCell 实现类似，均需实现 DataCell 接口
type daemonSetCell appsv1.DaemonSet

func(d daemonSetCell) GetCreation() time.Time {
	return d.CreationTimestamp.Time
}

func(d daemonSetCell) GetName() string {
	return d.Name
}

type statefulSetCell appsv1.StatefulSet

func(s statefulSetCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func(s statefulSetCell) GetName() string {
	return s.Name
}

type serviceCell corev1.Service

func(s serviceCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func(s serviceCell) GetName() string {
	return s.Name
}

type ingressCell nwv1.Ingress

func(i ingressCell) GetCreation() time.Time {
	return i.CreationTimestamp.Time
}

func(i ingressCell) GetName() string {
	return i.Name
}

type configMapCell corev1.ConfigMap

func(c configMapCell) GetCreation() time.Time {
	return c.CreationTimestamp.Time
}

func(c configMapCell) GetName() string {
	return c.Name
}

type secretCell corev1.Secret

func(s secretCell) GetCreation() time.Time {
	return s.CreationTimestamp.Time
}

func(s secretCell) GetName() string {
	return s.Name
}

type pvcCell corev1.PersistentVolumeClaim

func(p pvcCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func(p pvcCell) GetName() string {
	return p.Name
}

type nodeCell corev1.Node

func(n nodeCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func(n nodeCell) GetName() string {
	return n.Name
}

type namespaceCell corev1.Namespace

func(n namespaceCell) GetCreation() time.Time {
	return n.CreationTimestamp.Time
}

func(n namespaceCell) GetName() string {
	return n.Name
}

type pvCell corev1.PersistentVolume

func(p pvCell) GetCreation() time.Time {
	return p.CreationTimestamp.Time
}

func(p pvCell) GetName() string {
	return p.Name
}