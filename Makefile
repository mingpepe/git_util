EXE := main.exe

all:$(EXE)
	
$(EXE): app\summary\main.go
	go build app\summary\main.go
clean:
	del $(EXE)