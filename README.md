# algorithms

日常算法。

## 排序

排序种类：

- [冒泡排序](#冒泡排序)

### 冒泡排序

冒泡排序每次遍历结束，会得到无序区的最大值或者是最小值。

1. 比较相邻两个数据如果。第一个比第二个大，就交换两个数
2. 对每一个相邻的数做同样1的工作，这样从开始一队到结尾一队在最后的数就是最大的数。
3. 针对所有元素上面的操作，除了最后一个。
4. 重复1~3步骤，知道顺序完成。

### 选择排序

选择排序（Select Sort） 是直观的排序，通过确定一个 Key 最大或最小值，再从带排序的的数中找出最大或最小的交换到对应位置。再选择次之。双重循环时间复杂度为 O(n^2)

1. 在一个长度为 N 的无序数组中，第一次遍历 n-1 个数找到最小的和第一个数交换。
2. 第二次从下一个数开始遍历 n-2 个数，找到最小的数和第二个数交换。
3. 重复以上操作直到第 n-1 次遍历最小的数和第 n-1 个数交换，排序完成。
