# [GAN](https://www.cs.cmu.edu/~jeanoh/16-785/papers/goodfellow-nips2014-gans.pdf)

## abstract

对抗的形式, 训练两个模型, G 抓住数据的分布, D 辨别数据是否是被生成的还是来自于源数据.

## Introduction

discriminative models 效果好, 但 generative models 效果不好.

假设有造假者 G, 有警察 D, 通过对抗我们希望两者能相互学习对抗, 最终希望造假者 G 从中获胜, 从而生成出与源数据不尽相同的数据.

如果 D 与 一个以随机噪音(random noise)为输入的 G 都为 MLP (mutilayer perceptron 多层感知机), 则可以只使用通过 BP(back propagation), dropout, 等技术进行计算. (模型简单)

原文:

In this article, we explore the special case when the generative model generates samples
by passing random noise through a multilayer perceptron, and the discriminative model is also a
multilayer perceptron. We refer to this special case as adversarial nets. In this case, we can train both models using only the highly successful backpropagation and dropout algorithms [16] and
sample from the generative model using only forward propagation. No approximate inference or
Markov chains are necessary

## Related work

以往的工作期望学习完整原始数据的 probability distribution function, 学习原数据的方差, 均值等等.

近期有一些工作被称为"generative machines", 期待构造出近似的结果, 不需要具体了解原数据的分布到底是怎样的. 减少计算量.

[VAE](https://arxiv.org/abs/1906.02691): GAN 类似的工作

[NCE](http://proceedings.mlr.press/v9/gutmann10a/gutmann10a.pdf): 比 GAN 更复杂的损失函数

[predictability minimization](https://direct.mit.edu/neco/article-abstract/4/6/863/5678/Learning-Factorial-Codes-by-Predictability?redirectedFrom=fulltext): GAN 的前辈, GAN 是 PM 的变种

[adversarial example](https://arxiv.org/abs/1312.6199v1): 生成与源数据相似的数据, 来测试算法的稳定性

## Adversarial nets

生成器 G 与辨别器 D 均为 MLP, 生成器以一个随机噪声为输入, 输出一个与原始图片相近的数据. 辨别器以一个数据为输入并输出一个标量,辨别数据如果来自于生成器输出 1, 如果来自原始数据输出 0. 同时, 训练 D 时, 也会训练 G, 目的是最小化 log(1 - D( G(z) ) ).

分析:

GAN 的目的是希望造假者 G 骗过 D, 而生成与原始数据不同的数据.

假设 D 是百分之百正确的, 可以完全分辨出数据是否为生成数据.

如果数据 G(z)为原始数据, 则 D(G(z))为 0, log(1-0) = log(1) = 0

如果数据 G(z)为生成数据, 则 D(G(z))为 1, log(1-1) = -Infinity

所以我们期待, log(1 - D( G(z) ) ) 最小化, 这意味着我们成功的生成了与原始数据不同的数据

总结: GAN 的"loss function"

<div align=center><img src="https://tvax1.sinaimg.cn/large/006cK6rNgy1gwa1h7llyxj30i0015gm6.jpg">

我们总是期望 G 有最小值(-Infinity), D 有最大值(1), 都朝向可以生成数据的方向.</div>

### 图解 GAN

<div align=center><img src="https://tva2.sinaimg.cn/large/006cK6rNgy1gwa2ahi6gmj313e0ar0uo.jpg"> </div>

紫色虚线代表 D 的输出. 黑色虚线代表原始数据的分布, a 代表原始数据的平均值. 绿色实线代表生成数据的分布, b 代表生成数据的平均值. z 为 随机噪音, 映射到数据变量 x.

想要生成与原始数据相关的新数据, 只需要让生成数据的分布与原始数据的分布相似, 此例中为向左移动. 循环往复, D 逐渐分辨不出原始数据与生成数据的区别.

### GAN 的算法 (略)

<div align=center><img src="https://tva3.sinaimg.cn/large/006cK6rNgy1gwa2lg9eylj30pc0gg13n.jpg"></div>

有一个问题, 开始训练时, 由于 G 尚不完全, 很容易 D 太厉害, 而导致 log(1 - D( G(z) ) )为 0, 而需要对此式进行求导更新 G, 对 0 求导还是 0, 导致 G 之后的训练很困难.

GAN 的算法证明详情移步 [GAN](https://www.cs.cmu.edu/~jeanoh/16-785/papers/goodfellow-nips2014-gans.pdf) 较难收敛

## 贡献

使用有监督学习的损失函数(被分为原始数据与生成数据), 去训练无监督学习.
