fun main(args: Array<String>) {
  val title = "yukiyunahayusha"
  println(title.length)
  println(title.capitalize())
  println(title.isBlank())
  println(title.isNotBlank())
  println(title.isEmpty())
  println(title.isNotEmpty())

  val titleEmpty = ""
  println(titleEmpty.length)
  println(titleEmpty.capitalize())
  println(titleEmpty.isBlank())
  println(titleEmpty.isNotBlank())
  println(titleEmpty.isEmpty())
  println(titleEmpty.isNotEmpty())

  val titleBlank = "     "
  println(titleBlank.length)
  println(titleBlank.capitalize())
  println(titleBlank.isBlank())
  println(titleBlank.isNotBlank())
  println(titleBlank.isEmpty())
  println(titleBlank.isNotEmpty())

  val name = "琥珀"
  println("$name さん")
  val description = """
<HTML></HTML>
"""

  println(description)

  arraySample()

  listSample()

  setSample()

  mapSample()

  rangeSample()

  ifSample()

  whenSample()

  loopSample()
}

fun arraySample() {
  val ints = arrayOfNulls<Int>(5)
  println(ints.size)
  println(ints[0])
  ints[0] = 99
  println(ints[0])

  val strs = arrayOf("red", "green", "blue")
  println(strs[0])

  val ints2 = intArrayOf(1, 2, 3)
  println(ints2[0])
}

fun listSample() {
  // イミュータブル（変更不可能）
  val ints = listOf(1, 2, 3)
  println("$ints ${ints.size} ${ints[0]}")

  // ミュータブル（変更可能）
  val ints2 = mutableListOf(1, 2, 3)
  println("$ints2 ${ints2.size} ${ints2[0]}")
  ints2[1] = 99
  ints2.removeAt(2)
  println("$ints2 ${ints2.size} ${ints2[0]}")
}

fun setSample() {
  // イミュータブル（変更不可能）
  val ints = setOf(1, 2, 1, 3)
  println(ints)

  val chars: MutableSet<Char> = mutableSetOf('a', 'k', 'i', 'h', 'a')
  println(chars)
  chars -= 'a'
  println(chars)
}

fun mapSample() {
  val numberMap: MutableMap<String, Int> = mutableMapOf(
      "one" to 1, "two" to 2
  )
  println(numberMap)
  println(numberMap.size)
  println(numberMap["one"])
  println(numberMap["two"])
  numberMap += "three" to 3
  numberMap["for"] = 4
  println(numberMap)
}

fun rangeSample() {
  println(5 in 1..10)
  val range: IntRange = 12..15
  println(5 in range)
  println(5 !in range)
  println(range.toList())
  println(range.reversed().toList())
  println((5 downTo 1).toList())
  println((1..5 step 2).toList())
}

fun ifSample() {
  // 普通にifで使えるし、ブロック内で最後に評価される式が返ってくる
  var score = 78
  val message = if (score >= 60) "合格!" else "不合格"
  println(message)
}

fun whenSample() {
  // switchを強力にした感じ
  val x = 1
  val message = when (x) {
    1 -> "one"
    2, 3 -> "two or three"
    else -> {
      "unknown" // 値を返すwhen式ではelseが必須
    }
  }
  println(message)

  val y = ""
  val blank = when (y) {
    is String -> y.isBlank()
    else -> true
  }
}

fun loopSample() {
  var count = 3
  while (count-- > 0) {
    println("Hello")
  }

  for (x in arrayOf(1, 2, 3)) {
    println(x)
  }

  val names = listOf("foo", "bar", "baz")
  for (name in names) {
    println(name)
  }

  for (i in 1..10) {
    println(i)
  }
}


