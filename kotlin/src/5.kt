// デフォルト引数の指定が可能
fun hello(name: String = "Kiyohi"): String = "Hello, $name"

// varargで可変長引数をうけることができる
// 1つの関数で1つまで
//fun sum(vararg ints: Int): Int = ints.sum()

// 関数内に関数を定義することで外部からaccumulatorを0以外に指定されることを防ぐ
fun sum(numbers: List<Long>): Long {
  // TCO（Tail Call Optimization, 末尾呼び出し最適化）でスタックオーバーフローを回避できる
  tailrec fun go(numbers: List<Long>, accumulator: Long = 0): Long =
      if (numbers.isEmpty()) accumulator
      else go(numbers.drop(1), accumulator + numbers.first())
  return go(numbers, 0)
}

fun main(args: Array<String>) {
  println(hello())
  println(sum((1L..4).toList()))
  println(sum((1L..12345).toList()))
}

