object matrix{
  
  def main(args: Array[String]) {
    //creates random 5x5 matrix of ints
    val matrix = createMatrix()
    //prints matrix
    printMatrix(matrix)
    //transposes matrix
    val transposed = transposeMatrix(matrix)
    //prints transposed matrix
    printMatrix(transposed)
    //converts matrix to a list
    val list = convertMatrix(transposed)
    //prints list
    print("\n" + list.toString)
    //sorts list
    val sorted = list.sorted
    //prints sorted list
    print("\n" + sorted.toString)
    //prints sorted list in english words
    printWords(sorted)
  }
  
  def createMatrix(): Array[Array[Int]] = {
    var matrix = Array.ofDim[Int](5,5)
    var i = 0
    var j = 0
    var r = scala.util.Random
    for (i <-0 to 4){
        for (j <-0 to 4){
            matrix(i)(j) = r.nextInt(8) + 1
        }
    }
    return matrix
  }
  
  def transposeMatrix(matrix: Array[Array[Int]]): Array[Array[Int]] = {
    var transposed = Array.ofDim[Int](5,5)
    var i = 0
    var j = 0
    for (i <-0 to 4){
        for (j <-0 to 4){
            transposed(i)(j) = matrix(j)(i)
        }
    }
    return transposed
  }
  
  def convertMatrix(matrix: Array[Array[Int]]): List[Int] = {
    var list: List[Int] = List()
    for (i <- 4 to 0 by -1){
      list = list.:::(matrix(i).toList)
    }
    return list
  }
  
  def printMatrix(matrix: Array[Array[Int]]): Unit =  {
    var i = 0
    var j = 0
    for (i <-0 to 4){
      print("\n[")
        for (j <-0 to 4){
          print(" ")
          print(matrix(i)(j))
        }
      print(" ]")
    }
    print("\n")
  }
  
  def printWords(list: List[Int]): Unit = {
    val strings = List("one", "two", "three", "four", "five", "six", "seven", "eight", "nine")
    for (i<-0 to 24) {
      println(strings(list(i) - 1))
    }
  }
}