from threading import Thread, Lock
i = 0


def someThreadFunction1():
	for j in range (0, 1000000):
		l.acquire()
		i += 1 
		l.release()
	print("Hello from a thread 1!")

def someThreadFunction2():
	for j in range (0, 1000000):
		l.acquire()
		i -= 1
		l.release()
	print("Hello from a thread 2!")

# Potentially useful thing:
# In Python you "import" a global variable, instead of "export"ing it when you declare it
# (This is probably an effort to make you feel bad about typing the word "global")



def main():
	l = Lock()
	someThread1 = Thread(target = someThreadFunction1, args = (),)
	someThread2 = Thread(target = someThreadFunction2, args = (),)
	someThread1.start()
	someThread2.start()
	someThread1.join()
	someThread2.join()
	print(i, "Hello from main!")

main()
