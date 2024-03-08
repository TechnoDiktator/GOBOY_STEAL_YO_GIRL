package main
import (

	"fmt"
	"math/rand"
	"sync"
	"time"
)


const (
	matrixSize = 250
)

var (
	matrixA   = [matrixSize][matrixSize]int{}
	matrixB   = [matrixSize][matrixSize]int{}
	result    = [matrixSize][matrixSize]int{}
	lock      = sync.Mutex{} // Regular mutex for protecting result
	waitGroup = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5 // Fix: Assignment was missing, added "="
		}
	}
}

/*
In a function with multiple defer statements, they are executed in reverse order, 
meaning the defer statement that occurs last in the code will actually run first, and so on.
*/

func workOutRow(row int) {
	defer waitGroup.Done() // Move waitGroup.Done() here to ensure it's always called
	lock.Lock()            // Acquire lock before modifying result
	defer lock.Unlock()    // Release lock after modification

	for col := 0; col < matrixSize; col++ {
		for i := 0; i < matrixSize; i++ {
			result[row][col] += matrixA[row][i] * matrixB[i][col]
		}
	}
}

func main() {
	fmt.Println("Working...")
	waitGroup.Add(matrixSize)
	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
	}

	start := time.Now()

	for i := 0; i < 100; i++ { 
		
		waitGroup.Wait() // Wait for all rows to finish computation before regenerating matrices
		lock.Lock()     
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		lock.Unlock() // Release lock after modifying matrices

		//THIS IS THE FUCKING DEADLOCKER!!!!!
		waitGroup.Add(matrixSize)
		
	}

	elapsed := time.Since(start)
	fmt.Println("Done")
	// fmt.Println(result) // Printing the result might be too large, comment out if unnecessary
	fmt.Printf("Processing took %s\n", elapsed)
}


/*
const (
	matrix_size = 250
)

var (
	matrixA = [matrix_size][matrix_size]int{}
	matrixB = [matrix_size][matrix_size]int{}
	result = [matrix_size][matrix_size]int{}
	rwLock = sync.RWMutex{}
	cond = sync.NewCond(rwLock.RLocker())
	waitGroup = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrix_size][matrix_size]int) {
	for row:= 0 ; row < matrix_size ; row++ {
		for col:= 0 ; col<matrix_size ; col++{
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int){
	for col := 0 ; col < matrix_size ; col++ {
		rwLock.Lock()
		for i := 0 ; i < matrix_size ; i++ {
			result[row][col] += matrixA[row][i]*matrixB[i][col]
		}
		rwLock.Unlock()
	}
	cond.Wait()
	waitGroup.Done()
}

func main() {
	fmt.Println("Working ....")
	waitGroup.Add(matrix_size)

	for row := 0 ; row < matrix_size ; row++ {
		go workOutRow(row)
	}
	start := time.Now()
	for i := 0 ; i < 100 ; i++ {
		waitGroup.Wait()
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		waitGroup.Add(matrix_size)
		cond.Broadcast()
		rwLock.Unlock()
	}
	elapsed_time := time.Since(start)
	fmt.Println("DONE")
	fmt.Println("processing time :" , elapsed_time )
}
*/






