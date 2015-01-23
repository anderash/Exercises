#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t lock;

void* threadFunc1(){
	for(int j = 0; j<=1000000; j+=1){
		pthread_mutex_lock(&lock);
		i++;
		pthread_mutex_unlock(&lock);
	}
	return NULL;
}

void* threadFunc2(){
	for(int j = 0; j<=1000000; j+=1){
		pthread_mutex_lock(&lock);
		i--;
		pthread_mutex_unlock(&lock);
	}
	return NULL;
}

int main(){
	pthread_mutex_init(&lock, NULL);
	pthread_t someThread1;
	pthread_t someThread2;
    pthread_create(&someThread1, NULL, threadFunc1, NULL);
    pthread_create(&someThread2, NULL, threadFunc2, NULL);

    pthread_join(someThread1, NULL);
    pthread_join(someThread2, NULL);

    printf("%d\n", i);

    return 0;

}
