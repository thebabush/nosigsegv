package nosigsegv

// #include <sys/mman.h>
//
// int __nosigsegv() {
//     // **only** read/write, otherwise it would be dangerous
//     void *addr = mmap(NULL, 0x1000, PROT_READ | PROT_WRITE, MAP_ANONYMOUS | MAP_PRIVATE | MAP_FIXED, -1, 0);
//     if (addr == (void *)-1) {
//	       return 1;
//	   }
//     if (addr == NULL) {
//         return 0;
//     }
//     return 2;
// }
import "C"

var errorCode int = -1

func InitError() bool {
	return errorCode != 0
}

func ErrorCode() int {
	return errorCode
}

func init() {
	errorCode = int(C.__nosigsegv())
}

// Force init
func Init() {
	if errorCode == -1 {
		errorCode = int(C.__nosigsegv())
	}
}
