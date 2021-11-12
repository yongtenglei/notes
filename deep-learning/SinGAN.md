# SinGAN

## 创新性

1. 设计一种基于单张自然图像训练的非条件 GAN 网络（由噪声作为输入直接生成）；

2. 无需修改模型，直接应用于多种图像任务；

3. GAN 网络中的 G 网络与 D 网络具有相同的模型结构以及相同的感受野，并添加一种重建损失，保证 GAN 可以进行平稳训练；

4. 设计一种 coarse-to-fine 的金字塔型 GAN 网络，每一层学习到前一层缺失的细节；

## 架构

<div align=center><img src="https://pic3.zhimg.com/80/v2-607a47e9f3f2dfb78c519298814433d6_720w.jpg"> </div>

采用金字塔形状, 基于 Coarse-to-Fine 思想. **由下到上**,尺度逐渐由 **粗糙到精细**.

在 scale n 时, $G_n$ 的输出是第 n+1 个 G 生成出来的图像 $\tilde{x}_{n+1} 与$噪声$z_n$之和

D 辨别 G 生成的图像真假, 每一层都是基于 Patch 判断, 从最粗糙层$G_N$一直上升到最精细层$G_0$, 每个 D 的感受野固定,都是 11\*11,也就是说，在最粗糙的 $G_N$，patch 大小为图像的 1/2, 此时 GAN 网络可以学习到图像的全局结构，而在最精细的 $G_0$，GAN 网络学习的是局部细节.

## 训练

<div align=center><img src="https://pic3.zhimg.com/80/v2-e7acf312bd30b92fb4dcfcbfaf332d42_720w.jpg"></div>

随着 scale 的增加, 学习的特征更加抽象, 需要通道数更多, 网络更深.

感受野:

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gwcuaq2zfwj30ms0c0n0p.jpg"></div>

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gwcub109eij30mj0bngp9.jpg"></div>

## Loss

<div align=center><img src="https://pic4.zhimg.com/80/v2-669bca7cdd495ad084cda3dd1ed690e3_720w.jpg"></div>

[对抗损失](https://zhuanlan.zhihu.com/p/52799555): WGAN-gradient penalty 使我们不用太关注网络结构的设计，无论采用什么样的结构都能训练得比较好.

[重建函数(reconstruction function)](]): 保证最粗糙一层可以由噪声直接生成图像，保证了相近的图像风格.

## [结果(github)](https://github.com/tamarott/SinGAN)

## [原论文](https://arxiv.org/pdf/1905.01164.pdf)

## [知乎](https://zhuanlan.zhihu.com/p/92218525)
