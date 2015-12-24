#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    if (argc < 2) {
        printf("Needs args!");
        return 1;
    }

    printf("%s\n", argv[1]);
    char *filename = argv[1];
    FILE *fp = fopen(filename, "r");

    fclose(fp);
    return 0;
}
