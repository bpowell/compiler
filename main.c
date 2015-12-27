#include <stdio.h>
#include <stdlib.h>

typedef struct file_info {
    char *filename;
    FILE *fp;
    int linenumber;
    int charlocation;
} FileInfo;

FileInfo fileInfo;

void display_fileinfo(FileInfo fi) {
    printf("\nLine no: %i\tChar no: %i\n", fi.linenumber, fi.charlocation);
}

int main(int argc, char *argv[]) {
    if (argc < 2) {
        printf("Needs args!");
        return 1;
    }

    fileInfo.filename = argv[1];
    fileInfo.linenumber = fileInfo.charlocation = 0;

    fileInfo.fp = fopen(fileInfo.filename, "r");
    if (fileInfo.fp==NULL) {
        printf("Cannot open file!\n");
        return 2;
    }

    char c;
    for(;;) {
        c = fgetc(fileInfo.fp);
        if (feof(fileInfo.fp)) {
            break;
        }

        display_fileinfo(fileInfo);
        printf("%c", c);

        if (c=='\n') {
            fileInfo.charlocation = 0;
            fileInfo.linenumber++;
        }

        fileInfo.charlocation++;
    }

    fclose(fileInfo.fp);
    return 0;
}
