interface Bucket {
  fun fill()
  fun drainAway()
  fun pourTo(that: Bucket)

  fun getCapacity(): Int
  fun getQuantity(): Int
  fun setQuantity(quantity: Int)
}

fun createBucket(capacity: Int): Bucket = object : Bucket {
  //    val capacity: Int = 5
  var _quantity: Int = 0

  override fun fill() {
    setQuantity(getCapacity())
  }

  override fun drainAway() {
    setQuantity(0)
  }

  override fun pourTo(that: Bucket) {
    val thatVacuity = that.getCapacity() - that.getQuantity()
    if (getQuantity() <= thatVacuity) {
      that.setQuantity(that.getQuantity() + getQuantity())
      drainAway()
    } else {
      that.fill()
      setQuantity(getQuantity() - thatVacuity)
    }
  }

  override fun getCapacity(): Int = capacity

  override fun getQuantity(): Int = _quantity

  override fun setQuantity(quantity: Int) {
    _quantity = quantity
  }
}

fun main(args: Array<String>) {
  val bucket = createBucket(10)
  println(bucket.getQuantity())
  bucket.fill()
  println(bucket.getQuantity())
  bucket.drainAway()
  println(bucket.getQuantity())
}
