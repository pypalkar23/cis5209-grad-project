package memory_mgmt

import (
	"fmt"
	"os"
	"strings"
)

var strSlice string //variable of type string on heap

// Memory management in strings
func memoryLeakingStringSliceAndCopy(paramStr string) {
	strSlice = paramStr[:100]
	/*
		Since strSlice is created by slicing the paramStr, both the variables
		share the same underlying memory block. Even though paramStr's lifespan
		finishes after the function has returned, that memory can't be garbage
		collected as some chunk of it's memory(100 bytes) is in use.
	*/
}

func memoryEfficientStringSliceAndCopy1(paramStr string) {
	tempByteArray := []byte(paramStr[:100])
	strSlice = string(tempByteArray)
	/*
		A two step process where we first convert the string to a byte array whic
		then convert that intermediate byte array to a string again. This way the
		resulting slice gets created at new memory block.
	*/
}

func memoryEfficientStringSliceAndCopy2(paramStr string) {
	strSlice = (" " + paramStr[:100][:1])
	/*
		The use of empty string makes the compiler creates a new memory block for
		resulting string at the cost of 1 extra byte.
	*/
}

func memoryEfficientStringSliceAndCopy3(paramStr string) {
	var temp strings.Builder
	temp.Grow(100)
	temp.WriteString(paramStr[:100])
	strSlice = temp.String()

	/*
		A slightly verbose way of making sure that sliced string gets created at a new memory
		location and the source string can be collected by the garbage collector.
	*/
}

// Memory management in pointers
func memoryLeakingSlicePointer() []*string {
	new_slice := []*string{new(string), new(string), new(string), new(string), new(string), new(string)}
	//.....
	return new_slice[1:4:4]
	/*As long as this returned slice is been in use in the other parts of the program
	the underling slice can't be garbage collected. Subsequently preventing first and
	last elements of the slice from  being garbage collected*/
}
func memoryEfficientSlicePointer() []*string {
	new_slice := []*string{new(string), new(string), new(string), new(string), new(string), new(string)}
	//....
	new_slice[0], new_slice[len(new_slice)-1] = nil, nil
	/*Marking the unused elements of slice as nil, will make them available for garbage collection */
	return new_slice[1:4:4]
}

type FileDetails struct {
	Path    string
	Content string
}

// Memory management in defer calls
func memoryLeakingMultipleFileWrites(fileList []FileDetails) error {
	for _, entry := range fileList {
		file, err := os.Open(entry.Path)
		if err != nil {
			return err
		}

		defer file.Close()
		/* Such approach may result in large number of deferred calls in case
		of thousands of file writes. Subsequently delaying the release of system
		resources.
		*/
		_, err = file.WriteString(entry.Content)
		if err != nil {
			return err
		}

		err = file.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

func memoryEfficientMultipleFileWrites(fileList []FileDetails) error {
	for _, entry := range fileList {
		if err := func() error {
			file, err := os.Open(entry.Path)
			if err != nil {
				return err
			}

			/*
				Use of anonymous function makes sure that the every deferred call
				for file close gets called at the end of every iteration instead of
				processing every call at the end of the outer function.
			*/
			defer file.Close()

			_, err = file.WriteString(entry.Content)
			if err != nil {
				return err
			}

			return file.Sync()
		}(); err != nil {
			return err
		}

	}

	return nil
}

func main() {
	fmt.Println("Functions depicting kind of memory leaking scenarios")
}
