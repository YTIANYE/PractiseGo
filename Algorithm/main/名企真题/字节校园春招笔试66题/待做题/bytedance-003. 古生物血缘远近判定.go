/*
DNA 是由 ACGT 四种核苷酸组成，例如 AAAGTCTGAC，假定自然环境下 DNA 发生异变的情况有：

基因缺失一个核苷酸
基因新增一个核苷酸
基因替换一个核苷酸
且发生概率相同。
古生物学家 Sam 得到了若干条相似 DNA 序列，Sam 认为一个 DNA 序列向另外一个 DNA 序列转变所需的最小异变情况数可以代表其物种血缘相近程度，异变情况数越少，血缘越相近，请帮助 Sam 实现获取两条 DNA 序列的最小异变情况数的算法。

格式：


输入：
- 每个样例只有一行，两个 DNA 序列字符串以英文逗号“,”分割
输出：
- 输出转变所需的最少情况数，类型是数字
示例：


输入：ACT,AGCT
输出：1
提示：

每个 DNA 序列不超过 100 个字符

例子：
[186,419,83,408]
6249

[411,412,413,414,415,416,417,418,419,420,421,422]
9864

*/

package main
