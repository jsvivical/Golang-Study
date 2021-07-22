#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>
#include <windows.h>

#define WHITE SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE), 15);
#define YELLOW SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE), 14);
#define RED SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE), 12);
#define BLUE SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE), 9);

#define MAX_STACK_SIZE 100


typedef struct {
	int vert;
	int horiz;
}offset; 

typedef struct
{
	int row;
	int col;
	int dir;
}element;

enum direction {N, NE, E, SE, S, SW, W, NW};


offset move[8] =
{
	{-1, 0}, //N
	{-1, 1}, //NE
	{ 0, 1}, //E
	{ 1, 1}, //SE
	{ 1, 0}, //S
	{ 1,-1},//SW
	{ 0,-1}, //W
	{-1,-1}, //NW
};


char* fileName[] = { "maze1.txt", "maze2.txt", "maze3.txt", "maze4.txt" };
FILE* f;

int** maze; //미로 -> 초기값 변경 없음

int** mark; //한번 갔던 경로 표시 0으로 초기화 하고 갔던 길은 1로 변경

element stack[MAX_STACK_SIZE];
int top = -1;
int row; int col;
int ENROW; int ENCOL; int EXROW; int EXCOL;
char c;

void loadMaze();
void printMaze();
bool isEmpty();
element pop(int* top);
bool isFull();
void push(element position);
void path();


void loadMaze()
{
	int menu = 2;
	do
	{
		system("cls");
		if (menu != 1 && menu != 2 && menu != 3 && menu != 4) 
		{
			RED
			printf("\n\n\n\t잘못된 입력입니다!\n\n\n\t1 ~ 4의 정수를 입력하세요.\n\n\t잠시후 이전 화면으로 돌아갑니다.\n");
			clock_t wait = clock();
			while (clock() - wait < 1500){} //경고메시지를 일정시간 보여주고 이전 화면으로 돌아감
			system("cls");
		}
		WHITE
		puts("");
		printf("\t**************\n\n");
		printf("\t1 -> maze1.txt\n");
		printf("\t2 -> maze2.txt\n");
		printf("\t3 -> maze3.txt\n");
		printf("\t4 -> maze4.txt\n\n");
		printf("\t**************\n\n");
		printf("\tchoice : ");
		scanf("%d", &menu); //처음 화면에서 풀 미로를 선택 (1 ~ 4번까지)
	} while (menu < 1 || menu > 4); //데이터 유효성 검사
	

	f = fopen(fileName[menu -1], "r");
	if (!f)
	{
		fprintf(stderr, "ERROR01 : cannot open file\n");
		exit(1);
	}
	//미로에 대한 정보 불러오기(행, 열)
	fscanf(f, "%d %d", &row, &col);
	row += 2; col += 2;

	//maze와 mark를 int형 2차원 배열으로 동적할당

	maze = (int**)malloc(sizeof(int*) * row); 
	mark = (int**)malloc(sizeof(int*) * row);
	for (int i = 0; i < row; i++)
	{
		maze[i] = (int*)malloc(sizeof(int) * col);
		mark[i] = calloc(col, sizeof(int));
	}
	for (int i = 0; i < row ; i++)
	{
		for (int j = 0; j < col; j++)
		{
			if (i == 0 || i == (row - 1) || j == 0 || j == (col -1)) maze[i][j] = 1;
			else fscanf(f, "%d", &maze[i][j]);
		}
	}
	//입구 출구의 위치 저장, 주위를 1로 두르기때문에 1씩 더함.
	fscanf(f, "%d %d %d %d",
		&ENROW, &ENCOL, &EXROW, &EXCOL);
	ENROW += 1; ENCOL += 1; EXROW += 1; EXCOL += 1; 
}


void printMaze()
{
	//초기의 미로를 출력하는 함수
	for (int i = 0; i < row; i++)
	{
		for (int j = 0; j < col; j++)
		{
			if (maze[i][j] == 1)
			{
				BLUE
				printf("%3d",maze[i][j]);
				WHITE;
			}
			else printf("%3d",maze[i][j]);
		}
		puts("");
	}
	puts(""); puts("");
}

