EXE := main.exe

all:$(EXE)
	
$(EXE): main.go
	go build main.go
clean:
	del $(EXE)