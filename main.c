#include <stdio.h>
#include <stdlib.h>

typedef struct file_info {
    char *filename;
    FILE *fp;
    int linenumber;
    int charlocation;
} FileInfo;

FileInfo fileInfo;
int lookahead;

void display_fileinfo() {
    printf("\nLine no: %i\tChar no: %i\n", fileInfo.linenumber, fileInfo.charlocation);
}

int lexan() {
    int c;
    for(;;) {
        c = fgetc(fileInfo.fp);
        if (feof(fileInfo.fp)) {
            return -1;
        }

        fileInfo.charlocation++;

        if (c==' ' || c=='\t') {
            continue;
        }

        if (c=='\n') {
            fileInfo.charlocation = 1;
            fileInfo.linenumber++;
            printf("\n");
            continue;
        }

        return c;
    }

    return -1;
}

void error(const char *msg) {
    printf("%s\n", msg);
    display_fileinfo();
    exit(1);
}

void emit(int t) {
    printf("%c ", t);
}

void match(int t) {
    if (lookahead==t) {
        lookahead = lexan();
    } else {
        //printf("lookahead = %c\tt = %c\n", lookahead, t);
        error("syntax error");
    }
}

void term() {
    for(;;) {
        switch(lookahead) {
            case '0':
            case '1':
            case '2':
            case '3':
            case '4':
            case '5':
            case '6':
            case '7':
            case '8':
            case '9':
                emit(lookahead);
                match(lookahead);
                return;
            default:
                error("syntax error2:");
        }
    }
}

void expr() {
    int t;
    term();
    for(;;) {
        switch(lookahead) {
            case '+':
            case '-':
                t = lookahead;
                match(lookahead);
                term();
                emit(t);
                continue;
            default:
                return;
        }
    }
}

void parse() {
    lookahead = lexan();
    for(;;) {
        if (lookahead==-1) {
            break;
        }

        expr();
        match(';');
    }
}

int main(int argc, char *argv[]) {
    if (argc < 2) {
        printf("Needs args!");
        return 1;
    }

    fileInfo.filename = argv[1];
    fileInfo.linenumber = fileInfo.charlocation = 1;

    fileInfo.fp = fopen(fileInfo.filename, "r");
    if (fileInfo.fp==NULL) {
        printf("Cannot open file!\n");
        return 2;
    }

    parse();

    fclose(fileInfo.fp);
    return 0;
}