void printResult(int *top)
{
	//미로의 결과를 출력하는 함수

	while (*top > -1) //스택에 있는 위치에 있는 요소들을 2로 설정
	{
		maze[stack[*top].row][stack[(*top)].col] = 2;
		(*top)--;
	}
	maze[ENROW][ENCOL] = 3; //입구의 위치를 3으로 저장
	maze[EXROW][EXCOL] = 4; //출구의 위치를 4로 저장


	for (int i = 0; i < row; i++)
	{
		for (int j = 0; j < col; j++)
		{
			if (maze[i][j] == 2)  //스택에 저장된 요소->목적지까지의 경로
			{ 
				YELLOW
				printf("%3c", 'X'); 
				WHITE
			}
			else if (maze[i][j] == 3) //입구를 나타내는 요소에서는  'S'출력
			{ 
				RED
				printf("%3c", 'S');
				WHITE
			}

			else if (maze[i][j] == 4)  //출구를 나타내는 요소에서는 'F'출력
			{ 
				RED
				printf("%3c", 'F');
				WHITE
			}

			else 
			{
				if (maze[i][j] == 1)
				{
					BLUE
						printf("%3d", maze[i][j]);
					WHITE;
				}
				else printf("%3d", maze[i][j]);
			} //그 외에는 원래 미로의 요소대로 출력
		}
		puts("");
	}
		puts("");puts("");
}




bool isEmpty() 
{
	//스택이 비었는지를 검사하는 함수 
	//비었으면 true 아니면 false 리턴
	if (top == -1) return true;
	else return false;
}

element pop(int *top)
{
	//스택의 가장 위에 있는 요소를 pop;
	if (isEmpty()) return stack[*top];
	else 
	{ 
		return stack[(*top)--];
	}
}

bool isFull()
{
	//스택이 가득 찼는지를 검사하는 함수
	//가득 찼으면 true, 아니면 false를 출력
	if (top == MAX_STACK_SIZE) return true;
	else return false;
}

void push(int* top,element position)
{
	//경로를 스택에 push하는 함수
	if (isFull()) return;
	else 
	{
		stack[++(*top)] = position;
	}
}

void path()
{
	system("cls") ;
	printf("Maze with 1 boundaries\n\n");
	printMaze(); // 불러온 미로를 1로 둘러싼 후 출력
	int i, ROW, COL, nextRow, nextCol, dir;

	bool found = false; //길을 찾으면 true

	element position;

	mark[ENROW][ENCOL] = 1; top = 0;//시작위치

	stack[0].row = ENROW ; stack[0].col = ENCOL; stack[0].dir = N; //시작위치를 경로(스택)에 저장 처음에는 N위치부터 이동

	while (top > -1 && !found) //top = -1은 시작위치(시작위치보다 뒤로 갈 수 없음)
	{
		position = pop(&top); //position에 가장 최근의 위치 저장

		ROW = position.row; COL = position.col; dir = position.dir; //지역변수 ROW, COL,dir에 가장 최근 위치정보를 저장

		while (dir < 8 && !found)  //8방향 모두 갈 곳이 없거나 길을 찾으면 반복을 벗어남.
		{
			nextRow = ROW + move[dir].vert; //다음 이동 위치
			nextCol = COL + move[dir].horiz;

			if (nextRow == EXROW && nextCol == EXCOL)  //길을 찾음
			{
				found = true;
			}

			else if (!maze[nextRow][nextCol] && !mark[nextRow][nextCol]) //아직 가지 않은 경로이고 갈 수 있는 길이면 이동
			{
				mark[nextRow][nextCol] = 1; //이동했으므로 mark에 표시
				position.row = ROW; position.col = COL; position.dir = dir++;/*방향 저장*/
				push(&top, position);                /*스택에 삽입*/
				ROW = nextRow; COL = nextCol; dir = N; //최근 위치를 수정하고 탐색 방향을 N으로 초기화
			}
			else ++dir;
		}

	}
	position.row = ROW; position.col = COL; //다음 위치가 출구면 빠져나오기 때문에 stack에 저장하기 위해 다시 수정해줌
	push(&top, position);

	if (found) //길을 찾은 경우 경로를 표시한 미로를 출력
	{
		printf("The path is : \n");
		printResult(&top);
	}
	else //길을 못 찾은 경우 원래의 미로를 출력
	{
		printMaze();
		printf("The maze does not have any path!\n\n\n");
	}
}

int main()
{
	do
	{
		system("cls"); //콘솔 프로그램의 화면을 모두 지움
		loadMaze(); //*.txt 파일의 미로를 힙 메모리에 저장
		path(); //미로탐색하는 함수
		printf("처음으로 돌아가려면 아무 문자 입력(x : 종료) : "); 
		scanf(" %c", &c);
		system("cls");
	} while (c != 'x'); //x를 입력하면 프로그램 종료


	free(maze); //동적 메모리 해제
	free(mark);
	fclose(f); //열었던 파일 닫기
}