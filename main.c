#include <stdio.h>
#include <stdlib.h>

typedef struct file_info {
    char *filename;
    int linenumber;
    int charlocation;
    char *line;
} FileInfo;

int main(int argc, char *argv[]) {
    if (argc < 2) {
        printf("Needs args!");
        return 1;
    }

    FileInfo fileInfo;
    fileInfo.filename = argv[1];
    FILE *fp = fopen(fileInfo.filename, "r");

    fclose(fp);
    return 0;
}
