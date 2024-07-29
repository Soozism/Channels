# Golang Channels and Goroutines Example

This project demonstrates the use of channels and goroutines in Go. It includes a simple program that checks the availability of a list of websites and prints their status to the console.

## Overview

The program performs the following steps:
1. Initializes a list of website URLs.
2. Creates a channel to communicate between goroutines.
3. Launches a goroutine for each URL to check its status.
4. Continuously receives messages from the channel and relaunches the status check for each URL after a delay.

## Usage

To run the program, ensure you have Go installed on your system. Then, follow these steps:

1. Save the code to a file, for example, `main.go`.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to execute the program:

```bash
go run main.go
```

## Code Explanation

The main components of the code are:

1. **Initialization of Links:**
   ```go
   links := []string{
       "http://go.dev/",
       "http://www.google.com/",
       "http://fedoraproject.org/",
       "http://fedoraproject.org/",
       "http://matplotlib.org/",
       "http://numpy.org/",
   }
   ```

   This slice contains the list of URLs to be checked.

2. **Channel Creation:**
   ```go
   c := make(chan string)
   ```

   A channel is created to communicate between the main function and the goroutines.

3. **Launching Goroutines:**
   ```go
   for _, link := range links {
       go checkLink(link, c)
   }
   ```

   A goroutine is launched for each link to check its status.

4. **Receiving Messages and Relaunching Goroutines:**
   ```go
   for l := range c {
       go func(link string) {
           time.Sleep(time.Second * 5)
           checkLink(link, c)
       }(l)
   }
   ```

   The program continuously receives messages from the channel and relaunches the status check for each link after a 5-second delay.

5. **Checking Link Status:**
   ```go
   func checkLink(link string, c chan string) {
       _, err := http.Get(link)
       if err != nil {
           fmt.Println(link, "might be down: ", err)
           c <- link
           return
       }

       fmt.Println(link, "is OK")
       c <- link
   }
   ```

   This function performs an HTTP GET request to check if the link is accessible. If there is an error, it prints a message and sends the link back to the channel. If the link is accessible, it prints a success message and also sends the link back to the channel.

## Notes

- The program runs indefinitely, continuously checking the status of the websites every 5 seconds.
- To stop the program, you can interrupt it using `Ctrl+C` in the terminal.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
