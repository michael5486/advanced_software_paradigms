object matrix{
  def main(args: Array[String]) {
    println("testing")

      
  }
  
  def createMatrix():Array = {
    val matrix = Array.ofDim[Int](5,5)
    var i = 0
    var j = 0
    var r = scala.util.Random
    for (i <-0 to 5){
        for (j <-0 to 5){
            matrix(i)(j) = r.nextInt(9) + 1
        }
    }
    return matrix
  }
  
  def printMatrix(matrix: Array): Unit =  {
    //   val i, j = 0
    //   for (i<=5){
    //       for (j<-5){
    //           print
    //       }
    //   }
    println(matrix)
  }
}