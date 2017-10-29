class Rational(private val numerator: Int, val denominator: Int) {
  override fun toString(): String = "$numerator/$denominator"
}

fun main(args: Array<String>) {
  val half = Rational(1, 2)
  println(half.denominator)
  println(half)
  val half2 = Rational(2, 5)
  println(half2)
}

// varは変更可能
// valは変更不可能
