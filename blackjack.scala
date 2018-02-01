object blackjack {

  var picks = List("A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K")
  var colors = List("spade", "hearts", "plum", "diamonds")
  val cardValue: Map[String, Int] = Map("A" → 11, "2" → 2, "3" → 3, "4" → 4, "5" → 5, "6" → 6, "7" → 7, "8" → 8, "9" → 9, "10" → 10, "J" → 10, "Q" → 10, "K" → 10)

  def initCard(): (List[String], Map[String, Int]) = {
    var cards: List[String] = List()
    var cardValues: Map[String, Int] = Map()
    var color: String = ""
    var pick: String = ""
    var count = 0
    for (color: String ← colors) {
      for (pick ← picks) {
        var c = color + "_" + pick
        cards = cards :+ c
        cardValues = cardValues.+(c → cardValue(pick))
      }
    }
    return (cards, cardValues)
  }

  def shuffle(cards: List[String]): List[String] = {
    return  util.Random.shuffle(cards)
  }

  def main(args: Array[String]): Unit = {

    var (cards, cardValues) = initCard()
    cards.foreach { i =>
      println("card=" + i + " value=" + cardValues(i))
    }
    println(cards.length)
    println(cards(1))
    println(cards(2))
    cards = util.Random.shuffle(cards)
    println(cards(1))
    println(cards(2))
  }

}