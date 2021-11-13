# [ConSinGAN (Concurrent SinGAN)](https://arxiv.org/pdf/2003.11512.pdf)

## SinGAN 的局限性

1. 图像块差异大时, 无法学到很好的分布, 即生成的图像不真实. 当训练 scale 次数较多时, 又会出现与原图像差异小的问题.

2. 生成图像受限于原图像提供的语义信息, 可能出现"创造性"差的问题.

3. 训练时间较长 (使用 1080ti 单张图片需要 2 小时左右)

## ConSinGAN 优势

1. 并行的实现, 使得训练时间缩短 (使用 1080ti 单张图片 20 分钟左右)

2. 更好的学习图片的结构 (动态的学习率)

## 改进

### 改进 1

Adobe 与汉堡大学的研究人员发现, 在给定的时间内仅能训练一个生成器, 并将图像(而不是特征图)从一个生成器传输到下一个生成器, 这就限制了生成器之间的交互.

使用 end to end 的方法, 在一定时间内, 训练多个 G, 每个 G 的输入为前一个 G 的输出.

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gwdly6c9zfj311f0fxk06.jpg"></div>

Stage 0 根据噪声 z 生成 feature map, Stage 1 将 Stage 0 的输出 作为输入, 经过 upsampling + 噪声 z1 生成新的 feature map, 经过 n 个 Stage, 得到全图的 feature map.

这里有一个问题, 由于训练时都是训练 feature map, 而不是整个图像, 如果同时并行训练所有的 G, 所以可能会造成过拟合, 即训练的很好, 但是泛化性很低.

### 改进 2

1. 在给定的时间内仅训练一部分生成器 G (默认为 3 个, 同时训练更多, 会导致过拟合)

2. 对于不同的 G, 设置不同的学习率 learning rate

<div align=center><img src="https://imgconvert.csdnimg.cn/aHR0cHM6Ly9tbWJpei5xcGljLmNuL21tYml6X3BuZy9ZaWNVaGs1YUFHdENqdGY5dE4yZTFqTzR0b3pTaWIxS1BoUHVHOWxEbGJFcklsRDZUbTBTSWtFaWNtYUI3TkpCWHY4MW1RMkxlcTBmTmJpYmVMazhXTGc0YmcvNjQw?x-oss-process=image/format,png"></div>

同时训练前三层, 学习率分别为 0.1, 0.01， 0.001, 之后不进行并行, 学习率为 0.001.

## 个人小组实验结果

训练时间明显缩短 (20 min), 但是对于染色体的增广效果较差, 不能很好的学习到染色体的结构.

## Loss

与 SinGAN 相同

## Matrix

与 SinGAN 相同 (SIFID 越小越好)

## [ConSinGAN (Concurrent SinGAN)](https://arxiv.org/pdf/2003.11512.pdf)

## [ConSinGAN Github](https://github.com/tohinz/ConSinGAN)

## [SinGAN CSDN](https://blog.csdn.net/QbitAI/article/details/103675677?spm=1001.2101.3001.6650.4&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-4.no_search_link&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-4.no_search_link)

## [ConSinGAN CSDN](https://blog.csdn.net/QbitAI/article/details/105212623)
