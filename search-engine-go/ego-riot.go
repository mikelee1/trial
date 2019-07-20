package main

import (
	"fmt"
	"log"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"os"
)

var (
	// searcher is coroutine safe
	searcher = riot.Engine{}

	text  = "李原野"
	text1 = `公司`
	path  = "./search-engine-go/riot-index"

	opts = types.EngineOpts{
		Using: 1,
		IndexerOpts: &types.IndexerOpts{
			IndexType: types.DocIdsIndex,
		},
		UseStore: true,
		StoreFolder: path,
		StoreEngine: "bg", // bg: badger, lbd: leveldb, bolt: bolt
		// GseDict: "../../data/dict/dictionary.txt",
		GseDict:       "./search-engine-go/data/test_dict.txt",
		StopTokenFile: "./search-engine-go/data/stop_tokens.txt",
		StoreShards: 6,
	}
)

func initEngine() {
	// var path = "./riot-index"
	searcher.Init(opts)
	defer searcher.Close()
	os.MkdirAll(path, 0777)

	// Add the document to the index, docId starts at 1
	searcher.Index(1, types.DocData{Content: text})
	searcher.Index(2, types.DocData{Content: text1})


	//searcher.RemoveDoc(5)

	// Wait for the index to refresh
	searcher.Flush()

	log.Println("Created index number: ", searcher.NumDocsIndexed())
}

func restoreIndex() {
	// var path = "./riot-index"
	searcher.Init(opts)
	defer searcher.Close()
	// os.MkdirAll(path, 0777)

	// Wait for the index to refresh
	searcher.Flush()

	log.Println("recover index number: ", searcher.NumDocsIndexed())
}

func main() {
	initEngine()
	defer restoreIndex()

	sea := searcher.Search(types.SearchReq{
		Text: "原野，阿里巴巴",
		RankOpts: &types.RankOpts{
			OutputOffset: 0,
			MaxOutputs:   100,
		}})

	fmt.Println("search response: ", sea, "; docs = ", sea.Docs)

	// os.RemoveAll("riot-index")
}
