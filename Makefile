EXE := main.exe

all:$(EXE)
	
$(EXE):
	go build main.go
clean:
	del $(EXE)