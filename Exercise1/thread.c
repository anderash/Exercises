#include <pthread.h>
#include <stdio.h>

int i = 0;

void* threadFunc1(){
	for(int j = 0; j<=1000000; j+=1){
		i += 1;
	}
	return 0;
}

void* threadFunc2(){
	for(int j = 0; j<=1000000; j+=1){
		i -= 1;
	}
	return 0;
}

int main(){
	pthread_t someThread1;
	pthread_t someThread2;
    pthread_create(&someThread1, NULL, threadFunc1, NULL);
    pthread_create(&someThread2, NULL, threadFunc2, NULL);

    pthread_join(someThread1, NULL);

    printf("%d\n", i);

    return 0;

}
