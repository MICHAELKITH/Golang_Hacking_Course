package main 
 
import ( 
	"fmt" 
	"net" 
	"sort" 
) 

❶ func worker(ports, results chan int) { 
	for p := range ports { 
		address := fmt.Sprintf("scanme.nmap.org:%d", p) 
		conn, err := net.Dial("tcp", address) 
		if err != nil { 
		 ❷ results <- 0 
			continue 
		} 
		conn.Close() 
	 ❸ results <- p 
	} 
} 
func main() { 
	ports := make(chan int, 100) 
 ❹ results := make(chan int) 
 ❺ var openports []int 

	for i := 0; i < cap(ports); i++ { 
		go worker(ports, results) 
	} 

 ❻ go func() { 
	   for i := 1; i <= 1024; i++ { 
		   ports <- i 
		} 
	}() 

 ❼ for i := 0; i < 1024; i++ { 
		port := <-results 
		if port != 0 { 
			openports = append(openports, port) 
		} 
	} 

close(ports) 
close(results) 
❽ sort.Ints(openports) 
for _, port := range openports { 
fmt.Printf("%d open\n", port) 
} 
} 