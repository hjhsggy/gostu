package main

import (
	"fmt"
	"sync"
)

func main1() {
	channel := make(chan string, 5)
	go func() { channel <- doWork("test1") }()
	go func() { channel <- doWork("test2") }()
	go func() { channel <- doWork("test3") }()
	go func() { channel <- doWork("test4") }()
	go func() { channel <- doWork("test5") }()
	println(<-channel)
}

func doWork(str string) string {
	return str + "测试结果"
}

var x sync.Map

// var x map[int]string = make(map[int]string)

func f(s string, wg *sync.WaitGroup) {
	x.Store(0, s)
	wg.Done()
}

func g(s string, wg *sync.WaitGroup) {
	x.Store(1, s)
	wg.Done()
}

func main2() {

	for {
		var wg sync.WaitGroup
		wg.Add(2)
		go f("Hello", &wg)
		go g("Playground", &wg)
		wg.Wait()
	}
}

func kmp(pat string) [][]int {
	lenPat := len(pat)
	dp := make([][]int, 256)

	dp[0][pat[0]] = 1
	x := 0
	for j := 1; j < lenPat; j++ {
		for c := 0; c < lenPat; c++ {
			dp[j][c] = dp[x][c]
			dp[j][pat[j]] = j + 1
			x = dp[x][pat[j]]
		}
	}
	return dp
}

func main() {
	dest := "aadabcad"
	pat := "abca"

	l1 := len(dest)
	l2 := len(pat)

	dp := kmp(dest)
	j := 0
	tmp := -1
	for i := 0; i < l1; i++ {
		j = dp[j][pat[i]]
		if j == l2 {
			tmp = i - l2 + 1
		}
	}

	fmt.Println(tmp)

}

//     public class KMP {
//     private int[][] dp;
//     private String pat;

//     public KMP(String pat) {
//         this.pat = pat;
//         int M = pat.length();
//         // dp[状态][字符] = 下个状态
//         dp = new int[M][256];
//         // base case
//         dp[0][pat.charAt(0)] = 1;
//         // 影子状态 X 初始为 0
//         int X = 0;
//         // 构建状态转移图（稍改的更紧凑了）
//         for (int j = 1; j < M; j++) {
//             for (int c = 0; c < 256; c++) {
//                 dp[j][c] = dp[X][c];
//             dp[j][pat.charAt(j)] = j + 1;
//             // 更新影子状态
//             X = dp[X][pat.charAt(j)];
//         }
//     }

//     public int search(String txt) {
//         int M = pat.length();
//         int N = txt.length();
//         // pat 的初始态为 0
//         int j = 0;
//         for (int i = 0; i < N; i++) {
//             // 计算 pat 的下一个状态
//             j = dp[j][txt.charAt(i)];
//             // 到达终止态，返回结果
//             if (j == M) return i - M + 1;
//         }
//         // 没到达终止态，匹配失败
//         return -1;
//     }
// }
