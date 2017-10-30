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
}
