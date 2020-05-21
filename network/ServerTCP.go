package network

import (
    "net"
    "bytes"
    "strings"
    "strconv"
    "encoding/hex"
    "encoding/json"
    "../utils"
    "../models"
    "../settings"
    "../blockchain"
)

func ServerTCP() {
    listen, err := net.Listen("tcp", settings.User.Addr.Port)
    utils.CheckError(err)
    defer listen.Close()

    for {
        conn, err := listen.Accept()
        if err != nil { break }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    var (
        buffer = make([]byte, 512)
        message string
    )

    for {
        length, err := conn.Read(buffer)
        if err != nil || length == 0 { break }
        message += string(buffer[:length])
    }

    var pack models.PackageTCP
    json.Unmarshal([]byte(message), &pack)

    switch pack.Head.Title {
        case settings.TITLE_CONNECT:
            switch pack.Head.Mode {
                case settings.MODE_SAVE: connectSave(&pack)
            }
        case settings.TITLE_BLOCK:
            switch pack.Head.Mode {
                case settings.MODE_READ: blockRead(&pack)
                case settings.MODE_SAVE: blockSave(&pack)
                case settings.MODE_PUSH: blockPush(&pack)
                case settings.MODE_HARD: blockHard(&pack)
            }
    }
}

func blockRead(pack *models.PackageTCP) {
    var (
        my_blocks []models.Block
        another_blocks []models.Block
    )

    err := json.Unmarshal([]byte(getAllBlockChain(pack.Body.Branch)), &my_blocks)
    utils.CheckError(err)

    err = json.Unmarshal([]byte(pack.Body.Data), &another_blocks)
    utils.CheckError(err)

    var (
        my_length = len(my_blocks)
        another_length = len(another_blocks)
    )

    if my_length > another_length {
        connectSave(pack)
        return
    }

    if !blockchain.ValidateChain(another_blocks) {
        return
    }
    
    var index = 0
    for range my_blocks {
        if !bytes.Equal(my_blocks[index].CurrHash, another_blocks[index].CurrHash) {
            break
        }
        index++
    }

    if my_length == index && my_length < another_length {
        addBlocks(pack.Body.Branch, another_blocks[index:])
    } else if my_length < another_length {
        initBlocks(pack.Body.Branch, another_blocks)
    }
}

func addBlocks(branch string, blocks []models.Block) {
    for _, block := range blocks {
        block_json, err := json.MarshalIndent(block, "", "\t")
        utils.CheckError(err)
        _, err = settings.BlockChain.Exec(
            "INSERT INTO BlockChain (Branch, Hash, Block) VALUES ($1, $2, $3)",
            branch,
            hex.EncodeToString(block.CurrHash),
            string(block_json),
        )
    }
    if settings.MainBranch == branch {
        settings.LastHash = blocks[len(blocks)-1].CurrHash
    }
}

func initBlocks(branch string, blocks []models.Block) {
    _, err := settings.BlockChain.Exec(
        "DELETE FROM BlockChain WHERE Branch=$1",
        branch,
    )
    utils.CheckError(err)
    addBlocks(branch, blocks)
}      

func blockSave(pack *models.PackageTCP) {
    var block models.Block
    err := json.Unmarshal([]byte(pack.Body.Data), &block)
    utils.CheckError(err)
  
    if !blockchain.ValidateBlock(pack.Body.Branch, block) {
        return
    }

    _, err = settings.BlockChain.Exec(
        "INSERT INTO BlockChain (Branch, Hash, Block) VALUES ($1, $2, $3)",
        pack.Body.Branch,
        hex.EncodeToString(block.CurrHash),
        pack.Body.Data,
    )

    if pack.Body.Branch == settings.MainBranch {
        settings.LastHash = block.CurrHash
    }
}

func blockPush(pack *models.PackageTCP) {
    var splited = strings.Split(pack.Body.Data, settings.SEPARATOR)
    if len(splited) != 2 {
        return
    }

    var block = mineBlock(pack.Body.Branch, pack.From.Hash, splited[0], splited[1])
    if block == nil {
        return
    }

    var new_pack = convertBlockToPackage(pack.Body.Branch, block)
    SendGlobalPackage(new_pack)
}

func mineBlock(branch, from, to, value_str string) *models.Block {
    value, err := strconv.ParseUint(value_str, 10, 64)
    if err != nil {
        return nil
    }
    var trans = blockchain.NewTransaction(branch, from, to, value)
    if trans == nil {
        return nil
    }
    var block = blockchain.NewBlock(trans, settings.GetLastHash(branch))
    blockchain.PushBlock(branch, block)
    return block
}

func blockHard(pack *models.PackageTCP) {
    var (
        block models.Block
        splited = strings.Split(pack.Body.Branch, settings.SEPARATOR)
    )

    if len(splited) != 2 {
        return
    }

    var (
        root_branch = splited[0]
        new_branch = splited[1]
    )

    err := json.Unmarshal([]byte(pack.Body.Data), &block)
    utils.CheckError(err) 
    blockchain.CopyBranch(root_branch, new_branch, block.PrevHash)
    _, err = settings.BlockChain.Exec(
        "INSERT INTO BlockChain (Branch, Hash, Block) VALUES ($1, $2, $3)",
        new_branch,
        hex.EncodeToString(block.CurrHash),
        pack.Body.Data,
    )
    utils.CheckError(err)
}

func getAllBlockChain(branch string) string {
    var (
        block_str string
        block models.Block
        blocks []models.Block
    )
    rows, err := settings.BlockChain.Query(
        "SELECT Block FROM BlockChain WHERE Branch=$1",
        branch,
    )
    utils.CheckError(err)
    defer rows.Close()

    for rows.Next() {
        rows.Scan(&block_str)
        err := json.Unmarshal([]byte(block_str), &block)
        utils.CheckError(err)
        blocks = append(blocks, block)
    }

    blocks_json, err := json.MarshalIndent(blocks, "", "\t")
    utils.CheckError(err)
    return string(blocks_json)
}

func connectSave(pack *models.PackageTCP) {
    settings.Connections[pack.From.Addr] = true
    var new_pack = models.PackageTCP {
        To: models.To {
            Addr: pack.From.Addr,
        },
        Head: models.Head {
            Title: settings.TITLE_BLOCK,
            Mode: settings.MODE_READ,
        },
        Body: models.Body {
            Data: getAllBlockChain(pack.Body.Branch),
            Branch: pack.Body.Branch,
        },
    }
    SendPackage(&new_pack)
}
