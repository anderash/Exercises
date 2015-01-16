#include <pthread.h>
#include <studio.h>


void* thread1(int i){
	while(i <= 1000000){
		i += 1;
	}
}

void* thread2(int i){
	while(i >= -1000000){
		i -= 1;
	}
}

int main(){
	i = 0;



}
