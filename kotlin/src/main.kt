class Rational(n: Int, d: Int) {
    init {
        require(d != 0, { "denominator must not be null" })
    }

    private val g by lazy { gcd(Math.abs(n), Math.abs(d)) }
    val numberator: Int by lazy { n / g }
    val denominator: Int by lazy { d / g }
    operator fun plus(that: Rational): Rational =
        Rational(
            numberator * that.denominator + that.numberator * that.denominator,
            denominator * that.denominator
        )

    operator fun plus(n: Int): Rational =
        Rational(numberator + n * denominator, denominator)

    override fun toString(): String = "${numberator}/${denominator}"

    tailrec private fun gcd(a: Int, b: Int): Int =
        if (b == 0) a
        else gcd(b, a % b)
}

operator fun Int.plus(r: Rational): Rational = r + this

fun main(args: Array<String>) {
    println("Hello world " + (Rational(1, 10) + Rational(1, 10)))
}

