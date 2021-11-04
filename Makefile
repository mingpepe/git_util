ifeq ($(OS),Windows_NT)
    DEL := del
    EXT := exe
else
    DEL := rm
    EXT := elf
endif

RS := repoSummary.$(EXT)
LA := logAnalysis.$(EXT)

all:$(RS) $(LA)
	
$(RS): app/summary/main.go util/util.go repo/repo.go
	go build -o $(RS) app/summary/main.go 

$(LA): app/logAnalysis/main.go util/util.go repo/repo.go
	go build -o $(LA) app/logAnalysis/main.go 

clean:
	$(DEL) *.$(EXT)