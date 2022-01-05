# 数组中位数, 众数, 平均值

```rust
use std::collections::HashMap;

fn main() {
    let mut nums = [3, 4, 1, 1, 1];
    let s = bubble_sort(nums);

    println!("mean: {}", mean(&mut nums));
    println!("median: {}", median(&mut nums));
    println!("mode: {}", mode(&mut nums));
    print!("{:?}", s)
}

fn mean(nums: &mut [i32; 5]) -> f64 {
    let mut counter = 0;
    let mut sum = 0;

    for i in nums.iter() {
        counter += 1;
        sum += i;
    }

    let counter = counter as f64;
    let sum = sum as f64;
    let avg = sum / counter;

    return avg;
}

fn median(nums: &mut [i32; 5]) -> f64 {
    let s = bubble_sort(*nums);
    let m = nums.len() / 2;

    if s.len() % 2 == 0 {
        let a = s[m] as f64;
        let b = s[m + 1] as f64;
        return a / b;
    } else {
        return s[m] as f64;
    }
}

fn mode(mut nums: &[i32; 5]) -> i32 {
    let mut nb = HashMap::new();
    for i in nums.iter() {
        let counter = nb.entry(i).or_insert(0);
        *counter += 1;
    }

    let mut max = 0;
    let mut maxk = nums[0];
    for (k, v) in nb {
        if v > max {
            max = v;
            maxk = *k;
        }
    }

    return maxk;
}

fn bubble_sort(mut nums: [i32; 5]) -> [i32; 5] {
    for _i in 1..nums.len() {
        for i in 1..nums.len() {
            if nums[i - 1] > nums[i] {
                nums.swap(i - 1, i);
            }
        }
    }
    return nums;
}

```
