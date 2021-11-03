ifeq ($(OS),Windows_NT)
    DEL := del
else
    DEL := rm
endif

RS := repoSummary.exe
LA := logAnalysis.exe

all:$(RS) $(LA)
	
$(RS): app/summary/main.go util/util.go repo/repo.go
	go build -o $(RS) app/summary/main.go 

$(LA): app/logAnalysis/main.go util/util.go repo/repo.go
	go build -o $(LA) app/logAnalysis/main.go 

clean:
	$(DEL) *.exe