class Rational(private val n: Int, private val d: Int) {
  // インスタンスの初期化を行うさいにinitを呼び出す
  init {
    // 標準ライブラリが提供する: エラーをスローする
    require(d != 0, { "denominator must not be null" })
  }

  // TODO: by lazy?
  private val g by lazy { gcd(Math.abs(n), Math.abs(d)) }
  private val numerator: Int by lazy { n / g }
  val denominator: Int by lazy { d / g }

  // operator overload: 演算子オーバーロード
  operator fun plus(that: Rational): Rational = Rational(
      (numerator * that.denominator) + (that.numerator * denominator),
      denominator * that.denominator
  )

  operator fun plus(n: Int): Rational = Rational(
      numerator + (n * denominator),
      denominator
  )

  override fun toString(): String = "$numerator/$denominator"
  // TODO: 再帰のやつらしい。あとで
  tailrec private fun gcd(a: Int, b: Int): Int = if (b == 0) a else gcd(b, a % b)
}

// IntにはRationalを引数にとるメソッドがないので、拡張関数を定義してplusの演算子オーバーロード
operator fun Int.plus(r: Rational): Rational = r + this

fun main(args: Array<String>) {
  val half = Rational(1, 2)
  println(half.denominator)
  println(half)
  val half2 = Rational(2, 5)
  println(half2)

//  println(Rational(1, 0)) // TODO: throw error

  // 最大公約数
  println(Rational(16, 24))
  // 演算子オーバーロード
  println(Rational(16, 24).plus(Rational(1, 3)))
  println(Rational(16, 24) + (Rational(1, 3))) // 演算子オーバーロードで計算可能になる
  println(Rational(16, 24) + 1) // Int引数の演算子オーバーロードで計算可能になる
  println(1 + Rational(16, 24)) // 拡張関数で計算可能になる
}

// varは変更可能
// valは変更不可能
