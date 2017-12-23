interface Bucket {
  fun fill()
  fun drainAway()
  fun pourTo(that: Bucket)

  val capacity: Int
  var quantity: Int
}

class BucketImpl(_capacity: Int) : Bucket {
  override val capacity = _capacity
  override var quantity = 0

  override fun fill() {
    quantity = capacity
  }

  override fun drainAway() {
    quantity = 0
  }

  override fun pourTo(that: Bucket) {
    val thatVacuity = that.capacity - that.quantity
    if (quantity <= thatVacuity) {
      that.quantity = that.quantity + quantity
      drainAway()
    } else {
      that.fill()
      quantity -= thatVacuity
    }
  }
}

fun main(args: Array<String>) {
  val bucket = BucketImpl(10)
  println(bucket.quantity)
  bucket.fill()
  println(bucket.quantity)
  bucket.drainAway()
  println(bucket.quantity)
}
