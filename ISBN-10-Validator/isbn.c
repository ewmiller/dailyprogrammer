#include <stdio.h>
#include <string.h>
#include <stdlib.h>

//determines if given input is a valid ISBN

int main(int argc, char *argv[]){
	char *input = argv[1];

	if(argv[1] == '\0') {
		printf("Usage: <program name> <10-digit ISBN>\n");
		exit(1);
	}
	if((strlen(input)) != 10){
		printf("Incorrect input.\n");
		exit(1);
	}
	else {
		printf("Correct input. Verifying ISBN...\n");
		//algorithm
		int i;
		int sum = 0;
		for(i = 0; i < 10; i++){
			sum+=((10-i)*(int)input[i]);
		}
		//the check
		if((sum % 11) == 0){
			printf("Valid ISBN!\n");
			exit(0);
		}
		else{
			printf("Invalid ISBN.\n");
			return(0);
		}
	}
}
