package assignment02IBC

import (
"crypto/sha256"
"fmt"
)

const miningReward = 100
const rootUser = "Satoshi"


type BlockData struct {
Title    string
Sender   string
Receiver string
Amount   int
}
type Block struct {
Data        []BlockData
PrevPointer *Block
PrevHash    string
CurrentHash string
}


func CalculateBalance(userName string, chainHead *Block) int {
  var temp *Block
  temp = chainHead
  total_balance := 0
  for temp != nil {
    for i,_ := range temp.Data {
      if temp.Data[i].Sender == userName {
        total_balance -= temp.Data[i].Amount
      }
      if temp.Data[i].Receiver == userName {
        total_balance += temp.Data[i].Amount
      }
    }
    temp = temp.PrevPointer
  }
  return total_balance
}

func CalculateHash(inputBlock *Block) string {
  var block_hash string = fmt.Sprintf("%v", *inputBlock)
  hash := sha256.Sum256([]byte(block_hash))
  return fmt.Sprintf("%x", hash)

}

func VerifyTransaction(transaction *BlockData, chainHead *Block) bool {
  amount := CalculateBalance(transaction.Sender, chainHead)
  if amount >= transaction.Amount {
    return true
  } else {
    return false
  }
}

func InsertBlock(blockData []BlockData, chainHead *Block) *Block {
  if chainHead == nil {
      if blockData[0].Title == "Premined" {
        var temp_block Block
        chainHead = &temp_block
        chainHead.PrevPointer = nil
        chainHead.Data = blockData
        cur_hash := CalculateHash(chainHead)
        chainHead.CurrentHash = cur_hash
        chainHead.PrevHash = ""

      } else {
          var validity bool
          for i,_ := range blockData {
            validity = VerifyTransaction(&blockData[i], chainHead)
            if !validity {
              fmt.Println("------------Invalid Transaction found------------")
              fmt.Println(blockData[i].Sender, "has", CalculateBalance("Alice", chainHead), "coins", blockData[i].Amount, "were needed")
              break
            }
          }
          if validity {
         var temp_block Block
         chainHead = &temp_block
         chainHead.PrevPointer = nil
         transaction := append(blockData, BlockData{Title: "CoinBased", Sender: "System", Receiver: "Satoshi", Amount: miningReward})
         chainHead.Data = transaction
         fmt.Println("Transaction = ", chainHead.Data[0])
         cur_hash := CalculateHash(chainHead)
         chainHead.CurrentHash = cur_hash
         chainHead.PrevHash = ""
       }
     }
    }  else {
      if blockData[0].Title == "Premined" {
        var temp_Block Block
        temp_Block.PrevPointer = chainHead
        temp_Block.Data = blockData
        temp_Block.PrevHash = chainHead.CurrentHash
        cur_hash := CalculateHash(chainHead)
        temp_Block.CurrentHash = cur_hash
        chainHead = &temp_Block

      } else {
        var validity bool
        for i,_ := range blockData {
          validity = VerifyTransaction(&blockData[i], chainHead)
          if !validity {
            fmt.Println("------------Invalid Transaction found------------")
            break
          }
        }
        if validity {
          var temp_Block Block
          temp_Block.PrevPointer = chainHead
          transaction := append(blockData, BlockData{Title: "CoinBased", Sender: "System", Receiver: "Satoshi", Amount: miningReward})
          temp_Block.Data = transaction
          temp_Block.PrevHash = chainHead.CurrentHash
          cur_hash := CalculateHash(chainHead)
          temp_Block.CurrentHash = cur_hash
          chainHead = &temp_Block
      }
    }
  }

return chainHead

}
func ListBlocks(chainHead *Block) {
  fmt.Println("------------Listing Blocks-----------------")
  var temp *Block
  temp = chainHead
  for temp != nil {
    for i,_ := range temp.Data {
      fmt.Println("Title:",temp.Data[i].Title, "Sender:",temp.Data[i].Sender,"Receiver:",temp.Data[i].Receiver,"Amount:",temp.Data[i].Amount)
      i++
    }
    fmt.Println("||")
    temp = temp.PrevPointer
  }

}

func VerifyChain(chainHead *Block) {
  ptr := chainHead.PrevPointer
  for ptr.PrevPointer != nil {
    if ptr.CurrentHash != chainHead.PrevHash {
      fmt.Println("Block Chain not verified")
      break
    }
  }
  fmt.Println("Block Chain Verified")

}

func PremineChain(chainHead *Block, numBlocks int) *Block {
  var block_ptr *Block
  sysTosatoshi := []BlockData{{Title: "Premined", Sender: "nil", Receiver: "nil", Amount: 0}}
  sysTosatoshi = append(sysTosatoshi, BlockData{Title: "CoinBased", Sender: "System", Receiver: "Satoshi", Amount: miningReward})
  for i := 0; i < numBlocks; i++ {
  block_ptr = InsertBlock(sysTosatoshi, chainHead)
  chainHead = block_ptr
  }
  return chainHead
}
